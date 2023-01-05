package conv

import (
	"fmt"
	"math"
	"strconv"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func Int(src interface{}) int {
	return int(Int64(src))
}

func Int64(src interface{}) (dst int64) {
	switch dst := src.(type) {
	case int8: // byte
		return int64(dst)
	case uint8:
		return int64(dst)
	case int16:
		return int64(dst)
	case uint16:
		return int64(dst)
	case int32: // rune
		return int64(dst)
	case uint32:
		return int64(dst)
	case int:
		return int64(dst)
	case uint:
		return int64(dst)
	case int64:
		return dst
	case uint64:
		if dst > math.MaxInt64 {
			// TODO
			return 0
		}
		return int64(dst)
	case string:
		out, _ := strconv.ParseInt(dst, 10, 64)
		return out
	case []byte:
		out, _ := strconv.ParseInt(string(dst), 10, 64)
		return out
	case float32:
		return int64(math.Floor(float64(dst)))
	case float64:
		return int64(math.Floor(dst))
	case nil:
		return 0
	case bool:
		if dst {
			return 1
		}
		return 0
	default:
		str := fmt.Sprintf("%v", src)
		out, _ := strconv.ParseInt(str, 10, 64)
		return int64(out)
	}
}

func Boolean(src interface{}) bool {
	switch dst := src.(type) {
	case bool:
		return dst
	case string:
		switch dst {
		case "true", "True", "TRUE", "Yes", "yes", "YES", "T", "t", "y", "Y":
			return true
		}
		return false
	default:
		return Int64(src) != 0
	}
}

func Float64(src interface{}) (dst float64) {
	switch src := src.(type) {
	case float64:
		return src
	case float32:
		return float64(src)
	case nil:
		return 0
	case int64:
		return float64(src)
	case int:
		return float64(src)
	case int8:
		return float64(src)
	case uint8:
		return float64(src)
	case int16:
		return float64(src)
	case uint16:
		return float64(src)
	case int32:
		return float64(src)
	case uint32:
		return float64(src)
	case uint64:
		return float64(src)
	case uint:
		return float64(src)
	case []byte:
		dst, _ = strconv.ParseFloat(string(src), 64)
		return dst
	case string:
		dst, _ = strconv.ParseFloat(src, 64)
		return dst
	case bool:
		if src {
			return 1
		}
		return 0
	default:
		str := fmt.Sprintf("%v", src)
		dst, _ = strconv.ParseFloat(str, 64)
		return
	}
}

func String(src interface{}) string {
	switch val := src.(type) {
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case string:
		return val
	case []byte:
		return string(val)
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", src)
	}
}

func StringSlice(src []interface{}) (dst []string) {
	for _, val := range src {
		dst = append(dst, String(val))
	}
	return
}

// 尽量使用gb13080
func Gb130802U8(src string) string {
	result, _ := simplifiedchinese.GB18030.NewDecoder().String(src)
	return result
}

func U82Gb13080(src string) string {
	result, _ := simplifiedchinese.GB18030.NewEncoder().String(src)
	return result
}

// gbk 为固定字节数，应使用扩展的gb13080编码
func GBK2U8(src string) string {
	result, _ := simplifiedchinese.GBK.NewDecoder().String(src)
	return result
}

func U82GBK(src string) string {
	result, _ := simplifiedchinese.GBK.NewEncoder().String(src)
	return result
}
