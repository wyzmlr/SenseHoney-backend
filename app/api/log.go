package api

import (
	"SenseHoney/app/model"
	"SenseHoney/lib/error"
	"github.com/gin-gonic/gin"
)

func (s *Service) LogsHandler(c *gin.Context) (int, interface{}) {
	var data model.SpLog
	err := c.BindJSON(&data)
	if err != nil {
		return s.ErrJSON(500, error.ErrJsonParseCode, error.ErrJsonParseMsg)
	}

	spLogs := model.SpLog{
		Level:       data.Level,
		ClientIP:    c.ClientIP(),
		AccessToken: data.AccessToken,
		Count:       1,
	}

	if s.CheckIfLogExist(data.Level, data.AccessToken) {
		s.UpdateLog(spLogs)
		s.wsSend(s.DataInfo())
		//log.DoLogs("report")
		return s.SuccessJSON("")
	}
	s.InsertLogFirst(spLogs)
	s.wsSend(s.DataInfo())

	//s.updateLog(spLogs)
	return s.SuccessJSON("")
}
