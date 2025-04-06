package main

import (
	"golang_graphs/backend/internal/controller/task_check"
	"golang_graphs/backend/internal/models"
	"math"
	"testing"
)

// Тест для MinPath
func TestGraph_MinPath(t *testing.T) {
	graph := models.Graph{
		Nodes: []models.Node{
			{Id: 1, Label: "A"},
			{Id: 2, Label: "B"},
			{Id: 3, Label: "C"},
		},
		Edges: []models.Edge{
			{Source: models.Node{Id: 1, Label: "A"}, Target: models.Node{Id: 2, Label: "B"}, Weight: 2},
			{Source: models.Node{Id: 2, Label: "B"}, Target: models.Node{Id: 3, Label: "C"}, Weight: 3},
			{Source: models.Node{Id: 1, Label: "A"}, Target: models.Node{Id: 3, Label: "C"}, Weight: 5},
		},
	}

	// Positive case
	minPath, weightsPath := graph.MinPath(models.Node{Id: 1, Label: "A"}, models.Node{Id: 3, Label: "C"}, true)
	expectedMinPath := 5
	expectedWeights := map[string]int{"A": 0, "B": 2, "C": 5}

	if minPath != expectedMinPath {
		t.Errorf("expected %d, got %d", expectedMinPath, minPath)
	}

	for key, value := range expectedWeights {
		if weightsPath[key] != value {
			t.Errorf("expected weights[%s] = %d, got %d", key, value, weightsPath[key])
		}
	}

	// Negative case (unreachable nodes)
	graph2 := models.Graph{
		Nodes: []models.Node{
			{Id: 1, Label: "A"},
			{Id: 2, Label: "B"},
		},
		Edges: []models.Edge{},
	}

	minPath2, _ := graph2.MinPath(models.Node{Id: 1, Label: "A"}, models.Node{Id: 2, Label: "B"}, true)
	if minPath2 != math.MaxInt {
		t.Errorf("expected unreachable path, got %d", minPath2)
	}
}

// Тест для NodeLabelAdjacentMatrix
func TestGraph_NodeLabelAdjacentMatrix(t *testing.T) {
	graph := models.Graph{
		Nodes: []models.Node{
			{Label: "A"},
			{Label: "B"},
		},
		Edges: []models.Edge{
			{Source: models.Node{Label: "A"}, Target: models.Node{Label: "B"}},
		},
	}

	adjMatrix := graph.NodeLabelAdjacentMatrix()
	expectedMatrix := map[string]map[string]int{
		"A": {"A": 0, "B": 1},
		"B": {"A": 1, "B": 0},
	}

	for node1, row := range expectedMatrix {
		for node2, expectedValue := range row {
			if adjMatrix[node1][node2] != expectedValue {
				t.Errorf("expected adjMatrix[%s][%s] = %d, got %d", node1, node2, expectedValue, adjMatrix[node1][node2])
			}
		}
	}
}

// Тест для NodeIdAdjacentMatrix
func TestGraph_NodeIdAdjacentMatrix(t *testing.T) {
	graph := models.Graph{
		Nodes: []models.Node{
			{Id: 1},
			{Id: 2},
		},
		Edges: []models.Edge{
			{Source: models.Node{Id: 1}, Target: models.Node{Id: 2}},
		},
	}
	adjMatrix := graph.NodeIdAdjacentMatrix()
	expectedMatrix := map[int]map[int]int{
		1: {1: 0, 2: 1},
		2: {1: 1, 2: 0},
	}

	for id1, row := range expectedMatrix {
		for id2, expectedValue := range row {
			if adjMatrix[id1][id2] != expectedValue {
				t.Errorf("expected adjMatrix[%d][%d] = %d, got %d", id1, id2, expectedValue, adjMatrix[id1][id2])
			}
		}
	}
}

