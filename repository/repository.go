package repository

import (
	"errors"

	"aktai/domain"
)

type Repository struct {
	DB map[string]domain.College
}

func NewRepository() *Repository {
	return &Repository{
		DB: map[string]domain.College{},
	}
}

func (repo *Repository) GetAllColleges() ([]domain.College, error) {
	var colleges []domain.College

	for _, college := range repo.DB {
		colleges = append(colleges, college)
	}

	return colleges, nil
}

func (repo *Repository) GetCollege(id string) (domain.College, error) {
	if college, exist := repo.DB[id]; exist {
		return college, nil
	}

	return domain.College{}, errors.New("not found")
}

func (repo *Repository) CreateCollege(college domain.College) (domain.College, error) {
	if _, exist := repo.DB[college.ID]; exist {
		return domain.College{}, errors.New("already exist")
	}

	repo.DB[college.ID] = college

	return college, nil
}

func (repo *Repository) UpdateCollege(college domain.College) error {
	if _, exist := repo.DB[college.ID]; !exist {
		return errors.New("not found")
	}

	repo.DB[college.ID] = college

	return nil
}

func (repo *Repository) DeleteCollege(id string) error {
	if _, exist := repo.DB[id]; !exist {
		return errors.New("not found")
	}

	delete(repo.DB, id)

	return nil
}
