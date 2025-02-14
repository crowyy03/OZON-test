package config
import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Storage struct {
		Type string
	}
	Postgres struct {
		DSN string
	}
	Server struct {
		Port string
	}
}

func LoadConfig(path string) (Config, error) {
	var cfg Config

	viper.SetConfigFile(path)
	viper.AutomaticEnv() // Читает переменные окружения
	viper.SetEnvPrefix("")

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	log.Println("Конфигурация загружена:", cfg)
	return cfg, nil
}
