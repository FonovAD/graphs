package model

type Module struct {
	ModuleId    int64  `db:"module_id" json:"moduleId"`
	ModuleType  string `db:"type" json:"moduleType"`
	Description string `db:"description" json:"description"`
}
