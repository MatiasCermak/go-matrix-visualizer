package domain

type Response struct {
	Timestamp string   `json:"ts"`
	Rows      [][]uint `json:"rows,omitempty"`
	Message   string   `json:"message,omitempty"`
}
