package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net/url"
	"os"
	"path"
	"strings"
)

type Config struct {
	Host    string `yaml:"host" json:"host"`
	Port    int64  `yaml:"port" json:"port"`
	Timeout int64  `yaml:"timeout" json:"timeout"`
}

func Load(configFile *string) (*Config, error) {
	var err error

	if _, statErr := os.Stat(*configFile); statErr != nil {
		return nil, fmt.Errorf("%w", statErr)
	}

	file, fileErr := os.Open(*configFile)
	if fileErr != nil {
		return nil, errors.New("cannot open file")
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("close file error: %v", err)
		}
	}()

	var conf *Config

	switch strings.ToLower(path.Ext(*configFile))[1:] {
	case "yaml", "yml":
		err = yaml.NewDecoder(file).Decode(&conf)
		if err != nil {
			log.Printf("Error: %v", err)
			return nil, errors.New("cannot decode yaml file")
		}
	case "json":
		err = json.NewDecoder(file).Decode(&conf)
		if err != nil {
			log.Printf("Error: %v", err)
			return nil, errors.New("cannot decode json file")
		}
	default:
		return nil, errors.New("unknown config file")
	}

	isOk, err := conf.Validate()

	if isOk {
		return conf, nil
	}
	return conf, err
}

func (c *Config) Validate() (bool, error) {
	_, err := url.ParseRequestURI(c.Host)
	if err != nil {
		return false, err
	}

	return true, nil
}
