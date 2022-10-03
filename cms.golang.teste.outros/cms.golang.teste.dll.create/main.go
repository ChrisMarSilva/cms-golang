package main

// go mod init github.com/chrismarsilva/cms.golang.teste.dll.create
// go get -u XXXXXXXXXXXX
// go mod tidy

// go run main.go

// go build -buildmode=shared main.go
// Resultado: Erro = -buildmode=shared not supported on windows/amd64

// go build -o cgo/lib/lib.dll -buildmode=c-shared cgo/lib.go
// Resultado: Erro = package cgo/lib.go is not in GOROOT (C:\Program Files\go\src\cgo\lib.go)

// go build -buildmode=c-archive main.go
// Resultado: OK = porem gerou os arquivos "cmain.a" e "main.h"

// go build -buildmode=c-archive github.com/chrismarsilva/cms.golang.teste.dll.create
// Resultado: OK = porem gerou os arquivos "cms.golang.teste.dll.create.a" e "cms.golang.teste.dll.create.h"

// go build -ldflags "-s -w" -buildmode=c-shared -o gosum.dll
// Resultado: OK = gerou os arquivos "gosum.dll", "gosum.h"
// Delphi: Erro: #1 - Dll não carregada - gosum.dll 

// go build -buildmode=c-shared -o exportgo.dll main.go
// Resultado: OK = gerou os arquivos "exportgo.dll", "exportgo.h"
// Delphi: Erro: #1 - Dll não carregada - exportgo.dll 

// go build -o helloworld.dll -buildmode=c-shared
// Resultado: OK = gerou os arquivos "helloworld.dll", "helloworld.h"
// Delphi: Erro: #1 - Dll não carregada - helloworld.dll 

// go build -buildmode=c-shared -ldflags="-w -s -H=windowsgui" -o Updater.dll
// Resultado: OK = gerou os arquivos "Updater.dll", "Updater.h"
// Delphi: Erro: #1 - Dll não carregada - Updater.dll 

// GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -buildmode=c-shared -ldflags="-w -s -H=windowsgui" -o Updater.dll
// Resultado: OK = gerou os arquivos "Updater.dll", "Updater.h"
// Delphi: Erro: #1 - Dll não carregada - Updater.dll 

// go build -ldflags “-s -w” -buildmode=c-shared -o file.dll ./main.go
// gcc -shared -pthread -o goDLL.dll goDLL.c exportgo.a -lWinMM -lntdll -lWS2_32

/*
#include <stdlib.h>
#include <windows.h>
*/
import "C"
import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"fmt"
)

//export version
func version() C.int {
	return 2
}

//export PrintBye
func PrintBye() { 
    fmt.Println("From DLL: Bye!")
}

//export PrintHello
func PrintHello(name string) {
	fmt.Printf("From DLL: Hello, %s!\n", name)
}


//export Sum
func Sum(a int32, b int32) int32 {
    return a + b
}

//export Sum2
func Sum2(a int, b int) int {
    return a + b
}

//export getHelloWord
func getHelloWord(number int) int {
    return 7888 * number
}

//export rand64
func rand64() C.ulonglong {
	var buf [8]byte
	rand.Read(buf[:])
	r := binary.LittleEndian.Uint64(buf[:])
	return C.ulonglong(r)
}

//export dist
func dist(x, y C.float) C.float {
	return C.float(math.Sqrt(float64(x*x + y*y)))
}

// //export StartWorker
// func StartWorker {
//     C.MessageBox(nil, C.CString("Hello Worker"), C.CString("Warning:"), 0)
// }

func main() {
    //Need a main function to make CGO compile package as C shared library
}
