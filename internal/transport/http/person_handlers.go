package http

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /people [post]

func (h *Handler) createPerson(w http.ResponseWriter, r *http.Request) {
	var input models.PersonCreateInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logError("createPerson", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
	}

	if err := input.Validate(); err != nil {
		logError("createPerson", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
		return
	}

	person := &models.Person{
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Patronymic: input.Patronymic,
	}

	person, err := h.PersonService.Create(person)
	if err != nil {
		logError("createPerson", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	response, err := json.Marshal(person)
	if err != nil {
		logError("createPerson", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

// UpdatePerson ... Update person
// @Summary Update person
// @Description Update person
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Param order body PersonUpdateInput true "Person"
// @Success 200 {object} Person
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /people/{id} [put]

func (h *Handler) updatePerson(w http.ResponseWriter, r *http.Request) {
	var input models.PersonUpdateInput

	id, err := getIDFromRequest(r)
	if err != nil {
		logrus.Error("updatePerson error:", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logError("updatePerson", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
	}

	if err := input.Validate(); err != nil {
		logError("updatePerson", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
		return
	}

	person := &models.Person{
		ID:         id,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Patronymic: input.Patronymic,
	}

	person, err = h.PersonService.Update(person)
	if err != nil {
		logError("updatePerson", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	response, err := json.Marshal(person)
	if err != nil {
		logError("updatePerson", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

// DeletePerson ... Delete person
// @Summary Delete person
// @Description Delete person
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /people/{id} [delete]

func (h *Handler) deletePerson(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		logrus.Error("deletePerson error:", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
		return
	}

	if err := h.PersonService.Delete(id); err != nil {
		logError("deletePerson", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	WriteJSONToResponse(w, http.StatusOK, "success", "person deleted")
}

// GetPerson ... Get person
// @Summary Get person
// @Description Get person
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} Person
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /people/{id} [get]

func (h *Handler) getPersonByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		logrus.Error("getPersonByID error:", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
		return
	}

	person, err := h.PersonService.GetByID(id)
	if err != nil {
		logError("getPersonByID", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	response, err := json.Marshal(person)
	if err != nil {
		logError("getPersonByID", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

// GetPeopleByName ... Get people by name
// @Summary Get people by name
// @Description Get people by name
// @Tags person
// @Accept json
// @Produce json
// @Param name query string true "Person Name"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} PeopleResponse
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /people [get]

func (h *Handler) getPeopleByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	page, pageSize, err := getPageAndPageSizeFromQuery(r)
	if err != nil {
		logrus.Error("getPeopleByName error:", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
		return
	}

	people, err := h.PersonService.GetPeopleByName(name, page, pageSize)
	if err != nil {
		logError("getPeopleByName", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	response, err := json.Marshal(people)
	if err != nil {
		logError("getPeopleByName", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

// GetPeopleByAge ... Get people by age
// @Summary Get people by age
// @Description Get people by age
// @Tags person
// @Accept json
// @Produce json
// @Param age query int true "Person Age"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} PeopleResponse
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /people [get]

func (h *Handler) getPeopleByAge(w http.ResponseWriter, r *http.Request) {
	age, err := strconv.Atoi(r.URL.Query().Get("age"))
	if err != nil {
		logrus.Error("getPeopleByAge error:", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
		return
	}

	page, pageSize, err := getPageAndPageSizeFromQuery(r)
	if err != nil {
		logrus.Error("getPeopleByAge error:", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
		return
	}

	people, err := h.PersonService.GetPeopleByAge(age, page, pageSize)
	if err != nil {
		logError("getPeopleByAge", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	response, err := json.Marshal(people)
	if err != nil {
		logError("getPeopleByAge", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

// GetPeopleByGender ... Get people by gender
// @Summary Get people by gender
// @Description Get people by gender
// @Tags person
// @Accept json
// @Produce json
// @Param gender query string true "Person Gender"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} PeopleResponse
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 401 {object} Response
// @Failure 500 {object} Response
// @Router /people/gender [get]

func (h *Handler) getPeopleByGender(w http.ResponseWriter, r *http.Request) {
	gender := r.URL.Query().Get("gender")
	page, pageSize, err := getPageAndPageSizeFromQuery(r)
	if err != nil {
		logrus.Error("getPeopleByGender error:", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
		return
	}

	people, err := h.PersonService.GetPeopleByGender(gender, page, pageSize)
	if err != nil {
		logError("getPeopleByGender", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	response, err := json.Marshal(people)
	if err != nil {
		logError("getPeopleByGender", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

// GetPeopleByNationality ... Get people by nationality
// @Summary Get people by nationality
// @Description Get people by nationality
// @Tags person
// @Accept json
// @Produce json
// @Param nationality query string true "Person Nationality"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} PeopleResponse
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 401 {object} Response
// @Failure 500 {object} Response
// @Router /people/nationality [get]

func (h *Handler) getPeopleByNationality(w http.ResponseWriter, r *http.Request) {
	nationality := r.URL.Query().Get("nationality")
	page, pageSize, err := getPageAndPageSizeFromQuery(r)
	if err != nil {
		logrus.Error("getPeopleByNationality error:", err)
		WriteJSONToResponse(w, http.StatusBadRequest, "error", err.Error())
		return
	}

	people, err := h.PersonService.GetPeopleByNationality(nationality, page, pageSize)
	if err != nil {
		logError("getPeopleByNationality", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	response, err := json.Marshal(people)
	if err != nil {
		logError("getPeopleByNationality", err)
		WriteJSONToResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
