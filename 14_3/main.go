package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type contract struct {
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}
type contracts struct {
	List []contract `xml:"contract"`
}

func main() {
	// Задача №3. Необходимо распарсить xml

	str := `<?xml version="1.0" encoding="UTF-8"?>
	<contracts>
	<contract>
	<number>1</number>
	<sign_date>2023-09-02</sign_date>
	<landlord>Остап</landlord>
	<tenat>Шура</tenat>
	</contract>
	<contract>
	<number>2</number>
	<sign_date>2023-09-03</sign_date>
	<landlord>Бендер</landlord>
	<tenat>Балаганов</tenat>
	</contract>
	</contracts>`

	contracts := contracts{}

	if err := xml.Unmarshal([]byte(str), &contracts); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v", contracts)
}
