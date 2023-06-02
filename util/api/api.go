package api

import (
	"context"
)

type CallApi interface {
	GetPlatformName() string
	GetApiUrl() string
	Call(ctx context.Context, method, queryPath string, params map[string]string) ([]byte, error) // 接口调用实现 	//
}
