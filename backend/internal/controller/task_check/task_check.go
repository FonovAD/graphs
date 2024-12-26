package task_check

import (
	"golang_graphs/backend/internal/models"
	"math"
	"strconv"

	gograph "github.com/yourbasic/graph"
)

const (
	LINEAR_MODULE_COEFF              int    = 20
	RADIUS_AND_DIAMETER_MODULE_COEFF int    = 20
	ADJACENT_MATRIX_MODULE_COEFF     int    = 20
	MIN_PATH_MODULE_COEFF            int    = 20
	DEFAULT_COLOR                    string = ""
)

func createGographWithouInfo(graph models.Graph) *gograph.Mutable {
	if len(graph.Nodes) == 0 {
		return gograph.New(0)
	}

	m := make(map[int]int)

	for i := 0; i < len(graph.Nodes); i++ {
		m[graph.Nodes[i].Id] = i
	}

	g := gograph.New(len(graph.Nodes))

	for _, edge := range graph.Edges {
		g.AddBoth(m[edge.Source.Id], m[edge.Target.Id])
	}
	return g
}

type checker struct {
}

type Checker interface {
	CheckLinearToLine(task models.Graph, answer models.Graph) int
	CheckLinearFromLine(task models.Graph, answer models.Graph) int
	CheckRadiusAndDiameter(task models.Graph, radius_ans int, diameter_ans int, dist_matrix_ans map[string]map[string]int) int
	CheckAdjacentMatrix(task models.Graph, answer map[string]map[string]int) int
	CheckEulerGraph(task models.Graph, answer_quest bool, answer_graph models.Graph) int
	CheckMinPath(task models.Graph, source string, target string, min_path_ans int, weights_path_ans map[string]int, answer models.Graph) int
	CheckPlanarGraph(answer models.Graph) int
}

