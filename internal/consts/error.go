package consts

import "net/http"

var (
	ErrorDescriptions = map[int]string{
		http.StatusBadRequest:          "error in user request",
		http.StatusInternalServerError: "cant handle request, internal error",
		http.StatusNotFound:            "not found",
	}
)
