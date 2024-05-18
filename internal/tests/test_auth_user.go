package main

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	response = `{"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaXJzdF9uYW1lIjoic3RyaW5nIiwiaWQiOiIwIiwibGFzdF9uYW1lIjoic3RyaW5nIiwicm9sZSI6InN0dWRlbnQifQ.Oaoy0IN8QIMbvPb_8Kfhmn1MJTcWRajPFhGENv8zmIg"}`
	host     = "http://localhost:8080/"
	token    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaXJzdF9uYW1lIjoic3RyaW5nIiwiaWQiOiIwIiwibGFzdF9uYW1lIjoic3RyaW5nIiwicm9sZSI6InN0dWRlbnQifQ.Oaoy0IN8QIMbvPb_8Kfhmn1MJTcWRajPFhGENv8zmIg"
)

func main() {
	// Получить токен нужно в случае сброса базы данных
	// err := GetToken()
	// if err != nil {
	//	 panic(err)
	// }
	err := getTestsSend()
	if err != nil {
		panic(err)
	}
}

func checkToken() error {
	token, err := auth()
	if err != nil {
		return err
	}
	if token != response {
		return errors.New("invalid token")
	}
	return nil
}

func insertTest() error {
	client := resty.New().SetBaseURL(host + "insert_test")
	res, err := client.R().SetHeader("Content-Type", "application/json").SetBody(`{
  "test": {
    "description": "string",
    "end": "2022-01-02T00:00:00Z",
    "id": 0,
    "name": "string",
    "start": "2022-01-02T00:00:00Z"
  }
}
`).Post("")
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		fmt.Println(res.StatusCode())
		fmt.Println(res.String())
		return errors.New("code not 200")
	}
	return nil
}

func insertTask() error {
	client := resty.New().SetBaseURL(host + "insert_task")
	res, err := client.R().SetHeader("Content-Type", "application/json").SetBody(`
{
  "task": {
    "answer": "string",
    "data": "string",
    "description": "string",
    "id": 0,
    "maxGrade": 0,
    "name": "string",
    "testID": 1
  }
}

`).Post("")
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		fmt.Println(res.StatusCode())
		fmt.Println(res.String())
		return errors.New("code not 200")
	}
	return nil
}

func getTestsSend() error {
	client := resty.New().SetBaseURL(host + "get_tests")

	res, err := client.R().Get("")
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		fmt.Println(res.StatusCode())
		fmt.Println(res.String())
		return errors.New("code not 200")
	}
	return nil
}

func checkResult() error {
	client := resty.New().SetBaseURL(host + "check_results")
	res, err := client.R().SetHeader("Content-Type", "application/json").SetBody(
		fmt.Sprintf(
			`{
  "token": "%s"
}`, token)).Post("")
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		fmt.Println(res.StatusCode())
		fmt.Println(res.String())
		return errors.New("code not 200")
	}
	return nil
}

func getTasksFromTest() error {
	client := resty.New().SetBaseURL(host + "get_tasks_from_test")
	res, err := client.
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"test_id": 0}`).
		Post("")
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		fmt.Println(res.StatusCode())
		fmt.Println(res.String())
		return errors.New("code not 200")
	}
	return nil
}

func auth() (string, error) {
	client := resty.New().SetBaseURL(host+"auth_user").SetHeader("Content-Type", "application/json")

	res, err := client.R().SetBody(`{
  "email": "b@b.b",
  "password": "string"
}`).Post("")
	if err != nil {
		return "", err
	}

	return res.String(), nil
}

func GetToken() (string, error) {
	client := resty.New().SetBaseURL(host+"create_user").SetHeader("Content-Type", "application/json")

	res, err := client.R().SetBody(`{
  "email": "b@b.b",
  "firstName": "string",
  "lastName": "string",
  "password": "string"
}`).Post("")
	if err != nil {
		return "", err
	}

	return res.String(), nil
}
