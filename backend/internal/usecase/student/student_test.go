package usecase

// import (
// 	"encoding/json"
// 	"fmt"
// 	"golang_graphs/backend/internal/domain/model"
// 	service "golang_graphs/backend/internal/domain/student/service/graphconverter"
// 	"strconv"
// 	"testing"
// )

// func ConvertJSONStructsToGraph(nodes_json []service.NodeJSON, edges_json []service.EdgeJSON) (*model.Graph, error) {
// 	graph := new(model.Graph)
// 	node_id_map := make(map[string]int)
// 	curr_id := 0
// 	for _, node_json := range nodes_json {
// 		node_id_map[node_json.NodeData.Id] = curr_id
// 		weight, err := strconv.Atoi(node_json.NodeData.Weight)
// 		if err != nil {
// 			weight = 0
// 		}
// 		graph.AddNodeByInfo(curr_id,
// 			node_json.NodeData.Label,
// 			node_json.NodeData.Color,
// 			weight,
// 			node_json.Position.X,
// 			node_json.Position.Y,
// 		)
// 		curr_id++
// 	}
// 	for _, edge := range edges_json {
// 		src, err := graph.FindNodeById(node_id_map[edge.EdgeData.Source])
// 		if err != nil {
// 			return graph, err
// 		}
// 		trg, _ := graph.FindNodeById(node_id_map[edge.EdgeData.Target])
// 		if err != nil {
// 			return graph, err
// 		}
// 		weight_edge, err := strconv.Atoi(edge.EdgeData.Label)
// 		if err != nil {
// 			weight_edge = 0
// 		}
// 		graph.AddEdgeByInfo(
// 			src,
// 			trg,
// 			edge.EdgeData.Id,
// 			edge.EdgeData.Label,
// 			edge.EdgeData.Color,
// 			weight_edge)
// 	}
// 	return graph, nil
// }

// func TestGraph_MinPath1(t *testing.T) {
// 	payload := `{"labID":2,"modules":[{"type":2,"taskID":13,"data":[{"nodes":[{"data":{"id":"0","label":"0","color":"","weight":"0"},"position":{"x":886.2190272116819,"y":288.4123589060649}},{"data":{"id":"1","label":"1","color":"","weight":"10"},"position":{"x":708.3268590217984,"y":51.4543084288449}},{"data":{"id":"2","label":"2","color":"","weight":"5"},"position":{"x":708.3268590217984,"y":435}},{"data":{"id":"3","label":"3","color":"","weight":"9"},"position":{"x":169.99109987735656,"y":117.88170083485497}},{"data":{"id":"4","label":"4","color":"","weight":"8"},"position":{"x":240.67314097820153,"y":434.99999999999994}},{"data":{"id":"5","label":"5","color":"","weight":"6"},"position":{"x":390.347007140179,"y":217.64391154662636}}],"edges":[{"data":{"source":"3","target":"5","label":"4","color":"","id":"66a593ee-1158-4dba-9457-2da5c2106aee"}},{"data":{"source":"2","target":"0","label":"5","color":"#7e3a3a","id":"d04b77d4-a97e-4193-9a05-a130f02bdeb5"}},{"data":{"source":"4","target":"1","label":"6","color":"","id":"523e4226-dfdc-4f48-91b5-d5fc2fb17d9a"}},{"data":{"source":"1","target":"5","label":"6","color":"","id":"9d6cc6aa-ab7d-434d-8564-0533fcda472b"}},{"data":{"source":"5","target":"2","label":"1","color":"#7e3a3a","id":"3b84dee5-de87-43cb-ad61-fc49de7f1b88"}},{"data":{"source":"4","target":"2","label":"3","color":"","id":"02e5df5b-4a12-4670-a33f-888c714101fb"}},{"data":{"source":"2","target":"1","label":"5","color":"","id":"e50acd01-188e-4bd8-b6ad-aa451325c737"}},{"data":{"source":"1","target":"3","label":"1","color":"","id":"08b3a118-a2a1-4caf-920f-6ecaeb269638"}},{"data":{"source":"3","target":"4","label":"1","color":"","id":"073410bf-b625-401e-ae18-8968e98dcf36"}}]}],"minPathAns":6,"weightPathAns":{"0":0,"1":10,"2":5,"3":9,"4":8,"5":6}}]}`
// 	ans := new(SendAnswersDTOIn)
// 	// var parser_ans []ModulesDataJSON
// 	if err := json.Unmarshal([]byte(payload), &ans); err != nil {
// 		fmt.Println(err)
// 	}

// 	g, _ := ConvertJSONStructsToGraph(ans.Modules[0].DataModule[0].Nodes, ans.Modules[0].DataModule[0].Edges)
// 	fmt.Printf("g: %+v\n", ans.Modules[0].DataModule[0].Nodes)
// 	src_node := model.Node{
// 		Id: 1, Label: "0",
// 	}
// 	trg_node := model.Node{
// 		Id: 5, Label: "5",
// 	}
// 	i, w := g.MinPath(src_node, trg_node, true)
// 	fmt.Println(i, w)
// }
