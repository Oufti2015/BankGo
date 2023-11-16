package repo

import "github.com/oufti2015/html/model"

type Repository struct {
	Categories []model.Category
	Operations []model.Operation
}

// NewRepository - our constructor function
func NewRepository(categories []model.Category, operations []model.Operation) (*Repository, error) {
	svc := &Repository{
		Categories: categories,
		Operations: operations,
	}
	return svc, nil
}
