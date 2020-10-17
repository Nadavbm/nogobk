package env

import (
	"os"
	"strconv"
)

// db connection variables
var DatabaseUser = GetEnvVar("DATABASE_USER")
var DatabasePass = GetEnvVar("DATABASE_PASSWORD")
var DatabaseDB = GetEnvVar("DATABASE_DB")
var DatabaseHost = GetEnvVar("DATABASE_HOST")
var DatabasePort = GetEnvInt("DATABASE_PORT")

// html and css location
var HtmlPath = GetEnvVar("HTML_PATH")
var CssPath = GetEnvVar("CSS_PATH")

func GetEnvVar(key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return ""
}

func GetEnvInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return value
	}
	return value
}
