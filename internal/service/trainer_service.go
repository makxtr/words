package service

import (
	"fmt"
	"math/rand"
	"words/internal/domain"
	"words/pkg/utils"
)

const (
	OptionsCount = 4
	MinOption    = 1
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
		utils.ClearScreen()
		fmt.Printf("Word %d of %d\n\n", i+1, totalWords)
		fmt.Printf("%s\n\n", word.Original)
		options := s.generateOptions(words, i)
		for j, option := range options {
			fmt.Printf("%d. %s\n", j+1, option.Translation)
		}

		fmt.Printf("Choose the correct answer (%d-%d): ", MinOption, OptionsCount)
		answer := utils.ReadIntAnswer(MinOption, OptionsCount)

		if options[answer-1].Translation == word.Translation {
			correctAnswers++
			fmt.Println("\n✓ Right!")
		} else {
			fmt.Printf("\n✗ Wrong! Correct answer: %s\n", word.Translation)
		}

		progress := float64(correctAnswers) / float64(totalWords) * 100
		fmt.Printf("Progress: %d/%d (%.2f%%)\n", correctAnswers, totalWords, progress)
		fmt.Println("\nPress Enter to continue...")
		_, _ = fmt.Scanln()
	}

	return nil
}

func (s *TrainerService) generateOptions(words []domain.Word, correctIndex int) []domain.Word {
	options := make([]domain.Word, OptionsCount)
	options[0] = words[correctIndex]

	usedIndices := make(map[int]bool)
	usedIndices[correctIndex] = true

	for i := 1; i < OptionsCount; i++ {
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
