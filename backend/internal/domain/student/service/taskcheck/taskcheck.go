package service

import (
	model "golang_graphs/backend/internal/domain/model"
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

func createGographWithoutInfo(graph model.Graph) *gograph.Mutable {
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
	CheckLinearToLine(task, answer *model.Graph) int
	CheckLinearFromLine(task model.Graph, answer model.Graph) int
	CheckRadiusAndDiameter(task model.Graph, radius_ans int, diameter_ans int, dist_matrix_ans map[string]map[string]int) int
	CheckAdjacentMatrix(task model.Graph, answer map[string]map[string]int) int
	CheckEulerGraph(task model.Graph, is_euler_ans bool, answer_graph model.Graph) int
	CheckMinPath(task model.Graph, source string, target string, min_path_ans int, weights_path_ans map[string]int, answer model.Graph) int
	CheckPlanarGraph(answer model.Graph) int
	CheckIntersectionGraphs(answer, graph1, graph2 *model.Graph) int
	CheckUnionGraphs(answer, graph1, graph2 *model.Graph) int
	CheckJoinGraphs(answer, graph1, graph2 *model.Graph) int
	// Harary definition
	// CheckCartesianProduct(answer, graph1, graph2 *models.Graph) int
	// Gorbatov definition
	// CheckTensorProduct(answer, graph1, graph2 *models.Graph) int
	// Composition
	// CheckLexicographical(answer, graph1, graph2 *models.Graph) int
	CheckIntersectionMatrices(answer *model.Graph, matrix1, matrix2 map[string]map[string]int) int
	CheckUnionMatrices(answer *model.Graph, matrix1, matrix2 map[string]map[string]int) int
	CheckJoinMatrices(answer *model.Graph, matrix1, matrix2 map[string]map[string]int) int
}

// Харари - декартово, cartesian, box product
// Горбатов - тензорное, tensor, kroneker, categorial
// Композиция - лексикографическое, lexicographical

func NewChecker() Checker {
	return &checker{}
}

func Max_(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Проверка модуля "Реберный граф" (из графа в реберный)
func (ch *checker) CheckLinearToLine(task *model.Graph, answer *model.Graph) int {
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
	return Max_(0, 100-err_count*LINEAR_MODULE_COEFF)
}

// Проверка модуля "Реберный граф" (из реберного в граф)
func (ch *checker) CheckLinearFromLine(task model.Graph, answer model.Graph) int {
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
	return Max_(0, 100-err_count*LINEAR_MODULE_COEFF)
}

// Проверки модуля "Радиус и диметр"
func (ch *checker) CheckRadiusAndDiameter(task model.Graph, radius_ans int, diameter_ans int, dist_matrix_ans map[string]map[string]int) int {
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
	return Max_(0, 100-err_count*RADIUS_AND_DIAMETER_MODULE_COEFF)
}

func (ch *checker) CheckAdjacentMatrix(task model.Graph, answer map[string]map[string]int) int {
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
	return Max_(0, 100-err_count*ADJACENT_MATRIX_MODULE_COEFF)
}

// Проверка модуля "Эйлеров граф"
// Должно быть больше 2 ребер
func (ch *checker) CheckEulerGraph(task model.Graph, is_euler_ans bool, answer_graph model.Graph) int {
	task_gograph := createGographWithoutInfo(task)
	_, is_euler := gograph.EulerUndirected(task_gograph)
	if is_euler != is_euler_ans {
		return 0
	}
	if !is_euler && !is_euler_ans {
		return 100
	}
	walk_ans_labels := make(map[string]model.Edge)
	for _, edge := range answer_graph.Edges {
		walk_ans_labels[edge.Label] = edge
	}
	n_edges := len(answer_graph.Edges)
	if len(walk_ans_labels) != n_edges {
		return 0
	}
	walk_ans := make([]model.Edge, n_edges+1)
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
func (ch *checker) CheckMinPath(task model.Graph, source string, target string, min_path_ans int, weights_path_ans map[string]int, answer model.Graph) int {
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
	return Max_(0, 100-err_count*MIN_PATH_MODULE_COEFF)
}

func boundingBox(x1, x2, x3, x4 float64) bool {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if x3 > x4 {
		x3, x4 = x4, x3
	}
	return math.Max(x1, x3) > math.Max(x2, x4)
}

func pseudoScalar(node1, node2, node3 model.Node) float64 {
	return (node2.X-node1.X)*(node3.Y-node1.Y) - (node2.Y-node1.Y)*(node3.X-node1.X)
}

func isIntersect(edge1, edge2 model.Edge) bool {
	b1 := boundingBox(edge1.Source.X, edge1.Target.X, edge2.Source.X, edge2.Target.X)
	b2 := boundingBox(edge1.Source.Y, edge1.Target.Y, edge2.Source.Y, edge2.Target.Y)
	b3 := pseudoScalar(edge1.Source, edge1.Target, edge2.Source)*pseudoScalar(edge1.Source, edge1.Target, edge2.Target) <= 0
	b4 := pseudoScalar(edge2.Source, edge2.Target, edge1.Source)*pseudoScalar(edge2.Source, edge2.Target, edge1.Target) <= 0
	return b1 && b2 && b3 && b4
}

func (ch *checker) CheckPlanarGraph(answer model.Graph) int {
	for _, edge1 := range answer.Edges {
		for _, edge2 := range answer.Edges {
			if edge1.Id != edge2.Id && !answer.IsEdgesAdjacent(edge1, edge2) {
				if isIntersect(edge1, edge2) {
					return 0
				}
			}
		}
	}
	return 100
}

func (ch *checker) checkBinaryOperations(answer, true_answer *model.Graph) (int, int) {
	true_node_set := make(map[string]struct{})
	answer_node_set := make(map[string]struct{})
	for _, node := range true_answer.Nodes {
		true_node_set[node.Label] = struct{}{}
	}
	for _, node := range answer.Nodes {
		answer_node_set[node.Label] = struct{}{}
	}
	for label := range true_node_set {
		_, ok := answer_node_set[label]
		if !ok {
			return 0, 0
		}
	}
	for label := range answer_node_set {
		_, ok := true_node_set[label]
		if !ok {
			return 0, 0
		}
	}
	correct_edges := 0
	odd_edges := 0
	for _, edge_answer := range answer.Edges {
		if true_answer.FindEdge(edge_answer.Source.Label, edge_answer.Target.Label) {
			correct_edges++
		} else {
			odd_edges++
		}
	}
	return correct_edges, odd_edges
}

func (ch *checker) CheckIntersectionGraphs(answer, graph1, graph2 *model.Graph) int {
	true_answer := graph1.Intersect(graph2)
	correct_edges, odd_edges := ch.checkBinaryOperations(answer, true_answer)
	true_edges_count := len(true_answer.Edges)
	return Max_(0, int(math.Ceil(100.00*float64(correct_edges-odd_edges)/float64(true_edges_count))))
}

func (ch *checker) CheckUnionGraphs(answer, graph1, graph2 *model.Graph) int {
	true_answer := graph1.Union(graph2)
	correct_edges, odd_edges := ch.checkBinaryOperations(answer, true_answer)
	true_edges_count := len(true_answer.Edges)
	return Max_(0, int(math.Ceil(100.00*float64(correct_edges-odd_edges)/float64(true_edges_count))))
}

func (ch *checker) CheckJoinGraphs(answer, graph1, graph2 *model.Graph) int {
	true_answer := graph1.Union(graph2)
	correct_edges, odd_edges := ch.checkBinaryOperations(answer, true_answer)
	true_edges_count := len(true_answer.Edges)
	return Max_(0, int(math.Ceil(100.00*float64(correct_edges-odd_edges)/float64(true_edges_count))))
}

func (ch *checker) CheckIntersectionMatrices(answer *model.Graph, matrix1, matrix2 map[string]map[string]int) int {
	graph1 := model.MakeGraphFromAdjLabelMatrix(matrix1)
	graph2 := model.MakeGraphFromAdjLabelMatrix(matrix2)
	return ch.CheckIntersectionGraphs(answer, graph1, graph2)
}

func (ch *checker) CheckUnionMatrices(answer *model.Graph, matrix1, matrix2 map[string]map[string]int) int {
	graph1 := model.MakeGraphFromAdjLabelMatrix(matrix1)
	graph2 := model.MakeGraphFromAdjLabelMatrix(matrix2)
	return ch.CheckUnionGraphs(answer, graph1, graph2)
}

func (ch *checker) CheckJoinMatrices(answer *model.Graph, matrix1, matrix2 map[string]map[string]int) int {
	graph1 := model.MakeGraphFromAdjLabelMatrix(matrix1)
	graph2 := model.MakeGraphFromAdjLabelMatrix(matrix2)
	return ch.CheckJoinGraphs(answer, graph1, graph2)
}
