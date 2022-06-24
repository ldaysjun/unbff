package config

type DB struct {
	DBName    string `yaml:"db_name"`    // DB name
	DBHost    string `yaml:"db_host"`    // DB host
	DBPort    int    `yaml:"db_port"`    // DB port
	DBUser    string `yaml:"db_user"`    // DB user
	DBPwd     string `yaml:"db_pwd"`     // DB pwd
	DBCharset string `yaml:"db_charset"` // DB charset
}

// Config config info
type Config struct {
	PoolDB      DB `yaml:"pool_db"`
	ProcessorDB DB `yaml:"processor_db"`
}

func ForTest() *Config {
	return &Config{
		PoolDB: DB{
			DBName:    "",
			DBHost:    "",
			DBPort:    0,
			DBUser:    "",
			DBPwd:     "",
			DBCharset: "",
		},
		ProcessorDB: DB{
			DBName:    "entry_task",
			DBHost:    "127.0.0.1",
			DBPort:    3306,
			DBUser:    "ldaysjun",
			DBPwd:     "gv123456",
			DBCharset: "utf8",
		},
	}
}
