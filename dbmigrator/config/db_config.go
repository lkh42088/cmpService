package config

import "nubes/common/config"

// Mariadb of Customizing Contents Bridge
func GetNewDatabaseConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"nubes1510",
		"nubes",
		"192.168.10.115",
		3306,
	}
	return &config
}

// Mysql database of Contents Bridge
func GetOldDatabaseConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"Nubes1510!",
		"cdn_db_2020",
		"192.168.10.44",
		3306,
	}
	return &config
}

// Test Database of Contents Bridge
func GetTestCbDatabaseConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"Nubes1510!",
		"cb_db",
		"192.168.10.33",
		3306,
	}
	return &config
}
