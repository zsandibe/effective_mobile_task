package domain

type Person struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type SearchParams struct {
	Gender      string
	Nationality string
	PageStart   int
	PageLimit   int
}
