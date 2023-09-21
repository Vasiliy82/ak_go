package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputHistory map[string][]int
var data map[string]string

func main() {
	inputHistory = make(map[string][]int)
	data = make(map[string]string)

	fmt.Println("Счётчик готов к работе!")
	for {
		var str string
		fmt.Println("Введите действие: «название» счётчика_«значение»")
		fmt.Scanf("%s\n", &str)
		if str == "exit" {
			fmt.Println("Завершение работы")
			os.Exit(0)
		}
		if ctrName, ctrIncrementor, err := parse(str); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s: %d\n", ctrName, process(ctrName, ctrIncrementor))
			// второй вариант исполнения будет работать только с положительными числами, и
			// паниковать при вводе отрицательных
			// зато он более "честный", он не хранит явно даже первое введенное значение
			// fmt.Printf("%s: %d\n", ctrName, process2(ctrName, ctrIncrementor))
		}
	}
}

func sum(s []int) int {
	if len(s) > 1 {
		return s[0] + sum(s[1:])
	}
	return s[0]
}

func process(ctrName string, ctrIncrementor int) int {
	// проверять ключ на существование не требуется, т.к. append(nil, something) отлично работает
	inputHistory[ctrName] = append(inputHistory[ctrName], ctrIncrementor)
	return sum(inputHistory[ctrName])
}

func process2(ctrName string, ctrIncrementor int) int {
	// проверять ключ на существование не требуется, т.к. при его отсутствии вернется пустая строка
	data[ctrName] = data[ctrName] + strings.Repeat("*", ctrIncrementor)
	return len(data[ctrName])
}

func parse(s string) (string, int, error) {
	parts := strings.Split(s, "_")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("Некорректный формат")
	}
	ctrName := parts[0]
	ctrIncrementor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, fmt.Errorf("Значение должно быть целым числом")
	}
	return ctrName, ctrIncrementor, nil
}

func parse_wo_split(s string) (string, int, error) {
	ctrName := ""
	for idx, v := range s {
		// собираем имя счетчика, пока не встретится _
		if v == '_' {
			ctrName = s[:idx]
			s = s[idx+1:]
			break
		}
	}
	for _, v := range s {
		// проверяем значение на ошибку типа "некорректный формат"
		if v == '_' {
			return "", 0, fmt.Errorf("Некорректный формат")
		}
	}
	ctrIncrementor, err := strconv.Atoi(s)
	if err != nil {
		return "", 0, fmt.Errorf("Значение должно быть целым числом")
	}
	return ctrName, ctrIncrementor, nil
}
