package config

import (
	"os"
	"strings"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/memory"
)

var Conf = config.DefaultConfig

func init() {

	Conf.Load(memory.NewSource(
		memory.WithJSON(GetEtcdAddrConf()),
	))

}

func GetEtcdAddrConf() []byte {
	addrs := strings.Split(os.Getenv("gi"), ",")
	etcdConfig := []byte(`{"etcdAddr": ["` + strings.Join(addrs, `","`) + `"]}`)
	return etcdConfig
}
