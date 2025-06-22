package taskcheck

// import (
// 	// "fmt"
// 	// usecase "golang_graphs/backend/internal/usecase/student"
// 	"log"

// 	//
// 	// usecase "golang_graphs/backend/internal/usecase/student"

// 	// taskcheck "golang_graphs/backend/internal/domain/student/service/taskcheck"

// 	"encoding/json"
// 	"testing"
// )

// var (
// 	DEFAULT_COLOR = ""
// )



// // Тест для MinPath
// func TestGraph_MinPath(t *testing.T) {
// 	graph := model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"},
// 			{Id: 2, Label: "B"},
// 			{Id: 3, Label: "C"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Id: 1, Label: "A"}, Target: model.Node{Id: 2, Label: "B"}, Weight: 2},
// 			{Source: model.Node{Id: 2, Label: "B"}, Target: model.Node{Id: 3, Label: "C"}, Weight: 3},
// 			{Source: model.Node{Id: 1, Label: "A"}, Target: model.Node{Id: 3, Label: "C"}, Weight: 5},
// 		},
// 	}

// 	// Positive case
// 	minPath, weightsPath := graph.MinPath(model.Node{Id: 1, Label: "A"}, model.Node{Id: 3, Label: "C"}, true)
// 	expectedMinPath := 5
// 	expectedWeights := map[string]int{"A": 0, "B": 2, "C": 5}

// 	if minPath != expectedMinPath {
// 		t.Errorf("expected %d, got %d", expectedMinPath, minPath)
// 	}

// 	for key, value := range expectedWeights {
// 		if weightsPath[key] != value {
// 			t.Errorf("expected weights[%s] = %d, got %d", key, value, weightsPath[key])
// 		}
// 	}

// 	// Negative case (unreachable nodes)
// 	graph2 := model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"},
// 			{Id: 2, Label: "B"},
// 		},
// 		Edges: []model.Edge{},
// 	}

// 	minPath2, _ := graph2.MinPath(model.Node{Id: 1, Label: "A"}, model.Node{Id: 2, Label: "B"}, true)
// 	if minPath2 != math.MaxInt {
// 		t.Errorf("expected unreachable path, got %d", minPath2)
// 	}
// }

// // Тест для NodeLabelAdjacentMatrix
// func TestGraph_NodeLabelAdjacentMatrix(t *testing.T) {
// 	graph := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "A"},
// 			{Label: "B"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}},
// 		},
// 	}

// 	adjMatrix := graph.NodeLabelAdjacentMatrix()
// 	expectedMatrix := map[string]map[string]int{
// 		"A": {"A": 0, "B": 1},
// 		"B": {"A": 1, "B": 0},
// 	}

// 	for node1, row := range expectedMatrix {
// 		for node2, expectedValue := range row {
// 			if adjMatrix[node1][node2] != expectedValue {
// 				t.Errorf("expected adjMatrix[%s][%s] = %d, got %d", node1, node2, expectedValue, adjMatrix[node1][node2])
// 			}
// 		}
// 	}
// }

// // Тест для NodeIdAdjacentMatrix
// func TestGraph_NodeIdAdjacentMatrix(t *testing.T) {
// 	graph := model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1},
// 			{Id: 2},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Id: 1}, Target: model.Node{Id: 2}},
// 		},
// 	}
// 	adjMatrix := graph.NodeIdAdjacentMatrix()
// 	expectedMatrix := map[int]map[int]int{
// 		1: {1: 0, 2: 1},
// 		2: {1: 1, 2: 0},
// 	}

// 	for id1, row := range expectedMatrix {
// 		for id2, expectedValue := range row {
// 			if adjMatrix[id1][id2] != expectedValue {
// 				t.Errorf("expected adjMatrix[%d][%d] = %d, got %d", id1, id2, expectedValue, adjMatrix[id1][id2])
// 			}
// 		}
// 	}
// }

// // Тест для DistanceMatrix
// func TestGraph_DistanceMatrix(t *testing.T) {
// 	graph := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "A"},
// 			{Label: "B"},
// 			{Label: "C"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}, Weight: 1},
// 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}, Weight: 2},
// 		},
// 	}

// 	distMatrix := graph.DistanceMatrix(true)
// 	expectedMatrix := map[string]map[string]int{
// 		"A": {"A": 0, "B": 1, "C": 3},
// 		"B": {"A": 1, "B": 0, "C": 2},
// 		"C": {"A": 3, "B": 2, "C": 0},
// 	}

// 	for node1, row := range expectedMatrix {
// 		for node2, expectedValue := range row {
// 			if distMatrix[node1][node2] != expectedValue {
// 				t.Errorf("expected distMatrix[%s][%s] = %d, got %d", node1, node2, expectedValue, distMatrix[node1][node2])
// 			}
// 		}
// 	}
// }

// func TestCheckAdjacentMatrix(t *testing.T) {
// 	checker := NewChecker()
// 	task := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "A"},
// 			{Label: "B"},
// 			{Label: "C"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}},
// 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}},
// 		},
// 	}

// 	// Правильный ответ
// 	correctAnswer := map[string]map[string]int{
// 		"A": {"A": 0, "B": 1, "C": 0},
// 		"B": {"A": 1, "B": 0, "C": 1},
// 		"C": {"A": 0, "B": 1, "C": 0},
// 	}

// 	// Случай: ответ верный
// 	score := checker.CheckAdjacentMatrix(&task, correctAnswer)
// 	if score != 100 {
// 		t.Errorf("Expected score 100 for correct answer, got %d", score)
// 	}

// 	// Случай: частичный ответ
// 	partialAnswer := map[string]map[string]int{
// 		"A": {"A": 0, "B": 1, "C": 0},
// 		"B": {"A": 1, "B": 0, "C": 0}, // Ошибка здесь: должна быть 1 вместо 0
// 		"C": {"A": 0, "B": 0, "C": 0}, // Ошибка здесь: должна быть 1 вместо 0
// 	}
// 	score = checker.CheckAdjacentMatrix(&task, partialAnswer)
// 	if score >= 100 || score == 0 {
// 		t.Errorf("Expected partial score for partial answer, got %d", score)
// 	}

// 	// Случай: пустой ответ
// 	emptyAnswer := map[string]map[string]int{}
// 	score = checker.CheckAdjacentMatrix(&task, emptyAnswer)
// 	if score != 0 {
// 		t.Errorf("Expected score 0 for empty answer, got %d", score)
// 	}

// 	// Случай: неправильный размер матрицы
// 	wrongSizeAnswer := map[string]map[string]int{
// 		"A": {"A": 0, "B": 1},
// 		"B": {"A": 1, "B": 0},
// 	}
// 	score = checker.CheckAdjacentMatrix(&task, wrongSizeAnswer)
// 	if score != 0 {
// 		t.Errorf("Expected score 0 for wrong-sized matrix, got %d", score)
// 	}
// }

// // Тест для CheckRadiusAndDiameter
// func TestChecker_CheckRadiusAndDiameter(t *testing.T) {
// 	graph := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "A"},
// 			{Label: "B"},
// 			{Label: "C"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}, Weight: 1},
// 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}, Weight: 1},
// 		},
// 	}

// 	checker := NewChecker()

// 	distMatrixAns := map[string]map[string]int{
// 		"A": {"A": 0, "B": 1, "C": 2},
// 		"B": {"A": 1, "B": 0, "C": 1},
// 		"C": {"A": 2, "B": 1, "C": 0},
// 	}

// 	radiusAns := 1
// 	diameterAns := 2

// 	score := checker.CheckRadiusAndDiameter(&graph, radiusAns, diameterAns, distMatrixAns)
// 	if score == 0 {
// 		t.Errorf("Radius and Diameter check failed: expected non-zero score, got %d", score)
// 	}
// }

// // Тест для CheckMinPath
// func TestChecker_CheckMinPath(t *testing.T) {
// 	graph := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "A"},
// 			{Label: "B"},
// 			{Label: "C"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}, Weight: 1},
// 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}, Weight: 2},
// 		},
// 	}

// 	checker := NewChecker()

// 	source := "A"
// 	target := "C"
// 	minPathAns := 3
// 	weightsPathAns := map[string]int{
// 		"A": 0,
// 		"B": 1,
// 		"C": 3,
// 	}

// 	answer := model.Graph{
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}, Weight: 1, Color: "red"},
// 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}, Weight: 2, Color: "red"},
// 		},
// 	}

// 	score := checker.CheckMinPath(&graph, source, target, minPathAns, weightsPathAns, &answer)
// 	if score == 0 {
// 		t.Errorf("MinPath check failed: expected non-zero score, got %d", score)
// 	}
// }

// // Тест для CheckLinearToLine
// func TestChecker_CheckLinearToLine(t *testing.T) {
// 	task := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "A"},
// 			{Label: "B"},
// 			{Label: "C"},
// 			{Label: "D"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}, Label: "1"},
// 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}, Label: "2"},
// 			{Source: model.Node{Label: "C"}, Target: model.Node{Label: "D"}, Label: "3"},
// 		},
// 	}

// 	answer := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "1"},
// 			{Label: "2"},
// 			{Label: "3"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "1"}, Target: model.Node{Label: "2"}},
// 			{Source: model.Node{Label: "2"}, Target: model.Node{Label: "3"}},
// 		},
// 	}

// 	checker := NewChecker()
// 	score := checker.CheckLinearToLine(&task, &answer)
// 	fmt.Println(score)
// 	if score == 0 {
// 		t.Errorf("Linear to Line check failed: expected non-zero score, got %d", score)
// 	}
// }

// // Тест для CheckLinearFromLine
// func TestChecker_CheckLinearFromLine(t *testing.T) {
// 	task := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "1"},
// 			{Label: "2"},
// 			{Label: "3"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "1"}, Target: model.Node{Label: "2"}, Label: "a"},
// 			{Source: model.Node{Label: "2"}, Target: model.Node{Label: "3"}, Label: "b"},
// 		},
// 	}

// 	answer := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "A"},
// 			{Label: "B"},
// 			{Label: "C"},
// 			{Label: "D"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}, Label: "1"},
// 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}, Label: "2"},
// 			{Source: model.Node{Label: "C"}, Target: model.Node{Label: "D"}, Label: "3"},
// 		},
// 	}

// 	checker := NewChecker()
// 	score := checker.CheckLinearFromLine(&task, &answer)
// 	if score != 100 {
// 		t.Errorf("Linear from Line check failed: expected non-zero score, got %d", score)
// 	}
// }

// // Тест для max_
// func TestMax_(t *testing.T) {
// 	if result := Max_(5, 10); result != 10 {
// 		t.Errorf("Expected 10, got %d", result)
// 	}
// 	if result := Max_(10, 5); result != 10 {
// 		t.Errorf("Expected 10, got %d", result)
// 	}
// 	if result := Max_(7, 7); result != 7 {
// 		t.Errorf("Expected 7, got %d", result)
// 	}
// }

// func createGraphA() *model.Graph {
// 	g := new(model.Graph)
// 	a := model.Node{Id: 0, Label: "A", Color: "", Weight: 0, X: 0, Y: 0}
// 	b := model.Node{Id: 1, Label: "B", Color: "", Weight: 0, X: 0, Y: 0}
// 	c := model.Node{Id: 2, Label: "C", Color: "", Weight: 0, X: 0, Y: 0}
// 	d := model.Node{Id: 3, Label: "D", Color: "", Weight: 0, X: 0, Y: 0}
// 	e := model.Node{Id: 4, Label: "E", Color: "", Weight: 0, X: 0, Y: 0}
// 	g.AddNode(a)
// 	g.AddNode(b)
// 	g.AddNode(c)
// 	g.AddNode(d)
// 	g.AddNode(e)

// 	g.AddEdgeByInfo(a, b, "e1", "", "", 0)
// 	g.AddEdgeByInfo(b, c, "e2", "", "", 0)
// 	g.AddEdgeByInfo(c, d, "e3", "", "", 0)
// 	g.AddEdgeByInfo(a, d, "e4", "", "", 0)
// 	g.AddEdgeByInfo(d, e, "e5", "", "", 0)

// 	return g
// }

// func createGraphB() *model.Graph {
// 	g := new(model.Graph)
// 	c := model.Node{Id: 0, Label: "C", Color: "", Weight: 0, X: 0, Y: 0}
// 	d := model.Node{Id: 1, Label: "D", Color: "", Weight: 0, X: 0, Y: 0}
// 	e := model.Node{Id: 2, Label: "E", Color: "", Weight: 0, X: 0, Y: 0}
// 	f := model.Node{Id: 3, Label: "F", Color: "", Weight: 0, X: 0, Y: 0}
// 	g1 := model.Node{Id: 4, Label: "G", Color: "", Weight: 0, X: 0, Y: 0}
// 	g.AddNode(c)
// 	g.AddNode(d)
// 	g.AddNode(e)
// 	g.AddNode(f)
// 	g.AddNode(g1)

// 	g.AddEdgeByInfo(c, d, "e6", "", "", 0)
// 	g.AddEdgeByInfo(d, e, "e7", "", "", 0)
// 	g.AddEdgeByInfo(e, f, "e8", "", "", 0)
// 	g.AddEdgeByInfo(f, g1, "e9", "", "", 0)
// 	g.AddEdgeByInfo(c, f, "e10", "", "", 0)

// 	return g
// }

// func nodeExists(g *model.Graph, label string) bool {
// 	for _, node := range g.Nodes {
// 		if node.Label == label {
// 			return true
// 		}
// 	}
// 	return false
// }

// func edgeExists(g *model.Graph, srcLabel, trgLabel string) bool {
// 	for _, edge := range g.Edges {
// 		if (edge.Source.Label == srcLabel && edge.Target.Label == trgLabel) ||
// 			(edge.Source.Label == trgLabel && edge.Target.Label == srcLabel) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func TestIntersect(t *testing.T) {
// 	g1 := createGraphA()
// 	g2 := createGraphB()
// 	result := g1.Intersect(g2)
// 	expectedNodes := []string{"C", "D", "E"}
// 	for _, label := range expectedNodes {
// 		if !result.IsNodeByLabel(label) {
// 			t.Errorf("Expected node %s in intersection", label)
// 		}
// 	}

// 	if found, _ := result.FindEdge("C", "D"); !found {
// 		t.Errorf("Expected edge C-D in intersection")
// 	}
// 	if found, _ := result.FindEdge("D", "E"); !found {
// 		t.Errorf("Expected edge D-E in intersection")
// 	}
// }

