package domain

type Response struct {
	Message string `json:"message"`
}

type Sort struct {
	Field string
	// ASC DESC
	Order string
}

type Range struct {
	Start  int
	Length int
}
