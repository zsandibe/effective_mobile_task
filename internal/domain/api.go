package domain

type ApiAgeResponse struct {
	Age int `json:"age"`
}

type ApiGenderResponse struct {
	Gender string `json:"gender"`
}

type ApiNationalityResponse struct {
	Nationality []struct {
		CountryId string `json:"country_id"`
	} `json:"country"`
}
