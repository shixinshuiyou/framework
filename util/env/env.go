package env

import "os"

var envConf = "APP_ENV"

func GetEnv() string {
	result := os.Getenv(envConf)
	if result == "" {
		return "dev"
	}
	return result
}
