package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ResultPeopleList struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

type ResultPeople struct {
	List    []ResultPeopleList `json:"people"`
	Number  int                `json:"number"`
	Message string             `json:"message"`
}

func main() {

	url := "http://api.open-notify.org/astros.json"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")

	//myClient := http.Client{Timeout: time.Second * 10}
	myClient := &http.Client{Timeout: time.Second * 10}
	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	// target interface{}
	//  return json.NewDecoder(r.Body).Decode(target)

	// decoder := json.NewDecoder(res.Body)
	// var data Tracks
	// err = decoder.Decode(&data)

	// res, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	if res.Body != nil {
		defer res.Body.Close()
	}

	fmt.Println("res.StatusCode", res.StatusCode, " - ", res.Status)
	// if res.StatusCode != http.StatusOK {
	// 	fmt.Println("unexpected http GET status: %s", res.Status)
	// }

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//data := []ResultPeople{}
	var data ResultPeople

	//err = json.Unmarshal([]byte(body), &data)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	for i, people := range data.List {
		fmt.Printf("%d: %s %s\n", i, people.Name, people.Craft)
	}

	fmt.Println("FIM")
}
