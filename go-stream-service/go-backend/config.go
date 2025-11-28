package main

import "os"

func GetEnvString(key string, defaultValue string) string {
	// 環境変数の取得
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	} else {
		return defaultValue
	}
}