// func TestUnion(t *testing.T) {
// 	g1 := createGraphA()
// 	g2 := createGraphB()
// 	result := g1.Union(g2)
// 	expectedNodes := []string{"A", "B", "C", "D", "E", "F", "G"}
// 	for _, label := range expectedNodes {
// 		if !result.IsNodeByLabel(label) {
// 			t.Errorf("Expected node %s in union", label)
// 		}
// 	}

// 	expectedEdges := [][2]string{
// 		{"A", "B"}, {"B", "C"}, {"C", "D"}, {"A", "D"}, {"D", "E"},
// 		{"E", "F"}, {"F", "G"}, {"C", "F"},
// 	}
// 	for _, edge := range expectedEdges {
// 		if found, _ := result.FindEdge(edge[0], edge[1]); !found {
// 			t.Errorf("Expected edge %s-%s in union", edge[0], edge[1])
// 		}
// 	}
// }

// func TestJoin(t *testing.T) {
// 	g1 := createGraphA()
// 	g2 := createGraphB()
// 	result := g1.Join(g2)
// 	af, _ := result.FindEdge("A", "F")
// 	ag, _ := result.FindEdge("A", "G")
// 	if !af && !ag {
// 		t.Errorf("Expected join edge A-F or A-G")
// 	}
// 	bf, _ := result.FindEdge("A", "F")
// 	bg, _ := result.FindEdge("B", "G")
// 	if !bf && !bg {
// 		t.Errorf("Expected join edge B-F or B-G")
// 	}
// }

// func TestCartesianProduct(t *testing.T) {
// 	g1 := createGraphA()
// 	g2 := createGraphB()
// 	result := g1.CartesianProduct(g2)
// 	if !result.IsNodeByLabel("(A,C)") {
// 		t.Errorf("Expected node (A,C) in cartesian product")
// 	}

// 	if found, _ := result.FindEdge("(A,C)", "(B,C)"); !found {
// 		t.Errorf("Expected edge (A,C)-(B,C) in cartesian product")
// 	}
// 	if found, _ := result.FindEdge("(B,C)", "(B,D)"); !found {
// 		t.Errorf("Expected edge (B,C)-(B,D) in cartesian product")
// 	}
// 	if found, _ := result.FindEdge("(C,D)", "(C,E)"); !found {
// 		t.Errorf("Expected edge (C,D)-(C,E) in cartesian product")
// 	}
// 	if found, _ := result.FindEdge("(E,E)", "(E,F)"); !found {
// 		t.Errorf("Expected edge (E,E)-(E,E) in cartesian product")
// 	}
// 	if found, _ := result.FindEdge("(E,E)", "(D,E)"); !found {
// 		t.Errorf("Expected edge (E,E)-(D,E) in cartesian product")
// 	}
// }

// func TestTensorProduct(t *testing.T) {
// 	g1 := createGraphA()
// 	g2 := createGraphB()
// 	result := g1.TensorProduct(g2)

// 	if !result.IsNodeByLabel("(C,C)") {
// 		t.Errorf("Expected node (C,C) in tensor product")
// 	}

// 	if found, _ := result.FindEdge("(C,C)", "(D,D)"); !found {
// 		t.Errorf("Expected edge (C,C)-(D,D) in tensor product")
// 	}
// 	if found, _ := result.FindEdge("(A,C)", "(B,D)"); !found {
// 		t.Errorf("Expected edge (A,C)-(C,D) in tensor product")
// 	}
// 	if found, _ := result.FindEdge("(D,E)", "(E,F)"); !found {
// 		t.Errorf("Expected edge (D,E)-(E,F) in tensor product")
// 	}
// 	if found, _ := result.FindEdge("(D,C)", "(C,D)"); !found {
// 		t.Errorf("Expected edge (D,C)-(C,D) in tensor product")
// 	}
// 	if found, _ := result.FindEdge("(E,G)", "(D,F)"); !found {
// 		t.Errorf("Expected edge (E,G)-(D,F) in tensor product")
// 	}
// }

// func TestLexicographicalProduct(t *testing.T) {
// 	g1 := createGraphA()
// 	g2 := createGraphB()
// 	result := g1.LexicographicalProduct(g2)

// 	if !result.IsNodeByLabel("(A,F)") {
// 		t.Errorf("Expected node (A,F) in lex product")
// 	}

// 	if found, _ := result.FindEdge("(A,F)", "(B,F)"); !found {
// 		t.Errorf("Expected edge (A,F)-(B,F) in lexicographical product")
// 	}
// 	if found, _ := result.FindEdge("(B,C)", "(C,G)"); !found {
// 		t.Errorf("Expected edge (B,C)-(C,G) in lexicographical product")
// 	}
// 	if found, _ := result.FindEdge("(A,E)", "(A,F)"); !found {
// 		t.Errorf("Expected edge (A,E)-(A,F) in lexicographical product")
// 	}
// 	if found, _ := result.FindEdge("(D,D)", "(D,C)"); !found {
// 		t.Errorf("Expected edge (D,D)-(D,C) in lexicographical product")
// 	}
// 	if found, _ := result.FindEdge("(C,C)", "(D,E)"); !found {
// 		t.Errorf("Expected edge (C,C)-(D,E) in lexicographical product")
// 	}
// }

// func createTestGraph1() *model.Graph {
// 	g := new(model.Graph)
// 	n1 := model.Node{Id: 1, Label: "A"}
// 	n2 := model.Node{Id: 2, Label: "B"}
// 	n3 := model.Node{Id: 3, Label: "C"}
// 	g.AddNode(n1)
// 	g.AddNode(n2)
// 	g.AddNode(n3)
// 	g.AddEdgeByInfo(n1, n2, "1", "", "", 0)
// 	g.AddEdgeByInfo(n2, n3, "2", "", "", 0)
// 	return g
// }

// func createTestGraph2() *model.Graph {
// 	g := new(model.Graph)
// 	n1 := model.Node{Id: 4, Label: "B"}
// 	n2 := model.Node{Id: 5, Label: "C"}
// 	n3 := model.Node{Id: 6, Label: "D"}
// 	g.AddNode(n1)
// 	g.AddNode(n2)
// 	g.AddNode(n3)
// 	g.AddEdgeByInfo(n1, n2, "3", "", "", 0)
// 	g.AddEdgeByInfo(n2, n3, "4", "", "", 0)
// 	return g
// }

// func TestCheckIntersectionGraphs(t *testing.T) {
// 	g1 := createTestGraph1()
// 	g2 := createTestGraph2()
// 	expected := model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "B"}, {Id: 2, Label: "C"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Id: 1, Label: "B"}, Target: model.Node{Id: 2, Label: "C"}},
// 		},
// 	}

// 	ch := NewChecker()
// 	score := ch.CheckIntersectionGraphs(&expected, g1, g2)

// 	if score != 100 {
// 		t.Errorf("Expected 100, got %d", score)
// 	}
// }

// func TestCheckUnionGraphs(t *testing.T) {
// 	g1 := createTestGraph1()
// 	g2 := createTestGraph2()
// 	expected := model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"}, {Id: 2, Label: "B"},
// 			{Id: 3, Label: "C"}, {Id: 4, Label: "D"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "B", Id: 2}, Target: model.Node{Id: 3, Label: "C"}, Id: "1"},
// 			{Source: model.Node{Label: "A", Id: 1}, Target: model.Node{Id: 2, Label: "B"}, Id: "2"},
// 			{Source: model.Node{Label: "C", Id: 3}, Target: model.Node{Id: 4, Label: "D"}, Id: "3"},
// 		},
// 	}

// 	ch := NewChecker()
// 	score := ch.CheckUnionGraphs(&expected, g1, g2)

// 	if score != 100 {
// 		t.Errorf("Expected 100, got %d", score)
// 	}
// }

// func TestCheckJoinGraphs(t *testing.T) {
// 	g1 := createTestGraph1()
// 	g2 := createTestGraph2()
// 	expected := model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"}, {Id: 2, Label: "B"},
// 			{Id: 3, Label: "C"}, {Id: 4, Label: "D"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "B", Id: 2}, Target: model.Node{Id: 3, Label: "C"}, Id: "1"},
// 			{Source: model.Node{Label: "A", Id: 1}, Target: model.Node{Id: 2, Label: "B"}, Id: "2"},
// 			{Source: model.Node{Label: "C", Id: 3}, Target: model.Node{Id: 4, Label: "D"}, Id: "3"},
// 			{Source: model.Node{Label: "A", Id: 1}, Target: model.Node{Id: 4, Label: "D"}, Id: "4"},
// 		},
// 	}
// 	ch := NewChecker()
// 	score := ch.CheckJoinGraphs(&expected, g1, g2)

// 	if score != 100 {
// 		t.Errorf("Expected 100, got %d", score)
// 	}
// }

// func TestCheckCartesianProduct(t *testing.T) {
// 	g1 := createTestGraph1()
// 	g2 := createTestGraph2()
// 	nodes := []model.Node{
// 		{Id: 0, Label: "(A,B)"}, {Id: 1, Label: "(A,C)"}, {Id: 2, Label: "(A,D)"},
// 		{Id: 3, Label: "(B,B)"}, {Id: 4, Label: "(B,C)"}, {Id: 5, Label: "(B,D)"},
// 		{Id: 6, Label: "(C,B)"}, {Id: 7, Label: "(C,C)"}, {Id: 8, Label: "(C,D)"},
// 	}
// 	expected := model.Graph{
// 		Nodes: nodes,
// 		Edges: []model.Edge{
// 			{Source: nodes[0], Target: nodes[1], Id: "1"}, {Source: nodes[1], Target: nodes[2], Id: "2"},
// 			{Source: nodes[3], Target: nodes[4], Id: "3"}, {Source: nodes[4], Target: nodes[5], Id: "4"},
// 			{Source: nodes[6], Target: nodes[7], Id: "5"}, {Source: nodes[7], Target: nodes[8], Id: "6"},
// 			{Source: nodes[0], Target: nodes[3], Id: "7"}, {Source: nodes[3], Target: nodes[6], Id: "8"},
// 			{Source: nodes[1], Target: nodes[4], Id: "9"}, {Source: nodes[4], Target: nodes[7], Id: "10"},
// 			{Source: nodes[2], Target: nodes[5], Id: "11"}, {Source: nodes[5], Target: nodes[8], Id: "12"},
// 		},
// 	}

// 	ch := NewChecker()
// 	score := ch.CheckCartesianProduct(&expected, g1, g2)

// 	if score != 100 {
// 		t.Errorf("Expected 100, got %d", score)
// 	}
// }

// func TestCheckTensorProduct(t *testing.T) {
// 	g1 := createTestGraph1()
// 	g2 := createTestGraph2()
// 	nodes := []model.Node{
// 		{Id: 0, Label: "(A,B)"}, {Id: 1, Label: "(A,C)"}, {Id: 2, Label: "(A,D)"},
// 		{Id: 3, Label: "(B,B)"}, {Id: 4, Label: "(B,C)"}, {Id: 5, Label: "(B,D)"},
// 		{Id: 6, Label: "(C,B)"}, {Id: 7, Label: "(C,C)"}, {Id: 8, Label: "(C,D)"},
// 	}
// 	expected := model.Graph{
// 		Nodes: nodes,
// 		Edges: []model.Edge{
// 			{Source: nodes[0], Target: nodes[4], Id: "1"}, {Source: nodes[2], Target: nodes[4], Id: "2"},
// 			{Source: nodes[6], Target: nodes[4], Id: "3"}, {Source: nodes[8], Target: nodes[4], Id: "4"},
// 			{Source: nodes[1], Target: nodes[3], Id: "5"}, {Source: nodes[1], Target: nodes[5], Id: "6"},
// 			{Source: nodes[7], Target: nodes[3], Id: "7"}, {Source: nodes[7], Target: nodes[5], Id: "8"},
// 		},
// 	}
// 	ch := NewChecker()
// 	score := ch.CheckTensorProduct(&expected, g1, g2)

// 	if score != 100 {
// 		t.Errorf("Expected 100, got %d", score)
// 	}
// }

// func TestCheckLexicographicalProduct(t *testing.T) {
// 	g1 := createTestGraph1()
// 	g2 := createTestGraph2()
// 	nodes := []model.Node{
// 		{Id: 0, Label: "(A,B)"}, {Id: 1, Label: "(A,C)"}, {Id: 2, Label: "(A,D)"},
// 		{Id: 3, Label: "(B,B)"}, {Id: 4, Label: "(B,C)"}, {Id: 5, Label: "(B,D)"},
// 		{Id: 6, Label: "(C,B)"}, {Id: 7, Label: "(C,C)"}, {Id: 8, Label: "(C,D)"},
// 	}
// 	expected := model.Graph{
// 		Nodes: nodes,
// 		Edges: []model.Edge{
// 			{Source: nodes[0], Target: nodes[1], Id: "1"}, {Source: nodes[1], Target: nodes[2], Id: "2"},
// 			{Source: nodes[3], Target: nodes[4], Id: "3"}, {Source: nodes[4], Target: nodes[5], Id: "4"},
// 			{Source: nodes[6], Target: nodes[7], Id: "5"}, {Source: nodes[7], Target: nodes[8], Id: "6"},
// 			{Source: nodes[0], Target: nodes[3], Id: "7"}, {Source: nodes[3], Target: nodes[6], Id: "8"},
// 			{Source: nodes[1], Target: nodes[4], Id: "9"}, {Source: nodes[4], Target: nodes[7], Id: "10"},
// 			{Source: nodes[2], Target: nodes[5], Id: "11"}, {Source: nodes[5], Target: nodes[8], Id: "12"},

// 			{Source: nodes[0], Target: nodes[4], Id: "13"}, {Source: nodes[0], Target: nodes[5], Id: "14"},
// 			{Source: nodes[1], Target: nodes[3], Id: "15"}, {Source: nodes[1], Target: nodes[5], Id: "16"},
// 			{Source: nodes[2], Target: nodes[3], Id: "17"}, {Source: nodes[2], Target: nodes[4], Id: "18"},
// 			{Source: nodes[3], Target: nodes[7], Id: "19"}, {Source: nodes[3], Target: nodes[8], Id: "20"},
// 			{Source: nodes[4], Target: nodes[6], Id: "21"}, {Source: nodes[4], Target: nodes[8], Id: "22"},
// 			{Source: nodes[5], Target: nodes[6], Id: "23"}, {Source: nodes[5], Target: nodes[7], Id: "24"},
// 		},
// 	}
// 	ch := NewChecker()
// 	score := ch.CheckLexicographicalProduct(&expected, g1, g2)

