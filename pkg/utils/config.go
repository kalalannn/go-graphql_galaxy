package utils

import (
	"fmt"
	"os"

	"go-graphql_galaxy/pkg/log"

	"gopkg.in/yaml.v3"
)

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"sslmode"`
	Timezone string `yaml:"timezone"`
}

type Config struct {
	Env      string `yaml:"env"`
	Database `yaml:"database"`
}

func DSN(dbConfig *Database) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.SSLMode,
		dbConfig.Timezone)
}

// ENV -> default
func MustLoadConfig() *Config {
	configPath := "config/local.yaml"

	if fromEnv := os.Getenv("APP_CONFIG_PATH"); fromEnv != "" {
		configPath = fromEnv
	}

	config, err := loadConfig(configPath)
	if err != nil {
		log.Fatal("wrong config (%s): %v", configPath, err)
	}
	return config
}

func loadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
