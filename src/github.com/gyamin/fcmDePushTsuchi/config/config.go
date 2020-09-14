package config

type Config struct {
	DB database `toml:"database"`
}

type database struct {
	Server   string
	Port     int
	User     string
	Password string
	DbName   string
}