// 	if score != 100 {
// 		t.Errorf("Expected 100, got %d", score)
// 	}
// }

// func TestHamiltonianCycle(t *testing.T) {
// 	graph := model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"}, {Id: 2, Label: "B"},
// 			{Id: 3, Label: "C"}, {Id: 4, Label: "D"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "B", Id: 2}, Target: model.Node{Id: 3, Label: "C"}, Id: "1"},
// 			{Source: model.Node{Label: "A", Id: 1}, Target: model.Node{Id: 2, Label: "B"}, Id: "2"},
// 			{Source: model.Node{Label: "C", Id: 3}, Target: model.Node{Id: 4, Label: "D"}, Id: "3"},
// 			{Source: model.Node{Label: "A", Id: 1}, Target: model.Node{Id: 4, Label: "D"}, Id: "4"},
// 		},
// 	}
// 	ans_bool, _ := graph.HamiltonianCycle()
// 	if ans_bool == false {
// 		t.Errorf("Expected true answer")
// 	}
// }

// func TestCheckHamiltonianCycle(t *testing.T) {
// 	graph := model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"}, {Id: 2, Label: "B"},
// 			{Id: 3, Label: "C"}, {Id: 4, Label: "D"},
// 			{Id: 5, Label: "F"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "B", Id: 2}, Target: model.Node{Id: 3, Label: "C"}, Id: "1", Color: "r"},
// 			{Source: model.Node{Label: "A", Id: 1}, Target: model.Node{Id: 5, Label: "F"}, Id: "2", Color: "r"},
// 			{Source: model.Node{Label: "B", Id: 2}, Target: model.Node{Id: 5, Label: "F"}, Id: "5", Color: "r"},
// 			{Source: model.Node{Label: "C", Id: 3}, Target: model.Node{Id: 4, Label: "D"}, Id: "3", Color: "r"},
// 			{Source: model.Node{Label: "A", Id: 1}, Target: model.Node{Id: 4, Label: "D"}, Id: "4", Color: "r"},
// 			{Source: model.Node{Label: "D", Id: 4}, Target: model.Node{Id: 5, Label: "F"}, Id: "6"},
// 			{Source: model.Node{Label: "C", Id: 3}, Target: model.Node{Id: 5, Label: "F"}, Id: "7"},
// 		},
// 	}
// 	ch := NewChecker()
// 	score := ch.CheckHamiltonian(&graph, true, &graph)
// 	if score != 100 {
// 		t.Error("expected score 100")
// 	}
// }

// func TestMinimalSpanningTree(t *testing.T) {
// 	graph := model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"}, {Id: 2, Label: "B"},
// 			{Id: 3, Label: "C"}, {Id: 4, Label: "D"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "B", Id: 2}, Target: model.Node{Id: 3, Label: "C"}, Id: "1", Weight: 10},
// 			{Source: model.Node{Label: "A", Id: 1}, Target: model.Node{Id: 2, Label: "B"}, Id: "2", Weight: 5},
// 			{Source: model.Node{Label: "C", Id: 3}, Target: model.Node{Id: 4, Label: "D"}, Id: "3", Weight: 3},
// 			{Source: model.Node{Label: "A", Id: 1}, Target: model.Node{Id: 4, Label: "D"}, Id: "4", Weight: 7},
// 		},
// 	}
// 	expected := map[model.Edge]struct{}{
// 		{Source: model.Node{Label: "A", Id: 1}, Target: model.Node{Id: 2, Label: "B"}, Id: "2", Weight: 5}: {},
// 		{Source: model.Node{Label: "C", Id: 3}, Target: model.Node{Id: 4, Label: "D"}, Id: "3", Weight: 3}: {},
// 		{Source: model.Node{Label: "A", Id: 1}, Target: model.Node{Id: 4, Label: "D"}, Id: "4", Weight: 7}: {},
// 	}
// 	mst, ans := graph.MinimalSpanningTree()
// 	if ans != 15 {
// 		t.Errorf("Expected answer %d", ans)
// 	}
// 	for _, edge := range mst {
// 		if _, ok := expected[edge]; !ok {
// 			t.Errorf("Expected edge %s - %s", edge.Source.Label, edge.Target.Label)
// 		}
// 	}
// }

// // ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// // ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// // ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// func TestChecker_CheckLinearToLine1(t *testing.T) {
// 	checker := NewChecker()
// 	task := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "A"}, {Label: "B"}, {Label: "C"},
// 			{Label: "D"}, {Label: "E"}, {Label: "F"}, {Label: "G"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}, Label: "1"},
// 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}, Label: "2"},
// 			{Source: model.Node{Label: "C"}, Target: model.Node{Label: "D"}, Label: "3"},
// 			{Source: model.Node{Label: "D"}, Target: model.Node{Label: "E"}, Label: "4"},
// 			{Source: model.Node{Label: "E"}, Target: model.Node{Label: "F"}, Label: "5"},
// 			{Source: model.Node{Label: "F"}, Target: model.Node{Label: "G"}, Label: "6"},
// 		},
// 	}
// 	// Правильный ответ
// 	answer := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "1"}, {Label: "2"}, {Label: "3"},
// 			{Label: "4"}, {Label: "5"}, {Label: "6"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "1"}, Target: model.Node{Label: "2"}},
// 			{Source: model.Node{Label: "2"}, Target: model.Node{Label: "3"}},
// 			{Source: model.Node{Label: "3"}, Target: model.Node{Label: "4"}},
// 			{Source: model.Node{Label: "4"}, Target: model.Node{Label: "5"}},
// 			{Source: model.Node{Label: "5"}, Target: model.Node{Label: "6"}},
// 		},
// 	}

// 	score := checker.CheckLinearToLine(&task, &answer)
// 	if score != 100 {
// 		t.Errorf("CheckLinearToLine failed: expected non-zero score, got %d", score)
// 	}

// 	// Ответ с неправильным ребром
// 	answerWrongEdge := model.Graph{
// 		Nodes: answer.Nodes,
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "1"}, Target: model.Node{Label: "3"}}, // неверное ребро
// 			{Source: model.Node{Label: "1"}, Target: model.Node{Label: "2"}},
// 			{Source: model.Node{Label: "2"}, Target: model.Node{Label: "3"}},
// 			{Source: model.Node{Label: "3"}, Target: model.Node{Label: "4"}},
// 			{Source: model.Node{Label: "4"}, Target: model.Node{Label: "5"}},
// 			{Source: model.Node{Label: "5"}, Target: model.Node{Label: "6"}},
// 		},
// 	}
// 	score = checker.CheckLinearToLine(&task, &answerWrongEdge)
// 	if score != 60 {
// 		t.Errorf("CheckLinearToLine failed: expected score < 100 for incorrect edges, got %d", score)
// 	}

// 	// Краевой случай: графы разного размера
// 	answerSmall := model.Graph{
// 		Nodes: []model.Node{{Label: "1"}, {Label: "2"}},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "1"}, Target: model.Node{Label: "2"}},
// 		},
// 	}
// 	score = checker.CheckLinearToLine(&task, &answerSmall)
// 	if score != 0 {
// 		t.Errorf("CheckLinearToLine failed: expected score 0 for different sized graphs, got %d", score)
// 	}

// 	// Пустые графы
// 	score = checker.CheckLinearToLine(&model.Graph{}, &model.Graph{})
// 	if score != 0 {
// 		t.Errorf("CheckLinearToLine failed: expected 100 for empty graphs, got %d", score)
// 	}
// }

// func TestChecker_CheckLinearFromLine1(t *testing.T) {
// 	checker := NewChecker()
// 	task := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "A"}, {Label: "B"}, {Label: "C"},
// 			{Label: "D"}, {Label: "E"}, {Label: "F"}, {Label: "G"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}, Label: "1"},
// 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}, Label: "2"},
// 			{Source: model.Node{Label: "C"}, Target: model.Node{Label: "D"}, Label: "3"},
// 			{Source: model.Node{Label: "D"}, Target: model.Node{Label: "E"}, Label: "4"},
// 			{Source: model.Node{Label: "E"}, Target: model.Node{Label: "F"}, Label: "5"},
// 			{Source: model.Node{Label: "F"}, Target: model.Node{Label: "G"}, Label: "6"},
// 		},
// 	}

// 	// Правильный ответ
// 	answer := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "1"}, {Label: "2"}, {Label: "3"},
// 			{Label: "4"}, {Label: "5"}, {Label: "6"},
// 			{Label: "0"}, {Label: "7"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "1"}, Target: model.Node{Label: "2"}, Label: "B"},
// 			{Source: model.Node{Label: "2"}, Target: model.Node{Label: "3"}, Label: "C"},
// 			{Source: model.Node{Label: "3"}, Target: model.Node{Label: "4"}, Label: "D"},
// 			{Source: model.Node{Label: "4"}, Target: model.Node{Label: "5"}, Label: "E"},
// 			{Source: model.Node{Label: "5"}, Target: model.Node{Label: "6"}, Label: "F"},
// 			{Source: model.Node{Label: "6"}, Target: model.Node{Label: "7"}, Label: "G"},
// 			{Source: model.Node{Label: "0"}, Target: model.Node{Label: "1"}, Label: "A"},
// 		},
// 	}

// 	score := checker.CheckLinearFromLine(&task, &answer)
// 	if score != 100 {
// 		t.Errorf("CheckLinearFromLine failed: expected 100 for matching graphs, got %d", score)
// 	}

// 	// Краевой случай: неправильное ребро в ответе
// 	answerWrongEdge := model.Graph{
// 		Nodes: answer.Nodes,
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "1"}, Target: model.Node{Label: "2"}, Label: "B"},
// 			{Source: model.Node{Label: "2"}, Target: model.Node{Label: "3"}, Label: "C"},
// 			{Source: model.Node{Label: "3"}, Target: model.Node{Label: "4"}, Label: "D"},
// 			{Source: model.Node{Label: "4"}, Target: model.Node{Label: "5"}, Label: "E"},
// 			{Source: model.Node{Label: "5"}, Target: model.Node{Label: "6"}, Label: "F"},
// 			{Source: model.Node{Label: "6"}, Target: model.Node{Label: "7"}, Label: "G"},
// 			{Source: model.Node{Label: "0"}, Target: model.Node{Label: "1"}, Label: "A"},
// 			{Source: model.Node{Label: "0"}, Target: model.Node{Label: "2"}, Label: "A"},
// 		},
// 	}
// 	score = checker.CheckLinearFromLine(&task, &answerWrongEdge)
// 	if score != 60 {
// 		t.Errorf("CheckLinearFromLine failed: expected score 60 for incorrect edges, got %d", score)
// 	}

// 	// Краевой случай: графы разного размера
// 	answerSmall := model.Graph{
// 		Nodes: []model.Node{{Label: "1"}, {Label: "2"}},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}},
// 		},
// 	}
// 	score = checker.CheckLinearFromLine(&task, &answerSmall)
// 	if score != 0 {
// 		t.Errorf("CheckLinearFromLine failed: expected score 0 for different sized graphs, got %d", score)
// 	}

// 	// Пустые графы
// 	score = checker.CheckLinearFromLine(&model.Graph{}, &model.Graph{})
// 	if score != 0 {
// 		t.Errorf("CheckLinearFromLine failed: expected 100 for empty graphs, got %d", score)
// 	}
// }

// func TestChecker_CheckRadiusAndDiameter1(t *testing.T) {
// 	checker := NewChecker()

// 	// Задача: граф из 7 вершин
// 	task := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "A"}, {Label: "B"}, {Label: "C"},
// 			{Label: "D"}, {Label: "E"}, {Label: "F"}, {Label: "G"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}},
// 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}},
// 			{Source: model.Node{Label: "C"}, Target: model.Node{Label: "D"}},
// 			{Source: model.Node{Label: "D"}, Target: model.Node{Label: "E"}},
// 			{Source: model.Node{Label: "E"}, Target: model.Node{Label: "F"}},
// 			{Source: model.Node{Label: "F"}, Target: model.Node{Label: "G"}},
// 		},
// 	}

// 	// Ожидаемые результаты — рассчитанные вручную или с помощью функции
// 	// Здесь - расстояния как длина пути в цепочке
// 	dist_matrix_ans := map[string]map[string]int{
// 		"A": {"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6},
// 		"B": {"A": 1, "B": 0, "C": 1, "D": 2, "E": 3, "F": 4, "G": 5},
// 		"C": {"A": 2, "B": 1, "C": 0, "D": 1, "E": 2, "F": 3, "G": 4},
// 		"D": {"A": 3, "B": 2, "C": 1, "D": 0, "E": 1, "F": 2, "G": 3},
// 		"E": {"A": 4, "B": 3, "C": 2, "D": 1, "E": 0, "F": 1, "G": 2},
// 		"F": {"A": 5, "B": 4, "C": 3, "D": 2, "E": 1, "F": 0, "G": 1},
// 		"G": {"A": 6, "B": 5, "C": 4, "D": 3, "E": 2, "F": 1, "G": 0},
// 	}

// 	radius_ans := 3   // минимальный макс. расстояние — для вершины D (maxDist=3)
// 	diameter_ans := 6 // максимальный макс. расстояние — между A и G

// 	// Проверка правильного случая
// 	score := checker.CheckRadiusAndDiameter(&task, radius_ans, diameter_ans, dist_matrix_ans)
// 	if score != 100 {
// 		t.Errorf("CheckRadiusAndDiameter failed: expected 100, got %d", score)
// 	}

// 	// Краевой случай: ошибочные данные в матрице расстояний
// 	dist_matrix_wrong := map[string]map[string]int{
// 		"A": {"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 7}, // ошибка: 7 вместо 6
// 		"B": {"A": 1, "B": 0, "C": 1, "D": 2, "E": 3, "F": 4, "G": 5},
// 		"C": {"A": 2, "B": 1, "C": 0, "D": 1, "E": 2, "F": 3, "G": 4},
// 		"D": {"A": 3, "B": 2, "C": 1, "D": 0, "E": 1, "F": 2, "G": 3},
// 		"E": {"A": 4, "B": 3, "C": 2, "D": 1, "E": 0, "F": 1, "G": 2},
// 		"F": {"A": 5, "B": 4, "C": 3, "D": 2, "E": 1, "F": 0, "G": 1},
// 		"G": {"A": 6, "B": 5, "C": 4, "D": 3, "E": 2, "F": 1, "G": 0},
// 	}

