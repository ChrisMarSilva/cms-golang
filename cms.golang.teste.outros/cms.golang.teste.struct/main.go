package main

import (
  "context"
  "fmt"
  "math/rand"
  "time"
    "runtime"
    "fmt"
    "time"
    "context"
)

type person struct {
	name    string
	cpf     cpf
	friends *friends
}

type cpf int

func (c cpf) IsValid() error  {
    if c <= 0 {
        return errors.New("Id must be greater than 0")
    }

 return nil
}

func (c cpf) IsValid() bool {

}

func (c cpf) String() string {
    return fmt.Sprintf("Foo Says: %s", c)
}

func (v *cpf) Error() string {
    return "ddd"
}

type friends struct {
	data []string
}

func (f friends) Add(name string) string {

}
func (f friends) Remove(name string) string {

}

func (b *board) BoardRepresentation() string {
	var buffer = &bytes.Buffer{}
	for _, l :- range b.squares(){
		l.addTo(buffer)
	}
	return buffer.String()
}


função DoSomething (ctx context.Context, arg Arg) error {
	// ... use ctx ...
}

func main() {

	ctx, cancel := context.Background()
  ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	ctx, cancel := context.TODO()
	ctx := context.WithValue(context.Background(), key, "test")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(150)*time.Millisecond)


	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	select {
  case <-ctx.Done():
    fmt.Println("Time to return")
  case sleeptime := <-sleeptimeChan:
    fmt.Println("Slept for ", sleeptime, "ms")
  }


	req := &http.Request{URL: parsedURL}
	req = req.WithContext(ctx)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)


select {
		case r := <-ch:
			cancel()
			return r
		case <-time.After(21 * time.Millisecond):
		}



	fmt.Println("ok")
}


func monitor2(ctx context.Context, number int) {
    for {
        select {
        case v := <- ctx.Done():
            fmt.Printf("monitor: %v, the received channel value is: %v, ending\n", number,v)
            return
        default:
            fmt.Printf("monitor: %v in progress...\n", number)
            time.Sleep(2 * time.Second)
        }
    }
}
func monitor1(ctx context.Context, number int) {
    for {
        go monitor2(ctx, number)
        select {
        case v := <- ctx.Done():
            // this branch is only reached when the ch channel is closed, or when data is sent(either true or false)
            fmt.Printf("monitor: %v, the received channel value is: %v, ending\n", number, v)
            return
        default:
            fmt.Printf("monitor: %v in progress...\n", number)
            time.Sleep(2 * time.Second)
        }
    }
}
func main() {
    var ctx context.Context = nil
    var cancel context.CancelFunc = nil
    ctx, cancel = context.WithCancel(context.Background())
    for i := 1; i <= 5; i = i + 1 {
        go monitor1(ctx, i)
    }
    time.Sleep(1 * time.Second)
    // close all gourtines
    cancel()
    // waiting 10 seconds, if the screen does not display <monitor: xxxx in progress>, all goroutines have been shut down
    time.Sleep(10 * time.Second)
    println(runtime.NumGoroutine())
    println("main program exit!!!!")
}


package main
import (
    "runtime"
    "fmt"
    "time"
    "context"
)
func monitor2(ctx context.Context, index int) {
    for {
        select {
        case v := <- ctx.Done():
            fmt.Printf("monitor2: %v, the received channel value is: %v, ending\n", index, v)
            return
        default:
            fmt.Printf("monitor2: %v in progress...\n", index)
            time.Sleep(2 * time.Second)
        }
    }
}
func monitor1(ctx context.Context, index int) {
    for {
        go monitor2(ctx, index)
        select {
        case v := <- ctx.Done():
            // this branch is only reached when the ch channel is closed, or when data is sent(either true or false)
            fmt.Printf("monitor1: %v, the received channel value is: %v, ending\n", index, v)
            return
        default:
            fmt.Printf("monitor1: %v in progress...\n", index)
            time.Sleep(2 * time.Second)
        }
    }
}
func main() {
    var ctx01 context.Context = nil
    var ctx02 context.Context = nil
    var cancel context.CancelFunc = nil
    ctx01, cancel = context.WithCancel(context.Background())
    ctx02, cancel = context.WithDeadline(ctx01, time.Now().Add(1 * time.Second)) // If it's WithTimeout, just change this line to "ctx02, cancel = context.WithTimeout(ctx01, 1 * time.Second)"
    defer cancel()
    for i := 1; i <= 5; i = i + 1 {
        go monitor1(ctx02, i)
    }
    time.Sleep(5 * time.Second)
    if ctx02.Err() != nil {
        fmt.Println("the cause of cancel is: ", ctx02.Err())
    }
    println(runtime.NumGoroutine())
    println("main program exit!!!!")
}


package main
import (
    "runtime"
    "fmt"
    "time"
    "context"
)
func monitor(ctx context.Context, index int) {
    for {
        select {
        case <- ctx.Done():
            // this branch is only reached when the ch channel is closed, or when data is sent(either true or false)
            fmt.Printf("monitor %v, end of monitoring. \n", index)
            return
        default:
            var value interface{} = ctx.Value("Nets")
            fmt.Printf("monitor %v, is monitoring %v\n", index, value)
            time.Sleep(2 * time.Second)
        }
    }
}
func main() {
    var ctx01 context.Context = nil
    var ctx02 context.Context = nil
    var cancel context.CancelFunc = nil
    ctx01, cancel = context.WithCancel(context.Background())
    ctx02, cancel = context.WithTimeout(ctx01, 1 * time.Second)
    var ctx03 context.Context = context.WithValue(ctx02, "Nets", "Champion") // key: "Nets", value: "Champion"
  
    defer cancel()
    for i := 1; i <= 5; i = i + 1 {
        go monitor(ctx03, i)
    }
    time.Sleep(5 * time.Second)
    if ctx02.Err() != nil {
        fmt.Println("the cause of cancel is: ", ctx02.Err())
    }
    println(runtime.NumGoroutine())
    println("main program exit!!!!")
}