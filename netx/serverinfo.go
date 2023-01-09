package netx

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
)

type Idc string

const (
	IdcZZDT  = Idc("zzdt")
	IdcSHBT  = Idc("shbt")
	IdcZZZC  = Idc("zzzc")
	IdcSHYC  = Idc("shyc")
	IdcBJYT  = Idc("bjyt")  // 北京电信
	IdcBJMD  = Idc("bjmd")  // 北京联通
	IdcBJZDT = Idc("bjzdt") // 北京电信25G
	IdcBJPDC = Idc("bjpdc") // 北京联通25G
)

const (
	// 针对容器服务做特出处理
	// 容器服务需要注入NODE_HOSTNAME 环境变量
	ENV_NODE_HOSTNAME = "NODE_HOSTNAME"
)

// 服务器变量初始化
type serverInfo struct {
	Hostname   string `yaml:"hostname"`
	Idcname    string `yaml:"idc"`
	ServerIp   string `yaml:"serverip"`
	ServerOs   string `yaml:"os"`
	ServerArch string `yaml:"arch"`
}

var localSrv = serverInfo{}

func init() {
	localSrv.Hostname, _ = os.Hostname()
	localSrv.Idcname = idc()
	localSrv.ServerIp = internalIp()
	localSrv.ServerOs = runtime.GOOS
	localSrv.ServerArch = runtime.GOARCH
}

func IDC() Idc {
	return Idc(localSrv.Idcname)
}

func Os() string {
	return localSrv.ServerOs
}

func Hostname() string {
	return localSrv.Hostname
}

func Arch() string {
	return localSrv.ServerArch
}

func InternalIp() string {
	return localSrv.ServerIp
}

func idc() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return "shbt"
	}

	if nodeHostname := os.Getenv(ENV_NODE_HOSTNAME); nodeHostname != "" {
		hostname = nodeHostname
	}
	hnArr := strings.Split(hostname, ".")
	if len(hnArr) > 3 {
		return hnArr[len(hnArr)-3]
	}
	return hostname
}

func internalIp() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "0.0.0.0"
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "0.0.0.0"
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}
	return "0.0.0.0"
}

func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}
