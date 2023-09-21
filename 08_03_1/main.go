package main

import "fmt"

func main() {

	s := []int{1, 5, 3, 7, 8, 10, 27, 13, 22}

	delete := func(num int) {
		// если if уже можно использовать
		if num < 0 || num >= len(s) {
			fmt.Println("error: slice s:", s, ", ничего мы не удалим, т.к. элемент выходит за пределы слайса:", num)
			return
		}
		fmt.Println("debug: slice s:", s, ", сейчас мы удалим элемент", s[num])
		if num > 0 {
			copy(s[num:], s[num+1:])
			s = s[:len(s)-1]
		} else {
			s = s[1:]
		}
	}

	fmt.Println("slice s:", s)
	delete(3)
	fmt.Println("slice s:", s)
	delete(7)
	fmt.Println("slice s:", s)
	delete(0)
	fmt.Println("slice s: ", s)
	delete(100)
	fmt.Println("slice s:", s)
}
