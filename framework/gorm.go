package framework

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-template/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Gorm struct {
	Session *gorm.DB
}

var gormInstance = &Gorm{
	Session: newSession(),
}

func GormInstance() *Gorm {
	return gormInstance
}

func newSession() *gorm.DB {
	l := log.StandardLogger()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configuration.Configs["db.username"], configuration.Configs["db.password"], configuration.Configs["db.host"], configuration.Configs["db.port"], configuration.Configs["db.database"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(l, logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}),
	})
	if err != nil {
		panic(err)
	}
	return db
}
