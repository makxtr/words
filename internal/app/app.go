package app

import (
	"words/internal/repository"
	"words/internal/service"
)

type App struct {
	trainerService *service.TrainerService
}

func NewApp() *App {
	repo := repository.NewCSVRepository("data/words.csv")
	trainerService := service.NewTrainerService(repo)
	return &App{
		trainerService: trainerService,
	}
}

func (a *App) Run() error {
	return a.trainerService.StartTraining()
}
