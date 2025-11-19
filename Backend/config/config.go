package config

import (
	"flag"
	"fmt"

	"github.com/PlopyBlopy/notebot/internal/adapters/note"
	"github.com/PlopyBlopy/notebot/pkg/httpserver"
	"github.com/PlopyBlopy/notebot/pkg/tgbot"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type App struct {
	Name    string `env:"APP_NAME" env-default:"AppName"`
	Version string `env:"APP_VERSION" env-default:"AppVersion"`
}

type Config struct {
	Environment string                      `env:"ENVIRONMENT" envDefault:"development"`
	App         App                         `env-prefix:"APP_"`
	TgBot       tgbot.Config                `env-prefix:"BOT_"`
	Metadata    note.MetadataConfig         `env-prefix:"MD_"`
	HttpServer  httpserver.HttpServerConfig `env-prefix:"HTTP_"`
}

func InitConfig() (Config, error) {
	envFlag := flag.String("env", "development", "Environment: development|production")

	godotenv.Load(".env." + *envFlag)

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("incorrect format environment. %w", err)
	}

	return cfg, nil
}
