package create_random_string

import "math/rand"

type creator struct {
	n int
}

type Creator interface {
	RandomString() string
}

func New(n int) Creator {
	return &creator{n: n}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (c *creator) RandomString() string {
	b := make([]byte, c.n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
