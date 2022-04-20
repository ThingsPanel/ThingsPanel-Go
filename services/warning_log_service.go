package services

import (
	"ThingsPanel-Go/initialize/psql"
	"ThingsPanel-Go/models"
	uuid "ThingsPanel-Go/utils"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type WarningLogService struct {
	//可搜索字段
	SearchField []string
	//可作为条件的字段
	WhereField []string
	//可做为时间范围查询的字段
	TimeField []string
}

// Paginate 分页获取WarningLog数据
func (*WarningLogService) Paginate(name string, offset int, pageSize int, startDate string, endDate string) ([]models.WarningLog, int64) {
	var warningLogs []models.WarningLog

	//日期条件筛选
	sqlWhere := "1=1"
	if startDate != "" {
		//2021/12/01 15:12:37
		tmp := "2006/1/2 15:4:5"
		sDate, _ := time.ParseInLocation(tmp, startDate, time.Local)
		eDate, _ := time.ParseInLocation(tmp, endDate, time.Local)
		sqlWhere += " and created_at between " + strconv.FormatInt(sDate.Unix(), 10) + " and " + strconv.FormatInt(eDate.Unix(), 10)
	}
	var count int64
	countResult := psql.Mydb.Model(&warningLogs).Where(sqlWhere).Count(&count)
	if countResult.Error != nil {
		errors.Is(countResult.Error, gorm.ErrRecordNotFound)
	}
	offset = pageSize * (offset - 1)
	result := psql.Mydb.Where(sqlWhere).Limit(pageSize).Offset(offset).Order("created_at desc").Find(&warningLogs)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	if len(warningLogs) == 0 {
		warningLogs = []models.WarningLog{}
	}
	return warningLogs, count
}

// 根据id获取100条WarningLog数据
func (*WarningLogService) GetList(offset int, pageSize int) ([]models.WarningLog, int64) {
	var warningLogs []models.WarningLog
	result := psql.Mydb.Limit(pageSize).Offset(offset).Find(&warningLogs)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	if len(warningLogs) == 0 {
		warningLogs = []models.WarningLog{}
	}
	return warningLogs, result.RowsAffected
}

// Add新增一条WarningLogService数据
func (*WarningLogService) Add(t string, describe string, data_id string) (bool, string) {
	var uuid = uuid.GetUuid()
	warningLog := models.WarningLog{
		ID:        uuid,
		Type:      t,
		Describe:  describe,
		DataID:    data_id,
		CreatedAt: time.Now().Unix(),
	}
	result := psql.Mydb.Create(&warningLog)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
		return false, ""
	}
	return true, uuid
}

// 根据wid获取告警信息
func (*WarningLogService) WarningForWid(device_id string, limit int) []models.WarningLogView {
	var warningLogs []models.WarningLogView
	// 通过图表的wid去widget中找到device_id(设备id)，再去warning_log中匹配data_id
	sqlWhere := `SELECT warning_log.*,device.name as device_name  FROM warning_log left join device on 
	warning_log.data_id = device.id WHERE data_id = ? ORDER BY created_at desc LIMIT ?`
	result := psql.Mydb.Raw(sqlWhere, device_id, strconv.Itoa(limit)).Scan(&warningLogs)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	if len(warningLogs) == 0 {
		warningLogs = []models.WarningLogView{}
	}
	return warningLogs
}