// 	score = checker.CheckRadiusAndDiameter(&task, radius_ans, diameter_ans, dist_matrix_wrong)
// 	if score != 80 {
// 		t.Errorf("CheckRadiusAndDiameter failed: expected score < 100 for incorrect dist matrix, got %d", score)
// 	}

// 	// Краевой случай: неправильные radius или diameter
// 	score = checker.CheckRadiusAndDiameter(&task, radius_ans+1, diameter_ans, dist_matrix_ans)
// 	if score != 0 {
// 		t.Errorf("CheckRadiusAndDiameter failed: expected 0 for wrong radius, got %d", score)
// 	}

// 	score = checker.CheckRadiusAndDiameter(&task, radius_ans, diameter_ans+1, dist_matrix_ans)
// 	if score != 0 {
// 		t.Errorf("CheckRadiusAndDiameter failed: expected 0 for wrong diameter, got %d", score)
// 	}

// 	// Пустой граф
// 	emptyGraph := model.Graph{}
// 	score = checker.CheckRadiusAndDiameter(&emptyGraph, 0, 0, map[string]map[string]int{})
// 	if score != 0 {
// 		t.Errorf("CheckRadiusAndDiameter failed: expected 0 for empty graph, got %d", score)
// 	}
// 	score = checker.CheckRadiusAndDiameter(&task, radius_ans, diameter_ans, map[string]map[string]int{})
// 	if score != 0 {
// 		t.Errorf("CheckRadiusAndDiameter failed: expected 0 for empty graph, got %d", score)
// 	}
// }

// func TestChecker_CheckAdjacentMatrix(t *testing.T) {
// 	checker := NewChecker()

// 	// Неориентированный граф из 7 вершин
// 	task := model.Graph{
// 		Nodes: []model.Node{
// 			{Label: "A"}, {Label: "B"}, {Label: "C"},
// 			{Label: "D"}, {Label: "E"}, {Label: "F"}, {Label: "G"},
// 		},
// 		Edges: []model.Edge{
// 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}},
// 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}},
// 			{Source: model.Node{Label: "C"}, Target: model.Node{Label: "D"}},
// 			{Source: model.Node{Label: "D"}, Target: model.Node{Label: "E"}},
// 			{Source: model.Node{Label: "E"}, Target: model.Node{Label: "F"}},
// 			{Source: model.Node{Label: "F"}, Target: model.Node{Label: "G"}},
// 		},
// 	}

// 	// Правильная матрица смежности (неориентированная)
// 	correctAnswer := map[string]map[string]int{
// 		"A": {"A": 0, "B": 1, "C": 0, "D": 0, "E": 0, "F": 0, "G": 0},
// 		"B": {"A": 1, "B": 0, "C": 1, "D": 0, "E": 0, "F": 0, "G": 0},
// 		"C": {"A": 0, "B": 1, "C": 0, "D": 1, "E": 0, "F": 0, "G": 0},
// 		"D": {"A": 0, "B": 0, "C": 1, "D": 0, "E": 1, "F": 0, "G": 0},
// 		"E": {"A": 0, "B": 0, "C": 0, "D": 1, "E": 0, "F": 1, "G": 0},
// 		"F": {"A": 0, "B": 0, "C": 0, "D": 0, "E": 1, "F": 0, "G": 1},
// 		"G": {"A": 0, "B": 0, "C": 0, "D": 0, "E": 0, "F": 1, "G": 0},
// 	}

// 	score := checker.CheckAdjacentMatrix(&task, correctAnswer)
// 	if score != 100 {
// 		t.Errorf("CheckAdjacentMatrix failed: expected 100, got %d", score)
// 	}

// 	// Матрица с ошибками (например, отсутствует ребро D–E, лишнее ребро G–A)
// 	wrongAnswer := map[string]map[string]int{
// 		"A": {"A": 0, "B": 1, "C": 0, "D": 0, "E": 0, "F": 0, "G": 0},
// 		"B": {"A": 1, "B": 0, "C": 1, "D": 0, "E": 0, "F": 0, "G": 0},
// 		"C": {"A": 0, "B": 1, "C": 0, "D": 1, "E": 0, "F": 0, "G": 0},
// 		"D": {"A": 0, "B": 0, "C": 1, "D": 0, "E": 0, "F": 0, "G": 0},
// 		"E": {"A": 0, "B": 0, "C": 0, "D": 1, "E": 0, "F": 1, "G": 0},
// 		"F": {"A": 0, "B": 0, "C": 0, "D": 0, "E": 1, "F": 0, "G": 1},
// 		"G": {"A": 0, "B": 0, "C": 0, "D": 0, "E": 0, "F": 1, "G": 0},
// 	}

// 	score = checker.CheckAdjacentMatrix(&task, wrongAnswer)
// 	if score != 80 {
// 		t.Errorf("CheckAdjacentMatrix failed: expected score < 100 for incorrect matrix, got %d", score)
// 	}

// 	// Пустая матрица — нет ребер
// 	emptyAnswer := map[string]map[string]int{
// 		"A": {}, "B": {}, "C": {}, "D": {}, "E": {}, "F": {}, "G": {},
// 	}
// 	score = checker.CheckAdjacentMatrix(&task, emptyAnswer)
// 	if score == 100 {
// 		t.Errorf("CheckAdjacentMatrix failed: expected score < 100 for empty matrix, got %d", score)
// 	}

// 	// Несовпадение размеров (меньше узлов)
// 	smallerAnswer := map[string]map[string]int{
// 		"A": {"B": 1},
// 		"B": {"A": 1},
// 	}
// 	score = checker.CheckAdjacentMatrix(&task, smallerAnswer)
// 	if score != 0 {
// 		t.Errorf("CheckAdjacentMatrix failed: expected 0 for size mismatch, got %d", score)
// 	}

// 	// Пустой граф и пустой ответ
// 	emptyGraph := model.Graph{}
// 	score = checker.CheckAdjacentMatrix(&emptyGraph, map[string]map[string]int{})
// 	if score != 100 {
// 		t.Errorf("CheckAdjacentMatrix failed: expected 100 for empty graph and empty matrix, got %d", score)
// 	}
// 	score = checker.CheckAdjacentMatrix(&task, map[string]map[string]int{})
// 	if score != 0 {
// 		t.Errorf("CheckAdjacentMatrix failed: expected 0 for empty graph and empty matrix, got %d", score)
// 	}
// }

// func TestChecker_CheckEulerGraph1(t *testing.T) {
// 	checker := NewChecker()

// 	// Вспомогательная функция для создания узлов с id
// 	makeNode := func(id int, label string) model.Node {
// 		return model.Node{Id: id, Label: label}
// 	}

// 	// Граф с эйлеровым циклом (все вершины имеют четную степень)
// 	task := model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"), makeNode(2, "B"), makeNode(3, "C"),
// 			makeNode(4, "D"), makeNode(5, "E"), makeNode(6, "F"), makeNode(7, "G"),
// 		},
// 		Edges: []model.Edge{
// 			{Source: makeNode(1, "A"), Target: makeNode(2, "B")},
// 			{Source: makeNode(2, "B"), Target: makeNode(3, "C")},
// 			{Source: makeNode(3, "C"), Target: makeNode(4, "D")},
// 			{Source: makeNode(4, "D"), Target: makeNode(5, "E")},
// 			{Source: makeNode(5, "E"), Target: makeNode(6, "F")},
// 			{Source: makeNode(6, "F"), Target: makeNode(7, "G")},
// 			{Source: makeNode(7, "G"), Target: makeNode(1, "A")},
// 			{Source: makeNode(2, "B"), Target: makeNode(5, "E")},
// 			{Source: makeNode(3, "C"), Target: makeNode(6, "F")},
// 			{Source: makeNode(2, "B"), Target: makeNode(6, "F")},
// 			{Source: makeNode(3, "C"), Target: makeNode(5, "E")},
// 		},
// 	}

// 	// Правильный эйлеров цикл по ребрам в правильном порядке (маршрут)
// 	correctEulerPath := model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"), makeNode(2, "B"), makeNode(3, "C"),
// 			makeNode(4, "D"), makeNode(5, "E"), makeNode(6, "F"), makeNode(7, "G"),
// 		},
// 		Edges: []model.Edge{
// 			{Source: makeNode(1, "A"), Target: makeNode(2, "B"), Label: "1"},
// 			{Source: makeNode(2, "B"), Target: makeNode(3, "C"), Label: "2"},
// 			{Source: makeNode(3, "C"), Target: makeNode(4, "D"), Label: "7"},
// 			{Source: makeNode(4, "D"), Target: makeNode(5, "E"), Label: "8"},
// 			{Source: makeNode(5, "E"), Target: makeNode(6, "F"), Label: "9"},
// 			{Source: makeNode(6, "F"), Target: makeNode(7, "G"), Label: "10"},
// 			{Source: makeNode(7, "G"), Target: makeNode(1, "A"), Label: "11"},
// 			{Source: makeNode(2, "B"), Target: makeNode(5, "E"), Label: "5"},
// 			{Source: makeNode(3, "C"), Target: makeNode(6, "F"), Label: "3"},
// 			{Source: makeNode(2, "B"), Target: makeNode(6, "F"), Label: "4"},
// 			{Source: makeNode(3, "C"), Target: makeNode(5, "E"), Label: "6"},
// 		},
// 	}

// 	score := checker.CheckEulerGraph(&task, true, &correctEulerPath)
// 	if score != 100 {
// 		t.Errorf("CheckEulerGraph failed for correct Euler cycle, got %d", score)
// 	}

// 	// Неправильный путь — неверный порядок ребер
// 	incorrectEulerPath := model.Graph{
// 		Edges: []model.Edge{
// 			{Source: makeNode(1, "A"), Target: makeNode(2, "B"), Label: "1"},
// 			{Source: makeNode(2, "B"), Target: makeNode(3, "C"), Label: "2"},
// 			{Source: makeNode(3, "C"), Target: makeNode(4, "D"), Label: "7"},
// 			{Source: makeNode(4, "D"), Target: makeNode(5, "E"), Label: "8"},
// 			{Source: makeNode(5, "E"), Target: makeNode(6, "F"), Label: "9"},
// 			{Source: makeNode(6, "F"), Target: makeNode(7, "G"), Label: "10"},
// 			{Source: makeNode(7, "G"), Target: makeNode(1, "A"), Label: "11"},
// 			{Source: makeNode(2, "B"), Target: makeNode(5, "E"), Label: "4"},
// 			{Source: makeNode(3, "C"), Target: makeNode(6, "F"), Label: "3"},
// 			{Source: makeNode(2, "B"), Target: makeNode(6, "F"), Label: "5"},
// 			{Source: makeNode(3, "C"), Target: makeNode(5, "E"), Label: "6"},
// 		},
// 	}

// 	score = checker.CheckEulerGraph(&task, true, &incorrectEulerPath)
// 	if score != 0 {
// 		t.Errorf("CheckEulerGraph failed to detect wrong Euler path order, got %d", score)
// 	}

// 	// Граф без эйлерова цикла
// 	noEulerGraph := model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"), makeNode(2, "B"), makeNode(3, "C"),
// 		},
// 		Edges: []model.Edge{
// 			{Source: makeNode(1, "A"), Target: makeNode(2, "B"), Label: "1"},
// 			{Source: makeNode(2, "B"), Target: makeNode(3, "C"), Label: "2"},
// 		},
// 	}

// 	emptyAnswer := model.Graph{}

// 	score = checker.CheckEulerGraph(&noEulerGraph, false, &emptyAnswer)
// 	if score != 0 {
// 		t.Errorf("CheckEulerGraph failed for graph without Euler cycle, got %d", score)
// 	}

// 	// Ошибка, когда заявлено что есть эйлеров цикл, но он отсутствует
// 	score = checker.CheckEulerGraph(&noEulerGraph, true, &emptyAnswer)
// 	if score != 0 {
// 		t.Errorf("CheckEulerGraph failed to detect missing Euler cycle, got %d", score)
// 	}
// }

// func TestChecker_CheckMinPath1(t *testing.T) {
// 	checker := NewChecker()

// 	// Удобная функция для создания узлов с Id и Label
// 	makeNode := func(id int, label string) model.Node {
// 		return model.Node{Id: id, Label: label}
// 	}

// 	// Создаем граф из 7 вершин с весами ребер
// 	task := model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"), makeNode(2, "B"), makeNode(3, "C"),
// 			makeNode(4, "D"), makeNode(5, "E"), makeNode(6, "F"), makeNode(7, "G"),
// 		},
// 		Edges: []model.Edge{
// 			{Source: makeNode(1, "A"), Target: makeNode(2, "B"), Weight: 2},
// 			{Source: makeNode(1, "A"), Target: makeNode(3, "C"), Weight: 5},
// 			{Source: makeNode(2, "B"), Target: makeNode(4, "D"), Weight: 1},
// 			{Source: makeNode(3, "C"), Target: makeNode(4, "D"), Weight: 2},
// 			{Source: makeNode(4, "D"), Target: makeNode(5, "E"), Weight: 3},
// 			{Source: makeNode(5, "E"), Target: makeNode(6, "F"), Weight: 1},
// 			{Source: makeNode(6, "F"), Target: makeNode(7, "G"), Weight: 2},
// 			{Source: makeNode(3, "C"), Target: makeNode(7, "G"), Weight: 10},
// 		},
// 	}

// 	// Предполагаемый минимальный путь из A в G: A-B-D-E-F-G
// 	minPathLength := 2 + 1 + 3 + 1 + 2 // = 9

// 	weightsPathAns := map[string]int{
// 		"A": 0,
// 		"B": 2,
// 		"C": 5,
// 		"D": 3,
// 		"E": 6,
// 		"F": 7,
// 		"G": 9,
// 	}

