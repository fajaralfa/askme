package model

type Response struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
