package services

import (
	"aktai/domain"
	"aktai/repository"

	"github.com/google/uuid"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices() *Services {
	return &Services{
		Repository: repository.NewRepository(),
	}
}

func (s *Services) GetAllColleges() ([]domain.College, error) {
	return s.Repository.GetAllColleges()
}

func (s *Services) GetCollege(id string) (domain.College, error) {
	return s.Repository.GetCollege(id)
}

func (s *Services) CreateCollege(college domain.College) (domain.College, int, error) {
	college.ID = uuid.New().String()

	college, err := s.Repository.CreateCollege(college)
	if err != nil {
		return domain.College{}, 400, err
	}

	return college, 201, nil
}

func (s *Services) UpdateCollege(college domain.College) (domain.College, error) {
	if err := s.Repository.UpdateCollege(college); err != nil {
		return domain.College{}, err
	}

	return college, nil
}

func (s *Services) DeleteCollege(id string) error {
	if err := s.Repository.DeleteCollege(id); err != nil {
		return err
	}

	return nil
}
