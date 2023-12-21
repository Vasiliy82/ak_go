package main

import (
	"fmt"
	"runtime"
)

const numMessages = 1_000_000

var globalBool interface{}

var globalStruct interface{}
var globalCtr int = 0

func main() {
	// Измерение памяти для канала bool значений
	runtime.GC()
	memStatsBeforeBool := &runtime.MemStats{}
	runtime.ReadMemStats(memStatsBeforeBool)

	boolChan := make(chan bool, numMessages)
	for i := 0; i < numMessages; i++ {
		boolChan <- true
	}

	runtime.GC()
	memStatsAfterBool := &runtime.MemStats{}
	runtime.ReadMemStats(memStatsAfterBool)

	for i := 0; i < numMessages; i++ {
		x := <-boolChan
		globalBool = x
		globalCtr += 1
	}

	// Измерение памяти для канала struct{} значений
	runtime.GC()
	memStatsBeforeStruct := &runtime.MemStats{}
	runtime.ReadMemStats(memStatsBeforeStruct)

	structChan := make(chan struct{}, numMessages)
	for i := 0; i < numMessages; i++ {
		structChan <- struct{}{}
	}

	runtime.GC()
	memStatsAfterStruct := &runtime.MemStats{}
	runtime.ReadMemStats(memStatsAfterStruct)

	for i := 0; i < numMessages; i++ {
		x := <-structChan
		globalStruct = x
		globalCtr += 1
	}

	// Расчеты и вывод результатов
	memoryUsedBool := memStatsAfterBool.Alloc - memStatsBeforeBool.Alloc
	memoryUsedStruct := memStatsAfterStruct.Alloc - memStatsBeforeStruct.Alloc

	fmt.Printf("Memory used for bool channel: %d bytes\n", memoryUsedBool)
	fmt.Printf("Memory used for struct{} channel: %d bytes\n", memoryUsedStruct)

	percentageSaved := 100.0 * (float64(memoryUsedBool) - float64(memoryUsedStruct)) / float64(memoryUsedBool)
	fmt.Printf("Memory saved by using struct{}: %.2f%%\n", percentageSaved)

	sysMemoryUsedBool := memStatsAfterBool.Sys - memStatsBeforeBool.Sys
	sysMemoryUsedStruct := memStatsAfterStruct.Sys - memStatsBeforeStruct.Sys

	fmt.Printf("System memory used for bool channel: %d bytes\n", sysMemoryUsedBool)
	fmt.Printf("System memory used for struct{} channel: %d bytes\n", sysMemoryUsedStruct)

	sysPercentageSaved := 100.0 * (float64(sysMemoryUsedBool) - float64(sysMemoryUsedStruct)) / float64(sysMemoryUsedBool)
	fmt.Printf("System memory saved by using struct{}: %.2f%%\n", sysPercentageSaved)

	fmt.Println(globalCtr)
}
