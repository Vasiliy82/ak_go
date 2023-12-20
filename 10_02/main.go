package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// функция каждый раз будет возвращать накопленное в счетчике значение, а также новую версию функции, которая будет "помнить" это накопленное значение
type SelfReturningFunc func(int) (int, SelfReturningFunc)

var counters map[string]SelfReturningFunc

func newCounter(ctrIncrementor int) SelfReturningFunc {
	return func(value int) (int, SelfReturningFunc) {
		// это даже не рекурсия, ведь функция сама себя не вызывает, она каждый раз делает новую версию себя, если можно так выразиться
		return value + ctrIncrementor, newCounter(value + ctrIncrementor)
	}
}

func process(ctrName string, ctrIncrementor int) int {
	// если счетчик встретился впервые, надо инициализировать его нулевым значением
	if _, found := counters[ctrName]; !found {
		counters[ctrName] = newCounter(0)
	}
	// получаем накопленный результат и новую версию функции, в которую "вложены" все предыдущие функции
	result, newF := counters[ctrName](ctrIncrementor)
	counters[ctrName] = newF
	return result
}

// парсить строку по символу _ можно и без strings.Split, но я полагаю, что тут задача не на это
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

func main() {
	counters = make(map[string]SelfReturningFunc)

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
		}
	}
}
