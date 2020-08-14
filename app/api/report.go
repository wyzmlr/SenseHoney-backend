package api

import (
	"SenseHoney/app/model"
	"SenseHoney/lib/error"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
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
		Valid:       1,
		Invalid:     1,
	}
	fmt.Println("-----------------------------------------------------")
	fmt.Println(spInfos.AttackIP+"|", spInfos.ClientIP+"|", spInfos.AccessToken+"|", spInfos.Type+"|", spInfos.WebApp+"|", spInfos.Info+"|", spInfos.Country+"|", spInfos.Region) //打印报告过来的数据
	if s.CheckIfExist(data.AttackIP, data.Type, data.AccessToken) {                                                                                                               //存在则更新否则插入数据库
		s.UpdateInfo(spInfos)
		//s.wsSend(s.dataInfo())//前端数据
		//log.DoLogs("success")
		return s.SuccessJSON("")
	}

	s.InsertFirst(spInfos)
	//log.DoLogs("success")
	return s.SuccessJSON("")
}
