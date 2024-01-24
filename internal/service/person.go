package service

import (
	log "github.com/sirupsen/logrus"
	"task_peopler/internal/models"
	"task_peopler/pkg"
)

type PersonRepo interface {
	Create(p *models.Person) (*models.Person, error)
	Update(p *models.Person) (*models.Person, error)
	Delete(id int64) error
	GetByID(id int64) (*models.Person, error)
	GetPeople(filters map[string]interface{}, page, pageSize int) ([]*models.Person, error)
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

	log.Infof("Create person: %v", p)
	return s.repo.Create(p)
}

func (s *PersonService) Update(p *models.Person) (*models.Person, error) {
	log.Infof("Update person: %v", p)
	return s.repo.Update(p)
}

func (s *PersonService) Delete(id int64) error {
	log.Infof("Delete person with ID: %d", id)
	return s.repo.Delete(id)
}

func (s *PersonService) GetByID(id int64) (*models.Person, error) {
	log.Infof("Get person with ID: %d", id)
	return s.repo.GetByID(id)
}

func (s *PersonService) GetPeopleByName(name string, page, pageSize int) ([]*models.Person, error) {
	filters := map[string]interface{}{"first_name": name}
	return s.repo.GetPeople(filters, page, pageSize)
}

func (s *PersonService) GetPeopleByAge(age int, page, pageSize int) ([]*models.Person, error) {
	filters := map[string]interface{}{"age": age}
	return s.repo.GetPeople(filters, page, pageSize)
}

func (s *PersonService) GetPeopleByGender(gender string, page, pageSize int) ([]*models.Person, error) {
	filters := map[string]interface{}{"gender": gender}
	return s.repo.GetPeople(filters, page, pageSize)
}

func (s *PersonService) GetPeopleByNationality(nationality string, page, pageSize int) ([]*models.Person, error) {
	filters := map[string]interface{}{"nationality": nationality}
	return s.repo.GetPeople(filters, page, pageSize)
}
