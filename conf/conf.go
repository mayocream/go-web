package conf

import (
	"github.com/spf13/viper"
	"log"
)

type Configs struct {
	Version 	int
	Timezone 	string
	Environment EnvironmentConfig
	Database 	DatabaseConfig
	Redis 		RedisConfig
}

type EnvironmentConfig struct {
	Env 		string
	Debug 		bool
}

type RedisConfig struct {
	Host 		string
	Port		int
	Base 		int
}

type DatabaseConfig struct {
	Host 		string
	Port 		int
	Database 	string
	User 		string
	Password 	string
}

var Config *Configs

func init() {
	var err error
	// Read .env
	viper.SetConfigFile(".env")
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading .env file: %s \n", err)
	}

	// Read config.yaml
	viper.SetConfigFile("./config/config.yaml")
	if err = viper.MergeInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	viper.AutomaticEnv()

	Config = &Configs{
		Version: viper.GetInt("version"),
		Timezone: viper.GetString("timezone"),
		Environment: EnvironmentConfig{
			Env:   viper.GetString("ENV"),
			Debug: viper.GetBool("DEBUG"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			Database: viper.GetString("DB_DATABASE"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
		},
		Redis: RedisConfig{
			Host: viper.GetString("REDIS_HOST"),
			Port: viper.GetInt("REDIS_PORT"),
			Base: viper.GetInt("REDIS_BASE"),
		},
	}
}