// Тест для DistanceMatrix
func TestGraph_DistanceMatrix(t *testing.T) {
	graph := models.Graph{
		Nodes: []models.Node{
			{Label: "A"},
			{Label: "B"},
			{Label: "C"},
		},
		Edges: []models.Edge{
			{Source: models.Node{Label: "A"}, Target: models.Node{Label: "B"}, Weight: 1},
			{Source: models.Node{Label: "B"}, Target: models.Node{Label: "C"}, Weight: 2},
		},
	}

	distMatrix := graph.DistanceMatrix(true)
	expectedMatrix := map[string]map[string]int{
		"A": {"A": 0, "B": 1, "C": 3},
		"B": {"A": 1, "B": 0, "C": 2},
		"C": {"A": 3, "B": 2, "C": 0},
	}

	for node1, row := range expectedMatrix {
		for node2, expectedValue := range row {
			if distMatrix[node1][node2] != expectedValue {
				t.Errorf("expected distMatrix[%s][%s] = %d, got %d", node1, node2, expectedValue, distMatrix[node1][node2])
			}
		}
	}
}

func TestCheckAdjacentMatrix(t *testing.T) {
	checker := task_check.NewChecker()
	task := models.Graph{
		Nodes: []models.Node{
			{Label: "A"},
			{Label: "B"},
			{Label: "C"},
		},
		Edges: []models.Edge{
			{Source: models.Node{Label: "A"}, Target: models.Node{Label: "B"}},
			{Source: models.Node{Label: "B"}, Target: models.Node{Label: "C"}},
		},
	}

	// Правильный ответ
	correctAnswer := map[string]map[string]int{
		"A": {"A": 0, "B": 1, "C": 0},
		"B": {"A": 1, "B": 0, "C": 1},
		"C": {"A": 0, "B": 1, "C": 0},
	}

	// Случай: ответ верный
	score := checker.CheckAdjacentMatrix(task, correctAnswer)
	if score != 100 {
		t.Errorf("Expected score 100 for correct answer, got %d", score)
	}

	// Случай: частичный ответ
	partialAnswer := map[string]map[string]int{
		"A": {"A": 0, "B": 1, "C": 0},
		"B": {"A": 1, "B": 0, "C": 0}, // Ошибка здесь: должна быть 1 вместо 0
		"C": {"A": 0, "B": 0, "C": 0}, // Ошибка здесь: должна быть 1 вместо 0
	}
	score = checker.CheckAdjacentMatrix(task, partialAnswer)
	if score >= 100 || score == 0 {
		t.Errorf("Expected partial score for partial answer, got %d", score)
	}

	// Случай: пустой ответ
	emptyAnswer := map[string]map[string]int{}
	score = checker.CheckAdjacentMatrix(task, emptyAnswer)
	if score != 0 {
		t.Errorf("Expected score 0 for empty answer, got %d", score)
	}

	// Случай: неправильный размер матрицы
	wrongSizeAnswer := map[string]map[string]int{
		"A": {"A": 0, "B": 1},
		"B": {"A": 1, "B": 0},
	}
	score = checker.CheckAdjacentMatrix(task, wrongSizeAnswer)
	if score != 0 {
		t.Errorf("Expected score 0 for wrong-sized matrix, got %d", score)
	}
}

// Тест для CheckRadiusAndDiameter
func TestChecker_CheckRadiusAndDiameter(t *testing.T) {
	graph := models.Graph{
		Nodes: []models.Node{
			{Label: "A"},
			{Label: "B"},
			{Label: "C"},
		},
		Edges: []models.Edge{
			{Source: models.Node{Label: "A"}, Target: models.Node{Label: "B"}, Weight: 1},
			{Source: models.Node{Label: "B"}, Target: models.Node{Label: "C"}, Weight: 1},
		},
	}

	checker := task_check.NewChecker()

	distMatrixAns := map[string]map[string]int{
		"A": {"A": 0, "B": 1, "C": 2},
		"B": {"A": 1, "B": 0, "C": 1},
		"C": {"A": 2, "B": 1, "C": 0},
	}

	radiusAns := 1
	diameterAns := 2

	score := checker.CheckRadiusAndDiameter(graph, radiusAns, diameterAns, distMatrixAns)
	if score == 0 {
		t.Errorf("Radius and Diameter check failed: expected non-zero score, got %d", score)
	}
}

