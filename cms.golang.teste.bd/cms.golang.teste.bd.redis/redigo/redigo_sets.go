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

	vehicles1, err := client.Do("ZADD", "vehicles", 4, "car")
	if err != nil {
		log.Println("Redis.GoModule.ERRO.ZADD.vehicles.car: ", err)
		// return
	} else {
		log.Println("Redis.GoModule.OK.ZADD.vehicles.car:", vehicles1)
	}

	vehicles2, err := client.Do("ZADD", "vehicles", 2, "bike")
	if err != nil {
		log.Println("Redis.GoModule.ERRO.ZADD.vehicles.bike: ", err)
		// return
	} else {
		log.Println("Redis.GoModule.OK.ZADD.vehicles.bikes:", vehicles2)
	}

	vehicles3, err := client.Do("ZRANGE", "vehicles", 0, -1, "WITHSCORES")
	if err != nil {
		log.Println("Redis.GoModule.ERRO.ZRANGE.vehicles.WITHSCORES: ", err)
		// return
	} else {
		log.Println("Redis.GoModule.OK.ZRANGE.vehicles.WITHSCORES:", vehicles3)
		// for _, vehicle := range vehicles3 {
		// 	log.Println("Redis.GoModule.OK.ZRANGE.vehicles.WITHSCORES.vehicle:", vehicle)
		// }
	}

}