func max_(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Проверка модуля "Реберный граф" (из графа в реберный)
func (ch *checker) CheckLinearToLine(task models.Graph, answer models.Graph) int {
	task_adj := task.EdgeLabelAdjacentMatrix()
	answer_adj := answer.NodeLabelAdjacentMatrix()

	if len(task_adj) != len(answer_adj) {
		return 0
	}
	err_count := 0
	for label1, adj_arr := range task_adj {
		for label2, task_val := range adj_arr {
			ans_val, ok := answer_adj[label1][label2]
			if !ok || task_val != ans_val {
				err_count++
			}
		}
	}
	return max_(0, 100-err_count*LINEAR_MODULE_COEFF)
}

// Проверка модуля "Реберный граф" (из реберного в граф)
func (ch *checker) CheckLinearFromLine(task models.Graph, answer models.Graph) int {
	task_adj := task.NodeLabelAdjacentMatrix()
	answer_adj := answer.EdgeLabelAdjacentMatrix()

	if len(task_adj) != len(answer_adj) {
		return 0
	}
	err_count := 0
	for label1, adj_arr := range task_adj {
		for label2, task_val := range adj_arr {
			ans_val, ok := answer_adj[label1][label2]
			if !ok || task_val != ans_val {
				err_count++
			}
		}
	}
	return max_(0, 100-err_count*LINEAR_MODULE_COEFF)
}

// Проверки модуля "Радиус и диметр"
func (ch *checker) CheckRadiusAndDiameter(task models.Graph, radius_ans int, diameter_ans int, dist_matrix_ans map[string]map[string]int) int {
	dist_matrix := task.DistanceMatrix(false)
	max_dists := make([]int, 0, len(dist_matrix))
	err_count := 0
	for node1, node1_arr := range dist_matrix {
		max_dist := 0
		for node2, dist := range node1_arr {
			if dist > max_dist {
				max_dist = dist
			}
			dist_ans, ok := dist_matrix_ans[node1][node2]
			if !ok || dist_ans != dist {
				err_count++
			}
		}
		max_dists = append(max_dists, max_dist)
	}
	radius := math.MaxInt
	diameter := 0
	for _, dist := range max_dists {
		if dist < radius {
			radius = dist
		}
		if dist > diameter {
			diameter = dist
		}
	}
	if radius != radius_ans || diameter != diameter_ans {
		return 0
	}
	return max_(0, 100-err_count*RADIUS_AND_DIAMETER_MODULE_COEFF)
}

func (ch *checker) CheckAdjacentMatrix(task models.Graph, answer map[string]map[string]int) int {
	adj_matrix := task.NodeLabelAdjacentMatrix()
	if len(adj_matrix) != len(answer) {
		return 0
	}
	err_count := 0
	for node1, node1_arr := range adj_matrix {
		for node2, val := range node1_arr {
			val_ans, ok := answer[node1][node2]
			if !ok || val_ans != val {
				err_count++
			}
		}
	}
	return max_(0, 100-err_count*ADJACENT_MATRIX_MODULE_COEFF)
}

// Проыерка модуля "Эйлеров граф"
// Должно быть больше 2 ребер
func (ch *checker) CheckEulerGraph(task models.Graph, is_euler_ans bool, answer_graph models.Graph) int {
	task_gograph := createGographWithouInfo(task)
	_, is_euler := gograph.EulerUndirected(task_gograph)
	if is_euler != is_euler_ans {
		return 0
	}
	if !is_euler && !is_euler_ans {
		return 100
	}
	walk_ans_labels := make(map[string]models.Edge)
	for _, edge := range answer_graph.Edges {
		walk_ans_labels[edge.Label] = edge
	}
	n_edges := len(answer_graph.Edges)
	if len(walk_ans_labels) != n_edges {
		return 0
	}
	walk_ans := make([]models.Edge, n_edges+1)
	for label, edge := range walk_ans_labels {
		index, err := strconv.Atoi(label)
		if err != nil || index < 1 || index > n_edges {
			return 0
		}
		walk_ans[index] = edge
	}
	edge1 := walk_ans[1]
	edge2 := walk_ans[2]
	first_node := -1
	second_node := -1
	if edge1.Source.Id == edge2.Source.Id || edge1.Source.Id == edge2.Target.Id {
		first_node = edge1.Target.Id
		second_node = edge1.Source.Id
	} else if edge1.Target.Id == edge2.Source.Id || edge1.Target.Id == edge2.Target.Id {
		first_node = edge1.Source.Id
		second_node = edge2.Target.Id
	} else {
		return 0
	}
	prev_node := second_node
	for _, edge := range walk_ans[2:n_edges] {
		if edge.Source.Id == prev_node {
			prev_node = edge.Target.Id
		} else if edge.Target.Id == prev_node {
			prev_node = edge.Source.Id
		} else {
			return 0
		}
	}
	if (walk_ans[n_edges].Source.Id == prev_node && walk_ans[n_edges].Target.Id == first_node) || (walk_ans[n_edges].Target.Id == prev_node && walk_ans[n_edges].Source.Id == first_node) {
		return 100
	}
	return 0
}

// Проверка модуля "Кратчайший путь"
// Веса должны быть положительными
func (ch *checker) CheckMinPath(task models.Graph, source string, target string, min_path_ans int, weights_path_ans map[string]int, answer models.Graph) int {
	source_node := task.Nodes[0]
	target_node := task.Nodes[0]
	for _, node := range task.Nodes {
		if source == node.Label {
			source_node = node
		}
		if target == node.Label {
			target_node = node
		}
	}
	min_path, weights_path := task.MinPath(source_node, target_node, true)
	if min_path != min_path_ans {
		return 0
	}
	err_count := 0
	for node_label, weight := range weights_path {
		weight_ans, ok := weights_path_ans[node_label]
		if !ok || weight_ans != weight {
			err_count++
		}
	}
	min_path_check := 0
	for _, edge := range answer.Edges {
		if edge.Color != DEFAULT_COLOR {
			min_path_check += edge.Weight
		}
	}
	if min_path_check != min_path {
		return 0
	}
	return max_(0, 100-err_count*MIN_PATH_MODULE_COEFF)
}

// func (ch *checker) CheckPlanarGraph(answer models.Graph) int {
// 	for _, edge1 := range answer.Edges {
// 		for edge2 := range answer.Edges {

// 		}
// 	}
// }
