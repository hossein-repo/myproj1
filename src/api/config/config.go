package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
	Cors     CorsConfig
	Logger   LoggerConfig
	Otp      OtpConfig
	JWT      JWTConfig
}



type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}
type ServerConfig struct {
	InternalPort    string
	ExternalPort    string
	RunMode string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
}

type CorsConfig struct {
	AllowOrigins string
}

type OtpConfig struct {
	ExpireTime time.Duration
	Digits     int
	Limiter    time.Duration
}

type JWTConfig struct {
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
	Secret                     string
	RefreshSecret              string
}


func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("Error in load config %v", err)
	}

	cfg, err := ParseConfig(v)
	envPort := os.Getenv("PORT")
	if envPort != ""{
		cfg.Server.ExternalPort = envPort
		log.Printf("Set external port from environment -> %s", cfg.Server.ExternalPort)
	}else{
		cfg.Server.ExternalPort = cfg.Server.InternalPort
		log.Printf("Set external port from environment -> %s", cfg.Server.ExternalPort)
	}
	if err != nil {
		log.Fatalf("Error in parse config %v", err)
	}

	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}
func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
    v := viper.New()
    v.SetConfigType(fileType)
    v.SetConfigName(filepath.Base(filename))
    v.AddConfigPath(filepath.Dir(filename))
    v.AddConfigPath(".")
    v.AutomaticEnv()

    err := v.ReadInConfig()
    if err != nil {
        log.Printf("Unable to read config: %v", err)
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            return nil, errors.New("config file not found")
        }
        return nil, err
    }
    return v, nil
}
func getConfigPath(env string) string {
    if env == "docker" {
        return "./api/config/config-docker.yml"
    } else if env == "production" {
        return "./api/config/config-production.yml"
    } else {
        return "./api/config/config-development.yml"
    }
}