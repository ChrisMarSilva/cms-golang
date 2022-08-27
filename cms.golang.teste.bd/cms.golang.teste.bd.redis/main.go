package main

import (
	goredis "github.com/ChrisMarSilva/cms-golang-teste-redis/goredis"
	"log"
)

// go mod init github.com/chrismarsilva/cms-golang-teste-redis
// go get github.com/go-redis/redis
// go get github.com/go-redis/redis/v8
// go get github.com/gomodule/redigo/redis

// https://github.com/go-redis/redis/blob/master/example_test.go
// https://www.alexedwards.net/blog/working-with-redis

func main() {
	log.Println("Redis.INI")
	log.Println("")

	goredis.TesteGoRedis()

	log.Println("")
	log.Println("Redis.FIM")
}
