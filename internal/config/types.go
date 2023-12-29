package config

type Args struct {
	ConfigPath string
}

type Config struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Host        string `yaml:"host" env:"DB_HOST" env-description:"Database host"`
	Port        string `yaml:"port" env:"DB_PORT" env-description:"Database port"`
	Username    string `yaml:"username" env:"DB_USER" env-description:"Database user name"`
	Password    string `env:"DB_PASSWORD" env-description:"Database user password"`
	Name        string `yaml:"dbname" env:"DB_NAME" env-description:"Database name"`
	Connections int    `yaml:"connections" env:"DB_CONNECTIONS" env-description:"Total number of database connections"`
}
