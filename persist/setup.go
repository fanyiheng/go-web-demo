package persist

import (
	"fmt"
	"github.com/fanyiheng/go-web-demo/setting"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

var db *gorm.DB

func Setup() {
	//todo 参数验证和默认值处理
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		log.Fatalf("models.Setup er: %v", err)
	}

	//自动生成表时不会在表名后加s
	db.SingularTable(true)
	db.LogMode(setting.DatabaseSetting.EnableLog)
	db.SetLogger(dblog{})
	db.DB().SetMaxIdleConns(setting.DatabaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(setting.DatabaseSetting.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(setting.DatabaseSetting.ConnMaxLifetime) * time.Second)
}

type dblog struct {

}

func (l dblog) Print(v ...interface{}){
	logrus.Info(v)
}

