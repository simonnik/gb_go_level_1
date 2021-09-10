package config

import (
	"errors"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/url"
)

type Config struct {
	Port        string `envconfig:"PORT" default:"8080" required:"true"`
	DbUrl       string `envconfig:"DB_URL" default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" required:"true"`
	JaegerUrl   string `envconfig:"JAEGER_URL" default:"http://jaeger:16686" required:"true"`
	KafkaBroker string `envconfig:"KAFKA_BROKER" default:"kafka:9092" required:"true"`
}

func New() (*Config, error) {
	conf := &Config{}
	err := envconfig.Process("", conf)
	if err != nil {
		return nil, fmt.Errorf("can't process the config: %w", err)
	}
	isValid := conf.Validate()
	if isValid {
		return conf, nil
	}
	return conf, errors.New("Malformed Config")
}

func (c *Config) Validate() bool {
	if c.Port == "" {
		log.Fatal("Acceptable ports are 80 and 8080. Provided:", c.Port)
		return false
	}
	_, err := url.ParseRequestURI(c.JaegerUrl)
	if err != nil {
		log.Fatal("Invalid host:", c.JaegerUrl)
		return false
	}

	if c.KafkaBroker == "" {
		log.Fatal("Expected kafka broker host. Provided:", c.KafkaBroker)
		return false
	}
	return true
}
