package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("CPUs: ", runtime.NumCPU())

	// EXEMPLO #3
	var contador int64 = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched() // informar ao processador para PARAR e ir fazer outra rotina
			//fmt.Println("Contador:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Contador:", contador)

	// EXEMPLO #2
	// fmt.Println("CPUs: ", runtime.NumCPU())
	// fmt.Println("Goroutines: ", runtime.NumGoroutine())
	// contador := 0
	// totaldeoroutines := 1000
	// wg.Add(totaldeoroutines)
	// var mu sync.Mutex
	//	for i := 0; i < totaldeoroutines; i++ {
	//		//fmt.Println("Goroutine: ", (i+1), "Contador:", contador)
	//		go func() {
	//			//mu.Lock()
	//			v := contador
	//			runtime.Gosched() // informar ao processador para PARAR e ir fazer outra rotina
	//			v++
	//			contador = v
	//			//mu.Unlock()
	//			wg.Done()
	//		}()
	//	//fmt.Println("Goroutine: ", (i+1), "Contador:", contador, "NumGoroutine: ", runtime.NumGoroutine())
	//}
	//wg.Wait()
	//fmt.Println("Contador:", contador)
	//fmt.Println("Goroutines.Wait: ", runtime.NumGoroutine())

	// EXEMPLO #1
	// fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	// wg.Add(2) // o programa vai ter 1 Goroutines
	// fmt.Println("teste1()")
	// go teste1()
	// fmt.Println("NumGoroutine.teste1: ", runtime.NumGoroutine())
	// fmt.Println("teste2()")
	// go teste2()
	// fmt.Println("NumGoroutine.teste2: ", runtime.NumGoroutine())
	// fmt.Println("wg.Wait()")
	// wg.Wait() // esperar tds as Goroutines terminar
	// fmt.Println("NumGoroutine.Wait: ", runtime.NumGoroutine())
	//fmt.Println("FIM")

}

func teste1() {
	fmt.Printf("#1 - Hora Atual: %v\n", time.Now().UTC().Format(time.RFC3339))
	time.Sleep(time.Second * 3)
	fmt.Printf("#1 - Hora Atual: %v\n", time.Now().UTC().Format(time.RFC3339))
	wg.Done()
}

func teste2() {
	fmt.Printf("#2 - Hora Atual: %v\n", time.Now().UTC().Format(time.RFC3339))
	time.Sleep(time.Second * 1)
	fmt.Printf("#2 - Hora Atual: %v\n", time.Now().UTC().Format(time.RFC3339))
	wg.Done()
}
