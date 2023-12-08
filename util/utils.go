package util

import (
	"context"
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetEnvInt(key string, fallback int) int {
	env := GetEnv(key, "")
	if env == "" {
		return fallback
	}
	value, err := strconv.Atoi(env)
	if err != nil {
		return fallback
	}
	return value
}

func GetContextValue(ctx context.Context, key string) string {
	value := ctx.Value(key)
	if value == nil {
		return ""
	}
	ret, ok := value.(string)
	if !ok {
		return ""
	}
	return ret
}
