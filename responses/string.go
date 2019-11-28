package responses

type UppercaseResponse struct {
	Str string `json:"str"`
	Err error  `json:"err,omitempty"`
}

type CountResponse struct {
	Length int `json:"length"`
}
