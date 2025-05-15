package usecase

import (
	"context"
	model "golang_graphs/backend/internal/domain/model/user"
	teacherrepository "golang_graphs/backend/internal/domain/teacher/repository"
	teacherservice "golang_graphs/backend/internal/domain/teacher/service"
	userservice "golang_graphs/backend/internal/domain/user/service"

	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type TeacherUseCase interface {
	CreateUser(ctx context.Context, userDTO CreateUserDTO) (int64, error)
}

type teacherUseCase struct {
	teacherRepo    teacherrepository.TeacherRepository
	userService    userservice.UserService
	teacherService teacherservice.TeacherService
}

func NewTeacherUseCase(repo teacherrepository.TeacherRepository, userService userservice.UserService, teacherService teacherservice.TeacherService) TeacherUseCase {
	return &teacherUseCase{
		teacherRepo:    repo,
		userService:    userService,
		teacherService: teacherService,
	}
}

func (u *teacherUseCase) CreateUser(ctx context.Context, userDTO CreateUserDTO) (int64, error) {
	if err := validateCreateUser(userDTO); err != nil {
		return -1, err
	}

	salt := u.teacherService.RandomString()

	hash, err := hashPassword(userDTO.Password, salt)
	if err != nil {
		return -1, errors.Wrap(err, "hash password")
	}

	user := &model.User{
		DateRegistration: time.Now(),
		Email:            userDTO.Email,
		Password:         hash,
		FirstName:        userDTO.FirstName,
		LastName:         userDTO.LastName,
		Role:             "student",
		PasswordSalt:     salt,
	}

	userFromDB, err := u.teacherRepo.InsertUser(ctx, user)
	if err != nil {
		return -1, errors.Wrap(err, "insert user")
	}

	return userFromDB.Id, nil
}

// Hash password using the bcrypt hashing algorithm
func hashPassword(password, salt string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password + salt)

	// Hash password with bcrypt's default cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	return string(hashedPasswordBytes), err
}

func validateCreateUser(userDTO CreateUserDTO) error {
	if len(userDTO.FirstName) < 1 {
		return ErrShortFirstname
	}

	if len(userDTO.LastName) < 1 {
		return ErrShortLastname
	}

	if len(userDTO.Password) < 8 {
		return ErrShortPassword
	}

	if len(userDTO.Email) < 4 {
		return ErrShortEmail
	}

	if len(userDTO.Email) > 100 {
		return ErrLongEmail
	}

	return nil
}
