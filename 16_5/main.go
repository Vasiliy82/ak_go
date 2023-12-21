/*
Задача №5. «Метеоролог» Необходимо реализовать интерфейс
type Meteo interface {
ReadTemp() string
ChangeTemp(v string)
}
Речь про температуру окружающей среды. ReadTemp читает
температуру, ChangeTemp изменяет температуру. Код должен быть
потокобезопасным, т.е. при работе с температурой нескольких
параллельных горутин не должно возникать состояние гонки.
*/
package main

import (
	"fmt"
	"sync"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type Thermometer struct {
	tempValue string
	mu        sync.RWMutex
}

func (t *Thermometer) ReadTemp() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tempValue
}
func (t *Thermometer) ChangeTemp(v string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tempValue = v

}

func main() {

	var t Meteo = &Thermometer{}

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		i := i
		wg.Add(1)
		go func() {
			// множество горутин, каждая из которых максимально быстро будет менять значение температуры
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				t.ChangeTemp(fmt.Sprintf("%d", i*1000+j))
			}
			fmt.Println("routine changer ", i, "finished")
		}()
	}

	wg.Wait()

	fmt.Println("Final Temperature value is: ", t.ReadTemp())
}
