package http

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"task_peopler/internal/models"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type PeopleResponse struct {
	People []models.Person `json:"people"`
}

func WriteJSONToResponse(w http.ResponseWriter, httpStatus int, bodyStatus string, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	body := Response{
		Status:  bodyStatus,
		Message: message,
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		log.Error("Failed to Marshall json:", err)
	}

	w.Write(bodyJson)
}

func logFields(handler string) log.Fields {
	return log.Fields{
		"handler": handler,
	}
}

func logError(handler string, err error) {
	log.WithFields(logFields(handler)).Error(err)
}
