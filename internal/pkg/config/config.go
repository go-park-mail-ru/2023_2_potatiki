package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload" //Load enviroment from .env
)

// TODO : learn consulapi "github.com/hashicorp/consul/api"

type Config struct {
	ConfigPath string `env:"CONFIG_PATH" env-default:"config/config.yaml"`
	Enviroment string `env:"enviroment" env-default:"local" env-description:"avalible: local, dev, prod"`
	Version    string `yaml:"version" yaml-required:"true"`
	HTTPServer `yaml:"http_server"`
	Database
	Auther `yaml:"auther"`
}

type HTTPServer struct {
	Address           string        `yaml:"address" yaml-default:"localhost:8080"`
	Timeout           time.Duration `yaml:"timeout" yaml-default:"4s"`
	IdleTimeout       time.Duration `yaml:"idle_timeout" yaml-default:"60s"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" yaml-defualt:"10s"`
}

type Database struct {
	DBName string `env:"POSTGRES_DB" env-required:"true"`
	DBPass string `env:"POSTGRES_PASSWORD" env-required:"true"`
	DBHost string `env:"POSTGRES_HOST" env-default:""`
	DBPort int    `env:"POSTGRES_PORT" env-required:"true"`
	DBUser string `env:"POSTGRES_USER" env-required:"true"`
}

type Auther struct {
	JwtAccess            string        `env:"TOKEN_ACCESS" env-required:"true"`
	AccessExpirationTime time.Duration `yaml:"access_expiration_time" yaml-defualt:"6h"`
}

func MustLoad() *Config {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("cannot read .env file: %s\n (fix: you need to put .env file in main dir)", err)
	}

	// check if config file exists
	if _, err := os.Stat(cfg.ConfigPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", cfg.ConfigPath)
	}

	if err := cleanenv.ReadConfig(cfg.ConfigPath, &cfg); err != nil {
		log.Fatalf("cannot read %s: %v", cfg.ConfigPath, err)
	}
	return &cfg
}
