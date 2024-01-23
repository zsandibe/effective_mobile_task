package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zsandibe/effective_mobile_task/config"
	"github.com/zsandibe/effective_mobile_task/internal/domain"
	"github.com/zsandibe/effective_mobile_task/pkg"
)

type Enrichment interface {
	GetPersonAgeByName(name string) (int, error)
	GetPersonGenderByName(name string) (string, error)
	GetPersonNationalityByName(name string) (string, error)
}

type enrichment struct {
	config *config.Config
}

func NewEnrichment(config *config.Config) Enrichment {
	return &enrichment{
		config: config,
	}
}

const keyName = "name"

func (s enrichment) GetPersonAgeByName(name string) (int, error) {
	pkg.InfoLog.Println(s.config.Api.AgeURL, "laall")
	// pkg.InfoLog.Println("getResponseBody")
	body, err := getResponseBody(s.config.Api.AgeURL, keyName, name)
	if err != nil {
		return 0, err
	}
	var response domain.ApiAgeResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return 0, fmt.Errorf("problems with unmarshalling response: %v", err)
	}

	if response.Age < 0 {
		return 0, domain.AgeNotFound
	}

	return response.Age, nil
}

func (s enrichment) GetPersonGenderByName(name string) (string, error) {
	body, err := getResponseBody(s.config.Api.GenderURL, keyName, name)
	if err != nil {
		return "", err
	}

	var response domain.ApiGenderResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("problems with unmarshalling response: %v", err)
	}
	if response.Gender == "" {
		return "", domain.GenderNotFound
	}
	return response.Gender, nil
}

func (s enrichment) GetPersonNationalityByName(name string) (string, error) {
	body, err := getResponseBody(s.config.Api.NationalityURL, keyName, name)
	if err != nil {
		return "", err
	}

	var response domain.ApiNationalityResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("problems with unmarshalling response: %v", err)
	}
	if len(response.Nationality) == 0 {
		return "", domain.NationalityNotFound
	}
	return response.Nationality[0].CountryId, nil
}

func getResponseBody(url, key, name string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	query := req.URL.Query()
	query.Add(key, name)

	req.URL.RawQuery = query.Encode()

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to requesting: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	return body, nil

}
