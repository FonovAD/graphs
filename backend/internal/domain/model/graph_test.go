package model

import (
	"testing"
)

var (
	nodeA = Node{Id: 1, Label: "A"}
	nodeB = Node{Id: 2, Label: "B"}
	nodeC = Node{Id: 3, Label: "C"}
	nodeD = Node{Id: 4, Label: "D"}
	nodeE = Node{Id: 5, Label: "E"}
)

func setupBasicGraph() *Graph {
	g := &Graph{}
	_ = g.AddNode(nodeA)
	_ = g.AddNode(nodeB)
	_ = g.AddNode(nodeC)
	_ = g.AddNode(nodeD)
	_ = g.AddNode(nodeE)
	_ = g.AddEdge(Edge{Id: "1", Source: nodeA, Target: nodeB, Weight: 1, Label: "1"})
	_ = g.AddEdge(Edge{Id: "2", Source: nodeB, Target: nodeC, Weight: 2})
	_ = g.AddEdge(Edge{Id: "3", Source: nodeC, Target: nodeD, Weight: 3})
	_ = g.AddEdge(Edge{Id: "4", Source: nodeD, Target: nodeE, Weight: 4})
	_ = g.AddEdge(Edge{Id: "5", Source: nodeE, Target: nodeA, Weight: 5})
	return g
}

func TestGraphMethods(t *testing.T) {
	g := setupBasicGraph()

	// Check orientation
	if g.IsOriented() {
		t.Error("Expected graph to be undirected by default")
	}
	g.MakeOriented()
	if !g.IsOriented() {
		t.Error("Expected graph to be marked as oriented")
	}

	// Node checks
	if !g.IsNodeById(1) || !g.IsNodeByLabel("A") {
		t.Error("Expected node A to exist")
	}
	if g.IsNodeByLabel("Z") {
		t.Error("Expected node Z to not exist")
	}

	// Edge checks
	if !g.IsEdgeById("1") || !g.IsEdgeByLable("1") {
		// skip AB if not labeled, but make sure ID works
		t.Error("Expected edge with ID 1 to exist")
	}

	// Matrix checks
	labelMatrix := g.NodeLabelAdjacentMatrix()
	if labelMatrix["A"]["B"] != 1 {
		t.Error("Expected adjacency A-B in label matrix")
	}
	idMatrix := g.NodeIdAdjacentMatrix()
	if idMatrix[1][2] != 1 {
		t.Error("Expected adjacency 1-2 in ID matrix")
	}
	labelEdgeMatrix := g.EdgeLabelAdjacentMatrix()
	if len(labelEdgeMatrix) == 0 {
		t.Error("Expected edge label adjacency matrix")
	}
	idEdgeMatrix := g.EdgeIdAdjacentMatrix()
	if len(idEdgeMatrix) == 0 {
		t.Error("Expected edge ID adjacency matrix")
	}

	// MinPath
	dist, _ := g.MinPath(nodeA, nodeC, true)
	if dist != 3 {
		t.Errorf("Expected min path A-C = 3, got %d", dist)
	}

	// DistanceMatrix
	distMatrix := g.DistanceMatrix(true)
	if distMatrix["A"]["D"] != 6 {
		t.Errorf("Expected distance A-D = 6, got %d", distMatrix["A"]["D"])
	}

	// Edge adjacency
	e1 := Edge{Source: nodeA, Target: nodeB}
	e2 := Edge{Source: nodeB, Target: nodeC}
	if !g.IsEdgesAdjacent(e1, e2) {
		t.Error("Expected edges e1 and e2 to be adjacent")
	}
	e3 := Edge{Source: nodeD, Target: nodeE}
	e4 := Edge{Source: nodeA, Target: nodeC}
	if g.IsEdgesAdjacent(e3, e4) {
		t.Error("Expected edges e3 and e4 to not be adjacent")
	}

	// Intersect, Union, Join
	g2 := &Graph{}
	_ = g2.AddNode(Node{Id: 10, Label: "T"})
	_ = g2.AddNode(Node{Id: 10, Label: "R"})
	_ = g2.AddEdge(Edge{Id: "X", Source: nodeC, Target: nodeD, Weight: 3})

	inter := g.Intersect(g2)
	if len(inter.Nodes) == 0 {
		t.Error("Expected non-empty intersection")
	}
	union := g.Union(g2)
	if len(union.Nodes) < len(g.Nodes) {
		t.Error("Expected union to include all nodes")
	}
	joined := g.Join(g2)
	if len(joined.Edges) <= len(union.Edges) {
		t.Error("Expected joined graph to have more edges")
	}

	// Graph products
	cart := g.CartesianProduct(g2)
	if len(cart.Nodes) == 0 {
		t.Error("Expected Cartesian product nodes")
	}
	tens := g.TensorProduct(g2)
	if len(tens.Nodes) == 0 {
		t.Error("Expected Tensor product nodes")
	}
	lex := g.LexicographicalProduct(g2)
	if len(lex.Nodes) == 0 {
		t.Error("Expected Lexicographical product nodes")
	}

	// HamiltonianCycle
	hc, path := g.HamiltonianCycle()
	if !hc || len(path) != 5 {
		t.Errorf("Expected Hamiltonian cycle with 5 nodes, got %v", path)
	}

	// MST
	mstEdges, total := g.MinimalSpanningTree()
	if len(mstEdges) != 4 || total != 10 {
		t.Errorf("Expected MST with 4 edges and weight 10, got %d", total)
	}

	// MakeGraphFromAdjLabelMatrix
	adj := map[string]map[string]int{
		"X": {"Y": 1},
		"Y": {"X": 1},
	}
	newG := MakeGraphFromAdjLabelMatrix(adj)
	if len(newG.Nodes) != 2 || len(newG.Edges) != 1 {
		t.Error("Expected graph with 2 nodes and 1 edge")
	}
}
