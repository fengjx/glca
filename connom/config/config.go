package config

import (
	"os"

	"github.com/fengjx/go-halo/fs"
	"github.com/fengjx/luchen"
)

var appConfig AppConfig

func init() {
	var configFile string
	envConfigPath := os.Getenv("APP_CONFIG")
	if envConfigPath != "" {
		configFile = envConfigPath
	}
	if configFile == "" && len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	if configFile == "" {
		confFile, err := fs.Lookup("conf/app.yaml", 3)
		if err != nil {
			luchen.RootLogger().Panic("config file not found")
		}
		configFile = confFile
	}
	configFile, err := fs.Lookup(configFile, 3)
	if err != nil {
		panic(err)
	}
	appConfig = luchen.MustLoadConfig[AppConfig](configFile)
}

type AppConfig struct {
	Server Server                  `json:"server"`
	DB     map[string]*DbConfig    `json:"db"`
	Redis  map[string]*RedisConfig `json:"redis"`
}

type Server struct {
	HTTP HTTPServerConfig
}

type HTTPServerConfig struct {
	ServerName string `json:"server-name"`
	Listen     string `json:"listen"`
}

type DbConfig struct {
	Type    string `json:"type"`
	Dsn     string `json:"dsn"`
	MaxIdle int    `json:"max-idle"`
	MaxConn int    `json:"max-conn"`
	ShowSQL bool   `json:"show-sql"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func GetConfig() AppConfig {
	return appConfig
}
