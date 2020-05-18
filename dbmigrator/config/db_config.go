package config

import "cmpService/common/config"

// Mariadb of Customizing Contents Bridge
func AGetNewDatabaseConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"Nubes1510!",
		"nubes",
		"192.168.227.129",
		3306,
	}
	return &config
}

// Mysql database of Contents Bridge
func AGetOldDatabaseConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"Nubes1510!",
		"cdn_db_2020",
		"192.168.160.130",
		3306,
	}
	return &config
}

// Test Database of Contents Bridge
func GetTestCbDatabaseConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"root",
		"Nubes1510!",
		"test_cdn_db",
		"192.168.121.154",
		3306,
	}
	return &config
}

/*******************************************************************
* Jungbh
*******************************************************************/
// Mariadb of Customizing Contents Bridge
func GetNewDatabaseConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"Nubes1510!",
		"nubes",
		"192.168.122.214",
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
		"192.168.122.30",
		3306,
	}
	return &config
}
