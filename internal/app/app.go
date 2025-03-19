package app

import (
	"github.com/spf13/viper"
	"words/config"
	"words/internal/repository"
	"words/internal/service"
)

type App struct {
	trainerService *service.TrainerService
}

func NewApp() (*App, error) {
	if err := config.InitConfig(); err != nil {
		return nil, err
	}

	repo := repository.NewCSVRepository(viper.GetString("app.csv_path"))
	trainerService := service.NewTrainerService(repo)
	return &App{
		trainerService: trainerService,
	}, nil
}

func (a *App) Run() error {
	return a.trainerService.StartTraining()
}
