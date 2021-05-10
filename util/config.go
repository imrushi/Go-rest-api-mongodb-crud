package util

import "github.com/spf13/viper"

type Config struct {
	MONGO_URI        string `mapstructure:"MONGO_URI"`
	MONGO_DATABASE   string `mapstructure:"MONGO_DATABASE"`
	MONGO_COLLECTION string `mapstructure:"MONGO_COLLECTION"`
	API_PORT         string `mapstructure:"API_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
