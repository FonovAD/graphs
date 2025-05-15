package handler

type BadRequestResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type InternalServerErrorResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type UnauthorizedResponse struct {
	ErrorMsg string `json:"error_msg"`
}
