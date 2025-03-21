package service

import (
	"fmt"
	"math/rand"
	"words/internal/domain"
)

type TrainerService struct {
	repo domain.WordRepository
}

func NewTrainerService(repo domain.WordRepository) *TrainerService {
	return &TrainerService{
		repo: repo,
	}
}

func (s *TrainerService) StartTraining() error {
	words, err := s.repo.GetAllWords()
	if err != nil {
		return err
	}

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	correctAnswers := 0
	totalWords := len(words)

	for i, word := range words {
		fmt.Printf("%s\n", word.Original)
		options := s.generateOptions(words, i)
		for j, option := range options {
			fmt.Printf("%d. %s\n", j+1, option.Translation)
		}

		var answer int
		fmt.Print("Выберите правильный вариант: ")
		fmt.Scan(&answer)

		if options[answer-1].Translation == word.Translation {
			correctAnswers++
			fmt.Println("Правильно!")
		} else {
			fmt.Println("Неправильно!")
		}

		progress := float64(correctAnswers) / float64(totalWords) * 100
		fmt.Printf("Прогресс: %.2f%%\n", progress)
	}

	return nil
}

func (s *TrainerService) generateOptions(words []domain.Word, correctIndex int) []domain.Word {
	options := make([]domain.Word, 4)
	options[0] = words[correctIndex]

	usedIndices := make(map[int]bool)
	usedIndices[correctIndex] = true

	for i := 1; i < 4; i++ {
		var randomIndex int
		for {
			randomIndex = rand.Intn(len(words))
			if !usedIndices[randomIndex] {
				break
			}
		}
		options[i] = words[randomIndex]
		usedIndices[randomIndex] = true
	}

	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	return options
}
