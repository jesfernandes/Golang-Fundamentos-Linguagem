package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2) //especificar manualmente quantas gourotines serao executadas.

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done() // -1 waitGroup
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1 waitGroup
	}()

	waitGroup.Wait() // vai esperar as gourotines chegar a 0 na var waitGroup
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
