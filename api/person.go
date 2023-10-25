package api

type PersonInput struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type PersonAge struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonGender struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type PersonNationality struct {
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

type PeopleFilter struct {
	Gender string
}

type Pagination struct {
	Limit  uint
	Offset uint
}
