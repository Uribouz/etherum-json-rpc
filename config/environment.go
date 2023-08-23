package config

func GetDatabaseName() string {
	return "ethereum-block"
}
func GetDatabaseUser() string {
	return "ball-database-mongodb-usr"
}

func GetDatabasePassword() string {
	return "ball123"
}

func GetDatabaseHost() string {
	return "cluster0.rsu0js8.mongodb.net"
}

func GetEthNodeUrl() string {
	//TDOO: use wss client instead.
	return "https://rpc.ankr.com/eth"
}

func GetWorkerTotalNum() int {
	return DEFAULT_WORKER_TOTAL_NUM
}