package service

import (
	"testing"
	"words/internal/domain"
)

type mockRepository struct {
	words []domain.Word
	err   error
}

func (m *mockRepository) GetAllWords() ([]domain.Word, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.words, nil
}

func TestNewTrainerService(t *testing.T) {
	repo := &mockRepository{
		words: []domain.Word{
			{Original: "test", Translation: "тест"},
		},
	}

	service := NewTrainerService(repo)

	if service == nil {
		t.Fatal("NewTrainerService returned nil")
	}

	if service.repo == nil {
		t.Error("TrainerService repo is nil")
	}
}

func TestTrainerService_generateOptions(t *testing.T) {
	tests := []struct {
		name         string
		words        []domain.Word
		correctIndex int
		wantOptions  int
	}{
		{
			name: "generates 4 options",
			words: []domain.Word{
				{Original: "hello", Translation: "привет"},
				{Original: "world", Translation: "мир"},
				{Original: "test", Translation: "тест"},
				{Original: "apple", Translation: "яблоко"},
				{Original: "dog", Translation: "собака"},
			},
			correctIndex: 0,
			wantOptions:  4,
		},
		{
			name: "exactly 4 words",
			words: []domain.Word{
				{Original: "one", Translation: "один"},
				{Original: "two", Translation: "два"},
				{Original: "three", Translation: "три"},
				{Original: "four", Translation: "четыре"},
			},
			correctIndex: 2,
			wantOptions:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockRepository{words: tt.words}
			service := NewTrainerService(repo)

			options, err := service.generateOptions(tt.words, tt.correctIndex)
			if err != nil {
				t.Fatalf("generateOptions() returned error: %v", err)
			}

			if len(options) != tt.wantOptions {
				t.Errorf("generateOptions() returned %d options, want %d", len(options), tt.wantOptions)
			}

			// Check that correct answer is present
			correctWord := tt.words[tt.correctIndex]
			found := false
			for _, opt := range options {
				if opt.Translation == correctWord.Translation {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("generateOptions() doesn't contain correct answer %v", correctWord.Translation)
			}

			seen := make(map[string]bool)
			for _, opt := range options {
				if seen[opt.Translation] {
					t.Errorf("generateOptions() has duplicate translation: %v", opt.Translation)
				}
				seen[opt.Translation] = true
			}
		})
	}
}

func TestTrainerService_generateOptions_NotEnoughWords(t *testing.T) {
	words := []domain.Word{
		{Original: "one", Translation: "один"},
		{Original: "two", Translation: "два"},
	}

	repo := &mockRepository{words: words}
	service := NewTrainerService(repo)

	_, err := service.generateOptions(words, 0)
	if err == nil {
		t.Fatal("generateOptions() should return error when not enough words")
	}

	expectedError := "not enough words"
	if err.Error() != expectedError {
		t.Errorf("generateOptions() error = %v, want %v", err.Error(), expectedError)
	}
}
