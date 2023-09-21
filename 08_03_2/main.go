package main

import "fmt"

func main() {

	s := []int{1, 5, 3, 7, 8, 10, 27, 13, 22}

	delete := func(num int) {
		// если if нельзя
		_ = num > 0 && func() bool {
			copy(s[num-1:], s[num:])
			s = s[:len(s)-1]
			return true
		}()

		_ = num == 0 && func() bool {
			s = s[1:]
			return true
		}()
	}

	fmt.Println("slice s: ", s)
	delete(3)
	fmt.Println("slice s: ", s)
	delete(7)
	fmt.Println("slice s: ", s)
	delete(0)
	fmt.Println("slice s: ", s)

}
