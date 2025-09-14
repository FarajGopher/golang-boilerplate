package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

func InitEnvConfig(){
	EnvConfigs = loadEnvVariables()
	fmt.Println("Loaded env variables successfully")
}

type envConfigs struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	SecretKey       string `mapstructure:"SECRET_KEY"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          string `mapstructure:"DB_PORT"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPassword      string `mapstructure:"DB_PASSWORD"`
	DBName          string `mapstructure:"DB_NAME"`
}


func loadEnvVariables() (config *envConfigs) {
	// Tell viper the path/location of your env file, if its root just add "."
	viper.AddConfigPath("./config")

	// Tell viper the name of your file
	viper.SetConfigName("sample-app")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file: ", err)
	}

	// viper unmarshals the loaded env variables into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error while unmarshalling the config data: ",err)
	}

	return
}
