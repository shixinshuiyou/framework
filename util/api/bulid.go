package api

import (
	"context"
	"encoding/json"
	"github.com/shixinshuiyou/framework/conv"
	"github.com/shixinshuiyou/framework/log"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
	"time"
)

type BuildApi struct {
	CallApi
	Platform string
	Methods  map[string]RpcMethod
}

type RpcMethod struct {
	Method string
	Path   string
	Params []string
}

func NewRpc(api CallApi) *BuildApi {
	return &BuildApi{
		CallApi:  api,
		Platform: api.GetPlatformName(),
		Methods:  map[string]RpcMethod{},
	}
}

// Parse 解析Tag
func (api *BuildApi) Parse(field reflect.StructField) {
	tokens := strings.Split(field.Tag.Get("api"), ";")
	for _, v := range tokens {
		kv := strings.Split(strings.TrimSpace(v), ":")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "path":
			method := api.Methods[field.Name]
			method.Path = kv[1]
			api.Methods[field.Name] = method
		case "params":
			method := api.Methods[field.Name]
			method.Params = strings.Split(kv[1], ",")
			api.Methods[field.Name] = method
		case "method":
			method := api.Methods[field.Name]
			method.Method = kv[1]
			api.Methods[field.Name] = method
		}
	}
}

func (api *BuildApi) Method(field reflect.StructField) reflect.Value {
	return reflect.MakeFunc(field.Type, func(params []reflect.Value) []reflect.Value {
		reqParams := map[string]string{}
		method := api.Methods[field.Name]
		for i, v := range method.Params {
			reqParams[v] = conv.String(params[i+1].Interface())
		}

		start := time.Now()
		ctx := params[0].Interface().(context.Context)

		var (
			binResp []byte
			err     error
		)

		binResp, err = api.Call(ctx, method.Method, method.Path, reqParams)

		log.Logger.WithFields(logrus.Fields{
			"type":     "api_call",
			"cost":     float64(time.Since(start).Microseconds()) / 1000,
			"path":     method.Path,
			"method":   method.Method,
			"param":    reqParams,
			"platform": api.Platform,
		})

		val := reflect.New(field.Type.Out(0))
		if err != nil {
			return []reflect.Value{val.Elem(), reflect.ValueOf(err)}
		}

		err = json.Unmarshal(binResp, val.Interface())
		if err != nil {
			return []reflect.Value{val.Elem(), reflect.ValueOf(err)}
		}

		return []reflect.Value{val.Elem(), reflect.Zero(field.Type.Out(1))}
	})
}

func InitRpc(api CallApi) {
	rv := reflect.ValueOf(api).Elem()
	rt := reflect.TypeOf(api).Elem()
	rpc := NewRpc(api)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if !field.IsExported() {
			continue
		}
		if field.Type.Kind() == reflect.Func {
			rpc.Parse(rt.Field(i))
			rv.Field(i).Set(rpc.Method(rt.Field(i)))
		}
	}
}
