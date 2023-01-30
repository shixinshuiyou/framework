package database

import (
	"fmt"
	"github.com/shixinshuiyou/framework/database/my_logger"
	"github.com/shixinshuiyou/framework/log"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type DatabaseConfig struct {
	Type          string
	User          string
	Password      string
	Host          string
	Name          string
	MaxIdle       int
	MaxOpen       int
	LogLevel      logrus.Level
	SlowThreshold time.Duration
	Colorful      bool
}

func (conf *DatabaseConfig) InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Name)),
		&gorm.Config{ // 官方文档：https://gorm.io/docs/gorm_config.html
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: my_logger.New(log.Logger.WithField("", ""), my_logger.Config{
				SlowThreshold: conf.SlowThreshold,
				Colorful:      conf.Colorful,
				LogLevel:      gormLogger.LogLevel(conf.LogLevel),
			}),
		})
	if err != nil {
		return nil, err
	}
	sqlDB, err1 := db.DB()
	if err1 != nil {
		log.Logger.Warnf("db err: %v", err1)
	}
	sqlDB.SetMaxIdleConns(conf.MaxIdle)
	sqlDB.SetMaxOpenConns(conf.MaxOpen)
	sqlDB.SetConnMaxLifetime(100 * time.Second)

	return db, nil
}
