package handler

//	type SendAnswersResponse struct {
//		Result dto.Result `json:"result"`
//	}
type UnauthorizedResponse struct {
	ErrorMsg string `json:"error_msg"`
}
type BadRequestResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type InternalServerErrorResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type ForbiddenResponse struct {
	ErrorMsg string `json:"error_msg"`
}
