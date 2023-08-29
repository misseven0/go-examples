package main

import (
	"errors"
	"log"
	"net/http"
	"testgin/rpc"
	// "github.com/ethereum/go-ethereum/rpc"
)

type CalculatorService struct{}

func (s *CalculatorService) Add(a, b int) int {
	return a + b
}

func (s *CalculatorService) Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

// curl -X POST  -H 'Content-Type: application/json' -d '{"jsonrpc":"2.0","id":"id","method":"calculator_add","params":[1, 2]}'   http://127.0.0.1:8964/jsonrpc
// curl -X POST  -H 'Content-Type: application/json' -d '{"jsonrpc":"2.0","id":"id","method":"rpc_modules","params":[]}'   http://127.0.0.1:8964/jsonrpc
//{"jsonrpc":"2.0","id":"id","result":{"cac":"1.0","rpc":"1.0"}}

func main() {
	calculator := new(CalculatorService)
	server := rpc.NewServer()

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, req *http.Request) {
		log.Println("->>jsonrpc:")
		server.ServeHTTP(w, req)
	})
	server.RegisterName("calculator", calculator)
	err := http.ListenAndServe(":8964", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	// server.st
}
