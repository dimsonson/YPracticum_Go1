package main

import "fmt"

func main() {

	fmt.Println("test")
}

type APIClient interface {
	GetData(query string) (Response, error)
}

type Response struct {
	Text       string
	StatusCode int
}

type MockAPIClient interface {
	APIClient
	SetResponse(resp Response, err error)
}

type Mock struct {
	Resp Response
	Err  error
}

func (db *Mock) GetData(query string) (Response, error) {
	return db.Resp, db.Err
}

func (db *Mock) SetResponse(resp Response, err error) {
	db.Resp = resp
	db.Err = err
}

/*
Определена API со следующим интерфейсом:
type APIClient interface {
    GetData(query string) (Response, error)
}

type Response struct {
    Text       string
    StatusCode int
}

Реализуйте тип Mock с интерфейсом MockAPIClient:
type MockAPIClient interface {
    APIClient
    SetResponse(resp Response, err error)
}
*/
