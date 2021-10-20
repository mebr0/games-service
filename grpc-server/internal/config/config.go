package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	c *Config
)

type Config struct {
	GRPC struct {
		Port string `yaml:"port" envconfig:"GRPC_PORT"`
	} `yaml:"grpc"`

	WS struct {
		Host string `yaml:"host" envconfig:"WS_HOST"`
	} `yaml:"ws"`
}

func LoadConfig(configPath string) *Config {
	if c == nil {
		c = &Config{}

		c.readFile(configPath)
		c.readEnv()
	}

	return c
}

// File configs with values from configs file
func (c *Config) readFile(path string) {
	f, err := os.Open(path)

	if err != nil {
		processError(err)
	}

	defer f.Close()

	err = yaml.NewDecoder(f).Decode(c)

	if err != nil {
		log.Println(c)
		processError(err)
	}
}

// Read configs with values from env variables
func (c *Config) readEnv() {
	loadFromEnvFile()

	err := envconfig.Process("", c)

	if err != nil {
		processError(err)
	}
}

// Load values from .env file to system
func loadFromEnvFile() {
	if err := godotenv.Load(); err != nil {
		log.Warning("Error loading .env file")
	}
}

func processError(err error) {
	log.Error(err)
	os.Exit(2)
}
