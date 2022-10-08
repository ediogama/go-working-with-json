package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int
	Balance int
}

func main() {
	account := Account{Number: 1, Balance: 100}
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}
}
