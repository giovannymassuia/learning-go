package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// using everything we've seen so far
// > go run main.go 01001000

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		defer req.Body.Close()

		fmt.Printf("Status: %s\n", req.Status)

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		fmt.Println("CEP:", data)

		file, err := os.Create("cep.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))

		fmt.Println("File created successfully.")
		fmt.Println("Cidade:", data.Localidade)
	}
}
