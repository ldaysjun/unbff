package config

// Config config info
type Config struct {
	PoolDB struct {
		DBName    string `yaml:"db_name"`    // DB name
		DBHost    string `yaml:"db_host"`    // DB host
		DBPort    int    `yaml:"db_port"`    // DB port
		DBUser    string `yaml:"db_user"`    // DB user
		DBPwd     string `yaml:"db_pwd"`     // DB pwd
		DBCharset string `yaml:"db_charset"` // DB charset
	} `yaml:"pool_db"`

	ProcessorDB struct {
		DBName    string `yaml:"db_name"`    // DB name
		DBHost    string `yaml:"db_host"`    // DB host
		DBPort    int    `yaml:"db_port"`    // DB port
		DBUser    string `yaml:"db_user"`    // DB user
		DBPwd     string `yaml:"db_pwd"`     // DB pwd
		DBCharset string `yaml:"db_charset"` // DB charset
	} `yaml:"processor_db"`
}
