package config

import (
	"fmt"

	"github.com/PlopyBlopy/notebot/pkg/tgbot"
	"github.com/joho/godotenv"
)

type App struct {
	Name    string `env:"APP_NAME" env-default:"AppName"`
	Version string `env:"APP_VERSION" env-default:"AppVersion"`
}

type Config struct {
	App   App          `env-prefix:"APP_"`
	TgBot tgbot.Config `env-prefix:"BOT_"`
}

func InitConfig() (Config, error) {

	var cfg Config

	err := godotenv.Load("./.env.development")
	if err != nil {
		return cfg, fmt.Errorf("incorrect format environment. %w", err)
	}

	return cfg, nil
}
