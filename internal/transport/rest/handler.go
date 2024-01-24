package rest

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"task_peopler/internal/models"
)

type PersonService interface {
	Create(p *models.Person) (*models.Person, error)
	Update(p *models.Person) (*models.Person, error)
	Delete(id int64) error
	GetByID(id int64) (*models.Person, error)
	GetPeopleByName(name string, page, pageSize int) ([]*models.Person, error)
	GetPeopleByAge(age int, page, pageSize int) ([]*models.Person, error)
	GetPeopleByGender(gender string, page, pageSize int) ([]*models.Person, error)
	GetPeopleByNationality(nationality string, page, pageSize int) ([]*models.Person, error)
}

type Handler struct {
	PersonService PersonService
}

func NewHandler(personService PersonService) *Handler {
	return &Handler{PersonService: personService}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/people", h.createPerson).Methods("POST")
	router.HandleFunc("/people/{id}", h.updatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", h.deletePerson).Methods("DELETE")
	router.HandleFunc("/people/{id}", h.getPersonByID).Methods("GET")

	router.HandleFunc("/people/name", h.getPeopleByName).Methods("GET")
	router.HandleFunc("/people/age", h.getPeopleByAge).Methods("GET")
	router.HandleFunc("/people/gender", h.getPeopleByGender).Methods("GET")
	router.HandleFunc("/people/nationality", h.getPeopleByNationality).Methods("GET")

	return router
}

func getIDFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id can't be 0")
	}

	return id, nil
}

func getPageAndPageSizeFromQuery(r *http.Request) (int, int, error) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		return 0, 0, err
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		return 0, 0, err
	}

	return page, pageSize, nil
}
