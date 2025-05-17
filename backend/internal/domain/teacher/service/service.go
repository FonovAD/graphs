package teacherservice

import "github.com/samber/lo"

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

const n = 5

type TeacherService interface {
	RandomString() string
}

type teacherService struct {
	n int
}

func NewTeacherService() TeacherService {
	return &teacherService{n: n}
}

func (ts *teacherService) RandomString() string {
	return lo.RandomString(ts.n, charset)
}
