package main

import (
    "context"
    "fmt"
    "time"
    "errors"
    //"math/rand"
   // "runtime"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.struct
// go get -u XXXXXXXXXXXX
// go mod tidy

// go run main.go

func main() {

    user1 := Usuario{
        Nome:    "Pessoa #1",
        Idade:    10,
        Endereco: &Endereco{
            Rua:  "Rua Pessoa #1",
            Numero: 100,
            CEP:    "16.402-128",
        },
    }
	fmt.Println(user1)
	fmt.Println("")

    user2 := user1
    user2.Nome = "Pessoa #2"
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println("")

    user1.Endereco.Rua = "Rua Pessoa #11"
    user2.Endereco.Rua = "Rua Pessoa #22"
	fmt.Println(user1, user1.Endereco)
	fmt.Println(user2, user2.Endereco)
	fmt.Println("")

    user2.Endereco = &Endereco{Rua: user1.Endereco.Rua, Numero: user1.Endereco.Numero, CEP: user1.Endereco.CEP}
    user2.Endereco.Rua = "Rua Pessoa #33"
	fmt.Println(user1, user1.Endereco)
	fmt.Println(user2, user2.Endereco)
	fmt.Println("")

    user2.Endereco = user1.DeepCopyAddress()
    user2.Endereco.Rua = "Rua Pessoa #44"
	fmt.Println(user1, user1.Endereco)
	fmt.Println(user2, user2.Endereco)
	fmt.Println("")

    user3 := user1.DeepCopy()
    user3.Nome = "Pessoa #3"
    user3.Endereco.Rua = "Rua Pessoa #55"
	fmt.Println(user1, user1.Endereco)
	fmt.Println(user3, user3.Endereco)
	fmt.Println("")

}

type Usuario struct {
    Nome     string
    Idade    int
    Endereco *Endereco
}

func (u *Usuario) DeepCopy() *Usuario {
    return &Usuario{
        Nome:     u.Nome, 
        Idade:    u.Idade, 
        Endereco: u.DeepCopyAddress(),
    }
}

func (u *Usuario) DeepCopyAddress() *Endereco {
    return &Endereco{
        Rua:    u.Endereco.Rua, 
        Numero: u.Endereco.Numero, 
        CEP:    u.Endereco.CEP,
    }
}

type Endereco struct {
    Rua    string
    Numero int
    CEP    string
}

func main_old() {

	ctx := context.Background()
    // ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	// ctx, cancel := context.TODO()
	// ctx := context.WithValue(context.Background(), key, "test")
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))
	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(150)*time.Millisecond)
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	select {
    case <-ctx.Done():
        fmt.Println("Time to return")
    // case sleeptime := <-sleeptimeChan:
    //     fmt.Println("Slept for ", sleeptime, "ms")
    }

	// req := &http.Request{URL: parsedURL}
	// req = req.WithContext(ctx)
	// ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
    // select {
    // case r := <-ch:
    //     cancel()
    //     return r
    // case <-time.After(21 * time.Millisecond):
    // }

    // var ctx01 context.Context = nil
    // var ctx02 context.Context = nil
    // var cancel context.CancelFunc = nil
    // ctx01, cancel = context.WithCancel(context.Background())
    // ctx02, cancel = context.WithTimeout(ctx01, 1 * time.Second)
    // var ctx03 context.Context = context.WithValue(ctx02, "Nets", "Champion") // key: "Nets", value: "Champion"
    // defer cancel()
    // for i := 1; i <= 5; i = i + 1 {
    //     go monitor(ctx03, i)
    // }
    // time.Sleep(5 * time.Second)
    // if ctx02.Err() != nil {
    //     fmt.Println("the cause of cancel is: ", ctx02.Err())
    // }
    // println(runtime.NumGoroutine())
    // println("main program exit!!!!")
	fmt.Println("ok")
}

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
    return ""

}
func (f friends) Remove(name string) string {
    return ""
}


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
