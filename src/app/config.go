package app

type Config struct {
	Database struct {
		Name     string `yaml:"name" env:"POSTGRES_DB"`
		Host     string `yaml:"host" default:"127.0.0.1" env:"POSTGRES_SERVER"`
		Port     int    `yaml:"port" default:"5432" env:"POSTGRES_PORT"`
		User     string `yaml:"user" env:"POSTGRES_USER"`
		Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
	} `yaml:"database"`
	Http struct {
		Address string `yaml:"address" default:"127.0.0.1:8080"`
	} `yaml:"http"`
}
