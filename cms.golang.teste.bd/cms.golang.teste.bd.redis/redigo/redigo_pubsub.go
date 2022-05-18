package redigo

import (
	"fmt"
	//"context"
	"bytes"
	"encoding/gob"
	"encoding/json"
	redigo "github.com/gomodule/redigo/redis"
	"log"
	"reflect"
	"strconv"
	"time"
)

// import "github.com/gomodule/redigo/redis"

func TesteGoModule() {
	log.Println("Redis.GoModule.INI")
	defer log.Println("")
	defer log.Println("Redis.GoModule.FIM")

	var pool = newPool()
	client := pool.Get()
	defer client.Close()

	fmt.Println("PUBLISH")
	client.Do("PUBLISH", "key", "value")
	client.Do("PUBLISH", "key", "value")
	client.Do("PUBLISH", "key", "value")
	client.Do("PUBLISH", "key", "value")

	psc := redigo.PubSubConn{Conn: client}
	psc.PSubscribe("key")
	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redigo.Message:
				log.Println(v.Data)
			default:
				log.Println(v)
			}
		}
	}()

}

/*

const RMQ string = "mqtest"


func producer() {
	var i = 1
		_, err = redis_conn.Do("rpush", RMQ, strconv.Itoa(i))
		i++

func consumer() {
		ele, err := redis.String(redis_conn.Do("lpop", RMQ))
			fmt.Println("consume element:%s", ele)


*/
