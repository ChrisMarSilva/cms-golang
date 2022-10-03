package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.string
// go get -u XXXXXXXXXXXX
// go mod tidy

// go run main.go

func main() {
	s := [6]string{"otavio", "huncoding", "golang", "gopher", "performance", "code"}

	// Segundo (s)
	// Milissegundo (ms)
	// Microsegundo (µs)
	// Nanosegundo (ns)

	str0 := joinStrDirect(s) // 3.28us // 4º Lugar
	log.Println(str0)

	str1 := joinStrWithSprintf(s) // 1.13us // 3º Lugar
	log.Println(str1)

	str2 := joinStrWithStringsJoin(s) // 260ns // 1º Lugar
	log.Println(str2)

	str3 := joinStrWithBuilder(s) // 310ns // 2º Lugar
	log.Println(str3)
}

func joinStrDirect(str [6]string) (allStr string) {
	execTime := time.Now()

	for _, s := range str {
		allStr += s
	}

	log.Printf("joinStrDirect ExecTime is %22s %6s\n:", "->", time.Since(execTime))
	return
}

func joinStrWithSprintf(str [6]string) (allStr string) {
	execTime := time.Now()

	for _, s := range str {
		allStr += fmt.Sprintf("%s", s)
		// allStr = fmt.Sprintf("%s%s", allStr, s)
	}

	log.Printf("joinStrWithSprintf ExecTime is %22s %6s\n:", "->", time.Since(execTime))
	return
}

func joinStrWithStringsJoin(str [6]string) (allStr string) {
	execTime := time.Now()

	allStr = strings.Join(str[:], "")
	log.Printf("joinStrWithStringsJoin ExecTime is %22s %6s\n:", "->", time.Since(execTime))

	return
}

func joinStrWithBuilder(str [6]string) (allStr string) {
	execTime := time.Now()

	sb := strings.Builder{}
	for _, s := range str {
		sb.WriteString(s)
	}
	allStr = sb.String()

	log.Printf("joinStrWithBuilder ExecTime is %22s %6s\n:", "->", time.Since(execTime))
	return
}
