package service

type StudentService interface {
}

type studentService struct {
}

func NewStudentService() StudentService {
	return &studentService{}
}
