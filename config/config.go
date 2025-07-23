package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

type Config struct {
	LighthouseChainUrl        *string
	LighthouseContractAddress *string
	PrivateKey                *string
	LighthouseChainId         *uint64
	GasLimit                  *uint64
	LighthouseUrl             *string
	RpcNodeWsUrl              *string
	RpcNodeHttpUrl            *string
	AnvilUrl                  *string
	RollupId                  *string
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
