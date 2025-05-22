package usecase

import (
	"context"
	model "golang_graphs/backend/internal/domain/model"
	teacherrepository "golang_graphs/backend/internal/domain/teacher/repository"
	teacherservice "golang_graphs/backend/internal/domain/teacher/service"
	userservice "golang_graphs/backend/internal/domain/user/service"

	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const limit = 20

type TeacherUseCase interface {
	CreateStudent(ctx context.Context, userDTO *CreateUserDTOIn) (*CreateUserDTOOut, error)
	GetModules(ctx context.Context) (*GetModulesDTOOut, error)
	CreateLab(ctx context.Context, in *CreateLabDTOIn) (*CreateLabDTOOut, error)
	GetLabInfo(ctx context.Context, in *GetLabInfoDTOIn) (*GetLabInfoDTOOut, error)
	RemoveUserLab(ctx context.Context, in *RemoveUserLabDTOIn) (*RemoveUserLabDTOOut, error)
	UpdateLabInfo(ctx context.Context, in *UpdateLabInfoDTOIn) error
	AssignLab(ctx context.Context, in *AssignLabDTOIn) (*AssignLabDTOOut, error)
	AssignLabGroup(ctx context.Context, in *AssignLabGroupDTOIn) (*AssignLabGroupDTOOut, error)
	AddModuleLab(ctx context.Context, in *AddModuleLabDTOIn) (*AddModuleLabDTOOut, error)
	RemoveModuleLab(ctx context.Context, in *RemoveModuleLabDTOIn) (*RemoveModuleLabDTOOut, error)
	GetNonAssignedLabs(ctx context.Context, in *GetNonAssignedLabsDTOIn) (*GetNonAssignedLabsDTOOut, error)
	GetAssignedLabs(ctx context.Context, in *GetAssignedLabsDTOIn) (*GetAssignedLabsDTOOut, error)
	GetLabModules(ctx context.Context, in *GetLabModulesDTOIn) (*GetLabModulesDTOOut, error)
	GetGroups(ctx context.Context) (*GetGroupsDTOOut, error)
	GetTeacher(ctx context.Context, user *model.User) (*model.Teacher, error)
	AuthToken(ctx context.Context, token string) (*AuthTokenDTOOut, error)
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

func (u *teacherUseCase) CreateStudent(ctx context.Context, userDTO *CreateUserDTOIn) (*CreateUserDTOOut, error) {
	if err := validateCreateStudent(userDTO); err != nil {
		return nil, err
	}

	salt := u.teacherService.RandomString()

	hash, err := hashPassword(userDTO.Password, salt)
	if err != nil {
		return nil, errors.Wrap(err, "hash password")
	}

	user := &model.User{
		DateRegistration: time.Now(),
		Email:            userDTO.Email,
		Password:         hash,
		FirstName:        userDTO.FirstName,
		LastName:         userDTO.LastName,
		FatherName:       userDTO.FatherName,
		Role:             userDTO.Role,
		PasswordSalt:     salt,
	}

	userFromDB, err := u.teacherRepo.InsertUser(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "insert user")
	}

	return &CreateUserDTOOut{UserID: userFromDB.ID}, nil
}

func (u *teacherUseCase) GetModules(ctx context.Context) (*GetModulesDTOOut, error) {
	modules, err := u.teacherRepo.GetModules(ctx)
	if err != nil {
		return nil, err
	}
	return &GetModulesDTOOut{
		Modules: modules,
	}, nil
}

func (u *teacherUseCase) CreateLab(ctx context.Context, in *CreateLabDTOIn) (*CreateLabDTOOut, error) {
	lab := &model.Lab{
		Name:             in.LabName,
		Description:      in.Description,
		RegistrationDate: time.Now(),
		TeacherID:        in.TeacherID,
	}
	lab.SetDurationMinutes(in.Duration)

	out, err := u.teacherRepo.CreateLab(ctx, lab)
	if err != nil {
		return nil, err
	}

	return &CreateLabDTOOut{
		LabID: out.ID,
	}, nil
}

func (u *teacherUseCase) GetLabInfo(ctx context.Context, in *GetLabInfoDTOIn) (*GetLabInfoDTOOut, error) {
	lab := &model.Lab{
		ID: in.LabID,
	}
	out, err := u.teacherRepo.GetLabInfo(ctx, lab)
	if err != nil {
		return nil, err
	}

	return &GetLabInfoDTOOut{
		LabName:          out.Name,
		Description:      out.Description,
		Duration:         out.GetDurationMinutes(),
		RegistrationDate: out.RegistrationDate,
		TeacherID:        out.TeacherID,
		TeacherFIO:       out.TeacherFIO,
	}, nil
}

func (u *teacherUseCase) RemoveUserLab(ctx context.Context, in *RemoveUserLabDTOIn) (*RemoveUserLabDTOOut, error) {
	userLab := &model.UserLab{
		UserID: in.UserID,
		LabID:  in.LabID,
	}
	out, err := u.teacherRepo.RemoveUserLab(ctx, userLab)
	if err != nil {
		return nil, err
	}

	return &RemoveUserLabDTOOut{
		UserLabID: out.UserLabID,
	}, nil
}

func (u *teacherUseCase) UpdateLabInfo(ctx context.Context, in *UpdateLabInfoDTOIn) error {
	lab := &model.Lab{
		ID:          in.LabID,
		Name:        in.LabName,
		Description: in.Description,
	}
	lab.SetDurationMinutes(in.Duration)

	if err := u.teacherRepo.UpdateLab(ctx, lab); err != nil {
		return err
	}

	return nil
}

func (u *teacherUseCase) AssignLab(ctx context.Context, in *AssignLabDTOIn) (*AssignLabDTOOut, error) {
	userLab := &model.UserLab{
		UserID:         in.UserID,
		LabID:          in.LabID,
		AssignmentDate: time.Now(),
		StartTime:      in.StartTime,
		TeacherID:      in.AssigneID,
		Deadline:       in.Deadline,
	}
	out, err := u.teacherRepo.InsertUserLab(ctx, userLab)
	if err != nil {
		return nil, err
	}

	return &AssignLabDTOOut{
		UserLabID: out.UserLabID,
	}, nil
}

func (u *teacherUseCase) AssignLabGroup(ctx context.Context, in *AssignLabGroupDTOIn) (*AssignLabGroupDTOOut, error) {
	groupLab := &model.UserLabGroup{
		LabID:          in.LabID,
		AssignmentDate: in.AssignmentDate,
		StartTime:      in.StartTime,
		TeacherID:      in.AssigneID,
		Deadline:       in.Deadline,
	}
	out, err := u.teacherRepo.InsertLabToStudentGroup(ctx, groupLab)
	if err != nil {
		return nil, err
	}

	return &AssignLabGroupDTOOut{
		LabID: out.LabID,
	}, nil
}

func (u *teacherUseCase) AddModuleLab(ctx context.Context, in *AddModuleLabDTOIn) (*AddModuleLabDTOOut, error) {
	moduleLab := &model.ModuleLab{
		LabID:    in.LabID,
		ModuleID: in.ModuleID,
		Weight:   in.Weight,
	}
	out, err := u.teacherRepo.InsertModuleLab(ctx, moduleLab)
	if err != nil {
		return nil, err
	}

	return &AddModuleLabDTOOut{
		ModuleLabID: out.ModuleLabID,
	}, nil
}

func (u *teacherUseCase) RemoveModuleLab(ctx context.Context, in *RemoveModuleLabDTOIn) (*RemoveModuleLabDTOOut, error) {
	moduleLab := &model.ModuleLab{
		LabID:    in.LabID,
		ModuleID: in.ModuleID,
	}
	out, err := u.teacherRepo.RemoveModuleFromLab(ctx, moduleLab)
	if err != nil {
		return nil, err
	}

	return &RemoveModuleLabDTOOut{
		ModuleLabID: out.ModuleLabID,
	}, nil
}

func (u *teacherUseCase) GetNonAssignedLabs(ctx context.Context, in *GetNonAssignedLabsDTOIn) (*GetNonAssignedLabsDTOOut, error) {
	pagination := model.Pagination{
		Limit:  limit,
		Offset: limit * (in.Page - 1),
	}
	out, err := u.teacherRepo.SelectNonExistingUserLabs(ctx, pagination)
	if err != nil {
		return nil, err
	}

	return &GetNonAssignedLabsDTOOut{
		Labs: out,
	}, nil
}

func (u *teacherUseCase) GetAssignedLabs(ctx context.Context, in *GetAssignedLabsDTOIn) (*GetAssignedLabsDTOOut, error) {
	pagination := model.Pagination{
		Limit:  limit,
		Offset: limit * (in.Page - 1),
	}
	out, err := u.teacherRepo.SelectExistingUserLabs(ctx, pagination)
	if err != nil {
		return nil, err
	}

	return &GetAssignedLabsDTOOut{
		Labs: out,
	}, nil
}

func (u *teacherUseCase) GetLabModules(ctx context.Context, in *GetLabModulesDTOIn) (*GetLabModulesDTOOut, error) {
	lab := &model.Lab{
		ID: in.LabID,
	}
	out, err := u.teacherRepo.SelectModulesFromLab(ctx, lab)
	if err != nil {
		return nil, err
	}

	return &GetLabModulesDTOOut{
		LabID:   in.LabID,
		Modules: out,
	}, nil
}

func (u *teacherUseCase) AuthToken(ctx context.Context, token string) (*AuthTokenDTOOut, error) {
	user, err := u.userService.ParseToken(token)
	if err != nil {
		return nil, ErrParseToken
	}

	if user.Role != "teacher" {
		return nil, ErrNoPermissions
	}

	out, err := u.GetTeacher(ctx, user)
	if err != nil {
		return nil, err
	}

	return &AuthTokenDTOOut{
		UserID:    out.UserID,
		TeacherID: out.ID,
	}, nil
}

func (u *teacherUseCase) GetTeacher(ctx context.Context, user *model.User) (*model.Teacher, error) {
	out, err := u.teacherRepo.SelectTeacher(ctx, user)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (u *teacherUseCase) GetGroups(ctx context.Context) (*GetGroupsDTOOut, error) {
	groups, err := u.teacherRepo.SelectGroups(ctx)
	if err != nil {
		return nil, err
	}
	return &GetGroupsDTOOut{
		Groups: groups,
	}, nil
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

func validateCreateStudent(userDTO *CreateUserDTOIn) error {
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
