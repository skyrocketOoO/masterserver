package domain

type Response struct {
	Message string `json:"message"`
}

type Sort struct {
	Field string
	// ASC DESC
	Order string
}

type Pagination struct {
	Page    int
	PerPage int
}

type PageInfo struct {
	HasNextPage bool `json:"hasNextPage"`
	HasPrevPage bool `json:"hasPreviousPage"`
}
