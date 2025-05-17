package interactor

import (
	repository "golang_graphs/backend/internal/domain/student/repository"
	service "golang_graphs/backend/internal/domain/student/service"
	graphconverter "golang_graphs/backend/internal/domain/student/service/graphconverter"
	taskcheck "golang_graphs/backend/internal/domain/student/service/taskcheck"
	storage "golang_graphs/backend/internal/infrastructure/storage/pg/student"
	handler "golang_graphs/backend/internal/presenter/http/handler/student"
	usecase "golang_graphs/backend/internal/usecase/student"
)

func (i *interactor) NewStudentRepository() repository.StudentRepository {
	return storage.NewStudentRepository(i.conn)
}

func (i *interactor) NewStudentService() service.StudentService {
	return service.NewStudentService()
}

func (i *interactor) NewTaskCheckService() taskcheck.Checker {
	return taskcheck.NewChecker()
}

func (i *interactor) NewGraphConverterService() graphconverter.GraphConverter {
	return graphconverter.NewGraphConverter()
}

func (i *interactor) NewStudentUseCase() usecase.StudentUseCase {
	return usecase.NewStudentUseCase(
		i.NewStudentRepository(),
		i.NewStudentService(),
		i.NewTaskCheckService(),
		i.NewGraphConverterService(),
		i.NewUserService(),
	)
}

func (i *interactor) NewStudentHandler() handler.StudentHandler {
	return handler.NewStudentHandler(i.NewStudentUseCase())
}
