package main

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const response = `{"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaXJzdF9uYW1lIjoic3RyaW5nIiwiaWQiOiIwIiwibGFzdF9uYW1lIjoic3RyaW5nIiwicm9sZSI6InN0dWRlbnQifQ.Oaoy0IN8QIMbvPb_8Kfhmn1MJTcWRajPFhGENv8zmIg"}`

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

func getTestsSend() error {
	client := resty.New().SetBaseURL("http://localhost:8080/get_tests")

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

func auth() (string, error) {
	client := resty.New().SetBaseURL("http://localhost:8080/auth_user").SetHeader("Content-Type", "application/json")

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
	client := resty.New().SetBaseURL("http://localhost:8080/create_user").SetHeader("Content-Type", "application/json")

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
