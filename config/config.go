package config

// Config Get all config variables
type Config struct {
	DB *DBConfig
	Port string
}

// DBConfig database configuration
type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

// GetConfig Exported function to get all config related stuff
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "root!",
			Name:     "rivendell",
			Charset:  "utf8",
		},
		Port: ":5000",
	}
}