// Тест для CheckMinPath
func TestChecker_CheckMinPath(t *testing.T) {
	graph := models.Graph{
		Nodes: []models.Node{
			{Label: "A"},
			{Label: "B"},
			{Label: "C"},
		},
		Edges: []models.Edge{
			{Source: models.Node{Label: "A"}, Target: models.Node{Label: "B"}, Weight: 1},
			{Source: models.Node{Label: "B"}, Target: models.Node{Label: "C"}, Weight: 2},
		},
	}

	checker := task_check.NewChecker()

	source := "A"
	target := "C"
	minPathAns := 3
	weightsPathAns := map[string]int{
		"A": 0,
		"B": 1,
		"C": 3,
	}

	answer := models.Graph{
		Edges: []models.Edge{
			{Source: models.Node{Label: "A"}, Target: models.Node{Label: "B"}, Weight: 1, Color: "red"},
			{Source: models.Node{Label: "B"}, Target: models.Node{Label: "C"}, Weight: 2, Color: "red"},
		},
	}

	score := checker.CheckMinPath(graph, source, target, minPathAns, weightsPathAns, answer)
	if score == 0 {
		t.Errorf("MinPath check failed: expected non-zero score, got %d", score)
	}
}

// Тест для CheckLinearToLine
func TestChecker_CheckLinearToLine(t *testing.T) {
	task := models.Graph{
		Nodes: []models.Node{
			{Label: "A"},
			{Label: "B"},
			{Label: "C"},
			{Label: "D"},
		},
		Edges: []models.Edge{
			{Source: models.Node{Label: "A"}, Target: models.Node{Label: "B"}, Label: "1"},
			{Source: models.Node{Label: "B"}, Target: models.Node{Label: "C"}, Label: "2"},
			{Source: models.Node{Label: "C"}, Target: models.Node{Label: "D"}, Label: "3"},
		},
	}

	answer := models.Graph{
		Nodes: []models.Node{
			{Label: "1"},
			{Label: "2"},
			{Label: "3"},
		},
		Edges: []models.Edge{
			{Source: models.Node{Label: "1"}, Target: models.Node{Label: "2"}},
			{Source: models.Node{Label: "2"}, Target: models.Node{Label: "3"}},
		},
	}

	checker := task_check.NewChecker()
	score := checker.CheckLinearToLine(&task, &answer)
	if score == 0 {
		t.Errorf("Linear to Line check failed: expected non-zero score, got %d", score)
	}
}

// Тест для CheckLinearFromLine
func TestChecker_CheckLinearFromLine(t *testing.T) {
	task := models.Graph{
		Nodes: []models.Node{
			{Label: "1"},
			{Label: "2"},
			{Label: "3"},
		},
		Edges: []models.Edge{
			{Source: models.Node{Label: "1"}, Target: models.Node{Label: "2"}, Label: "a"},
			{Source: models.Node{Label: "2"}, Target: models.Node{Label: "3"}, Label: "b"},
		},
	}

	answer := models.Graph{
		Nodes: []models.Node{
			{Label: "A"},
			{Label: "B"},
			{Label: "C"},
			{Label: "D"},
		},
		Edges: []models.Edge{
			{Source: models.Node{Label: "A"}, Target: models.Node{Label: "B"}, Label: "1"},
			{Source: models.Node{Label: "B"}, Target: models.Node{Label: "C"}, Label: "2"},
			{Source: models.Node{Label: "C"}, Target: models.Node{Label: "D"}, Label: "3"},
		},
	}

	checker := task_check.NewChecker()
	score := checker.CheckLinearFromLine(task, answer)
	if score == 0 {
		t.Errorf("Linear from Line check failed: expected non-zero score, got %d", score)
	}
}

// Тест для max_
func TestMax_(t *testing.T) {
	if result := task_check.Max_(5, 10); result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
	if result := task_check.Max_(10, 5); result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
	if result := task_check.Max_(7, 7); result != 7 {
		t.Errorf("Expected 7, got %d", result)
	}
}
