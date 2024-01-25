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
	ageCh := make(chan int, 1)
	genderCh := make(chan string, 1)
	nationalityCh := make(chan string, 1)
	errCh := make(chan error, 3)

	// Fetch age concurrently
	go func() {
		age, err := s.APIClient.FetchAge(p.FirstName)
		if err != nil {
			errCh <- err
			return
		}
		ageCh <- age
	}()

	// Fetch gender concurrently
	go func() {
		gender, err := s.APIClient.FetchGender(p.FirstName)
		if err != nil {
			errCh <- err
			return
		}
		genderCh <- gender
	}()

	// Fetch nationality concurrently
	go func() {
		nationality, err := s.APIClient.FetchNationality(p.FirstName)
		if err != nil {
			errCh <- err
			return
		}
		nationalityCh <- nationality
	}()

	// Wait for all concurrent fetches to complete
	for i := 0; i < 3; i++ {
		select {
		case age := <-ageCh:
			p.Age = age
		case gender := <-genderCh:
			p.Gender = gender
		case nationality := <-nationalityCh:
			p.Nationality = nationality
		case err := <-errCh:
			log.Errorf("failed to fetch data: %v", err)
			return nil, err
		}
	}

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
