package api

import (
	"SenseHoney/app/model"
	"SenseHoney/lib/error"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

// TODO: 处理蜜罐上报信息 并存入数据库
var data model.SpInfo

type Service struct {
	model.Service
}

func (s *Service) ReportHandler(c *gin.Context) (int, interface{}) {
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("JSON解析失败！")
		os.Exit(1)
		return s.ErrJSON(500, error.ErrJsonParseCode, error.ErrJsonParseMsg)
	}
	isValid := data.Valid
	isInvalid := 0
	if isValid != 1 {
		isInvalid = 1
	}

	spInfos := model.SpInfo{
		AttackIP:    data.AttackIP,
		ClientIP:    c.ClientIP(),
		AccessToken: data.AccessToken,
		Type:        data.Type,
		WebApp:      data.WebApp,
		Info:        data.Info,
		Count:       1,
		Country:     data.Country,
		City:        data.City,
		Region:      data.Region,
		Valid:       isValid,
		Invalid:     uint(isInvalid),
	}
	fmt.Println("-----------------------------------------------------")
	fmt.Println(spInfos.AttackIP+"|", spInfos.ClientIP+"|", spInfos.AccessToken+"|", spInfos.Type+"|", spInfos.WebApp+"|", spInfos.Info+"|", spInfos.Country+"|", spInfos.Region) //打印报告过来的数据
	if s.CheckIfExist(data.AttackIP, data.Type, data.AccessToken) {                                                                                                               //存在则更新否则插入数据库
		s.UpdateInfo(spInfos)
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		s.wsSend(s.DataInfo(spInfos.Type, spInfos.AttackIP, spInfos.Info, spInfos.City, timeStr)) //前端数据
		//log.DoLogs("success")
		return s.SuccessJSON("")
	}

	s.InsertFirst(spInfos)
	//log.DoLogs("success")
	return s.SuccessJSON("")
}
