package interactor

import (
	repository "golang_graphs/backend/internal/domain/user/repository"
	userservice "golang_graphs/backend/internal/domain/user/service"
	storage "golang_graphs/backend/internal/infrastructure/storage/pg/user"
	userhandler "golang_graphs/backend/internal/presenter/http/handler/user"
	usecase "golang_graphs/backend/internal/usecase/user"
)

func (i *interactor) NewUserRepository() repository.UserRepository {
	return storage.NewUserRepository(i.conn)
}

func (i *interactor) NewUserService() userservice.UserService {
	return userservice.NewUserService()
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository(), i.NewUserService())
}

func (i *interactor) NewUserHandler() userhandler.UserHandler {
	return userhandler.NewUserHandler(i.NewUserUseCase())
}
