package storage

const (
	opInsertStudent             = "infra.storage.pg.teacher.repo.InsertStudent"
	opGetModules                = "infra.storage.pg.teacher.repo.GetModules"
	opCreateLab                 = "infra.storage.pg.teacher.repo.CreateLab"
	opGetLabInfo                = "infra.storage.pg.teacher.repo.GetLabInfo"
	opRemoveUserLab             = "infra.storage.pg.teacher.repo.RemoveUserLab"
	opUpdateLab                 = "infra.storage.pg.teacher.repo.UpdateLab"
	opInsertUserLab             = "infra.storage.pg.teacher.repo.InsertUserLab"
	opInsertModuleLab           = "infra.storage.pg.teacher.repo.InsertModuleLab"
	opRemoveModuleFromLab       = "infra.storage.pg.teacher.repo.RemoveModuleFromLab"
	opInsertLabStudentGroup     = "infra.storage.pg.teacher.repo.InsertLabToStudentGroup"
	opSelectNonExistingUserLabs = "infra.storage.pg.teacher.repo.SelectNonExistingUserLabs"
	opSelectExistingUserLabs    = "infra.storage.pg.teacher.repo.SelectExistingUserLabs"
	opSelectModulesFromLab      = "infra.storage.pg.teacher.repo.SelectModulesFromLab"
	opSelectGroups              = "infra.storage.pg.teacher.repo.SelectGroups"
	opSelectTeacher             = "infra.storage.pg.teacher.repo.SelectTeacher"
	opCreateTask                = "infra.storage.pg.teacher.repo.InsertTask"
	opUpdateTask                = "infra.storage.pg.teacher.repo.UpdateTask"
	opSelectTasksByModule       = "infra.storage.pg.teacher.repo.SelectTasksByModule"
)
