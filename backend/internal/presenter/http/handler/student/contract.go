package studenthandler

import "golang_graphs/backend/internal/domain/student/service/graphconverter"

type Module struct {
	TaskID     int64      `json:"type"`
	DataModule DataAnswer `json:"data"`
}

type DataAnswer struct {
	Nodes []graphconverter.NodeJSON `json:"nodes"`
	Edges []graphconverter.EdgeJSON `json:"edges"`
}

type SendAnswersRequest struct {
	Modules []Module `json:"modules"`
}
