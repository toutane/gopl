package util

import "github.com/spf13/viper"

type Config struct {
	GitHubUser  string `mapstructure:"GITHUB_USER"`
	GitHubToken string `mapstructure:"GITHUB_TOKEN"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
