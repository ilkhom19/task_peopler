package pkg

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type APIClient struct{}

func NewAPIClient() *APIClient {
	return &APIClient{}
}

func (c *APIClient) FetchAge(name string) (int, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		log.Error("Failed to fetch age:", err)
		return 0, nil
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Error("Failed to decode age:", err)
		return 0, nil
	}

	age, ok := data["age"].(float64)
	if !ok {
		log.Error("Failed to parse age:", err)
		return 0, nil
	}

	return int(age), nil
}

func (c *APIClient) FetchGender(name string) (string, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	unknown := "unknown"
	resp, err := http.Get(url)
	if err != nil {
		log.Error("Failed to fetch the gender:", err)
		return unknown, nil
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Error("Failed to decode the gender:", err)
		return unknown, nil
	}

	gender, ok := data["gender"].(string)
	if !ok {
		log.Error("Failed to parse the gender:", err)
		return unknown, nil
	}

	return gender, nil
}

func (c *APIClient) FetchNationality(name string) (string, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	unknown := "unknown"
	resp, err := http.Get(url)
	if err != nil {
		log.Error("Failed to fetch the nationality:", err)
		return unknown, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Error("Failed to decode the nationality:", err)
		return unknown, err
	}

	countries, ok := data["country"].([]interface{})
	if !ok || len(countries) == 0 {
		log.Error("Failed to parse the nationality:", err)
		return unknown, nil
	}

	nationality := countries[0].(map[string]interface{})["country_id"].(string)
	return nationality, nil
}
