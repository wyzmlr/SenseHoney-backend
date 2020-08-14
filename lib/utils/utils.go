package utils

import (
	"SenseHoney/lib/error"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/websocket"
	"log"
)

type Service struct {
	Conf   Config
	Router *gin.Engine
	Mysql  *gorm.DB
	Ws     *websocket.Conn
}

type Config struct {
	Logs struct {
		ErrorLog string `toml:"errorLog"`
		InfoLog  string `toml:"infoLog"`
		DebugLog string `toml:"debugLog"`
	} `toml:"logs"`

	API struct {
		APIAddr string `toml:"apiAddr"`
	} `toml:"api"`

	Database struct {
		DbType string `toml:"dbType"`
		DbHost string `toml:"dbHost"`
		DbUser string `toml:"dbUser"`
		DbPass string `toml:"dbPass"`
		DbName string `toml:"dbName"`
	} `toml:"database"`
}

func GetConfig() Config {
	var fg Config
	var fgPath string = "config/config.toml"
	if _, err := toml.DecodeFile(fgPath, &fg); err != nil {
		log.Fatal(err)
	}
	return fg
}

func (s *Service) InitConfig() {
	s.Conf = GetConfig()
}

func (s *Service) SuccessJSON(data interface{}) (int, interface{}) {
	return 200, gin.H{"error": error.SuccessCode, "msg": error.SuccessMsg, "data": data}
}
func (s *Service) ErrJSON(httpStatusCode int, errCode int, msg interface{}) (int, interface{}) {
	return httpStatusCode, gin.H{"error": errCode, "msg": fmt.Sprint(msg)}
}
