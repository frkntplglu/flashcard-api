package models

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type FailedResponse struct {
	Success bool  `json:"success"`
	Error   Error `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
