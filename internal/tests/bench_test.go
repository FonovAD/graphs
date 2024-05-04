package main

import (
	"testing"
)

func BenchmarkAuthUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		authUser(b)
	}
}

func BenchmarkGetTests(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getTests(b)
	}
}

func authUser(b *testing.B) {
	err := checkToken()
	if err != nil {
		panic(err)
	}
}

func getTests(b *testing.B) {
	err := getTestsSend()
	if err != nil {
		panic(err)
	}
}
