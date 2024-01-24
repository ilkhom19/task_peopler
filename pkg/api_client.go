package pkg

import (
	"encoding/json"
	"fmt"
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
		return 0, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	age, ok := data["age"].(float64)
	if !ok {
		return 0, fmt.Errorf("age not found in API response")
	}

	return int(age), nil
}

func (c *APIClient) FetchGender(name string) (string, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	gender, ok := data["gender"].(string)
	if !ok {
		return "", fmt.Errorf("gender not found in API response")
	}

	return gender, nil
}

func (c *APIClient) FetchNationality(name string) (string, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	countries, ok := data["country"].([]interface{})
	if !ok || len(countries) == 0 {
		return "", fmt.Errorf("nationality not found in API response")
	}

	nationality := countries[0].(map[string]interface{})["country_id"].(string)
	return nationality, nil
}
