package cfg

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseAuth struct {
	Username string
	Password string
}
type DatabaseConfig struct {
	Type       string
	Host       string
	Port       string
	Name       string
	TlsSupport bool
	Auth       DatabaseAuth
}

type ApiConfig struct {
	Port string
}

type FrontConfig struct {
	Url string
}

type FileManagerConfig struct {
	UploadPath string
}

type ServerConfig struct {
	Instance    ApiConfig
	Front       FrontConfig
	Database    DatabaseConfig
	FileManager FileManagerConfig
}

func SetServerConfig() ServerConfig {
	viper.SetConfigName("server")
	viper.SetConfigType("toml")
	/* Normal runtime */
	viper.AddConfigPath("./configs")
	/* Tests runtime */
	viper.AddConfigPath("../configs")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error : unable to read config file: %w", err))
	}

	return ServerConfig{
		Instance: ApiConfig{
			Port: viper.GetString("api.listening_port"),
		},
		Front: FrontConfig{
			Url: viper.GetString("front.url"),
		},
		Database: DatabaseConfig{
			Type:       viper.GetString("database.type"),
			Host:       viper.GetString("database.host"),
			Port:       viper.GetString("database.port"),
			Name:       viper.GetString("database.name"),
			TlsSupport: false,
			Auth: DatabaseAuth{
				Username: viper.GetString("database.username"),
				Password: viper.GetString("database.password"),
			},
		},
		FileManager: FileManagerConfig{
			UploadPath: viper.GetString("fileManager.upload_path"),
		},
	}
}
