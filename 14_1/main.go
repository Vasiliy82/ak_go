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
	// Задача №1 Необходимо распарсить json
	str := `{"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов"}`

	c := contract{}

	if err := json.Unmarshal([]byte(str), &c); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v", c)
}
