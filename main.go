package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  sync.Mutex
		wg sync.WaitGroup
		i  int
	)

	const total = 10000

	wg.Add(total)

	for x := 0; x < total; x++ {
		go func() {
			defer wg.Done()
			m.Lock()
			i++
			m.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Valor final de i:", i)
}
