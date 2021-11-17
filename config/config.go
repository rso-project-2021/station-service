package config

import "github.com/spf13/viper"

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	GinMode       string `mapstructure:"GIN_MODE"`
}

// Reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	// Some small change
	err = viper.Unmarshal(&config)
	return
}
