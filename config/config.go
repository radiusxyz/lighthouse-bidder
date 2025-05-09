package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

type Config struct {
	LighthouseUrl string
	RpcNodeUrl    string
}

func New() *Config {
	configPath := "config/config.toml"
	configFile, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := configFile.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()
	config := &Config{}
	if err = toml.NewDecoder(configFile).Decode(config); err != nil {
		panic(err)
	}
	return config
}
