package settings

import (
	"SenseHoney/app/api"
	"SenseHoney/router"
	"fmt"
)

type Service struct {
	api.Service
}

func Init(s *Service) {
	s.InitConfig()
	s.InitMysql()
	s.Router = router.InitRouter((*router.Service)(s))
	//s.GetValidAttack()
	//s.DataInfo()
	//s.GetServiceCount()
	//s.genApiToken()
	//s.getServiceCount()
	//s.GetAttackInfo()
	fmt.Println("--------------------------------")
	panic(s.Router.Run(s.Conf.API.APIAddr))
}
