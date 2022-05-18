package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

func (row Users) Print() {
	for i := 0; i < len(row.Users); i++ {
		fmt.Println(row.Users[i])
	}
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

func (row User) String() string {
	return fmt.Sprintf("type=%v, name=%v, %v", row.Type, row.Name, row.Social)
}

type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func (row Social) String() string {
	return fmt.Sprintf("Facebook=%v, Twitter=%v, Youtube=%v", row.Facebook, row.Twitter, row.Youtube)
}

func main() {

	xmlFile, err := os.Open("users.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	fmt.Println("Successfully Opened users.xml")

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var users Users
	xml.Unmarshal(byteValue, &users)

	// note := &Notes{}
	// _ = xml.Unmarshal([]byte(data), &note)

	// for i := 0; i < len(users.Users); i++ {
	// 	fmt.Println("User Type: " + users.Users[i].Type)
	// 	fmt.Println("User Name: " + users.Users[i].Name)
	// 	fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	// }

	users.Print()

	fmt.Print("FIM")
}
