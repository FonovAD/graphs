package models

import "math"

// type Node struct {
// 	Id int `json:"id"`
// }

// type Link struct {
// 	Source Node `json:"source"`
// 	Target Node `json:"target"`
// }

type Component struct {
	Component int `json:"component"`
}

type IsEuler struct {
	IsEuler bool `json:"isEuler"`
}

type IsBipartition struct {
	IsBipartition bool `json:"isBipartition"`
}

type Node struct {
	Id     int
	Label  string
	Color  string
	Weight int

	x float64
	y float64
}

type Edge struct {
	Id     int
	Source Node
	Target Node
	Label  string
	Color string
	Weight int
}

type Graph struct {
	Nodes []Node
	Edges []Edge
}

// Матрица смежности вершин, но вместо индексов - Label
func (g *Graph) NodeLabelAdjacentMatrix() map[string]map[string]int {
	matrix := make(map[string]map[string]int)
	for _, node1 := range g.Nodes {
		for _, node2 := range g.Nodes {
			matrix[node1.Label][node2.Label] = 0
		}
	}
	for _, edge := range g.Edges {
		matrix[edge.Source.Label][edge.Target.Label] = 1
		matrix[edge.Target.Label][edge.Source.Label] = 1
	}
	return matrix
}

// Матрица смежности вершин, но вместо индексов - Id
func (g *Graph) NodeIdAdjacentMatrix() map[int]map[int]int {
	matrix := make(map[int]map[int]int)
	for _, node1 := range g.Nodes {
		for _, node2 := range g.Nodes {
			matrix[node1.Id][node2.Id] = 0
		}
	}
	for _, edge := range g.Edges {
		matrix[edge.Source.Id][edge.Target.Id] = 1
		matrix[edge.Target.Id][edge.Source.Id] = 1
	}
	return matrix
}

// Матрица смежности ребер, но вместо индексов - Label
func (g *Graph) EdgeLabelAdjacentMatrix() map[string]map[string]int {
	matrix := make(map[string]map[string]int)
	for _, edge1 := range g.Edges {
		for _, edge2 := range g.Edges {
			matrix[edge1.Label][edge2.Label] = 0
		}
	}
	for _, edge1 := range g.Edges {
		for _, edge2 := range g.Edges {
			if (edge1.Source.Id == edge2.Source.Id || edge1.Source.Id == edge2.Target.Id || edge1.Target.Id == edge2.Source.Id || edge1.Target.Id == edge2.Target.Id) && (edge1.Id != edge2.Id) {
				matrix[edge1.Label][edge2.Label] = 1
				matrix[edge2.Label][edge1.Label] = 1
			}
		}
	}
	return matrix
}

// Матрица смежности ребер, но вместо индексов - Id
func (g *Graph) EdgeIdAdjacentMatrix() map[int]map[int]int {
	matrix := make(map[int]map[int]int)
	for _, edge1 := range g.Edges {
		for _, edge2 := range g.Edges {
			matrix[edge1.Id][edge2.Id] = 0
		}
	}
	for _, edge1 := range g.Edges {
		for _, edge2 := range g.Edges {
			if (edge1.Source.Id == edge2.Source.Id || edge1.Source.Id == edge2.Target.Id || edge1.Target.Id == edge2.Source.Id || edge1.Target.Id == edge2.Target.Id) && (edge1.Id != edge2.Id) {
				matrix[edge1.Id][edge2.Id] = 1
				matrix[edge2.Id][edge1.Id] = 1
			}
		}
	}
	return matrix
}

// Реализация алгоритма Дейкстры, возвращает (минимальный путь, матрицу расстояний)
// Все вершины должны быть достижимы, расстояния не меньше 0
func (g *Graph) MinPath(source Node, target Node, edges_have_weights bool) (int, map[string]int) {
	visited := make(map[string]bool)
	weights_path := make(map[string]int)
	for _, node := range g.Nodes {
		visited[node.Label] = false
		weights_path[node.Label] = math.MaxInt
	}
	weights_path[source.Label] = 0
	weights_matrix := make(map[string]map[string]int)
	for _, node1 := range g.Nodes {
		for _, node2 := range g.Nodes {
			weights_matrix[node1.Label][node2.Label] = 0
		}
	}
	for _, edge := range g.Edges {
		if edges_have_weights {
			weights_matrix[edge.Source.Label][edge.Target.Label] = edge.Weight
			weights_matrix[edge.Target.Label][edge.Source.Label] = edge.Weight
		} else {
			weights_matrix[edge.Source.Label][edge.Target.Label] = 1
			weights_matrix[edge.Target.Label][edge.Source.Label] = 1
		}
	}
	min_node := "1"
	for min_node != "" {
		min_val := math.MaxInt
		min_node = ""
		for _, node := range g.Nodes {
			if !visited[node.Label] && weights_path[node.Label] < min_val {
				min_val = weights_path[node.Label]
				min_node = node.Label
			}
		}
		if min_node != "" {
			for _, node := range g.Nodes {
				if weights_matrix[min_node][node.Label] > 0 {
					temp := min_val + weights_matrix[min_node][node.Label]
					if temp < weights_path[node.Label] {
						weights_path[node.Label] = temp
					}
				}
			}
			visited[min_node] = true
		}
	}
	return weights_path[target.Label], weights_path
}

// Возвращает матрицу кратчайших расстояний между каждой парой вершин, на диагонали MaxInt
func (g *Graph) DistanceMatrix(edges_have_weights bool) map[string]map[string]int {
	matrix := make(map[string]map[string]int)
	for _, node := range g.Nodes {
		_, matrix[node.Label] = g.MinPath(node, node, edges_have_weights)
		matrix[node.Label][node.Label] = 0
	}
	return matrix
}
