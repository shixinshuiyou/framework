package appconfig

import "os"

func GetEnv() string {
	result := os.Getenv("GO_ENV")
	if result == "" {
		return "dev"
	}
	return result
}

func IsProduct() bool {
	return GetEnv() == "prod"
}
func IsOnline() bool {
	return GetEnv() == "online"
}
func IsDev() bool {
	result := GetEnv() == "dev"
	if result {
		return true
	}
	if !IsProduct() && !IsTest() && !IsOnline() {
		return true
	}
	return false
}

func IsTest() bool {
	return GetEnv() == "test"
}

func GetGoPath() string {
	return os.Getenv("GOPATH")
}
