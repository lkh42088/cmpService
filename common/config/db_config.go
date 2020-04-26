package config

type MongoDbConfig struct {
	// MongoDB
	MongoIp         string `json:"mongo_ip"`
	MongoDb         string `json:"mongo_db"`
	MongoCollection string `json:"mongo_collection"`
}

type InfluxDbConfig struct {
	// InfluxDB
	InfluxIp       	string `json:"influx_ip"`
	InfluxDb       	string `json:"influx_db"`
	InfluxUser     	string `json:"influx_user"`
	InfluxPassword 	string `json:"influx_password"`
}

type MariaDbConfig struct {
	MariaIp 		string `json:"mariadb_ip"`
	MariaDb 		string `json:"mariadb_db"`
	MariaUser 		string `json:"mariadb_user"`
	MariaPassword 	string `json:"mariadb_password"`
}

// Database Configuration
type DBConfig struct {
	// database driver
	DBDriver string
	// db user
	Username string
	// db password
	Password string
	// database name
	DBName string
	// IP address of Database server
	Address string
	// Port number of Database server
	Port int
}

func NewDBConfig(driver, user, passwd, dbName, addr string, port int) (*DBConfig, error) {
	return &DBConfig{
		driver,
		user,
		passwd,
		dbName,
		addr,
		port,
	}, nil
}

