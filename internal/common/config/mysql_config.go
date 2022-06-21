package config

func MySQL() mysqlConfig {
	if !hasMySQLConfigLoaded {
		loadMySQLConfig()
		hasMySQLConfigLoaded = true
	}
	return mysql
}

var mysql mysqlConfig
var hasMySQLConfigLoaded = false

type mysqlConfig struct {
	DSN string
}

func loadMySQLConfig() {
	mysql = mysqlConfig{
		DSN: getENV("MYSQL_DB_CONNECTION_LINK", "admin:password@tcp(127.0.0.1:3306)/shape?charset=utf8mb4&parseTime=True&loc=Local"),
	}
}
