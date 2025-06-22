package service

import (
	"context"
	"encoding/json"
	model "golang_graphs/backend/internal/domain/model"
	"strconv"
)

type GraphConverter interface {
	ConvertJSONStructsToGraph(ctx context.Context, nodes_json []NodeJSON, edges_json []EdgeJSON) (*model.Graph, error)
}

type graphConverter struct {
}

func NewGraphConverter() GraphConverter {
	return &graphConverter{}
}

func (gc *graphConverter) parserNodeStruct(ctx context.Context, json_str string) (*NodesJSON, error) {
	ans := new(NodesJSON)
	var parser_ans []NodeJSON
	if err := json.Unmarshal([]byte(json_str), &parser_ans); err != nil {
		return nil, err
	}
	ans.NodeArr = parser_ans
	return ans, nil
}

func (gc *graphConverter) parserEdgeStruct(ctx context.Context, json_str string) (*EdgesJSON, error) {
	ans := new(EdgesJSON)
	var parser_ans []EdgeJSON
	if err := json.Unmarshal([]byte(json_str), &parser_ans); err != nil {
		return nil, err
	}
	ans.EdgeArr = parser_ans
	return ans, nil
}

func (gc *graphConverter) ConvertJSONStructsToGraph(ctx context.Context, nodes_json []NodeJSON, edges_json []EdgeJSON) (*model.Graph, error) {
	graph := new(model.Graph)
	node_id_map := make(map[string]int)
	curr_id := 0
	for _, node_json := range nodes_json {
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
	for _, edge := range edges_json {
		src, err := graph.FindNodeById(node_id_map[edge.EdgeData.Source])
		if err != nil {
			return graph, err
		}
		trg, _ := graph.FindNodeById(node_id_map[edge.EdgeData.Target])
		if err != nil {
			return graph, err
		}
		weight_edge, err := strconv.Atoi(edge.EdgeData.Label)
		if err != nil {
			weight_edge = 0
		}
		graph.AddEdgeByInfo(
			src,
			trg,
			edge.EdgeData.Id,
			edge.EdgeData.Label,
			edge.EdgeData.Color,
			weight_edge)
	}
	return graph, nil
}

func (gc *graphConverter) parserModuleStruct(ctx context.Context, json_str string) (*ModulesJSON, error) {
	ans := new(ModulesJSON)
	// var parser_ans []ModulesDataJSON
	if err := json.Unmarshal([]byte(json_str), &ans); err != nil {
		return nil, err
	}
	// ans.Modules = parser_ans
	return ans, nil
}

// func (c *controller) parserGraphStruct(ctx context.Context, request  SendAnswersRequest) (* Graph, error) {
// 	graph := new( Graph)
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
