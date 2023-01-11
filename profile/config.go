package profile

// import "github.com/pyroscope-io/client/pyroscope"

type Config struct {
	Enable       bool
	Name         string
	Tags         string
	Addr         string
	SampleRate   int64  // 每秒上传的频次
	ProfileTypes string // cpu,inuse_objects,alloc_objects,inuse_space,alloc_space
	AuthToken    string
	WhiteList    string // ip 白名单，逗号分隔
}
