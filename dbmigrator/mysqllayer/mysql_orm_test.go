package mysqllayer

import "nubes/common/config"

func getMysqlConfig() *config.DBConfig {
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