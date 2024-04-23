package create_random_string

import (
	"github.com/samber/lo"
)

type creator struct {
	n int
}

type Creator interface {
	RandomString() string
}

func New(n int) Creator {
	return &creator{n: n}
}

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func (c *creator) RandomString() string {
	return lo.RandomString(c.n, charset)
}
