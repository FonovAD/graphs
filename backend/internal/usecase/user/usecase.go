package usecase

import (
	"context"
	"fmt"
	userrepository "golang_graphs/backend/internal/domain/user/repository"
	userservice "golang_graphs/backend/internal/domain/user/service"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	AuthUser(ctx context.Context, userDTO AuthUserDTO) (string, error)
}

type userUseCase struct {
	userRepo    userrepository.UserRepository
	userService userservice.UserService
}

func NewUserUseCase(repo userrepository.UserRepository, userService userservice.UserService) UserUseCase {
	return &userUseCase{
		userRepo:    repo,
		userService: userService,
	}
}

func (u *userUseCase) AuthUser(ctx context.Context, userDTO AuthUserDTO) (string, error) {

	if err := validateAuthUser(userDTO.Email, userDTO.Password); err != nil {
		return "", err
	}
	user, err := u.userRepo.SelectUserByEmail(ctx, userDTO.Email)
	if err != nil {
		return "", err
	}

	passwd := userDTO.Password + user.PasswordSalt

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwd)); err != nil {
		return "", fmt.Errorf("incorrect nick or password %w", err)
	}

	token, err := u.userService.CreateToken(user)
	if err != nil {
		return "", errors.Wrap(err, "error creating token")
	}

	return token, nil

}

func validateAuthUser(email, password string) error {
	if len(password) < 8 {
		return ErrShortPassword
	}

	if len(password) > 50 {
		return ErrLongPassword
	}

	if len(email) < 4 {
		return ErrShortEmail
	}

	if len(email) > 100 {
		return ErrLongEmail
	}

	return nil
}
