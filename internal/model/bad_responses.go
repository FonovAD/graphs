package model

type BadRequestResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type InternalServerErrorResponse struct {
	ErrorMsg string `json:"error_msg"`
}
