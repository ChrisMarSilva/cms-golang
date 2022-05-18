package main

import (
	"fmt"

	"github.com/huydang284/fixedwidth"
)

// go get github.com/ianlopshire/go-fixedwidth
// go get github.com/huydang284/fixedwidth

type people struct {
	Name string `fixed:"10"`
	Age  int    `fixed:"3"`
}

func main() {

	me := people{Name: "Huy", Age: 25}
	data, _ := fixedwidth.Marshal(me)
	fmt.Println(string(data))

	var me1 people
	data1 := []byte("Huy       25 ")
	fixedwidth.Unmarshal(data1, &me1)
	fmt.Printf("%+v", me1)

	fmt.Println("FIM")
}
