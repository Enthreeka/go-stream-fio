package faker

type Data struct {
	Total int        `json:"total"`
	Data  []FakeUser `json:"data"`
}

type FakeUser struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`

	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`

	Address Address `json:"address"`
}

type Address struct {
	CountryCode string `json:"county_code"`
	Probability string `json:"probability,omitempty"`
}