// 	// Правильный ответ - выделены ребра минимального пути цветом, вес соответствует весу ребер
// 	answer := model.Graph{
// 		Edges: []model.Edge{
// 			{Source: makeNode(1, "A"), Target: makeNode(2, "B"), Weight: 2, Color: "red"},
// 			{Source: makeNode(1, "A"), Target: makeNode(3, "C"), Weight: 5, Color: DEFAULT_COLOR},
// 			{Source: makeNode(2, "B"), Target: makeNode(4, "D"), Weight: 1, Color: "red"},
// 			{Source: makeNode(3, "C"), Target: makeNode(4, "D"), Weight: 2, Color: DEFAULT_COLOR},
// 			{Source: makeNode(4, "D"), Target: makeNode(5, "E"), Weight: 3, Color: "red"},
// 			{Source: makeNode(5, "E"), Target: makeNode(6, "F"), Weight: 1, Color: "red"},
// 			{Source: makeNode(6, "F"), Target: makeNode(7, "G"), Weight: 2, Color: "red"},
// 			{Source: makeNode(3, "C"), Target: makeNode(7, "G"), Weight: 10, Color: DEFAULT_COLOR},
// 		},
// 	}

// 	score := checker.CheckMinPath(&task, "A", "G", minPathLength, weightsPathAns, &answer)
// 	if score != 100 {
// 		t.Errorf("CheckMinPath failed for valid shortest path")
// 	}

// 	score = checker.CheckMinPath(&task, "A", "G", minPathLength+1, weightsPathAns, &answer)
// 	if score != 0 {
// 		t.Errorf("CheckMinPath failed for valid shortest path")
// 	}

// 	// Ошибка в весах кратчайшего пути
// 	wrongWeightsAns := map[string]int{
// 		"A": 0,
// 		"B": 3, // неверно
// 	}
// 	score = checker.CheckMinPath(&task, "A", "B", 2, wrongWeightsAns, &answer)
// 	if score != 0 {
// 		t.Errorf("CheckMinPath failed to detect wrong weights")
// 	}

// 	// Ошибка в сумме весов выделенных ребер в ответе
// 	wrongAnswer := answer
// 	wrongAnswer.Edges[0].Color = "" // убрать цвет с первого ребра, сумма весов будет меньше
// 	score = checker.CheckMinPath(&task, "A", "G", minPathLength, weightsPathAns, &wrongAnswer)
// 	if score != 0 {
// 		t.Errorf("CheckMinPath failed to detect wrong highlighted path weights")
// 	}
// }

// func TestChecker_CheckPlanarGraph(t *testing.T) {
// 	checker := NewChecker()

// 	makeNode := func(id int, label string, x, y float64) model.Node {
// 		return model.Node{Id: id, Label: label, X: x, Y: y}
// 	}

// 	// 1) Планарный граф: треугольник
// 	planarGraph := model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A", 0, 0),
// 			makeNode(2, "B", 1, 0),
// 			makeNode(3, "C", 0.5, 1),
// 		},
// 		Edges: []model.Edge{
// 			{Id: "1", Source: makeNode(1, "A", 0, 0), Target: makeNode(2, "B", 1, 0)},
// 			{Id: "2", Source: makeNode(2, "B", 1, 0), Target: makeNode(3, "C", 0.5, 1)},
// 			{Id: "3", Source: makeNode(3, "C", 0.5, 1), Target: makeNode(1, "A", 0, 0)},
// 		},
// 	}
// 	if score := checker.CheckPlanarGraph(&planarGraph); score != 100 {
// 		t.Errorf("Planar triangle graph expected 100, got %d", score)
// 	}

// 	// 2) Непланарный граф: пересекающиеся ребра
// 	nodes := []model.Node{
// 		makeNode(1, "A", 0, 0),
// 		makeNode(2, "B", 2, 0),
// 		makeNode(3, "C", 2, 2),
// 		makeNode(4, "D", 0, 2),
// 	}
// 	nonPlanarGraph := model.Graph{
// 		Nodes: nodes,
// 		Edges: []model.Edge{
// 			{Id: "1", Source: nodes[0], Target: nodes[2]},
// 			{Id: "2", Source: nodes[1], Target: nodes[3]},
// 			{Id: "3", Source: nodes[0], Target: nodes[3]},
// 			{Id: "4", Source: nodes[1], Target: nodes[2]},
// 		},
// 	}
// 	if score := checker.CheckPlanarGraph(&nonPlanarGraph); score != 0 {
// 		t.Errorf("Non-planar graph expected 0, got %d", score)
// 	}

// 	// 3) Краевой случай: граф без ребер
// 	noEdgesGraph := model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A", 0, 0),
// 			makeNode(2, "B", 1, 1),
// 		},
// 		Edges: []model.Edge{},
// 	}
// 	if score := checker.CheckPlanarGraph(&noEdgesGraph); score != 0 {
// 		t.Errorf("Graph without edges expected 0, got %d", score)
// 	}

// 	// 4) Граф с соседними ребрами
// 	adjacentEdgesGraph := model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A", 0, 0),
// 			makeNode(2, "B", 1, 0),
// 			makeNode(3, "C", 1, 1),
// 		},
// 		Edges: []model.Edge{
// 			{Id: "1", Source: makeNode(1, "A", 0, 0), Target: makeNode(2, "B", 1, 0)},
// 			{Id: "2", Source: makeNode(2, "B", 1, 0), Target: makeNode(3, "C", 1, 1)},
// 		},
// 	}
// 	if score := checker.CheckPlanarGraph(&adjacentEdgesGraph); score != 100 {
// 		t.Errorf("Graph with adjacent edges expected 100, got %d", score)
// 	}

// 	// 5) Непланарный граф: квадрат с диагоналями
// 	k4PlanarGraph := model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A", 0, 0),
// 			makeNode(2, "B", 2, 0),
// 			makeNode(3, "C", 2, 2),
// 			makeNode(4, "D", 0, 2),
// 		},
// 		Edges: []model.Edge{
// 			{Id: "1", Source: makeNode(1, "A", 0, 0), Target: makeNode(2, "B", 2, 0)},
// 			{Id: "2", Source: makeNode(2, "B", 2, 0), Target: makeNode(3, "C", 2, 2)},
// 			{Id: "3", Source: makeNode(3, "C", 2, 2), Target: makeNode(4, "D", 0, 2)},
// 			{Id: "4", Source: makeNode(4, "D", 0, 2), Target: makeNode(1, "A", 0, 0)},
// 			// Диагонали, не пересекающиеся
// 			{Id: "5", Source: makeNode(1, "A", 0, 0), Target: makeNode(3, "C", 2, 2)},
// 			{Id: "6", Source: makeNode(2, "B", 2, 0), Target: makeNode(4, "D", 0, 2)},
// 		},
// 	}
// 	if score := checker.CheckPlanarGraph(&k4PlanarGraph); score != 0 {
// 		t.Errorf("K4 planar graph with crossing diagonals expected 0 (because edges intersect), got %d", score)
// 	}

// 	// 6) Планарный граф: звезда с центром
// 	starGraph := model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "Center", 1, 1),
// 			makeNode(2, "A", 0, 0),
// 			makeNode(3, "B", 2, 0),
// 			makeNode(4, "C", 2, 2),
// 			makeNode(5, "D", 0, 2),
// 		},
// 		Edges: []model.Edge{
// 			{Id: "1", Source: makeNode(1, "Center", 1, 1), Target: makeNode(2, "A", 0, 0)},
// 			{Id: "2", Source: makeNode(1, "Center", 1, 1), Target: makeNode(3, "B", 2, 0)},
// 			{Id: "3", Source: makeNode(1, "Center", 1, 1), Target: makeNode(4, "C", 2, 2)},
// 			{Id: "4", Source: makeNode(1, "Center", 1, 1), Target: makeNode(5, "D", 0, 2)},
// 		},
// 	}
// 	if score := checker.CheckPlanarGraph(&starGraph); score != 100 {
// 		t.Errorf("Star graph expected 100, got %d", score)
// 	}

// 	// 7) Краевой случай: ребро с самим собой (петля) - считаем планарным (потому что оно не пересекается с другими)
// 	loopGraph := model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A", 0, 0),
// 		},
// 		Edges: []model.Edge{
// 			{Id: "1", Source: makeNode(1, "A", 0, 0), Target: makeNode(1, "A", 0, 0)},
// 		},
// 	}
// 	if score := checker.CheckPlanarGraph(&loopGraph); score != 100 {
// 		t.Errorf("Graph with loop expected 100, got %d", score)
// 	}

// 	// 8) Краевой случай: пересекающиеся ребра, но с общим концом (не считаются пересечением)
// 	intersectAtVertexGraph := model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A", 0, 0),
// 			makeNode(2, "B", 2, 0),
// 			makeNode(3, "C", 1, 1),
// 		},
// 		Edges: []model.Edge{
// 			{Id: "1", Source: makeNode(1, "A", 0, 0), Target: makeNode(3, "C", 1, 1)},
// 			{Id: "2", Source: makeNode(3, "C", 1, 1), Target: makeNode(2, "B", 2, 0)},
// 		},
// 	}
// 	if score := checker.CheckPlanarGraph(&intersectAtVertexGraph); score != 100 {
// 		t.Errorf("Graph with edges intersecting at vertex expected 100, got %d", score)
// 	}
// }

// func TestChecker_CheckIntersectionGraphs(t *testing.T) {
// 	checker := NewChecker()

// 	makeNode := func(id int, label string) model.Node {
// 		return model.Node{Id: id, Label: label}
// 	}
// 	makeEdge := func(id int, src, tgt model.Node) model.Edge {
// 		return model.Edge{Id: strconv.Itoa(id), Source: src, Target: tgt}
// 	}

// 	// Граф 1
// 	graph1 := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "C")),
// 		},
// 	}

// 	// Граф 2
// 	graph2 := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "D"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "D")),
// 		},
// 	}

// 	// Правильный ответ - пересечение graph1 и graph2: ребро A-B
// 	correctAnswer := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(0, makeNode(1, "A"), makeNode(2, "B")),
// 		},
// 	}

// 	// Тест 1: полный правильный ответ
// 	score := checker.CheckIntersectionGraphs(correctAnswer, graph1, graph2)
// 	if score != 100 {
// 		t.Errorf("Expected 100 for correct intersection, got %d", score)
// 	}

// 	// Тест 2: ответ с лишним ребром (ребро B-C, которого нет в пересечении)
// 	wrongAnswerExtraEdge := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "C")), // лишнее ребро
// 		},
// 	}
// 	score = checker.CheckIntersectionGraphs(wrongAnswerExtraEdge, graph1, graph2)
// 	if score >= 100 {
// 		t.Errorf("Expected less than 100 for answer with extra edge, got %d", score)
// 	}

// 	// Тест 3: ответ с отсутствующим ребром (пустой ответ)
// 	emptyAnswer := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 		},
// 		Edges: []model.Edge{}, // нет ребер
// 	}
// 	score = checker.CheckIntersectionGraphs(emptyAnswer, graph1, graph2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for empty answer, got %d", score)
// 	}

// 	// Тест 4: ответ с правильным ребром, но лишними вершинами (вершины есть, ребро - правильное)
// 	extraNodesAnswer := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"), // лишняя вершина
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 		},
// 	}
// 	score = checker.CheckIntersectionGraphs(extraNodesAnswer, graph1, graph2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for answer with extra nodes but no corresponding edges, got %d", score)
// 	}

// 	// Тест 5: пересечение пустое (графы без общих ребер)
// 	graph3 := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "X"),
// 			makeNode(2, "Y"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "X"), makeNode(2, "Y")),
// 		},
// 	}
// 	score = checker.CheckIntersectionGraphs(correctAnswer, graph1, graph3)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for graphs with empty intersection, got %d", score)
// 	}
// }

// func TestChecker_CheckIntersectionGraphs_SixNodes(t *testing.T) {
// 	checker := NewChecker()

// 	makeNode := func(id int, label string) model.Node {
// 		return model.Node{Id: id, Label: label}
// 	}
// 	makeEdge := func(id int, src, tgt model.Node) model.Edge {
// 		return model.Edge{Id: strconv.Itoa(id), Source: src, Target: tgt}
// 	}

// 	// Граф 1: 6 вершин, несколько ребер
// 	graph1 := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 			makeNode(4, "D"),
// 			makeNode(5, "E"),
// 			makeNode(6, "F"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "C")),
// 			makeEdge(3, makeNode(3, "C"), makeNode(4, "D")),
// 			makeEdge(4, makeNode(4, "D"), makeNode(5, "E")),
// 			makeEdge(5, makeNode(5, "E"), makeNode(6, "F")),
// 		},
// 	}

// 	// Граф 2: 6 вершин, пересечение с graph1 частичное
// 	graph2 := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 			makeNode(4, "D"),
// 			makeNode(5, "X"),
// 			makeNode(6, "Y"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")), // общая грань
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "C")), // общая грань
// 			makeEdge(3, makeNode(3, "C"), makeNode(6, "Y")), // нет в graph1
// 			makeEdge(4, makeNode(5, "X"), makeNode(6, "Y")), // нет в graph1
// 		},
// 	}

// 	// Правильный ответ — пересечение graph1 и graph2
// 	correctAnswer := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 			makeNode(4, "D"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "C")),
// 		},
// 	}

// 	// Тест 1: полный правильный ответ
// 	score := checker.CheckIntersectionGraphs(correctAnswer, graph1, graph2)
// 	if score != 100 {
// 		t.Errorf("Expected 100 for correct intersection, got %d", score)
// 	}

// 	// Тест 2: ответ с лишним ребром (например, ребро C-D, отсутствующее в пересечении)
// 	wrongAnswerExtraEdge := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 			makeNode(4, "D"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "C")),
// 			makeEdge(3, makeNode(3, "C"), makeNode(4, "D")), // лишнее ребро
// 		},
// 	}
// 	score = checker.CheckIntersectionGraphs(wrongAnswerExtraEdge, graph1, graph2)
// 	if score >= 100 {
// 		t.Errorf("Expected less than 100 for answer with extra edge, got %d", score)
// 	}

// 	// Тест 3: ответ без ребер (пустое пересечение)
// 	emptyAnswer := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 		},
// 		Edges: []model.Edge{},
// 	}
// 	score = checker.CheckIntersectionGraphs(emptyAnswer, graph1, graph2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for empty answer, got %d", score)
// 	}

// 	// Тест 4: ответ с правильными ребрами, но лишними вершинами
// 	extraNodesAnswer := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 			makeNode(5, "E"), // лишняя вершина без ребер
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "C")),
// 		},
// 	}
// 	score = checker.CheckIntersectionGraphs(extraNodesAnswer, graph1, graph2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for answer with extra nodes but no corresponding edges, got %d", score)
// 	}

