package env

import (
	"os"
	"strconv"
)

/**
 * 環境変数から、boolで取得
 * @param name string 定義名
 * @param def bool デフォルト値
 */
func GetBool(name string, def bool) bool {
	var res bool
	res, _ = strconv.ParseBool(getenv(name, strconv.FormatBool(def)))
	return res
}

/**
 * 環境変数から、intで取得
 * @param name string 定義名
 * @param def int デフォルト値
 */
func GetInt(name string, def int) int {
	var res int
	var defStr string
	defStr = strconv.Itoa(def)
	res, err := strconv.Atoi(getenv(name, defStr))
	if err != nil {
		return def
	}
	return res
}

/**
 * 環境変数から、stringで取得
 * @param name string 定義名
 * @param def string デフォルト値
 */
func GetStr(name string, def string) string {
	return getenv(name, def)
}

/**
 * 環境変数から取得
 * @param name string 定義名
 * @param def string デフォルト値
 */
func getenv(key, def string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return def
}
