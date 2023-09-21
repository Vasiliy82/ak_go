package main

import (
	"fmt"
)

func main() {

	for str := ""; ; {
		fmt.Println("Введите строку:")
		// ограничение: если хотим читать строки, содержащие пробелы, придется поискать другой способ
		fmt.Scanf("%s\n", &str)
		// из условия задачи не указано явно, нужно ли применять алгоритм валидации и вывода результата к строке exit.
		// Предполагаем, что не нужно.
		if str == "exit" {
			break
		}
		// fmt.Println("DEBUG: получена строка", str)
		if err := valid(str); err != nil {
			// Если нам интересно, что случилось на самом деле
			fmt.Println("DEBUG: ", err)
			fmt.Println("Строка содержит не только символы A-Z")
			continue
		}
		fmt.Println("Строка содержит только символы A-Z")
	}
	fmt.Println("Завершение работы")
}

func valid(v string) error {
	runePosition := 0
	for _, char := range v {
		// на индекс, выдаваемый оператором range, ориентироваться нельзя, т.к. в UNICODE бывают двух- и более байтовые символы
		runePosition++
		if char < 'A' || char > 'Z' {
			return fmt.Errorf("Символ %q, не принадлежащий множеству [A-Z], найден в позиции %d", char, runePosition)
		}
	}
	return nil
}
