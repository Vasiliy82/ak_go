/*Задача №4. Нужно создать контекст со значениями
some key1: some value1
some key2: some value2
Контекст следует передать в функцию, которая выведет значения some
key1 и some key2 в stdout.*/

package main

import (
	"context"
	"fmt"
)

type ctxKey string

type ctxKey2 int
type myContextData map[string]string

func main() {
	// вариант 1
	ctx := context.Background()
	var key ctxKey = "some key1"
	var value string = "some value1"
	ctx = context.WithValue(ctx, key, value)
	key = "some key2"
	value = "some value2"
	ctx = context.WithValue(ctx, key, value)
	fmt.Println("Вариант 1")
	do(ctx)

	// вариант 2
	var key2 ctxKey2 = 1
	value2 := myContextData{"some key1": "some value1", "some key2": "some value2"}
	ctx2 := context.Background()
	ctx2 = context.WithValue(ctx2, key2, value2)
	fmt.Println("Вариант 2")
	do2(ctx2)
}

func do(ctx context.Context) {
	var key ctxKey = "some key1"
	fmt.Println(key, ":", ctx.Value(key))
	key = "some key2"
	fmt.Println(key, ":", ctx.Value(key))
}

func do2(ctx context.Context) {
	var key ctxKey2 = 1
	value := ctx.Value(key)
	data, ok := value.(myContextData)
	if ok == true {
		fmt.Println("some key1", ":", data["some key1"])
		fmt.Println("some key2", ":", data["some key2"])
	} else {
		fmt.Println("что-то пошло не так")
	}
}
