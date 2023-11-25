package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload" // Load enviroment from .env
)

// TODO : learn consulapi "github.com/hashicorp/consul/api"

type Config struct {
	ConfigPath string `env:"CONFIG_PATH" env-default:"config/config.yaml"`
	HTTPServer `yaml:"httpServer"`
	AuthJWT    `yaml:"authJwt"`
	CSRFJWT    `yaml:"csrfJwt"`
	Grpc       `yaml:"grpc"`

	Database
	Enviroment     string `env:"ENVIROMENT" env-default:"prod" env-description:"avalible: local, dev, prod"`
	LogFilePath    string `env:"LOG_FILE_PATH" env-default:"zuzu.log"`
	PhotosFilePath string `env:"PHOTOS_FILE_PATH" env-default:"photos/"`
}

type HTTPServer struct {
	Address           string        `yaml:"address" yaml-default:"localhost:8080"`
	Timeout           time.Duration `yaml:"timeout" yaml-default:"4s"`
	IdleTimeout       time.Duration `yaml:"idleTimeout" yaml-default:"60s"`
	ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout" yaml-defualt:"10s"`
}

type AuthJWT struct {
	JwtAccess            string        `env:"AUTH_JWT_SECRET_KEY" env-required:"true"`
	AccessExpirationTime time.Duration `yaml:"accessExpirationTime" yaml-defualt:"6h"`
	Issuer               string
}
type Grpc struct {
	AuthPort  int `yaml:"authPort" yaml-defualt:"8011"`
	OrderPort int `yaml:"orderPort" yaml-defualt:"8012"`
}

func (a AuthJWT) GetTTL() time.Duration {
	return a.AccessExpirationTime
}
func (a AuthJWT) GetSecret() string {
	return a.JwtAccess
}
func (a AuthJWT) GetIssuer() string {
	return "auth"
}

type CSRFJWT struct {
	JwtAccess            string        `env:"CSRF_JWT_SECRET_KEY" env-required:"true"`
	AccessExpirationTime time.Duration `yaml:"accessExpirationTime" yaml-defualt:"6h"`
	Issuer               string
}

func (a CSRFJWT) GetTTL() time.Duration {
	return a.AccessExpirationTime
}
func (a CSRFJWT) GetSecret() string {
	return a.JwtAccess
}
func (a CSRFJWT) GetIssuer() string {
	return "csrf"
}

type Database struct {
	DBName string `env:"POSTGRES_DB" env-required:"true"`
	DBPass string `env:"POSTGRES_PASSWORD" env-required:"true"`
	DBHost string `env:"DB_HOST" env-default:"0.0.0.0"`
	DBPort int    `env:"DB_PORT" env-required:"true"`
	DBUser string `env:"POSTGRES_USER" env-required:"true"`
}

func (c Config) GetPhotosFilePath() string {
	// need for mock interface

	return c.PhotosFilePath

}

func MustLoad() *Config {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Printf("cannot read .env file: %s\n (fix: you need to put .env file in main dir)", err)
		os.Exit(1)
	}

	// check if config file exists
	if _, err := os.Stat(cfg.ConfigPath); os.IsNotExist(err) {
		log.Printf("config file does not exist: %s", cfg.ConfigPath)
		os.Exit(1)
	}

	if err := cleanenv.ReadConfig(cfg.ConfigPath, &cfg); err != nil {
		log.Printf("cannot read %s: %v", cfg.ConfigPath, err)
		os.Exit(1)
	}

	return &cfg
}
