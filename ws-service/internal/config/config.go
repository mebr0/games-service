package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

var (
	c *Config
)

type Config struct {
	HTTP struct {
		Port               string        `yaml:"port" envconfig:"HTTP_PORT"`
		ReadTimeout        time.Duration `yaml:"read-timeout" envconfig:"HTTP_READ_TIMEOUT"`
		WriteTimeout       time.Duration `yaml:"write-timeout" envconfig:"HTTP_WRITE_TIMEOUT"`
		MaxHeaderMegabytes int           `yaml:"max-header-megabytes" envconfig:"HTTP_MAX_HEADER_MEGABYTES"`
	} `yaml:"http"`
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
