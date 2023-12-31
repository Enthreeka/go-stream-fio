package entity

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`

	Age    []Age    `json:"birthday"`
	Gender []Gender `json:"gender"`

	Address []Address `json:"address"`
}

type Gender struct {
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

type Age struct {
	Age         int     `json:"age"`
	Probability float32 `json:"probability"`
}

type Address struct {
	CountryCode string `json:"county_code"`
	Probability string `json:"probability,omitempty"`
}
