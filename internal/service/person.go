package service

import (
	log "github.com/sirupsen/logrus"
	"task_peopler/internal/models"
	"task_peopler/pkg"
)

type PersonRepo interface {
	Create(p *models.Person) (*models.Person, error)
}

type PersonService struct {
	repo      PersonRepo
	APIClient *pkg.APIClient
}

func NewPersonService(repo PersonRepo, APIClient *pkg.APIClient) *PersonService {
	return &PersonService{repo: repo, APIClient: APIClient}
}

func (s *PersonService) Create(p *models.Person) (*models.Person, error) {
	age, err := s.APIClient.FetchAge(p.FirstName)
	if err != nil {
		log.Errorf("failed to fetch age: %v", err)
		return nil, err
	}
	p.Age = age

	gender, err := s.APIClient.FetchGender(p.FirstName)
	if err != nil {
		log.Errorf("failed to fetch gender %v", err)
		return nil, err
	}
	p.Gender = gender

	nationality, err := s.APIClient.FetchNationality(p.FirstName)
	if err != nil {
		log.Errorf("failed to fetch nationality %v", err)
		return nil, err
	}
	p.Nationality = nationality

	return s.repo.Create(p)
}
