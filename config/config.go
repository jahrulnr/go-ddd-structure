package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"strconv"
)

type (
	Config struct {
		App        `yaml:"app"`
		Server     `yaml:"server"`
		Log        `yaml:"log"`
		Mysql      `yaml:"mysql"`
		Redis      `yaml:"redis"`
		Additional `yaml:"additional"`
	}

	App struct {
		Name        string `yaml:"name" env:"APP_NAME"`
		Environment string `yaml:"environment" env:"APP_ENVIRONMENT"`
		BaseDir     string `yaml:"base_dir" env:"APP_BASE_DIR"`
		OnPrime     bool   `yaml:"on_prime" env:"APP_ON_PRIME"`
	}

	Server struct {
		Host     string `yaml:"host" env:"SERVER_HOST"`
		Port     string `yaml:"port" env:"SERVER_PORT"`
		UseSSL   bool   `yaml:"use_ssl" env:"SERVER_USE_SSL"`
		SSLCert  string `yaml:"ssl_cert" env:"SERVER_SSL_CERT"`
		SSLKey   string `yaml:"ssl_key" env:"SERVER_SSL_KEY"`
		UseHTTP2 bool   `yaml:"use_http2" env:"SERVER_USE_HTTP2"`
	}

	Log struct {
		Name  string `yaml:"name" env:"LOG_NAME"`
		Level string `yaml:"level" env:"LOG_LEVEL"`
		Path  string `yaml:"path" env:"LOG_PATH"`
	}

	// Mysql Database
	Mysql struct {
		Driver           string `yaml:"driver" env:"MYSQL_DRIVER"`
		Host             string `yaml:"host" env:"MYSQL_HOST"`
		MaxOpenConns     int    `yaml:"max_open_conns" env:"MYSQL_MAX_OPEN_CONNS"`
		MaxIdleConns     int    `yaml:"max_idle_conns" env:"MYSQL_MAX_IDLE_CONNS"`
		MaxLifetimeConns int    `yaml:"max_lifetime" env:"MYSQL_MAX_LIFETIME_CONNS"`
		MaxIdleTime      int    `yaml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`
	}

	// Redis cache
	Redis struct {
		Address        string `yaml:"address" env:"REDIS_ADDRESS"`
		DB             int    `yaml:"db" env:"REDIS_DB"`
		Password       string `yaml:"password" env:"REDIS_PASSWORD"`
		ContextTimeout int    `yaml:"context_timeout" env:"REDIS_CONTEXT_TIMEOUT"`
	}

	Additional struct {
		AdminFee int64
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)

	isOn, _ := strconv.ParseBool(os.Getenv("APP_ON_PRIME"))
	if !isOn {
		err = cleanenv.ReadConfig("config/config.yml", cfg)
	}

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
