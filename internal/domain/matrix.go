package domain

//RequestPayload represent the body of the request
type RequestPayload struct {
	Matrix [][]int `json:"matrix"`
}

//Matrix represent the Matrix service
type Matrix interface {
	Rotate([][]int) ([][]int, error)
}
