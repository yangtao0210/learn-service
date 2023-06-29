package db

import (
	"errors"
	"fmt"
	"time"
)

type Apis struct {
	ActionName    string
	ServiceName   string
	FirstOperator string
	IsDeleted     int
}

func OffsetGetNameList(offset, limit int64) ([]string, error) {
	db := getDB()
	if db == nil {
		fmt.Println("获取DB失败")
		return nil, nil
	}
	end := time.Now().Unix()
	start := time.Now().AddDate(0, 0, -5).Unix()
	fmt.Println(start, end)
	apiList := make([]*Apis, 0)
	db = db.Debug().
		Distinct("action_name,first_operator").
		Where("service_name", "fulfill-service")
	orders := []string{"action_name", "first_operator"}
	for _, v := range orders {
		db = db.Order(v)
	}
	db = db.Offset(int(offset)).Limit(int(limit))
	if db.Find(&apiList).Error != nil {
		fmt.Println("查询数据错误")
		return nil, nil
	}
	orgIdList := make([]string, len(apiList))
	for i, api := range apiList {
		fmt.Println(api)
		orgIdList[i] = api.ActionName
	}
	return orgIdList, nil
}

func SaveApi(api *Apis) error {
	db := getDB()
	if db == nil {
		fmt.Println("获取DB失败")
		return errors.New("获取DB失败")
	}
	db.Debug().Create(api)
	return nil
}
func DeleteApiRecord(api *Apis) error {
	db := getDB()
	if db == nil {
		fmt.Println("获取DB失败")
		return errors.New("获取DB失败")
	}

	searchMp := map[string]interface{}{
		"action_name":  api.ActionName,
		"service_name": api.ServiceName,
		"is_deleted":   0,
	}

	if err := db.Debug().Where(searchMp).Delete(&Apis{}).Error; err != nil {
		return err
	}
	return nil
}
func UpdateApi(api *Apis) error {
	db := getDB()
	if db == nil {
		fmt.Println("获取DB失败")
		return errors.New("获取DB失败")
	}
	filters := make(map[string]interface{})
	filters["action_name"] = "CreateEventNotify"
	filters["is_deleted"] = 0
	db = db.Debug().Model(&Apis{})
	for k, v := range filters {
		db = db.Where(fmt.Sprintf("%v = ?", k), v)
	}
	params := make(map[string]interface{})
	params["service_name"] = api.ServiceName
	params["first_operator"] = api.FirstOperator
	params["is_deleted"] = 1
	db.Updates(params)
	return nil
}
