package cfg

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Type string
	Host string
	Port string
}

type ApiConfig struct {
	Port string
}

type ServerConfig struct {
	Instance ApiConfig
	Database DatabaseConfig
}

func SetServerConfig() ServerConfig {
	viper.SetConfigName("server")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./cfg")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error : unable to read config file: %w", err))
	}

	return ServerConfig{
		Instance: ApiConfig{
			Port: viper.GetString("config.listening_port"),
		},
		Database: DatabaseConfig{
			Type: viper.GetString("config.database_type"),
			Host: viper.GetString("config.database_host"),
			Port: viper.GetString("config.database_port"),
		},
	}
}
