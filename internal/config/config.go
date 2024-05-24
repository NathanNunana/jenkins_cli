package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	JenkinsURL string `json:"url"`
	Username   string `json:"username"`
	ApiToken   string `json:"token"`
}

func GetConfig(path string) (*Config, error) {
	viper.SetConfigName("credentials")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	if _, err := os.Stat(path); err == nil {
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
		return nil, err
	}

	config := Config{
		JenkinsURL: viper.GetString("url"),
		Username:   viper.GetString("username"),
		ApiToken:   viper.GetString("apiToken"),
	}

	return &config, nil
}
