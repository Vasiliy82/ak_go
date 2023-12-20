package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type contract struct {
	Number   int    `json:"number"`
	SignDate string `json:"sign_date,omitempty"`
	Landlord string `json:"landlord"`
	Tenat    string `json:"tenat"`
}

func main() {
	// Задача №2. Необходимо представить в виде json структуру contract

	contract := contract{
		Number:   2,
		Landlord: "Остап",
		Tenat:    "Шура",
	}

	if res, err := json.Marshal(contract); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%+v", string(res))
	}

}
