package util

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Hosts struct {
	GitHubUser        string `mapstructure:"GITHUB_USER"`
	GitHubAccessToken string `mapstructure:"GITHUB_ACCESS_TOKEN"`
}

func IsInitialized() bool {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return false
		}

	}
	return true
}

func InitializeConfig() error {
	fmt.Println("INITIALAZING...")
	_, err := os.OpenFile("app.env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	Reset(".")
	return nil
}

func LoadHosts() (*Hosts, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var hosts Hosts
	if err := viper.Unmarshal(&hosts); err != nil {
		return nil, err
	}

	// Returns nil if GitHub token is empty.
	if hosts.GitHubAccessToken != "" {
		return &hosts, nil
	}
	return nil, nil
}

func Write(key, value string) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.Set(key, value)

	viper.WriteConfig()

}

func Reset(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.Set("GITHUB_USER", "")
	viper.Set("GITHUB_ACCESS_TOKEN", "")
	viper.WriteConfig()
}
