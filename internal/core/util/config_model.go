package util

import "time"

type Config struct {
	Database Database   `mapstructure:"database"`
	Jwt      Jwt        `mapstructure:"jwt"`
	Server   Server     `mapstructure:"server"`
	Cors     CorsConfig `mapstructure:"cors"`
	Argon2   Argon2     `mapstructure:"argon2"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type CorsConfig struct {
	AllowOrigins string `mapstructure:"allow_origins"`
}

type Argon2 struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

type Database struct {
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	Name            string        `mapstructure:"dbName"`
	SSLMode         string        `mapstructure:"sslmode"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

type Jwt struct {
	Secret                    string        `mapstructure:"secret"`
	AccessTokenExpireDuration time.Duration `mapstructure:"access_token_expire_duration"`
}
