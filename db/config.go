package db

import "time"

var sqlConfig = &MySQLConfig{

}

type MySQLConfig struct {
	Host string `yaml:"host" json:"host" mapstructure:"host"`
	Port int    `yaml:"port" json:"port" mapstructure:"port"`
	User string `yaml:"user" json:"user" mapstructure:"user"`
	Pass string `yaml:"pass" json:"pass" mapstructure:"pass"`

	// default utf8
	Charset  string `yaml:"charset" json:"charset" mapstructure:"charset"`
	Database string `yaml:"database" json:"database" mapstructure:"database"`

	// default 0, 不保留空闲连接
	MaxIdle int `yaml:"max_idle" json:"max_idle" mapstructure:"max_idle"`

	// default 30s
	Timeout time.Duration `yaml:"timeout" json:"timeout" mapstructure:"timeout"` // s

	// default 300s
	ReadTimeout int `yaml:"read_timeout" json:"read_timeout" mapstructure:"read_timeout"` // s

	// default 0, 默认不限制最大连接数
	MaxDbConn int `yaml:"max_connection" json:"max_connection" mapstructure:"max_connection"`

	// default 1s
	SlowThreshold time.Duration `yaml:"slow_threshold" json:"slow_threshold" mapstructure:"slow_threshold"`
}

type MongoConfig struct {
	Url           string `yaml:"url" json:"url" mapstructure:"url"` // mongodb://10.172.195.159:7748
	User          string `yaml:"user" json:"user" mapstructure:"user"`
	Pass          string `yaml:"pass" json:"pass" mapstructure:"pass"`
	Database      string `yaml:"database" json:"database" mapstructure:"database"`
	MaxIdle       int    `yaml:"max_idle" json:"max_idle" mapstructure:"max_idle"`
	Timeout       int    `yaml:"timeout" json:"timeout" mapstructure:"timeout"`                      // ms
	SocketTimeout int    `yaml:"socket_timeout" json:"socket_timeout" mapstructure:"socket_timeout"` // ms  read or write timeout
	MaxPoolSize   uint64 `yaml:"max_pool_size" json:"max_pool_size" mapstructure:"max_pool_size"`
	MinPoolSize   uint64 `yaml:"min_pool_size" json:"min_pool_size" mapstructure:"min_pool_size"`
	MaxIdleTime   int    `yaml:"max_idle_time" json:"max_idle_time" mapstructure:"max_idle_time"` // ms
}