// 	// Тест 5: пересечение пустое — графы не имеют общих ребер
// 	graph3 := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "X"),
// 			makeNode(2, "Y"),
// 			makeNode(3, "Z"),
// 			makeNode(4, "W"),
// 			makeNode(5, "Q"),
// 			makeNode(6, "R"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "X"), makeNode(2, "Y")),
// 			makeEdge(2, makeNode(3, "Z"), makeNode(4, "W")),
// 		},
// 	}
// 	score = checker.CheckIntersectionGraphs(correctAnswer, graph1, graph3)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for graphs with empty intersection, got %d", score)
// 	}
// }

// func TestChecker_CheckUnionGraphs_SixNodes(t *testing.T) {
// 	checker := NewChecker()

// 	makeNode := func(id int, label string) model.Node {
// 		return model.Node{Id: id, Label: label}
// 	}
// 	makeEdge := func(id int, src, tgt model.Node) model.Edge {
// 		return model.Edge{Id: strconv.Itoa(id), Source: src, Target: tgt}
// 	}

// 	// Граф 1: 6 вершин, несколько ребер
// 	graph1 := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 			makeNode(4, "D"),
// 			makeNode(5, "E"),
// 			makeNode(6, "F"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "C")),
// 			makeEdge(3, makeNode(3, "C"), makeNode(4, "D")),
// 		},
// 	}

// 	// Граф 2: 6 вершин, с другими ребрами
// 	graph2 := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 			makeNode(4, "D"),
// 			makeNode(5, "E"),
// 			makeNode(6, "F"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(4, makeNode(4, "D"), makeNode(5, "E")),
// 			makeEdge(5, makeNode(5, "E"), makeNode(6, "F")),
// 			makeEdge(6, makeNode(1, "A"), makeNode(6, "F")),
// 		},
// 	}

// 	// Правильный ответ — объединение graph1 и graph2
// 	correctAnswer := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 			makeNode(4, "D"),
// 			makeNode(5, "E"),
// 			makeNode(6, "F"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "C")),
// 			makeEdge(3, makeNode(3, "C"), makeNode(4, "D")),
// 			makeEdge(4, makeNode(4, "D"), makeNode(5, "E")),
// 			makeEdge(5, makeNode(5, "E"), makeNode(6, "F")),
// 			makeEdge(6, makeNode(1, "A"), makeNode(6, "F")),
// 		},
// 	}

// 	// Тест 1: полный правильный ответ
// 	score := checker.CheckUnionGraphs(correctAnswer, graph1, graph2)
// 	if score != 100 {
// 		t.Errorf("Expected 100 for correct union, got %d", score)
// 	}

// 	// Тест 2: ответ с пропущенным ребром
// 	missingEdgeAnswer := &model.Graph{
// 		Nodes: correctAnswer.Nodes,
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "C")),
// 			makeEdge(3, makeNode(3, "C"), makeNode(4, "D")),
// 			makeEdge(4, makeNode(4, "D"), makeNode(5, "E")),
// 			makeEdge(6, makeNode(1, "A"), makeNode(6, "F")), // ребро (5,6) пропущено
// 		},
// 	}
// 	score = checker.CheckUnionGraphs(missingEdgeAnswer, graph1, graph2)
// 	if score >= 100 {
// 		t.Errorf("Expected less than 100 for answer with missing edge, got %d", score)
// 	}

// 	// Тест 3: ответ с лишним ребром
// 	extraEdgeAnswer := &model.Graph{
// 		Nodes: correctAnswer.Nodes,
// 		Edges: append(correctAnswer.Edges, makeEdge(7, makeNode(2, "B"), makeNode(5, "E"))), // лишнее ребро
// 	}
// 	score = checker.CheckUnionGraphs(extraEdgeAnswer, graph1, graph2)
// 	if score >= 100 {
// 		t.Errorf("Expected less than 100 for answer with extra edge, got %d", score)
// 	}

// 	// Тест 4: пустой ответ (нет ребер)
// 	emptyAnswer := &model.Graph{
// 		Nodes: correctAnswer.Nodes,
// 		Edges: []model.Edge{},
// 	}
// 	score = checker.CheckUnionGraphs(emptyAnswer, graph1, graph2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for empty answer, got %d", score)
// 	}
// }

// func TestChecker_CheckJoinGraphs_SixNodes(t *testing.T) {
// 	checker := NewChecker()

// 	makeNode := func(id int, label string) model.Node {
// 		return model.Node{Id: id, Label: label}
// 	}
// 	makeEdge := func(id int, src, tgt model.Node) model.Edge {
// 		return model.Edge{Id: strconv.Itoa(id), Source: src, Target: tgt}
// 	}

// 	// Граф 1: 6 вершин, ребра в первой части
// 	graph1 := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 			makeNode(4, "D"),
// 			makeNode(5, "E"),
// 			makeNode(6, "F"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(1, makeNode(1, "A"), makeNode(2, "B")),
// 			makeEdge(2, makeNode(2, "B"), makeNode(3, "C")),
// 			makeEdge(3, makeNode(3, "C"), makeNode(4, "D")),
// 		},
// 	}

// 	// Граф 2: 6 вершин, ребра во второй части
// 	graph2 := &model.Graph{
// 		Nodes: []model.Node{
// 			makeNode(1, "A"),
// 			makeNode(2, "B"),
// 			makeNode(3, "C"),
// 			makeNode(4, "D"),
// 			makeNode(5, "E"),
// 			makeNode(6, "F"),
// 		},
// 		Edges: []model.Edge{
// 			makeEdge(4, makeNode(4, "D"), makeNode(5, "E")),
// 			makeEdge(5, makeNode(5, "E"), makeNode(6, "F")),
// 			makeEdge(6, makeNode(1, "A"), makeNode(6, "F")),
// 		},
// 	}

// 	// Правильный ответ — соединение graph1 и graph2
// 	correctAnswer := graph1.Join(graph2)

// 	// Тест 1: полный правильный ответ
// 	score := checker.CheckJoinGraphs(correctAnswer, graph1, graph2)
// 	if score != 100 {
// 		t.Errorf("Expected 100 for correct join, got %d", score)
// 	}

// 	// Тест 2: ответ с пропущенным ребром
// 	missingEdgeAnswer := &model.Graph{
// 		Nodes: correctAnswer.Nodes,
// 		Edges: correctAnswer.Edges[:len(correctAnswer.Edges)-1], // одно ребро пропущено
// 	}
// 	score = checker.CheckJoinGraphs(missingEdgeAnswer, graph1, graph2)
// 	if score >= 100 {
// 		t.Errorf("Expected less than 100 for answer with missing edge, got %d", score)
// 	}

// 	// Тест 3: ответ с лишним ребром
// 	extraEdgeAnswer := &model.Graph{
// 		Nodes: correctAnswer.Nodes,
// 		Edges: append(correctAnswer.Edges, makeEdge(7, makeNode(2, "B"), makeNode(5, "E"))), // лишнее ребро
// 	}
// 	score = checker.CheckJoinGraphs(extraEdgeAnswer, graph1, graph2)
// 	if score >= 100 {
// 		t.Errorf("Expected less than 100 for answer with extra edge, got %d", score)
// 	}

// 	// Тест 4: пустой ответ (нет ребер)
// 	emptyAnswer := &model.Graph{
// 		Nodes: correctAnswer.Nodes,
// 		Edges: []model.Edge{},
// 	}
// 	score = checker.CheckJoinGraphs(emptyAnswer, graph1, graph2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for empty answer, got %d", score)
// 	}
// }

// func TestChecker_CheckIntersectionMatrices_SixNodes(t *testing.T) {
// 	checker := NewChecker()

// 	// Удобная функция для создания матрицы смежности
// 	makeAdjMatrix := func() map[string]map[string]int {
// 		return map[string]map[string]int{
// 			"A": {"B": 1, "C": 0, "D": 0, "E": 0, "F": 0},
// 			"B": {"A": 1, "C": 1, "D": 0, "E": 0, "F": 0},
// 			"C": {"A": 0, "B": 1, "D": 1, "E": 0, "F": 0},
// 			"D": {"A": 0, "B": 0, "C": 1, "E": 1, "F": 0},
// 			"E": {"A": 0, "B": 0, "C": 0, "D": 1, "F": 1},
// 			"F": {"A": 0, "B": 0, "C": 0, "D": 0, "E": 1},
// 		}
// 	}

// 	matrix1 := makeAdjMatrix()
// 	matrix2 := map[string]map[string]int{
// 		"A": {"B": 1, "C": 0, "D": 0, "E": 0, "F": 1},
// 		"B": {"A": 1, "C": 1, "D": 0, "E": 0, "F": 0},
// 		"C": {"A": 0, "B": 1, "D": 0, "E": 0, "F": 0},
// 		"D": {"A": 0, "B": 0, "C": 0, "E": 1, "F": 0},
// 		"E": {"A": 0, "B": 0, "C": 0, "D": 1, "F": 1},
// 		"F": {"A": 1, "B": 0, "C": 0, "D": 0, "E": 1},
// 	}

// 	// true_answer будет пересечением graph1 и graph2
// 	graph1 := model.MakeGraphFromAdjLabelMatrix(matrix1)
// 	graph2 := model.MakeGraphFromAdjLabelMatrix(matrix2)
// 	true_answer := graph1.Intersect(graph2)

// 	// Тест 1: правильный ответ (пересечение)
// 	score := checker.CheckIntersectionMatrices(true_answer, matrix1, matrix2)
// 	if score != 100 {
// 		t.Errorf("Expected 100 for correct intersection, got %d", score)
// 	}

// 	// Тест 2: ответ с пропущенным ребром
// 	missingEdgeAnswer := *true_answer
// 	if len(missingEdgeAnswer.Edges) > 0 {
// 		missingEdgeAnswer.Edges = missingEdgeAnswer.Edges[:len(missingEdgeAnswer.Edges)-1]
// 	}
// 	score = checker.CheckIntersectionMatrices(&missingEdgeAnswer, matrix1, matrix2)
// 	if score != 75 {
// 		t.Errorf("Expected less than 100 for answer with missing edge, got %d", score)
// 	}

// 	// Тест 3: ответ с лишним ребром
// 	extraEdgeAnswer := *true_answer
// 	extraEdgeAnswer.Edges = append(extraEdgeAnswer.Edges,
// 		model.Edge{
// 			Id:     "999",
// 			Source: model.Node{Label: "A"},
// 			Target: model.Node{Label: "D"},
// 		})
// 	score = checker.CheckIntersectionMatrices(&extraEdgeAnswer, matrix1, matrix2)
// 	if score >= 100 {
// 		t.Errorf("Expected less than 100 for answer with extra edge, got %d", score)
// 	}

// 	// Тест 4: пустой граф
// 	emptyGraph := &model.Graph{
// 		Nodes: true_answer.Nodes,
// 		Edges: []model.Edge{},
// 	}
// 	score = checker.CheckIntersectionMatrices(emptyGraph, matrix1, matrix2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for empty graph answer, got %d", score)
// 	}
// }

// func TestChecker_CheckUnionMatrices_SixNodes(t *testing.T) {
// 	checker := NewChecker()

// 	makeAdjMatrix1 := func() map[string]map[string]int {
// 		return map[string]map[string]int{
// 			"A": {"B": 1, "C": 0, "D": 0, "E": 0, "F": 0},
// 			"B": {"A": 1, "C": 1, "D": 0, "E": 0, "F": 0},
// 			"C": {"A": 0, "B": 1, "D": 1, "E": 0, "F": 0},
// 			"D": {"A": 0, "B": 0, "C": 1, "E": 1, "F": 0},
// 			"E": {"A": 0, "B": 0, "C": 0, "D": 1, "F": 1},
// 			"F": {"A": 0, "B": 0, "C": 0, "D": 0, "E": 1},
// 		}
// 	}

// 	makeAdjMatrix2 := func() map[string]map[string]int {
// 		return map[string]map[string]int{
// 			"A": {"B": 1, "C": 0, "D": 0, "E": 0, "F": 1},
// 			"B": {"A": 1, "C": 1, "D": 0, "E": 0, "F": 0},
// 			"C": {"A": 0, "B": 1, "D": 0, "E": 0, "F": 0},
// 			"D": {"A": 0, "B": 0, "C": 0, "E": 1, "F": 0},
// 			"E": {"A": 0, "B": 0, "C": 0, "D": 1, "F": 1},
// 			"F": {"A": 1, "B": 0, "C": 0, "D": 0, "E": 1},
// 		}
// 	}

// 	matrix1 := makeAdjMatrix1()
// 	matrix2 := makeAdjMatrix2()

// 	graph1 := model.MakeGraphFromAdjLabelMatrix(matrix1)
// 	graph2 := model.MakeGraphFromAdjLabelMatrix(matrix2)
// 	true_answer := graph1.Union(graph2)

// 	// Тест 1: правильный ответ (объединение)
// 	score := checker.CheckUnionMatrices(true_answer, matrix1, matrix2)
// 	if score != 100 {
// 		t.Errorf("Expected 100 for correct union, got %d", score)
// 	}

// 	// Тест 2: ответ с пропущенным ребром
// 	missingEdgeAnswer := *true_answer
// 	if len(missingEdgeAnswer.Edges) > 0 {
// 		missingEdgeAnswer.Edges = missingEdgeAnswer.Edges[:len(missingEdgeAnswer.Edges)-1]
// 	}
// 	score = checker.CheckUnionMatrices(&missingEdgeAnswer, matrix1, matrix2)
// 	if score >= 100 {
// 		t.Errorf("Expected less than 100 for answer with missing edge, got %d", score)
// 	}

// 	// Тест 3: ответ с лишним ребром
// 	extraEdgeAnswer := *true_answer
// 	extraEdgeAnswer.Edges = append(extraEdgeAnswer.Edges,
// 		model.Edge{
// 			Id:     "999",
// 			Source: model.Node{Label: "A"},
// 			Target: model.Node{Label: "D"},
// 		})
// 	score = checker.CheckUnionMatrices(&extraEdgeAnswer, matrix1, matrix2)
// 	if score >= 100 {
// 		t.Errorf("Expected less than 100 for answer with extra edge, got %d", score)
// 	}

// 	// Тест 4: пустой граф
// 	emptyGraph := &model.Graph{
// 		Nodes: true_answer.Nodes,
// 		Edges: []model.Edge{},
// 	}
// 	score = checker.CheckUnionMatrices(emptyGraph, matrix1, matrix2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for empty graph answer, got %d", score)
// 	}
// }

