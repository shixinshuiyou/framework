package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/shixinshuiyou/framework/log"
)

func init() {
	newLogger := logger.New(log.Logger, logger.Config{
		SlowThreshold:             sqlConfig.SlowThreshold, // 慢 SQL 阈值
		Colorful:                  true,                    // 禁用彩色打印
		IgnoreRecordNotFoundError: true,
		LogLevel:                  logger.Warn, // log level
	})
	dbLink := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		sqlConfig.User,
		sqlConfig.Pass,
		sqlConfig.Host,
		sqlConfig.Database)

	Db, err := gorm.Open(mysql.Open(dbLink),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
			Logger:                 newLogger,
			SkipDefaultTransaction: true,
		})
	if err != nil {
		log.Logger.Error("db err: %v", err)
	}

	db, err := Db.DB()
	if err != nil {
		log.Logger.Error("db err: %v", err)
	}
	db.SetMaxIdleConns(sqlConfig.MaxIdle)
	db.SetMaxOpenConns(sqlConfig.MaxDbConn)
	db.SetConnMaxLifetime(sqlConfig.Timeout)
}
