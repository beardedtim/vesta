package env

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)

	if !exists {
		return value, fmt.Errorf("cannot get env %s from environment - please ensure it is set", key)
	}

	return value, nil
}

func GetEnvStr(key string) string {
	value, err := GetEnv(key)

	if err != nil {
		panic(err)
	}

	return value
}

func GetEnvInt(key string) int {
	value, err := GetEnv(key)

	if err != nil {
		panic(err)
	}

	parsed, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		panic(err)
	}

	return int(parsed)
}
