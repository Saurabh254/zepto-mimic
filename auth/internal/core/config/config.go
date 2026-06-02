package config

import (
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Host string `mapstructure:"host" json:"host" yaml:"host"`
}

func LoadConfig() *Config {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("[config] Error loading .env file: " + err.Error())
	}

	// Get PORT
	port, err := strconv.Atoi(GetEnvOrDefault("PORT", "8000"))
	if err != nil {
		panic("[config] Invalid PORT value: " + err.Error())
	}

	return &Config{
		Port: port,
		Host: GetEnvOrDefault("HOST", "localhost"),
	}
}

func (c *Config) Address() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}
