package config

import (
	"flag"
	"os"

	"github.com/fengjx/go-halo/fs"
	"github.com/fengjx/luchen"
)

var appConfig AppConfig

func init() {
	configPath := flag.String("c", "", "custom config file path")
	flag.Parse()

	appConfigFile, err := fs.Lookup("conf/app.yml", 5)
	if err != nil {
		luchen.RootLogger().Panic("config file not found")
	}
	configs := []string{appConfigFile}
	var configFile string
	if configPath != nil {
		configFile = *configPath
	}
	if configFile == "" {
		envConfig := os.Getenv("APP_CONFIG")
		if envConfig != "" {
			configFile = envConfig
		}
	}
	if configFile != "" {
		configFile, err = fs.Lookup(configFile, 5)
		if err != nil {
			panic(err)
		}
		luchen.RootLogger().Infof("load custom config file: %s", configFile)
		configs = append(configs, configFile)
	}
	appConfig = luchen.MustLoadConfig[AppConfig](configs...)
}

type AppConfig struct {
	Server Server                  `json:"server"`
	Auth   AuthConfig              `json:"auth"`
	DB     map[string]*DbConfig    `json:"db"`
	Redis  map[string]*RedisConfig `json:"redis"`
}

type Server struct {
	HTTP HTTPServerConfig
}

type AuthConfig struct {
	Version string `json:"version"`
	Secret  string `json:"secret"`
}

// CorsConfig 跨域配置
type CorsConfig struct {
	AllowOrigins []string `json:"allow-origins"`
}

type HTTPServerConfig struct {
	ServerName string     `json:"server-name"`
	Listen     string     `json:"listen"`
	Cors       CorsConfig `json:"cors"`
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