// func TestChecker_CheckJoinMatrices_SixNodes(t *testing.T) {
// 	checker := NewChecker()

// 	// Матрица смежности для графа 1
// 	matrix1 := map[string]map[string]int{
// 		"A": {"B": 1, "C": 0, "D": 0, "E": 0, "F": 0},
// 		"B": {"A": 1, "C": 1, "D": 0, "E": 0, "F": 0},
// 		"C": {"A": 0, "B": 1, "D": 1, "E": 0, "F": 0},
// 		"D": {"A": 0, "B": 0, "C": 1, "E": 1, "F": 0},
// 		"E": {"A": 0, "B": 0, "C": 0, "D": 1, "F": 1},
// 		"F": {"A": 0, "B": 0, "C": 0, "D": 0, "E": 1},
// 	}

// 	// Матрица смежности для графа 2
// 	matrix2 := map[string]map[string]int{
// 		"A": {"B": 0, "C": 1, "D": 0, "E": 1, "F": 0},
// 		"B": {"A": 0, "C": 1, "D": 1, "E": 0, "F": 0},
// 		"C": {"A": 1, "B": 1, "D": 0, "E": 0, "F": 0},
// 		"D": {"A": 0, "B": 1, "C": 0, "E": 1, "F": 0},
// 		"E": {"A": 1, "B": 0, "C": 0, "D": 1, "F": 1},
// 		"F": {"A": 0, "B": 0, "C": 0, "D": 0, "E": 1},
// 	}

// 	graph1 := model.MakeGraphFromAdjLabelMatrix(matrix1)
// 	graph2 := model.MakeGraphFromAdjLabelMatrix(matrix2)
// 	true_answer := graph1.Join(graph2)

// 	// Тест 1: корректный ответ (соединение)
// 	score := checker.CheckJoinMatrices(true_answer, matrix1, matrix2)
// 	if score != 100 {
// 		t.Errorf("Expected 100 for correct join, got %d", score)
// 	}

// 	// Тест 2: ответ с пропущенным ребром
// 	missingEdgeAnswer := *true_answer
// 	if len(missingEdgeAnswer.Edges) > 0 {
// 		missingEdgeAnswer.Edges = missingEdgeAnswer.Edges[:len(missingEdgeAnswer.Edges)-1]
// 	}
// 	score = checker.CheckJoinMatrices(&missingEdgeAnswer, matrix1, matrix2)
// 	if score != 88 {
// 		t.Errorf("Expected less than 100 for answer with missing edge, got %d", score)
// 	}

// 	// Тест 3: ответ с лишним ребром
// 	extraEdgeAnswer := *true_answer
// 	extraEdgeAnswer.Edges = append(extraEdgeAnswer.Edges,
// 		model.Edge{
// 			Id:     "999",
// 			Source: model.Node{Label: "A"},
// 			Target: model.Node{Label: "F"},
// 		})
// 	score = checker.CheckJoinMatrices(&extraEdgeAnswer, matrix1, matrix2)
// 	if score >= 100 {
// 		t.Errorf("Expected less than 100 for answer with extra edge, got %d", score)
// 	}

// 	// Тест 4: пустой граф
// 	emptyGraph := &model.Graph{
// 		Nodes: true_answer.Nodes,
// 		Edges: []model.Edge{},
// 	}
// 	score = checker.CheckJoinMatrices(emptyGraph, matrix1, matrix2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for empty graph answer, got %d", score)
// 	}
// }

// func TestChecker_CheckCartesianProduct_ManualAnswer(t *testing.T) {
// 	checker := NewChecker()

// 	graph1 := &model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"},
// 			{Id: 2, Label: "B"},
// 			{Id: 3, Label: "C"},
// 		},
// 		Edges: []model.Edge{
// 			{Id: "1", Source: model.Node{Id: 1, Label: "A"}, Target: model.Node{Id: 2, Label: "B"}},
// 			{Id: "2", Source: model.Node{Id: 2, Label: "B"}, Target: model.Node{Id: 3, Label: "C"}},
// 		},
// 	}

// 	graph2 := &model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 10, Label: "1"},
// 			{Id: 11, Label: "2"},
// 		},
// 		Edges: []model.Edge{
// 			{Id: "10", Source: model.Node{Id: 10, Label: "1"}, Target: model.Node{Id: 11, Label: "2"}},
// 		},
// 	}

// 	true_answer := &model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 100, Label: "(A,1)"},
// 			{Id: 101, Label: "(A,2)"},
// 			{Id: 102, Label: "(B,1)"},
// 			{Id: 103, Label: "(B,2)"},
// 			{Id: 104, Label: "(C,1)"},
// 			{Id: 105, Label: "(C,2)"},
// 		},
// 		Edges: []model.Edge{
// 			// Ребра по graph2 при фиксированном graph1
// 			{Id: "1000", Source: model.Node{Id: 100, Label: "(A,1)"}, Target: model.Node{Id: 101, Label: "(A,2)"}},
// 			{Id: "1001", Source: model.Node{Id: 102, Label: "(B,1)"}, Target: model.Node{Id: 103, Label: "(B,2)"}},
// 			{Id: "1002", Source: model.Node{Id: 104, Label: "(C,1)"}, Target: model.Node{Id: 105, Label: "(C,2)"}},
// 			// Ребра по graph1 при фиксированном graph2
// 			{Id: "1003", Source: model.Node{Id: 100, Label: "(A,1)"}, Target: model.Node{Id: 102, Label: "(B,1)"}},
// 			{Id: "1004", Source: model.Node{Id: 101, Label: "(A,2)"}, Target: model.Node{Id: 103, Label: "(B,2)"}},
// 			{Id: "1005", Source: model.Node{Id: 102, Label: "(B,1)"}, Target: model.Node{Id: 104, Label: "(C,1)"}},
// 			{Id: "1006", Source: model.Node{Id: 103, Label: "(B,2)"}, Target: model.Node{Id: 105, Label: "(C,2)"}},
// 		},
// 	}

// 	score := checker.CheckCartesianProduct(true_answer, graph1, graph2)
// 	if score != 100 {
// 		t.Errorf("Expected 100 for correct manual Cartesian product, got %d", score)
// 	}

// 	// Пример с пропущенным ребром
// 	missingEdgeAnswer := *true_answer
// 	missingEdgeAnswer.Edges = missingEdgeAnswer.Edges[:len(missingEdgeAnswer.Edges)-1]
// 	score = checker.CheckCartesianProduct(&missingEdgeAnswer, graph1, graph2)
// 	if score != 86 {
// 		t.Errorf("Expected less than 100 for answer with missing edge, got %d", score)
// 	}

// 	// Пример с лишним ребром
// 	extraEdgeAnswer := *true_answer
// 	extraEdgeAnswer.Edges = append(extraEdgeAnswer.Edges,
// 		model.Edge{
// 			Id:     "9999",
// 			Source: model.Node{Id: 100, Label: "(A,1)"},
// 			Target: model.Node{Id: 105, Label: "(C,2)"},
// 		})
// 	score = checker.CheckCartesianProduct(&extraEdgeAnswer, graph1, graph2)
// 	if score >= 100 {
// 		t.Errorf("Expected less than 100 for answer with extra edge, got %d", score)
// 	}

// 	// Пример с пустым графом
// 	emptyGraph := &model.Graph{
// 		Nodes: true_answer.Nodes,
// 		Edges: []model.Edge{},
// 	}
// 	score = checker.CheckCartesianProduct(emptyGraph, graph1, graph2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for empty graph answer, got %d", score)
// 	}
// }

// func TestChecker_CheckTensorProduct_ManualAnswer(t *testing.T) {
// 	checker := NewChecker()

// 	graph1 := &model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"},
// 			{Id: 2, Label: "B"},
// 			{Id: 3, Label: "C"},
// 		},
// 		Edges: []model.Edge{
// 			{Id: "1", Source: model.Node{Id: 1, Label: "A"}, Target: model.Node{Id: 2, Label: "B"}},
// 			{Id: "2", Source: model.Node{Id: 2, Label: "B"}, Target: model.Node{Id: 3, Label: "C"}},
// 		},
// 	}

// 	graph2 := &model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 10, Label: "1"},
// 			{Id: 11, Label: "2"},
// 		},
// 		Edges: []model.Edge{
// 			{Id: "10", Source: model.Node{Id: 10, Label: "1"}, Target: model.Node{Id: 11, Label: "2"}},
// 		},
// 	}

// 	// Ручной эталонный граф тензорного произведения graph1 и graph2
// 	true_answer := &model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 100, Label: "(A,1)"},
// 			{Id: 101, Label: "(A,2)"},
// 			{Id: 102, Label: "(B,1)"},
// 			{Id: 103, Label: "(B,2)"},
// 			{Id: 104, Label: "(C,1)"},
// 			{Id: 105, Label: "(C,2)"},
// 		},
// 		Edges: []model.Edge{
// 			// Ребра существуют, если ребро в graph1 и ребро в graph2 существуют между соответствующими вершинами
// 			{Id: "1000", Source: model.Node{Id: 100, Label: "(A,1)"}, Target: model.Node{Id: 103, Label: "(B,2)"}},
// 			{Id: "1001", Source: model.Node{Id: 101, Label: "(A,2)"}, Target: model.Node{Id: 102, Label: "(B,1)"}},
// 			{Id: "1002", Source: model.Node{Id: 102, Label: "(B,1)"}, Target: model.Node{Id: 105, Label: "(C,2)"}},
// 			{Id: "1003", Source: model.Node{Id: 103, Label: "(B,2)"}, Target: model.Node{Id: 104, Label: "(C,1)"}},
// 		},
// 	}

// 	score := checker.CheckTensorProduct(true_answer, graph1, graph2)
// 	if score != 100 {
// 		t.Errorf("Expected 100 for correct manual Tensor product, got %d", score)
// 	}

// 	// Пример с отсутствующим ребром
// 	missingEdgeAnswer := *true_answer
// 	missingEdgeAnswer.Edges = missingEdgeAnswer.Edges[:len(missingEdgeAnswer.Edges)-1]
// 	score = checker.CheckTensorProduct(&missingEdgeAnswer, graph1, graph2)
// 	if score != 75 {
// 		t.Errorf("Expected less than 100 for answer with missing edge, got %d", score)
// 	}

// 	// Пример с лишним ребром
// 	extraEdgeAnswer := *true_answer
// 	extraEdgeAnswer.Edges = append(extraEdgeAnswer.Edges,
// 		model.Edge{
// 			Id:     "9999",
// 			Source: model.Node{Id: 100, Label: "(A,1)"},
// 			Target: model.Node{Id: 105, Label: "(C,2)"},
// 		})
// 	score = checker.CheckTensorProduct(&extraEdgeAnswer, graph1, graph2)
// 	if score != 75 {
// 		t.Errorf("Expected less than 100 for answer with extra edge, got %d", score)
// 	}

// 	// Пример с пустым графом
// 	emptyGraph := &model.Graph{
// 		Nodes: true_answer.Nodes,
// 		Edges: []model.Edge{},
// 	}
// 	score = checker.CheckTensorProduct(emptyGraph, graph1, graph2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for empty graph answer, got %d", score)
// 	}
// }

// func TestChecker_CheckLexicographicalProduct_ManualAnswer(t *testing.T) {
// 	checker := NewChecker()

// 	graph1 := &model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"},
// 			{Id: 2, Label: "B"},
// 		},
// 		Edges: []model.Edge{
// 			{Id: "1", Source: model.Node{Id: 1, Label: "A"}, Target: model.Node{Id: 2, Label: "B"}},
// 		},
// 	}

// 	graph2 := &model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 10, Label: "1"},
// 			{Id: 11, Label: "2"},
// 		},
// 		Edges: []model.Edge{
// 			{Id: "10", Source: model.Node{Id: 10, Label: "1"}, Target: model.Node{Id: 11, Label: "2"}},
// 		},
// 	}

// 	// Ручной эталонный граф лексикографического произведения graph1 и graph2
// 	true_answer := &model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 100, Label: "(A,1)"},
// 			{Id: 101, Label: "(A,2)"},
// 			{Id: 102, Label: "(B,1)"},
// 			{Id: 103, Label: "(B,2)"},
// 		},
// 		Edges: []model.Edge{
// 			// Ребра из graph1 между "A" и "B", со всеми вершинами graph2
// 			{Id: "1000", Source: model.Node{Id: 100, Label: "(A,1)"}, Target: model.Node{Id: 102, Label: "(B,1)"}},
// 			{Id: "1001", Source: model.Node{Id: 100, Label: "(A,1)"}, Target: model.Node{Id: 103, Label: "(B,2)"}},
// 			{Id: "1002", Source: model.Node{Id: 101, Label: "(A,2)"}, Target: model.Node{Id: 102, Label: "(B,1)"}},
// 			{Id: "1003", Source: model.Node{Id: 101, Label: "(A,2)"}, Target: model.Node{Id: 103, Label: "(B,2)"}},
// 			// Ребро из graph2 между "1" и "2" для вершины "A"
// 			{Id: "1004", Source: model.Node{Id: 100, Label: "(A,1)"}, Target: model.Node{Id: 101, Label: "(A,2)"}},
// 			// Ребро из graph2 между "1" и "2" для вершины "B"
// 			{Id: "1005", Source: model.Node{Id: 102, Label: "(B,1)"}, Target: model.Node{Id: 103, Label: "(B,2)"}},
// 		},
// 	}

// 	score := checker.CheckLexicographicalProduct(true_answer, graph1, graph2)
// 	if score != 100 {
// 		t.Errorf("Expected 100 for correct manual Lexicographical product, got %d", score)
// 	}

// 	// Пример с отсутствующим ребром
// 	missingEdgeAnswer := *true_answer
// 	missingEdgeAnswer.Edges = missingEdgeAnswer.Edges[:len(missingEdgeAnswer.Edges)-1]
// 	score = checker.CheckLexicographicalProduct(&missingEdgeAnswer, graph1, graph2)
// 	if score != 84 {
// 		t.Errorf("Expected less than 100 for answer with missing edge, got %d", score)
// 	}

