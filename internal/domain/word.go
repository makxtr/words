package domain

type Word struct {
	Original    string
	Translation string
}

type WordRepository interface {
	GetAllWords() ([]Word, error)
}
