package gconv

import (
	"reflect"
	"strings"
)

// 任意类型转换为 map[string]interface{} 类型,
// 如果给定的输入参数i不是map类型，那么转换会失败，返回nil.
// 当i为struct对象时，第二个参数noTagCheck表示不检测json标签，否则将会使用json tag作为map的键名。
func Map(value interface{}, noTagCheck ...bool) map[string]interface{} {
	if value == nil {
		return nil
	}
	if r, ok := value.(map[string]interface{}); ok {
		return r
	} else {
		// 仅对常见的几种map组合进行断言，最后才会使用反射
		m := make(map[string]interface{})
		switch value.(type) {
		case map[interface{}]interface{}:
			for k, v := range value.(map[interface{}]interface{}) {
				m[String(k)] = v
			}
		case map[interface{}]string:
			for k, v := range value.(map[interface{}]string) {
				m[String(k)] = v
			}
		case map[interface{}]int:
			for k, v := range value.(map[interface{}]int) {
				m[String(k)] = v
			}
		case map[interface{}]uint:
			for k, v := range value.(map[interface{}]uint) {
				m[String(k)] = v
			}
		case map[interface{}]float32:
			for k, v := range value.(map[interface{}]float32) {
				m[String(k)] = v
			}
		case map[interface{}]float64:
			for k, v := range value.(map[interface{}]float64) {
				m[String(k)] = v
			}

		case map[string]bool:
			for k, v := range value.(map[string]bool) {
				m[k] = v
			}
		case map[string]int:
			for k, v := range value.(map[string]int) {
				m[k] = v
			}
		case map[string]uint:
			for k, v := range value.(map[string]uint) {
				m[k] = v
			}
		case map[string]float32:
			for k, v := range value.(map[string]float32) {
				m[k] = v
			}
		case map[string]float64:
			for k, v := range value.(map[string]float64) {
				m[k] = v
			}

		case map[int]interface{}:
			for k, v := range value.(map[int]interface{}) {
				m[String(k)] = v
			}
		case map[int]string:
			for k, v := range value.(map[int]string) {
				m[String(k)] = v
			}
		case map[uint]string:
			for k, v := range value.(map[uint]string) {
				m[String(k)] = v
			}
			// 不是常见类型，则使用反射
		default:
			rv := reflect.ValueOf(value)
			kind := rv.Kind()
			// 如果是指针，那么需要转换到指针对应的数据项，以便识别真实的类型
			if kind == reflect.Ptr {
				rv = rv.Elem()
				kind = rv.Kind()
			}
			switch kind {
			case reflect.Map:
				ks := rv.MapKeys()
				for _, k := range ks {
					m[String(k.Interface())] = rv.MapIndex(k).Interface()
				}
			case reflect.Struct:
				rt := rv.Type()
				name := ""
				for i := 0; i < rv.NumField(); i++ {
					// 只转换公开属性
					fieldName := rt.Field(i).Name
					if isLetterUpper(fieldName[0]) {
						continue
					}
					name = ""
					// 检查tag, 支持gconv, json标签, 优先使用gconv
					if len(noTagCheck) == 0 || !noTagCheck[0] {
						tag := rt.Field(i).Tag
						if name = tag.Get("gconv"); name == "" {
							name = tag.Get("json")
						}
					}
					if name == "" {
						name = strings.TrimSpace(fieldName)
					} else {
						// 支持标准库json特性: -, omitempty
						name = strings.TrimSpace(name)
						if name == "-" {
							continue
						}
						array := strings.Split(name, ",")
						if len(array) > 1 {
							switch strings.TrimSpace(array[1]) {
							case "omitempty":
								if isEmpty(rv.Field(i).Interface()) {
									continue
								} else {
									name = strings.TrimSpace(array[0])
								}
							default:
								name = strings.TrimSpace(array[0])
							}
						}
					}
					m[name] = rv.Field(i).Interface()
				}
			default:
				return nil
			}
		}
		return m
	}
}

func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	// 优先通过断言来进行常用类型判断
	switch value := value.(type) {
	case int:
		return value == 0
	case int8:
		return value == 0
	case int16:
		return value == 0
	case int32:
		return value == 0
	case int64:
		return value == 0
	case uint:
		return value == 0
	case uint8:
		return value == 0
	case uint16:
		return value == 0
	case uint32:
		return value == 0
	case uint64:
		return value == 0
	case float32:
		return value == 0
	case float64:
		return value == 0
	case bool:
		return value == false
	case string:
		return value == ""
	case []byte:
		return len(value) == 0
	default:
		// 最后通过反射来判断
		rv := reflect.ValueOf(value)
		if rv.IsNil() {
			return true
		}
		kind := rv.Kind()
		switch kind {
		case reflect.Map:
			fallthrough
		case reflect.Slice:
			fallthrough
		case reflect.Array:
			return rv.Len() == 0
		}
	}
	return false
}
