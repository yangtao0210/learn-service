package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

const openKey = "openKey"

var redisCli *redis.Client

// 普通连接模式
func getNormalCon() {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //密码
		DB:       0,  //数据库
		PoolSize: 20, //连接池大小
	})
}

// getTLSCon TLS连接模式
func getTLSCon() {
	redisCli = redis.NewClient(&redis.Options{
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	})
}

// getSentinelCon 哨兵模式
func getSentinelCon() {
	redisCli = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "redis507",
		SentinelAddrs: []string{":9126", ":9127", ":9128"},
	})
}

// getClusterCon 集群模式
func getClusterCon() *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":9126", ":9127", ":9128"},
	})
}

// 执行命令
func doCommand() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 执行命令获取结果
	val, err := redisCli.Get(ctx, "key1").Result()
	fmt.Println(val, err)

	// 先获取到命令对象
	cmder := redisCli.Get(ctx, "key")
	fmt.Println(cmder.Val()) // 获取值
	fmt.Println(cmder.Err()) // 获取错误

	// 直接执行命令获取错误
	err = redisCli.Set(ctx, "key", 10, time.Hour).Err()

	// 直接执行命令获取值
	value := redisCli.Get(ctx, "key").Val()
	fmt.Println(value)
}

// 执行任意命令
func doDemo() {
	// 执行命令获取结果
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	// 直接执行命令获取错误
	err := redisCli.Do(ctx, "set", "gk", 100, "EX", 3600).Err()
	fmt.Println(err)

	// 执行命令获取结果
	val, err := redisCli.Do(ctx, "get", "gk").Result()
	fmt.Println(val, err)
}

// zset操作
func zsetDemo() {
	zsetKey := "language_rank"
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"}}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// zadd
	err := redisCli.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed. err : %v\n", err)
		return
	}
	fmt.Println("zadd success!")

	//Golang分数+10
	result, err := redisCli.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("ZIncrBy failed. err : %v\n", err)
		return
	}
	fmt.Printf("Golang`s score is %f now.\n", result)

	//取分数最高的三个
	fmt.Println("======== zset top 3: ========")
	zres := redisCli.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range zres {
		fmt.Println(z.Member, z.Score)
	}

	//取95-100区间元素
	zs, err := redisCli.ZRangeByScoreWithScores(ctx, zsetKey, &redis.ZRangeBy{Min: "95", Max: "100"}).Result()
	if err != nil {
		fmt.Printf("ZRangeByScoreWithScores failed. err : %v\n", err)
		return
	}
	fmt.Println("======== zset score [95-100]: ========")
	for _, z := range zs {
		fmt.Println(z.Member, z.Score)
	}
}

// 获取键值并给定默认值
func getValueFromRedis(key, defaultValue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	val, err := redisCli.Get(ctx, key).Result()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			return defaultValue, nil
		}
		return "", err
	}
	return val, nil
}

func is_open_task(is_run bool) error {
	if !is_run {
		// 停止定时任务：设置永久key，保证pod拿不到锁
		res, err := redisCli.Set(context.Background(), openKey, "stop", -1).Result()
		if err != nil {
			log.Fatalln("定时任务停止失败", err)
			return err
		}
		log.Println("定时任务是否成功停止:", res)
	} else {
		// 开启定时任务：删除key,保证pod可以成功拿到锁
		res, err := redisCli.Del(context.Background(), openKey).Result()
		if err != nil {
			log.Fatalln("定时任务开启失败", err)
			return err
		}
		log.Println("定时任务是否成功重启:", res)
	}
	return nil
}

func main() {
	getNormalCon()
	// false-停止，true-开启
	err := is_open_task(false)
	if err != nil {
		log.Fatalln("设置定时任务开关失败")
		return
	}
	// 定时任务
	c := cron.New(cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron:", log.LstdFlags))))
	c.AddFunc("@every 5s", func() {
		ok := redisCli.SetNX(context.Background(), openKey, "run", time.Second*4)
		if ok.Val() {
			log.Println("获取redis锁成功")
		} else {
			log.Println("获取redis锁失败")
		}
	})
	c.Start()
	time.Sleep(time.Second * 30)
}
