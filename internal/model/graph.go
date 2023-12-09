package model

type Node struct {
	Id int `json:"id"`
}

type Link struct {
	Source Node `json:"source"`
	Target Node `json:"target"`
}

type Graph struct {
	Nodes []Node
	Links []Link
}

type Component struct {
	Component int `json:"component"`
}

type IsEuler struct {
	IsEuler bool `json:"isEuler"`
}

type IsBipartition struct {
	IsBipartition bool `json:"isBipartition"`
}
