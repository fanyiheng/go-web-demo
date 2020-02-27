package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"time"
)

type logConfig struct {
	OutputFile string
	Level      string
	Stdout     bool
}

type dbConfig struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	MaxIdleConns int
	MaxOpenConns int
	EnableLog bool
	ConnMaxLifetime int64
}


type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}


var cfg *ini.File
var LogSetting =&logConfig{}
var DatabaseSetting =&dbConfig{}
var ServerSetting = &Server{}

func Setup(configFile string) {
	if configFile == "" {
		configFile = "conf/app.ini"
	}
	var err error
	cfg, err = ini.Load(configFile)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	mapTo("log", LogSetting)
	fmt.Printf("log conf: %+v\n",*LogSetting)
	mapTo("database", DatabaseSetting)
	fmt.Printf("database conf: %+v\n",*DatabaseSetting)
	mapTo("server", ServerSetting)
	fmt.Printf("server conf: %+v\n",*ServerSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting er: %v", err)
	}
}
