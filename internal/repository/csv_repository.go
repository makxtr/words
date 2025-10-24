package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"words/internal/domain"
)

type CSVRepository struct {
	filePath string
}

func NewCSVRepository(filePath string) *CSVRepository {
	return &CSVRepository{
		filePath: filePath,
	}
}

func (r *CSVRepository) GetAllWords() ([]domain.Word, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("invalid CSV format: %w", err)
	}

	words := make([]domain.Word, 0, len(records))
	for _, record := range records {
		words = append(words, domain.Word{
			Original:    record[0],
			Translation: record[1],
		})
	}

	return words, nil
}
