package config

import (
	"github.com/BurntSushi/toml"
)

// NetworkConfig represents the configulation of the network
type NetworkConfig struct {
	Address string `toml:"address"`
}

// EndpointConfig represents the configulation of the endpoint
type EndpointConfig struct {
	Path        string   `toml:"path"`
	AllowMethod []string `toml:"allow_method"`
	Plugin      string   `toml:"plugin"`
}

// AppConfig represents the configulation of the application
type AppConfig struct {
	Network  NetworkConfig    `toml:"network"`
	Endpoint []EndpointConfig `toml:"endpoint"`
}

var appConfig AppConfig

func validate() error {
	return nil
}

// Config returns Application config
func Config() *AppConfig {
	return &appConfig
}

// Load config file
func Load(filename string) error {
	if _, err := toml.DecodeFile(filename, &appConfig); err != nil {
		return err
	}

	if err := validate(); err != nil {
		return err
	}

	return nil
}
