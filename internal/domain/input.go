package domain

type UrlParams struct {
	Gender      string
	Nationality string
	Page        int
	Size        int
}

type UserInput struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}
