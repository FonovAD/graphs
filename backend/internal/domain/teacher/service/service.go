package teacherservice

import "github.com/samber/lo"

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type TeacherService interface {
	RandomString() string
}

type teacherService struct {
	n int
}

func NewTeacherService(n int) TeacherService {
	return &teacherService{n: n}
}

func (ts *teacherService) RandomString() string {
	return lo.RandomString(ts.n, charset)
}
