package config

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Env      string `mapstructure:"env"`
		HTPP     HTTPConfig
		Postgres PostgresConfig
		Jwt      JWTConfig
	}

	HTTPConfig struct {
		Host         string        `mapstructure:"host"`
		Port         uint16        `mapstructure:"port"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
	}

	PostgresConfig struct {
		Host     string `mapstructure:"host"`
		Port     uint16 `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"dbname"`
		Url      string `mapstructure:"url"`
	}

	JWTConfig struct {
		AccessTokenTTL time.Duration `mapstructure:"accessTokenTTL"`
		//RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SecretKey string `mapstructure:"secretKey"`
	}
)

func Init(cfgPath string) (*Config, error) {

	if cfgPath == "" {
		return nil, errors.New("config path is not set")
	}

	viper.AddConfigPath(cfgPath)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := unmarshalConfig(&cfg); err != nil {
		return nil, err
	}

	// the configuration has been successfully initialized
	return &cfg, nil
}

// unmarshal the configuration to fill in the structural data
func unmarshalConfig(cfg *Config) error {
	if err := viper.UnmarshalKey("env", &cfg.Env); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("http", &cfg.HTPP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("jwt", &cfg.Jwt); err != nil {
		return err
	}
	return nil
}
