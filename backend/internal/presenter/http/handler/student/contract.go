package handler

import service "golang_graphs/backend/internal/domain/student/service/graphconverter"

type Module struct {
	TaskID     int64      `json:"type"`
	DataModule DataAnswer `json:"data"`
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
