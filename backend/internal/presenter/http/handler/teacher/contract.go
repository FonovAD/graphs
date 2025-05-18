package handler

type CreateUserRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	FatherName string `json:"father_name"`
}

type CreateUserResponse struct {
	Token string
}

type AuthUserResponse struct {
	Token string `json:"token"`
}

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
