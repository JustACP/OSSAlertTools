package model

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type APP struct {
	AccessKey      string `yaml:"AccessKey"`
	AccessSecret   string `yaml:"AccessSecret"`
	RequestKey     string `yaml:"RequestKey"`
	BucketEndPoint string `yaml:"BucketEndPoint"`
	BucketName     string `yaml:"BucketName"`
}

type Config struct {
	App APP `yaml:"app"`
}

var ApplicationConfig Config

func GetConfig() {

	viperConfig := viper.New()
	viperConfig.AddConfigPath(".")
	viperConfig.SetConfigName("config")
	viperConfig.SetConfigType("yaml")

	if err := viperConfig.ReadInConfig(); err != nil {
		log.Fatal("Read Application Config File Error!")
		os.Exit(0)
	}

	if err := viperConfig.Unmarshal(&ApplicationConfig); err != nil {
		log.Fatal("Unmarshal Application Config Error!")
		os.Exit(0)
	}

}

func SetEnv() {
	os.Setenv("OSS_ACCESS_KEY_ID", ApplicationConfig.App.AccessKey)
	os.Setenv("OSS_ACCESS_KEY_SECRET", ApplicationConfig.App.AccessSecret)
}
