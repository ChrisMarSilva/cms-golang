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

	values, err := redigo.Values(client.Do("HGETALL", "album:1"))
	if err != nil {
		log.Println("Redis.GoModule.ERRO.HGETALL: ", err)
		// return
	} else {
		var album Album
		err = redigo.ScanStruct(values, &album)
		if err != nil {
			log.Println("Redis.GoModule.ERRO.HGETALL.ScanStruct: ", err)
			// return
		} else {
			//log.Println("%+v", album)
			log.Println("Redis.GoModule.ERRO.HGETALL.ScanStruct: ", album)
		}
	}

	n, err := client.Do("hset", "key", "field1", "value1") //write
	n, err = client.Do("hset", "key", "field2", "value2")  //write
	n, err = client.Do("hset", "key", "field3", "value3")  //write
	log.Println("hset: ", n)

	result, err := redigo.Values(client.Do("hgetall", "key")) //read
	log.Println("hgetall: ", result)

	_, err = client.Do("hset", "myhash", "bike1", "mobike")
	if err != nil {
		fmt.Println("haset failed", err.Error())
	}

	res, err := client.Do("hget", "myhash", "bike1")
	fmt.Println(reflect.TypeOf(res))
	if err != nil {
		fmt.Println("hget failed", err.Error())
	} else {
		fmt.Println("hget value :%s\n", res.([]byte))
	}

	_, err = client.Do("hmset", "myhash", "bike2", "bluegogo", "bike3", "xiaoming", "bike4", "xiaolan")
	if err != nil {
		fmt.Println("hmset error", err.Error())
	} else {
		value, err := redigo.Values(client.Do("hmget", "myhash", "bike1", "bike2", "bike3", "bike4"))
		if err != nil {
			fmt.Println("hmget failed", err.Error())
		} else {
			fmt.Println("hmget myhash's element :")
			for _, v := range value {
				fmt.Println("%s ", v.([]byte))
			}
			fmt.Println("\n")
		}
	}

	_, err = client.Do("hmset", "myhash", "bike2", "bluegogo", "bike3", "xiaoming", "bike4", "xiaolan")
	if err != nil {
		fmt.Println("hmset error", err.Error())
	} else {
		value, err := redigo.Values(client.Do("hmget", "myhash", "bike1", "bike2", "bike3", "bike4"))
		if err != nil {
			fmt.Println("hmget failed", err.Error())
		} else {
			fmt.Println("hmget myhash's element :")
			for _, v := range value {
				fmt.Println("%s ", v.([]byte))
			}
			fmt.Println("\n")
		}
	}

	isExist, err := client.Do("hexists", "myhash", "tmpnum")
	if err != nil {
		fmt.Println("hexist failed", err.Error())
	} else {
		fmt.Println("exist or not:", isExist)
	}

	ilen, err := client.Do("hlen", "myhash")
	if err != nil {
		fmt.Println("hlen failed", err.Error())
	} else {
		fmt.Println("myhash's len is :", ilen)
	}

	resKeys, err := redigo.Values(client.Do("hkeys", "myhash"))
	if err != nil {
		fmt.Println("hkeys failed", err.Error())
	} else {
		fmt.Println("myhash's keys is :")
		for _, v := range resKeys {
			fmt.Println("%s ", v.([]byte))
		}
		fmt.Println()
	}

	resValues, err := redigo.Values(client.Do("hvals", "myhash"))
	if err != nil {
		fmt.Println("hvals failed", err.Error())
	} else {
		fmt.Println("myhash's values is:")
		for _, v := range resValues {
			fmt.Println("%s ", v.([]byte))
		}
		fmt.Println()
	}

	_, err = client.Do("HDEL", "myhash", "tmpnum")
	if err != nil {
		fmt.Println("hdel failed", err.Error())
	}

	result, err = redigo.Values(client.Do("hgetall", "myhash"))
	if err != nil {
		fmt.Println("hgetall failed", err.Error())
	} else {
		fmt.Println("all keys and values are:")
		for _, v := range result {
			fmt.Println("%s ", v.([]byte))
		}
		fmt.Println()
	}

}
