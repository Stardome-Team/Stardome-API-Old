package config

import (
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"

	"gopkg.in/yaml.v2"
)

// Config represents an application configuration.
type Config struct {
	Database struct {
		Engine   string `yaml:"engine" envconfig:"DB_ENGINE"`
		Host     string `yaml:"host" envconfig:"DB_HOST"`
		Port     string `yaml:"port" envconfig:"DB_PORT"`
		Name     string `yaml:"name" envconfig:"DB_DATABASE"`
		Username string `yaml:"username" envconfig:"DB_USERNAME"`
		Password string `yaml:"password" envconfig:"DB_PASSWORD"`
	} `yaml:"Database"`
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string) (*Config, error) {

	var cfg Config

	err := readFromFile(file, &cfg)

	if err != nil {
		return nil, err
	}

	err = readFromEnvironment(file, &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func readFromFile(file string, cfg *Config) error {

	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bytes, cfg)

	return err
}

func readFromEnvironment(file string, cfg *Config) error {

	err := envconfig.Process("", cfg)

	return err
}
