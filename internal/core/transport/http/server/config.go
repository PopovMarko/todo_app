package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr            string        `envconfig:"ADDRESS" required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("HTTP", &config); err != nil {
		return Config{}, fmt.Errorf("server config init: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}
	return config
}
