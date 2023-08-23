package config

import (
	"os"
	"strconv"
)
const (
	KEY_DATABASE_HOST="DATABASE_HOST"
	KEY_DATABASE_NAME="DATABASE_NAME"
	KEY_DATABASE_USER="DATABASE_USER"
	KEY_DATABASE_PASSWORD="DATABASE_PASSWORD"
	KEY_ETHEREUM_NODE_URL="ETHEREUM_NODE_URL"
	KEY_WORKER_TOTAL_NUM="WORKER_TOTAL_NUM"
)
func GetDatabaseHost() string {
	return os.Getenv(KEY_DATABASE_HOST)
}

func GetDatabaseName() string {
	return os.Getenv(KEY_DATABASE_NAME)
}

func GetDatabaseUser() string {
	return os.Getenv(KEY_DATABASE_USER)
}

func GetDatabasePassword() string {
	return os.Getenv(KEY_DATABASE_PASSWORD)
}

func GetEthNodeUrl() string {
	return os.Getenv(KEY_ETHEREUM_NODE_URL)
}

func GetWorkerTotalNum() int {
	val, err := strconv.Atoi(os.Getenv(KEY_WORKER_TOTAL_NUM))
	if err != nil {
		return DEFAULT_WORKER_TOTAL_NUM
	}
	if val <= 1 {
		return DEFAULT_WORKER_TOTAL_NUM
	}
	return val
}