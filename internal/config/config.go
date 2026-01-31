package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	Env        string `mapstructure:"ENV"`
	DB         DBConfig
}

type DBConfig struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

func LoadConfig() *Config {
	viper.AddConfigPath(".")    // look in project root
	viper.SetConfigName(".env") // name of file
	viper.SetConfigType("env")  // format
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables only")
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatal("cannot load config:", err)
	}

	if cfg.ServerPort == "" {
		cfg.ServerPort = "8080"
	}

	cfg.DB = DBConfig{
		Host: viper.GetString("DB_HOST"),
		Port: viper.GetString("DB_PORT"),
		Name: viper.GetString("DB_NAME"),
		User: viper.GetString("DB_USER"),
		Pass: viper.GetString("DB_PASS"),
	}

	return cfg
}
