package http

import (
	"encoding/json"
	"net/http"
	"task_peopler/internal/models"
)

// CreatePerson ... Create person
// @Summary Create person
// @Description Create person
// @Tags person
// @Accept json
// @Produce json
// @Param order body models.PersonCreateInput true "Person"
// @Success 200 {object} models.Person
// @Failure 404
// @Failure 401
// @Failure 500
// @Router /people [post]

func (h *Handler) createPerson(w http.ResponseWriter, r *http.Request) {
	var input models.PersonCreateInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.respond(w, r, http.StatusBadRequest, err)
		return
	}

	if err := input.Validate(); err != nil {
		h.respond(w, r, http.StatusBadRequest, err)
		return
	}

	person := &models.Person{
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Patronymic: input.Patronymic,
	}

	person, err := h.PersonService.Create(person)
	if err != nil {
		h.respond(w, r, http.StatusBadRequest, err)
		return
	}

	h.respond(w, r, http.StatusOK, person)
}
