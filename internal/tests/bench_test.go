package main

import (
	"testing"
)

func BenchmarkAuthUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := checkToken()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGetTests(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := getTestsSend()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkCheckResults(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := checkResult()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGetTasksFromTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := getTasksFromTest()
		if err != nil {
			panic(err)
		}
	}
}
