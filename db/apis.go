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
	result := db.Debug().
		Distinct("action_name,first_operator").
		Where("service_name", "fulfill-service").
		Order("first_operator").
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&apiList)
	if result.Error != nil {
		fmt.Println("查询数据错误")
		return nil, nil
	}
	orgIdList := make([]string, len(apiList))
	for i, api := range apiList {
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

func UpdateApi(api *Apis) error {
	db := getDB()
	if db == nil {
		fmt.Println("获取DB失败")
		return errors.New("获取DB失败")
	}
	params := make(map[string]interface{})
	params["service_name"] = api.ServiceName
	params["first_operator"] = api.FirstOperator
	params["is_deleted"] = 0
	db.Debug().Model(&Apis{}).Where("action_name = ?", "CreateEventExport").Updates(params)
	return nil
}
