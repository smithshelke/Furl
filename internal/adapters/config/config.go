package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Neo4jConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Addr     string `yaml:"addr"`
}

type ServerConfig struct {
	Addr         string `yaml:"addr"`
	WriteTimeout int    `yaml:"write_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	Neo4j  Neo4jConfig  `yaml:"neo4j"`
}

func Initialize() *Config {
	config := &Config{}
	configBytes := readConfigFile()
	err := yaml.Unmarshal(configBytes, config)
	if err != nil {
		fmt.Printf("Err: %+v\n", err)
	}
	return config
}

func readConfigFile() []byte {
	config, err := os.ReadFile("config.yml")
	if err != nil {
		fmt.Printf("Err: %+v\n", err)
	}
	return config
}
