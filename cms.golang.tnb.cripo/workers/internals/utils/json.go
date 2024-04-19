package utils

import (
	"encoding/json"
	"fmt"
)

func prettyPrint(data interface{}) {
	str, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println("json.MarshalIndent()", err)
		return
	}
	fmt.Println(string(str)) // fmt.Printf("%s \n", str)
}
