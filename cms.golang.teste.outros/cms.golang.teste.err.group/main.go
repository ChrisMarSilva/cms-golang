package main

import (
    "fmt"
    "sync"
    "log"
    "math/rand"
	"golang.org/x/sync/errgroup"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.err.group
// go get -u golang.org/x/sync/errgroup
// go mod tidy

// go run main.go

func main() {
	teste1()
	teste2()
}

func teste1() {
	wg := &sync.WaitGroup{}
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func (task int)  {
            defer wg.Done()
            fmt.Println("WaitGroup - Processing task: ", task)

        }(i)
    }
    wg.Wait()
}

func teste2() {
	eg := &errgroup.Group{}
    for i := 0; i < 10; i++ {
        task := i
        eg.Go(func() error {
            return Task(task)
        })
    }
    if err := eg.Wait(); err != nil {
        log.Fatal("Error", err)
    }
    fmt.Println("Completed successfully!")
}

func Task(task int) error {
    if rand.Intn(10) == task {
        return fmt.Errorf("Task %v failed", task)
    }
    fmt.Printf("ErrGroup - Task %v completed\n", task)
    return nil
}