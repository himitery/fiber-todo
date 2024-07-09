package config

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	confOnce  sync.Once
	appConfig *Config
)

type Config struct {
	Host     string         `json:"host"`
	Port     string         `json:"port"`
	Develop  bool           `json:"develop"`
	Cors     CorsConfig     `json:"cors"`
	Log      LoggerConfig   `json:"log"`
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Database string `json:"database"`
	Password string `json:"password"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Db   int    `json:"db"`
}

type CorsConfig struct {
	Origins     []string `json:"origins"`
	Methods     []string `json:"methods"`
	Headers     []string `json:"headers"`
	Credentials bool     `json:"credentials"`
}

type LoggerConfig struct {
	Level    string `json:"level"` // debug, info, error...
	Type     string `json:"type"`  // options: file, stdout
	Filename string `json:"filename"`
}

func LoadConfigFile(filename string) *Config {
	confOnce.Do(func() {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(filename)
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			log.Fatal(err)
		}
		err = viper.Unmarshal(&appConfig)
		if err != nil {
			panic(err)
		}
		data, _ := json.MarshalIndent(appConfig, "", "  ")
		log.Printf("Config Load %s", data)
	})

	return appConfig
}
