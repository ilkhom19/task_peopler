package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"task_peopler/internal/models"
)

type PersonService interface {
	Create(p *models.Person) (*models.Person, error)
}

type Handler struct {
	PersonService PersonService
}

func NewHandler(personService PersonService) *Handler {
	return &Handler{PersonService: personService}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/person", h.createPerson).Methods("POST")

	return router
}

func (h *Handler) respond(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
	}
}

func (h *Handler) error(w http.ResponseWriter, r *http.Request, serverError int, err error) {
	h.respond(w, r, serverError, map[string]string{"error": err.Error()})
}