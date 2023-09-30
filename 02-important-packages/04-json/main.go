package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number       int    `json:"-"` // ignore this field
	Balance      int    `json:"b"`
	customerName string // not exported
}

func main() {
	account := Account{Number: 123, Balance: 1000, customerName: "John"}

	// convert to json using Marshal
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(res))

	// convert to json using encoder
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(account)
	if err != nil {
		panic(err)
	}

	// convert from json using Unmarshal
	jsonString := `{"Number":123,"b":1000,"customerName":"John"}` // note: customerName is not imported, because it's not exported
	var account2 Account
	err = json.Unmarshal([]byte(jsonString), &account2) // note: it's using `&` to pass the pointer
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", account2)
}
