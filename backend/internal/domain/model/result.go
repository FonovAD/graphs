package model

type GroupLabResult struct {
	LabID    int64              `json:"lab_id"`
	Students []StudentLabResult `json:"students"`
}

type StudentLabResult struct {
	UserID        int64          `json:"user_id"`
	FIO           string         `json:"fio"`
	OverallScore  int            `json:"overall_score"`
	ModuleResults []ModuleResult `json:"module_results"`
}

type ModuleResult struct {
	ModuleID    int64  `json:"module_id"`
	ModuleName  string `json:"module_name"`
	ModuleScore int    `json:"module_score"`
}
