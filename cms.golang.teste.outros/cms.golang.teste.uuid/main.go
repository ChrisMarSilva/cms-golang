package main

import (
	"github.com/google/uuid"
	uuid2 "github.com/satori/go.uuid"
	// - "github.com/edwingeng/wuid/redis/wuid"
	"log"
	"strings"
)

// go mod init github.com/chrismarsilva/cms-golang-teste-uuid
// go mod tidy

// go run main.go
// go run .

func main() {

	log.Println("")
	log.Println("github.com/google/uuid")
	log.Println(uuid.New().String())
	log.Println(uuid.NewString())
	r, _ := uuid.NewRandom()
	log.Println(r.String())
	uuidWithHyphen := uuid.New()
	log.Println(uuidWithHyphen, " = ", strings.Replace(uuidWithHyphen.String(), "-", "", -1))

	// log.Println("")
	// log.Println("github.com/satori/go.uuid")

	// // Creating UUID Version 4
	// // panic on error
	// u1 := uuid2.Must(uuid2.NewV4())
	// log.Printf("UUIDv4: %s\n", u1)

	// // or error handling
	// u2, err := uuid2.NewV4()
	// if err != nil {
	// 	log.Printf("Something went wrong: %s", err)
	// 	return
	// }
	// log.Printf("UUIDv4: %s\n", u2)

	// // Parsing UUID from string input
	// u2, err := uuid2.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	// if err != nil {
	// 	log.Printf("Something went wrong: %s", err)
	// 	return
	// }
	// log.Printf("Successfully parsed: %s", u2)

	// --------------------------------------------------
	// --------------------------------------------------

	// log.Println("")
	// log.Println("github.com/edwingeng/wuid/redis/wuid")

	// newClient := func() (redis.Cmdable, bool, error) {
	// 	var client redis.Cmdable
	// 	return client, true, nil
	// }

	// // Setup
	// g := NewWUID("default", nil)
	// _ = g.LoadH28FromRedis(newClient, "wuid")

	// // Generate
	// for i := 0; i < 10; i++ {
	// 	log.Printf("%#016x\n", g.Next())
	// }

	// --------------------------------------------------
	// --------------------------------------------------

	log.Println("")
	log.Println("ok")
}
