package repository

import (
	"encoding/csv"
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
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var words []domain.Word
	for _, record := range records {
		words = append(words, domain.Word{
			Original:    record[0],
			Translation: record[1],
		})
	}

	return words, nil
}
