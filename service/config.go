package service

import (
	"strings"

	"github.com/spf13/viper"
)

// Config holds the configuration values for the backend.
type Config struct {
}

// NewConfig loads the config file into the Config struct.
func NewConfig() *Config {
	config := viper.New()
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)
	config.AutomaticEnv()

	return &Config{
	}
}
