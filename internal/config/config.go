package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	HTTP ServerConfig
	DB   DBConfig
}

type ServerConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
	Username string `yaml:"username"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	err = cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
