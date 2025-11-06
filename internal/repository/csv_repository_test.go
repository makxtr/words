package repository

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCSVRepository_GetAllWords(t *testing.T) {
	tests := []struct {
		name       string
		csvContent string
		wantCount  int
		wantErr    bool
		checkFirst bool
		firstWord  string
		firstTrans string
	}{
		{
			name: "valid csv with multiple words",
			csvContent: `hello,привет
world,мир
test,тест`,
			wantCount:  3,
			wantErr:    false,
			checkFirst: true,
			firstWord:  "hello",
			firstTrans: "привет",
		},
		{
			name:       "valid csv with single word",
			csvContent: `apple,яблоко`,
			wantCount:  1,
			wantErr:    false,
			checkFirst: true,
			firstWord:  "apple",
			firstTrans: "яблоко",
		},
		{
			name:       "empty csv",
			csvContent: "",
			wantCount:  0,
			wantErr:    false,
		},
		{
			name: "invalid csv - wrong number of fields",
			csvContent: `hello,привет,extra
world,мир`,
			wantCount: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			csvPath := filepath.Join(tmpDir, "test.csv")

			if err := os.WriteFile(csvPath, []byte(tt.csvContent), 0644); err != nil {
				t.Fatalf("failed to create test file: %v", err)
			}

			repo := NewCSVRepository(csvPath)

			words, err := repo.GetAllWords()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(words) != tt.wantCount {
				t.Errorf("GetAllWords() got %d words, want %d", len(words), tt.wantCount)
			}

			if tt.checkFirst && len(words) > 0 {
				if words[0].Original != tt.firstWord {
					t.Errorf("First word Original = %v, want %v", words[0].Original, tt.firstWord)
				}
				if words[0].Translation != tt.firstTrans {
					t.Errorf("First word Translation = %v, want %v", words[0].Translation, tt.firstTrans)
				}
			}
		})
	}
}

func TestCSVRepository_GetAllWords_FileNotFound(t *testing.T) {
	repo := NewCSVRepository("/nonexistent/path/file.csv")

	words, err := repo.GetAllWords()

	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}

	if words != nil {
		t.Errorf("Expected nil words, got %v", words)
	}
}
