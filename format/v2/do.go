package v2

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"sort"
)

type DocRecord struct {
	Name  string `json:"number" xml:"Name"`
	Age   int    `json:"age" xml:"Age"`
	Email string `json:"email" xml:"Email"`
}

type patients struct {
	Recs []DocRecord `xml:"Patient"`
}

func Do(fnInput, fnOutput string) error {
	// Версия 2.1.0
	//
	/* Модуль должен прочитать файл со следующим
	   содержимым:
	   {"name":"Ёжик","age":10,"email":"ezh@mail.ru"}
	   {"name":"Зайчик","age":2,"email":"zayac@mail.ru"}
	   {"name":"Лисичка","age":3,"email":"alice@mail.ru"}

	   v1.0.0 должна создавать файл с содержимым:
	   [{"name":"Ёжик","age":10,"email":"ezh@mail.ru"},
	   {"name":"Зайчик","age":2,"email":"zayac@mail.ru"},
	   {"name":"Лисичка","age":3,"email":"alice@mail.ru"}]

		v1.1.0 должна сортировать данные по полю age по возрастанию:
		[{"name":"Зайчик","age":2,"email":"zayac@mail.ru"},
		{"name":"Лисичка","age":3,"email":"alice@mail.ru"},
		{"name":"Ёжик","age":10,"email":"ezh@mail.ru"}]

		v2.0.0 должна создавать файл с содержимым:
		<?xml version="1.0" encoding="UTF-8"?>
		<patients>
		<Patient>
		<Name>Ёжик</Name>
		<Age>10</Age>
		<Email>ezh@mail.ru</Email>
		</Patient>
		<Patient>
		<Name>Зайчик</Name>
		<Age>2</Age>
		<Email>zayac@mail.ru</Email>
		</Patient>
		<Patient>
		<Name>Лисичка</Name>
		<Age>3</Age>
		<Email>alice@mail.ru</Email>
		</Patient>
		</patients>

		v2.1.0 должна сортировать данные по полю age по возрастанию:
		<?xml version="1.0" encoding="UTF-8"?>
		<patients>
		<Patient>
		<Name>Зайчик</Name>
		<Age>2</Age>
		<Email>zayac@mail.ru</Email>
		</Patient>
		<Patient>
		<Name>Лисичка</Name>
		<Age>3</Age>
		<Email>alice@mail.ru</Email>
		</Patient>
		<Patient>
		<Name>Ёжик</Name>
		<Age>10</Age>
		<Email>ezh@mail.ru</Email>
		</Patient>
		</patients>
	*/

	// первым делом открываем входной и выходной файлы. Если возникнет ошибка на этом этапе, то и читать ничего не придется,
	// при большом объеме данных юзер раньше получит ошибку
	fin, err := os.Open(fnInput)

	if err != nil {
		return fmt.Errorf("Ошибка открытия файла: %w", err)
	}
	defer fin.Close()

	fout, err := os.Create(fnOutput)
	if err != nil {
		return fmt.Errorf("Ошибка создания файла: %w", err)
	}
	defer fout.Close()

	dec := json.NewDecoder(fin)
	records := make([]DocRecord, 0, 10)

	for dec.More() {
		var d DocRecord
		err := dec.Decode(&d)
		if err != nil {
			return fmt.Errorf("Ошибка чтения записи из файла: %w", err)
		}
		records = append(records, d)
	}
	log.Printf("%+v", records)

	// v1.1.0 добавилась сортировка
	// v2.0.0 пока убрал сортировку
	// v2.1.0 вернул сортировку
	sort.Slice(records, func(i, j int) bool { return records[i].Age < records[j].Age })

	p := patients{Recs: records}

	enc := xml.NewEncoder(fout)
	enc.Indent("", "    ")
	if err := enc.Encode(p); err != nil {
		return fmt.Errorf("Ошибка записи в файл: %w", err)
	}

	return nil
}
