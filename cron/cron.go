package cron

import (
	"ThingsPanel-Go/initialize/psql"
	"ThingsPanel-Go/models"
	"ThingsPanel-Go/services"
	"encoding/json"
	"errors"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/bitly/go-simplejson"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

func init() {
	logs.Info("定时任务初始化")

	//自动化控制策略（时间条件类型-每日某个时间或者单次）
	//interval为0的时候触发一次；为1的时候每天触发
	//conditions表status属性1-启动 0-关闭;type属性1-设备触发 2-定时触发
	//config->{"rules":[{"time":"14:20","interval":1}],"apply":[{"asset_id":"xxx","field":"pm10","device_id":"xxx","value":"10"}]}
	//config->{"rules":[{"time":"2022/5/23 14:21","interval":0}],"apply":[{"asset_id":"xxx","field":"pm10","device_id":"xxx","value":"10"}]}
	//0-创建一分钟一次的定时器
	logs.Info("创建一分钟一次的定时调度")
	crontab := cron.New()
	spec := "*/60 * * * * ?"
	task := func() {
		//获取当前系统时间
		format0 := "2006/1/2 15:04"
		format1 := "15:04"
		now0, _ := time.Parse(format0, time.Now().Format(format0))
		now1, _ := time.Parse(format1, time.Now().Format(format1))
		logs.Info("当前时间-", now0)
		//1-根据status->1&&type->2&&config的rules存在interval查询出匹配的所有定时任务
		logs.Info("1-根据status->1&&type->2&&config的rules存在interval查询出匹配的所有定时任务")
		var conditionConfigs []models.Condition
		result := psql.Mydb.Model(&models.Condition{}).Where("type = 2 and status='1' order by sort asc").Find(&conditionConfigs)
		if result.Error != nil {
			errors.Is(result.Error, gorm.ErrRecordNotFound)
		} else {
			//2-循环判断是否触发
			logs.Info("2-循环判断是否触发")
			for _, row := range conditionConfigs {
				res, err := simplejson.NewJson([]byte(row.Config))
				if err != nil {
					logs.Error("解析出错", err)
					continue
				}
				logs.Info("2-1-解析config为json", res)
				rulesRows, _ := res.Get("rules").Array()
				for _, rulesRow := range rulesRows {
					if rulesMap, ok := rulesRow.(map[string]interface{}); ok {
						if rulesMap["interval"] != nil {
							interval, _ := rulesMap["interval"].(json.Number).Int64()
							if interval == int64(0) {
								//单次触发
								ruleTime, err := time.Parse(format0, rulesMap["time"].(string))
								logs.Info("单次触发", ruleTime, "比对", now0)
								if err == nil && ruleTime.Equal(now0) {
									//触发
									var DeviceService services.DeviceService
									DeviceService.ApplyControl(res)
								}
							} else if interval == int64(1) {
								//每天触发
								ruleTime, err := time.Parse(format1, rulesMap["time"].(string))
								logs.Info("每天触发", ruleTime, "比对", now1)
								if err == nil && ruleTime.Equal(now1) {
									//触发
									var DeviceService services.DeviceService
									DeviceService.ApplyControl(res)
								}
							}
						}
					}
				}
			}
		}

		//2-1触发
	}

	crontab.AddFunc(spec, task)
	crontab.Start()
	logs.Info("定时调度启动")

}
