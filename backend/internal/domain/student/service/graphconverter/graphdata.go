package service

type NodeDataJSON struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Color  string `json:"color"`
	Weight string `json:"weight"`
}

type PositionJSON struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type NodeJSON struct {
	NodeData NodeDataJSON `json:"data"`
	Position PositionJSON `json:"position"`
}

type NodesJSON struct {
	NodeArr []NodeJSON
}

type EdgeDataJSON struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Label  string `json:"label"`
	Color  string `json:"color"`
	Id     string `json:"id"`
}

type EdgeJSON struct {
	EdgeData EdgeDataJSON `json:"data"`
	Source   NodeJSON     `json:"source"`
	Target   NodeJSON     `json:"target"`
}

type EdgesJSON struct {
	EdgeArr []EdgeJSON
}
