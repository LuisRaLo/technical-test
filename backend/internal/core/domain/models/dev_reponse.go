package models

type DevResponse struct {
	StatusCode int `json:"status"`
	Response   any `json:"response"`
}
