// author by lipengfei5 @2022-03-30

package profile

// import "github.com/pyroscope-io/client/pyroscope"

type Config struct {
	Enable       bool   `json:"enable" yaml:"enable" mapstructure:"enable"`
	Name         string `json:"name" yaml:"name" mapstructure:"name"`
	Tags         string `json:"tags" yaml:"tags" mapstructure:"tags"` // a:b;c:d
	Addr         string `json:"addr" yaml:"addr" mapstructure:"addr"`
	SampleRate   int64  `json:"sample_rate" yaml:"sample_rate" mapstructure:"sample_rate"`       // 每秒上传的频次
	ProfileTypes string `json:"profile_types" yaml:"profile_types" mapstructure:"profile_types"` // cpu,inuse_objects,alloc_objects,inuse_space,alloc_space
	AuthToken    string `json:"auth_token" yaml:"auth_token" mapstructure:"auth_token"`
	WhiteList    string `json:"white_list" yaml:"white_list" mapstructure:"white_list"` // ip 白名单，逗号分隔
}
