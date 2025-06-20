package handler

import service "golang_graphs/backend/internal/domain/student/service/graphconverter"

type Module struct {
	TypeID      int        `json:"type"`
	DataModule  DataAnswer `json:"data"`
	InputValue1 string     `json:"inputValue1"`
	InputValue2 string     `json:"inputValue2"`
	InputValue3 string     `json:"inputValue3"`
	InputValue4 string     `json:"inputValue4"`
	InputValue5 string     `json:"inputValue5"`
	InputValue6 string     `json:"inputValue6"`
}

type DataAnswer struct {
	Nodes []service.NodeJSON `json:"nodes"`
	Edges []service.EdgeJSON `json:"edges"`
}

type SendAnswersRequest struct {
	Modules []Module `json:"modules"`
}

// type SendAnswersResponse struct {
// 	Result dto.Result `json:"result"`
// }
