package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(3)

	go callDatabase(&wg)
	go callApi(&wg)
	go processInternal(&wg)

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Tempo total de execução: %v\n", elapsed)
}

func callDatabase(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Println("Finalizado callDatabase")
	wg.Done()
}

func callApi(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	fmt.Println("Finalizado callApi")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Println("Finalizado processInternal")
	wg.Done()
}