// 	// Пример с лишним ребром
// 	extraEdgeAnswer := *true_answer
// 	extraEdgeAnswer.Edges = append(extraEdgeAnswer.Edges,
// 		model.Edge{
// 			Id:     "9999",
// 			Source: model.Node{Id: 100, Label: "(A,1)"},
// 			Target: model.Node{Id: 103, Label: "(B,2)"},
// 		})
// 	score = checker.CheckLexicographicalProduct(&extraEdgeAnswer, graph1, graph2)
// 	if score != 84 {
// 		t.Errorf("Expected less than 100 for answer with extra edge, got %d", score)
// 	}

// 	// Пример с пустым графом
// 	emptyGraph := &model.Graph{
// 		Nodes: true_answer.Nodes,
// 		Edges: []model.Edge{},
// 	}
// 	score = checker.CheckLexicographicalProduct(emptyGraph, graph1, graph2)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for empty graph answer, got %d", score)
// 	}
// }

// func TestChecker_CheckHamiltonian(t *testing.T) {
// 	checker := NewChecker()

// 	// Создаем граф с гамильтоновым циклом: квадрат (4 вершины)
// 	graph := &model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"},
// 			{Id: 2, Label: "B"},
// 			{Id: 3, Label: "C"},
// 			{Id: 4, Label: "D"},
// 		},
// 		Edges: []model.Edge{
// 			{Id: "1", Source: model.Node{Id: 1, Label: "A"}, Target: model.Node{Id: 2, Label: "B"}},
// 			{Id: "2", Source: model.Node{Id: 2, Label: "B"}, Target: model.Node{Id: 3, Label: "C"}},
// 			{Id: "3", Source: model.Node{Id: 3, Label: "C"}, Target: model.Node{Id: 4, Label: "D"}},
// 			{Id: "4", Source: model.Node{Id: 4, Label: "D"}, Target: model.Node{Id: 1, Label: "A"}},
// 		},
// 	}

// 	// Ответ с корректным гамильтоновым циклом (цвет у ребер задан)
// 	answerGraph := &model.Graph{
// 		Nodes: graph.Nodes,
// 		Edges: []model.Edge{
// 			{Id: "1", Source: graph.Nodes[0], Target: graph.Nodes[1], Color: "red"},
// 			{Id: "2", Source: graph.Nodes[1], Target: graph.Nodes[2], Color: "red"},
// 			{Id: "3", Source: graph.Nodes[2], Target: graph.Nodes[3], Color: "red"},
// 			{Id: "4", Source: graph.Nodes[3], Target: graph.Nodes[0], Color: "red"},
// 		},
// 	}

// 	// Граф, который не имеет гамильтонова цикла
// 	graphNoHamilton := &model.Graph{
// 		Nodes: []model.Node{
// 			{Id: 1, Label: "A"},
// 			{Id: 2, Label: "B"},
// 			{Id: 3, Label: "C"},
// 		},
// 		Edges: []model.Edge{
// 			{Id: "1", Source: model.Node{Id: 1, Label: "A"}, Target: model.Node{Id: 2, Label: "B"}},
// 			{Id: "2", Source: model.Node{Id: 2, Label: "B"}, Target: model.Node{Id: 3, Label: "C"}},
// 		},
// 	}

// 	// Проверка корректного гамильтонова цикла
// 	score := checker.CheckHamiltonian(graph, true, answerGraph)
// 	if score != 100 {
// 		t.Errorf("Expected 100 for correct Hamiltonian cycle, got %d", score)
// 	}

// 	// Проверка при неверном is_hamiltonian_ans (ответ говорит, что цикл есть, а в графе его нет)
// 	score = checker.CheckHamiltonian(graphNoHamilton, true, answerGraph)
// 	if score != 0 {
// 		t.Errorf("Expected 0 when Hamiltonian answer does not match graph, got %d", score)
// 	}

// 	// Проверка правильного отрицательного результата (нет цикла, ответ говорит нет)
// 	score = checker.CheckHamiltonian(graphNoHamilton, false, nil)
// 	if score != 0 {
// 		t.Errorf("Expected 100 for correct negative Hamiltonian answer, got %d", score)
// 	}

// 	// Проверка неправильного ответа: путь не гамильтонов (не все вершины с двумя инцидентными ребрами)
// 	badAnswerGraph := &model.Graph{
// 		Nodes: graph.Nodes,
// 		Edges: []model.Edge{
// 			{Id: "1", Source: graph.Nodes[0], Target: graph.Nodes[1], Color: "red"},
// 			{Id: "2", Source: graph.Nodes[1], Target: graph.Nodes[2], Color: "red"},
// 		},
// 	}
// 	score = checker.CheckHamiltonian(graph, true, badAnswerGraph)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for incorrect Hamiltonian cycle answer, got %d", score)
// 	}

// 	// Проверка nil answer_graph
// 	score = checker.CheckHamiltonian(graph, true, nil)
// 	if score != 0 {
// 		t.Errorf("Expected 0 for nil answer_graph, got %d", score)
// 	}
// }

// // func TestMinimumSpanningTree(t *testing.T) {
// // 	// Исходный граф
// // 	graph := &model.Graph{
// // 		Nodes: []model.Node{
// // 			{Label: "A"}, {Label: "B"}, {Label: "C"}, {Label: "D"},
// // 		},
// // 		Edges: []model.Edge{
// // 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}, Weight: 1},
// // 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "C"}, Weight: 3},
// // 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}, Weight: 1},
// // 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "D"}, Weight: 4},
// // 			{Source: model.Node{Label: "C"}, Target: model.Node{Label: "D"}, Weight: 2},
// // 		},
// // 	}

// // 	// Ожидаемое минимальное остовное дерево
// // 	trueMST := &model.Graph{
// // 		Nodes: []model.Node{
// // 			{Label: "A"}, {Label: "B"}, {Label: "C"}, {Label: "D"},
// // 		},
// // 		Edges: []model.Edge{
// // 			{Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}, Weight: 1},
// // 			{Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}, Weight: 1},
// // 			{Source: model.Node{Label: "C"}, Target: model.Node{Label: "D"}, Weight: 2},
// // 		},
// // 	}

// // 	// Вызов алгоритма
// // 	edges, result := graph.MinimalSpanningTree()

// // 	// Проверка корректности
// // 	ch := NewChecker()
// // 	ch.CheckM
// // 	if score := expected; score != 100 {
// // 		t.Errorf("expected score 100, got %d", score)
// // 	}
// // }
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////

// func createTestGraph() *model.Graph {
// 	g := &model.Graph{}
// 	for i := 1; i <= 5; i++ {
// 		label := string(rune('A' + i - 1))
// 		g.AddNode(model.Node{Id: i, Label: label})
// 	}
// 	g.AddEdge(model.Edge{Id: "1", Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}, Weight: 1})
// 	g.AddEdge(model.Edge{Id: "2", Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}, Weight: 2})
// 	g.AddEdge(model.Edge{Id: "3", Source: model.Node{Label: "C"}, Target: model.Node{Label: "D"}, Weight: 1})
// 	g.AddEdge(model.Edge{Id: "4", Source: model.Node{Label: "D"}, Target: model.Node{Label: "A"}, Weight: 1})
// 	return g
// }

// func TestAddNode(t *testing.T) {
// 	g := &model.Graph{}
// 	err := g.AddNode(model.Node{Id: 1, Label: "A"})
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}
// 	err = g.AddNode(model.Node{Id: 1, Label: "B"})
// 	if err == nil {
// 		t.Errorf("expected error for duplicate ID")
// 	}
// }

// func TestAddEdge(t *testing.T) {
// 	g := createTestGraph()
// 	err := g.AddEdge(model.Edge{Id: "5", Source: model.Node{Label: "E"}, Target: model.Node{Label: "A"}})
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}
// }

// // func TestUnion1(t *testing.T) {
// // 	g1 := createTestGraph()
// // 	g2 := &model.Graph{}
// // 	for i := 6; i <= 10; i++ {
// // 		label := string(rune('A' + i - 1))
// // 		g2.AddNode(model.Node{Id: i, Label: label})
// // 	}
// // 	g2.AddEdge(model.Edge{Id: "6", Source: model.Node{Label: "F"}, Target: model.Node{Label: "G"}})
// // 	union := g1.Union(g2)
// // 	if len(union.Nodes) != 10 {
// // 		t.Errorf("expected 10 nodes, got %d", len(union.Nodes))
// // 	}
// // }

// // func TestIntersect1(t *testing.T) {
// // 	g1 := createTestGraph()
// // 	g2 := createTestGraph()
// // 	inter := g1.Intersect(g2)
// // 	if len(inter.Nodes) == 0 {
// // 		t.Errorf("expected non-zero intersection")
// // 	}
// // }

// // func TestJoin1(t *testing.T) {
// // 	g1 := createTestGraph()
// // 	g2 := &model.Graph{}
// // 	for i := 6; i <= 10; i++ {
// // 		label := string(rune('A' + i - 1))
// // 		g2.AddNode(model.Node{Id: i, Label: label})
// // 	}
// // 	joined := g1.Join(g2)
// // 	expectedEdges := len(g1.Edges) + len(g2.Nodes)*len(g1.Nodes)
// // 	if len(joined.Edges) < expectedEdges {
// // 		t.Errorf("expected at least %d edges, got %d", expectedEdges, len(joined.Edges))
// // 	}
// // }

// // func TestCartesianProduct1(t *testing.T) {
// // 	g1 := createTestGraph()
// // 	g2 := createTestGraph()
// // 	product := g1.CartesianProduct(g2)
// // 	expectedNodes := len(g1.Nodes) * len(g2.Nodes)
// // 	if len(product.Nodes) != expectedNodes {
// // 		t.Errorf("expected %d nodes, got %d", expectedNodes, len(product.Nodes))
// // 	}
// // }

// // func TestTensorProduct1(t *testing.T) {
// // 	g1 := createTestGraph()
// // 	g2 := createTestGraph()
// // 	product := g1.TensorProduct(g2)
// // 	expectedNodes := len(g1.Nodes) * len(g2.Nodes)
// // 	if len(product.Nodes) != expectedNodes {
// // 		t.Errorf("expected %d nodes, got %d", expectedNodes, len(product.Nodes))
// // 	}
// // }

// // func TestLexicographicalProduct1(t *testing.T) {
// // 	g1 := createTestGraph()
// // 	g2 := createTestGraph()
// // 	product := g1.LexicographicalProduct(g2)
// // 	expectedNodes := len(g1.Nodes) * len(g2.Nodes)
// // 	if len(product.Nodes) != expectedNodes {
// // 		t.Errorf("expected %d nodes, got %d", expectedNodes, len(product.Nodes))
// // 	}
// // }

// func TestEdgeCasesEmptyGraph1(t *testing.T) {
// 	empty := &model.Graph{}
// 	union := empty.Union(empty)
// 	if len(union.Nodes) != 0 {
// 		t.Errorf("expected 0 nodes in union of empty graphs")
// 	}
// 	inter := empty.Intersect(empty)
// 	if len(inter.Nodes) != 0 {
// 		t.Errorf("expected 0 nodes in intersection of empty graphs")
// 	}
// 	join := empty.Join(empty)
// 	if len(join.Nodes) != 0 {
// 		t.Errorf("expected 0 nodes in join of empty graphs")
// 	}
// }

// // func TestHamiltonianCycle1(t *testing.T) {
// // 	g := &model.Graph{}
// // 	for i := 0; i < 5; i++ {
// // 		label := string(rune('A' + i))
// // 		g.AddNode(model.Node{Id: i, Label: label})
// // 	}
// // 	g.AddEdge(model.Edge{Id: "1", Source: model.Node{Label: "A"}, Target: model.Node{Label: "B"}})
// // 	g.AddEdge(model.Edge{Id: "2", Source: model.Node{Label: "B"}, Target: model.Node{Label: "C"}})
// // 	g.AddEdge(model.Edge{Id: "3", Source: model.Node{Label: "C"}, Target: model.Node{Label: "D"}})
// // 	g.AddEdge(model.Edge{Id: "4", Source: model.Node{Label: "D"}, Target: model.Node{Label: "E"}})
// // 	g.AddEdge(model.Edge{Id: "5", Source: model.Node{Label: "E"}, Target: model.Node{Label: "A"}})
// // 	g.AddEdge(model.Edge{Id: "6", Source: model.Node{Label: "A"}, Target: model.Node{Label: "C"}})

// // 	found, path := g.HamiltonianCycle()
// // 	if !found {
// // 		t.Errorf("expected Hamiltonian cycle to exist")
// // 	}
// // 	if len(path) != 5 {
// // 		t.Errorf("expected cycle of length 5, got %d", len(path))
// // 	}
// // }

// func TestMinimalSpanningTree1(t *testing.T) {
// 	g := createTestGraph()
// 	edges, totalWeight := g.MinimalSpanningTree()
// 	if len(edges) != 3 {
// 		t.Errorf("expected 4 edges in MST, got %d", len(edges))
// 	}
// 	if totalWeight != 3 {
// 		t.Errorf("expected MST total weight 3, got %d", totalWeight)
// 	}
// }

// // func TestFindIndependentSets(t *testing.T) {
// // 	g := createTestGraph()
// // 	sets := g.FindIndependentSets()
// // 	if len(sets) == 0 {
// // 		t.Errorf("expected at least one independent set")
// // 	}
// // 	for _, set := range sets {
// // 		for i := 0; i < len(set); i++ {
// // 			for j := i + 1; j < len(set); j++ {
// // 				found, _ := g.FindEdge(set[i], set[j])
// // 				if found {
// // 					t.Errorf("set %v is not independent", set)
// // 				}
// // 			}
// // 		}
// // 	}
// // }

// // func TestFindDominatingSets(t *testing.T) {
// // 	g := createTestGraph()
// // 	sets := g.FindDominatingSets()
// // 	if len(sets) == 0 {
// // 		t.Errorf("expected at least one dominating set")
// // 	}
// // 	// Optional deeper validation
// // }

// // func TestMakeGraphFromAdjLabelMatrix(t *testing.T) {
// // 	matrix := map[string]map[string]int{
// // 		"A": {"B": 1, "C": 0},
// // 		"B": {"A": 1, "C": 1},
// // 		"C": {"B": 1, "A": 0},
// // 	}
// // 	g := model.MakeGraphFromAdjLabelMatrix(matrix)
// // 	if len(g.Nodes) != 3 {
// // 		t.Errorf("expected 3 nodes, got %d", len(g.Nodes))
// // 	}
// // 	if len(g.Edges) != 2 {
// // 		t.Errorf("expected 2 edges, got %d", len(g.Edges))
// // 	}
// // }
