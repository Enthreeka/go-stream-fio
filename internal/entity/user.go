package entity

//type User struct {
//	Firstname  string `json:"firstname"`
//	Lastname   string `json:"lastname"`
//	Patronymic string `json:"patronymic"`
//
//	Age     Age       `json:"age,omitempty"`
//	Gender  Gender    `json:"gender,omitempty"`
//	Country []Country `json:"country,omitempty"`
//}
//
//type Age struct {
//	Age string `json:"age"`
//}
//
//type Gender struct {
//	Gender      string  `json:"gender"`
//	Probability float32 `json:"probability"`
//}
//
//type Country struct {
//	CountryCode string  `json:"country_code"`
//	Probability float32 `json:"probability"`
//}

type Data struct {
	Total int      `json:"total"`
	Data  []Person `json:"data"`
}

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`

	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`

	Address Address `json:"address"`
}

type Address struct {
	CountryCode string `json:"county_code"`
}
