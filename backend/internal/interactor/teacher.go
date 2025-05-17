package interactor

import (
	repository "golang_graphs/backend/internal/domain/teacher/repository"
	service "golang_graphs/backend/internal/domain/teacher/service"
	storage "golang_graphs/backend/internal/infrastructure/storage/pg/teacher"
	handler "golang_graphs/backend/internal/presenter/http/handler/teacher"
	usecase "golang_graphs/backend/internal/usecase/teacher"
)

func (i *interactor) NewTeacherRepository() repository.TeacherRepository {
	return storage.NewTeacherRepository(i.conn)
}

func (i *interactor) NewTeacherService() service.TeacherService {
	return service.NewTeacherService()
}

func (i *interactor) NewTeacherUseCase() usecase.TeacherUseCase {
	return usecase.NewTeacherUseCase(i.NewTeacherRepository(), i.NewUserService(), i.NewTeacherService())
}

func (i *interactor) NewTeacherHandler() handler.TeacherHandler {
	return handler.NewTeacherHandler(i.NewTeacherUseCase())
}
