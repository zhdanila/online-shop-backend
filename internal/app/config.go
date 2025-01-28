package app

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	HTTPPort string `mapstructure:"HTTP_PORT" validate:"required"`
}

func (cnf *Config) GetPort() string {
	return cnf.HTTPPort
}

func NewConfig() (*Config, error) {
	baseCnf := &Config{}

	if err := baseCnf.Load(); err != nil {
		return nil, err
	}

	return baseCnf, nil
}

func (cnf *Config) Load() error {
	viper.SetConfigFile(`.env`)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	if err := viper.Unmarshal(cnf); err != nil {
		return fmt.Errorf("failed to unmarshal config: %s", err)
	}

	zap.L().Info("Config successfully loaded.")
	return nil
}
