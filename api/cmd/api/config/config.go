package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int16
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	Server struct {
		Address      int16
		CookieSecret string
	}
	Cloudinary struct {
		CloudName string
		ApiKey    string
		ApiSecret string
	}
}

func (c *Config) Parse() error {
	e := os.Getenv("ENVIRONMENT")
	switch e {
	case "development":
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.dev")
	default:
		return fmt.Errorf("environment variable ENVIRONMENT is not set")
	}
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("could not read config, err: %s", err.Error())
	}
	if err := viper.Unmarshal(&c); err != nil {
		return fmt.Errorf("could not unmarshal config, err: %s", err.Error())
	}
	return nil
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
