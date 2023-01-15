package configs

import "github.com/spf13/viper"

type DBConfig struct {
	Port     string `mapstructure:"API_PORT"`
	Type     string `mapstructure:"DB_TYPE"`
	Host     string `mapstructure:"DB_HOST"`
	DBPort   string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Database string `mapstructure:"DB_NAME"`
}

func init() {
	viper.SetDefault("port", "8000")
	viper.SetDefault("type", "postgres")
	viper.SetDefault("host", "localhost")
	viper.SetDefault("dbPort", "5432")
	viper.SetDefault("user", "postgres")
	viper.SetDefault("password", "password")
	viper.SetDefault("database", "postgres")
}
