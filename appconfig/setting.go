package appconfig

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

var cfg *ini.File

func init() {
	println(GetEnv())
	var err error
	var configPath string
	if IsProduct() {
		configPath = "conf/prod/app.ini"
	} else if IsOnline() {
		configPath = "conf/online/app.ini"
	} else if IsTest() {
		configPath = "conf/test/app.ini"
	} else {
		configPath = "conf/dev/app.ini"
	}
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		configPath = GetGoPath() + "/src/showapi/conf/dev/app.ini"
	}
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		configPath = GetGoPath() + "/src/adgit.src.corp.qihoo.net/MvWeb/showapi/conf/dev/app.ini"
	}
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatalf(" configPath %s not find ", configPath)
	}
	println("----configPath-----")
	println(configPath)
	println()

	cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
		return
	}

	mapTo("dbMediav", DbSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
