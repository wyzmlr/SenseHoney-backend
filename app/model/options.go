package model

import (
	"SenseHoney/lib/utils/log"
	"fmt"
	"github.com/jinzhu/gorm"
)

// 首次插入
func (s *Service) InsertFirst(info SpInfo) {
	s.Mysql.Create(&info)
}

// 更新数据
func (s *Service) UpdateInfo(info SpInfo) {
	var oldInfo SpInfo
	// 旧数据拼接
	s.Mysql.Where(map[string]interface{}{"attack_ip": info.AttackIP, "type": info.Type}).Find(&oldInfo)
	if info.Valid == 1 { //此处注意bug
		s.Mysql.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("valid", oldInfo.Valid+1)
	} else {
		s.Mysql.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("invalid", oldInfo.Invalid+1)
	}
	s.Mysql.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("info", oldInfo.Info+"^^"+info.Info)
	// 更新攻击次数
	s.Mysql.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("count", oldInfo.Count+1)
}

func CheckMysql(user string, pass string, host string, name string) (bool, error) {
	_, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8",
		user,
		pass,
		host,
		name))
	if err != nil {
		return false, err
	}
	return true, err
}
func (s *Service) InitMysql() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8",
		s.Conf.Database.DbUser,
		s.Conf.Database.DbPass,
		s.Conf.Database.DbHost,
		s.Conf.Database.DbName,
	))
	if err != nil {
		log.Err("zh-CN", "", err)
	}

	s.Mysql = db
	// 创建表自动迁移
	s.Mysql.AutoMigrate(&SpAdmin{}, &SpInfo{}, &SpUser{}, &SpLog{})
}

// 检查是否已经存在该攻击者记录
func (s *Service) CheckIfExist(attackIp string, attackType string, token string) bool {
	var dataExi SpInfo
	res := s.Mysql.Where(map[string]interface{}{"attack_ip": attackIp, "type": attackType, "access_token": token}).Find(&dataExi).RowsAffected

	// 攻击者单次攻击记录存在则返回true 否则返回false
	if res > 0 {
		return true
	} else {
		return false
	}
}

func (s *Service) InsertLogFirst(info SpLog) {
	s.Mysql.Create(&info)
}

func (s *Service) UpdateLog(info SpLog) {
	var oldlogs SpLog
	// 旧数据拼接
	s.Mysql.Where(map[string]interface{}{"level": info.Level, "access_token": info.AccessToken}).Find(&oldlogs)
	// 更新攻击次数
	s.Mysql.Model(&info).Where("level = ? AND access_token = ?", info.Level, info.AccessToken).Update("count", oldlogs.Count+1)
}

func (s *Service) CheckIfLogExist(level string, token string) bool {
	var data SpLog
	res := s.Mysql.Where(map[string]interface{}{"level": level, "access_token": token}).Find(&data).RowsAffected
	// 攻击者单次攻击记录存在则返回true 否则返回false
	if res > 0 {
		return true
	} else {
		return false
	}
}

//func (s *Service) insertReportFirst(info SpLog) {
//	s.Mysql.Create(&info)
//}
//
//func (s *Service) insertReportCount(accessToken string) {
//	var dataLog SpLog
//	var oldData SpLog
//	r := s.Mysql.Where(map[string]interface{}{"level": "report", "access_token": accessToken}).Find(&dataLog)
//	fmt.Println(r.RowsAffected)
//	if r.RowsAffected > 0 {
//		s.Mysql.Where(map[string]interface{}{"level": "report", "access_token": accessToken}).Find(&oldData)
//		s.Mysql.Model(&dataLog).Where("level = ? AND access_token = ?", "report", accessToken).Update("count", oldData.Count+1)
//	} else {
//		s.Mysql.Create(&dataLog)
//	}
//}
