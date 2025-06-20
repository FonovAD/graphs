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
	Weight string `json:"weight"`
}

type EdgeJSON struct {
	EdgeData EdgeDataJSON `json:"data"`
	Source   NodeJSON     `json:"source"`
	Target   NodeJSON     `json:"target"`
	Color    string       `json:"color"`
}

type EdgesJSON struct {
	EdgeArr []EdgeJSON
}

type GraphDataJSON struct {
	NodeArr []NodeJSON `json:"nodes"`
	EdgeArr []EdgeJSON `json:"edges"`
}

type ModulesDataJSON struct {
	Type        int           `json:"type"`
	GraphData   GraphDataJSON `json:"data"`
	InputValue1 string        `json:"inputValue1"`
	InputValue2 string        `json:"inputValue2"`
	InputValue3 string        `json:"inputValue3"`
	InputValue4 string        `json:"inputValue4"`
	InputValue5 string        `json:"inputValue5"`
	InputValue6 string        `json:"inputValue6"`
}

type ModulesJSON struct {
	Modules []ModulesDataJSON `json:"modules"`
}
