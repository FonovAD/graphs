package main

import (
	"sync"
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

func BenchmarkInsertTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := insertTest()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkInsertTask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := insertTask()
		if err != nil {
			panic(err)
		}
	}
}

///// PARALLEL

func BenchmarkPAuthUser(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			err := checkToken()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkPGetTests(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			err := getTestsSend()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkPCheckResults(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			err := checkResult()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkPGetTasksFromTest(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			err := getTasksFromTest()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkPInsertTest(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			err := insertTest()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkPInsertTask(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			err := insertTask()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
