package studentservice

import (
	"context"
	"encoding/json"
	"golang_graphs/backend/internal/models"
	"strconv"
)

type StudentService interface {
}

type studentService struct {
}

func NewStudentService() StudentService {
	return &studentService{}
}

func (c *controller) parserNodeStruct(ctx context.Context, json_str string) (*models.NodesJSON, error) {
	ans := new(models.NodesJSON)
	var parser_ans []models.NodeJSON
	if err := json.Unmarshal([]byte(json_str), &parser_ans); err != nil {
		return nil, err
	}
	ans.NodeArr = parser_ans
	return ans, nil
}

func (c *controller) parserEdgeStruct(ctx context.Context, json_str string) (*models.EdgesJSON, error) {
	ans := new(models.EdgesJSON)
	var parser_ans []models.EdgeJSON
	if err := json.Unmarshal([]byte(json_str), &parser_ans); err != nil {
		return nil, err
	}
	ans.EdgeArr = parser_ans
	return ans, nil
}

func (c *controller) convertJSONStructsToGraph(ctx context.Context, nodes_json *models.NodesJSON, edges_json *models.EdgesJSON) (*models.Graph, error) {
	graph := new(models.Graph)
	node_id_map := make(map[string]int)
	curr_id := 0
	for _, node_json := range nodes_json.NodeArr {
		node_id_map[node_json.NodeData.Id] = curr_id
		curr_id++
		weight, err := strconv.Atoi(node_json.NodeData.Weight)
		if err != nil {
			weight = 0
		}
		graph.AddNodeByInfo(curr_id,
			node_json.NodeData.Label,
			node_json.NodeData.Color,
			weight,
			node_json.Position.X,
			node_json.Position.Y)
	}

	for _, edge := range edges_json.EdgeArr {
		src, err := graph.FindNodeById(node_id_map[edge.EdgeData.Source])
		if err != nil {
			return graph, err
		}
		trg, _ := graph.FindNodeById(node_id_map[edge.EdgeData.Target])
		if err != nil {
			return graph, err
		}
		graph.AddEdgeByInfo(
			src,
			trg,
			edge.EdgeData.Id,
			edge.EdgeData.Label,
			edge.EdgeData.Color,
			0)
	}
	return graph, nil
}

// func (c *controller) parserGraphStruct(ctx context.Context, request models.SendAnswersRequest) (*models.Graph, error) {
// 	graph := new(models.Graph)
// 	var json_struct
// 	// пробегаться по циклу request.Answers?
// 	if err := json.Unmarshal([]byte(request.Answers[0].Answer), &json_struct); err != nil {
// 		return nil, err
// 	}
// 	for _, data := range json_struct {
// 		if data.source == nil {

// 		}
// 	}
// }
// нужны парсеры для разных заданий
// узнать структуру всего json, приходящего с фронта, чтобы вытащить всю инфу по заданию
// в зависимости от задания (модуля) вызывать нужную функцию
