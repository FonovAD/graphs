package model

import "github.com/shopspring/decimal"

type ModuleLab struct {
	ModuleLabID int64           `db:"module_lab_id"`
	Weight      decimal.Decimal `db:"weight"`
	LabID       int64           `db:"lab_id"`
	ModuleID    int64           `db:"module_id"`
}

type LabWithModules struct {
	LabID   int64
	Modules []ModulesInLab
}

type ModulesInLab struct {
	ModuleLabID int64  `db:"module_lab_id" json:"moduleLabId"`
	LabID       int64  `db:"lab_id" json:"labId"`
	ModuleId    int64  `db:"module_id" json:"moduleId"`
	ModuleType  string `db:"type" json:"moduleType"`
}
