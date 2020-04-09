package config

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

