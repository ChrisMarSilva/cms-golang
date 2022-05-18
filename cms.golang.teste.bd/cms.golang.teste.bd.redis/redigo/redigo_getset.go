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

	_, err := client.Do("SET", "mykey", "Hello from redigo!")
	if err != nil {
		log.Println("Redis.GoModule.ERRO.Set.mykey: ", err)
		// return
	} else {
		log.Println("Redis.GoModule.OK.Set.mykey:", "Hello from redigo!")
	}

	//val, err := client.Do("GET", "mykey")
	val, err := redigo.String(client.Do("GET", "mykey"))
	if err != nil {
		log.Println("Redis.GoModule.ERRO.Get.mykey: ", err)
		// return
	} else {
		log.Println("Redis.GoModule.OK.Get.mykey:", val)
	}

	valOld, err := client.Do("GET", "mykeyOld")
	if err == redigo.ErrNil {
		log.Println("Redis.GoModule.ERRO.Get.mykeyOld: this Key does not exist", err)
		// return
	} else if err != nil {
		log.Println("Redis.GoModule.ERRO.Get.mykeyOld: ", err)
		// return
	} else {
		log.Println("Redis.GoModule.OK.Get.mykeyOld:", valOld)
	}

	// keys, err := redigo.Strings(client.Do("KEYS", "*"))
	// if err != nil {
	// 	log.Println("Redis.GoModule.ERRO.KEYS: ", err)
	// 	// return
	// } else {
	// 	for indx, key := range keys {
	// 		log.Println("Redis.GoModule.OK.KEYS:", indx, " -", key)
	// 	}
	// }

}
