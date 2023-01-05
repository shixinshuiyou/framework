package appconfig

import "time"

var DbSetting = Database{}

type Database struct {
	Type         string
	User         string
	Password     string
	Host         string
	Name         string
	TablePrefix  string
	MaxIdleConns int
	MaxOpenConns int
	LogMode      int
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	Db          int
	IdleTimeout time.Duration
}
