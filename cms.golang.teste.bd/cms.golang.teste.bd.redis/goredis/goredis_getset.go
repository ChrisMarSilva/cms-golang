package goredis

import (
	"encoding/json"
	"log"
)

func TesteGoRedis() {

	log.Println("Redis.GoRedis.INI")
	defer log.Println("")
	defer log.Println("Redis.GoRedis.FIM")

	redis, err := NewRedis(0)
	if err != nil {
		log.Println("Redis.GoRedis.ERRO.NewRedis: ", err)
		return
	}
	defer redis.client.Close()

	log.Println("Redis.GoRedis.FlushDB.INI")
	err = redis.client.FlushDB().Err()
	if err != nil {
		log.Println("Redis.GoRedis.FlushDB.ERRO:", err)
	} else {
		log.Println("Redis.GoRedis.FlushDB.FIM")
	}

	log.Println("Redis.GoRedis.TESTE.INI")

	nameKey := "KeyName1"
	nameVal := "Chris"
	nameSet := redis.client.Set(nameKey, nameVal, 0)
	if nameSet.Err() != nil {
		log.Println("Redis.GoRedis.ERRO.Set.Name: ", nameSet.Err())
		// return
	} else {
		log.Println("Redis.GoRedis.OK.Set.Name:", nameSet.Val())
		log.Println(nameSet)
	}

	nameSet = redis.client.Set(nameKey, nameVal, 0)
	if nameSet.Err() != nil {
		log.Println("Redis.GoRedis.ERRO.Set.Name: ", nameSet.Err())
		// return
	} else {
		log.Println("Redis.GoRedis.OK.Set.Name:", nameSet.Val())
		log.Println(nameSet)
	}

	log.Println("Redis.GoRedis.TESTE.FIM")

	val, err := redis.client.Get(nameKey).Result()
	if err != nil {
		log.Println("Redis.GoRedis.ERRO.Get.Name: ", err)
		// return
	} else {
		log.Println("Redis.GoRedis.OK.Get.Name:", val)
	}

	redis.client.Del("name")

	val, err = redis.client.Get("name").Result()
	if err != nil {
		log.Println("Redis.GoRedis.ERRO.Get.Name: ", err)
		// return
	} else {
		log.Println("Redis.GoRedis.OK.Get.Name:", val)
	}

	pessoa1 := Pessoa{"0001", "Chris MarSil"}
	err = redis.SetPessoa(pessoa1)
	if err != nil {
		log.Println("Redis.GoRedis.ERRO.Set.Pessoa.0001: ", err)
		// return
	} else {
		log.Println("Redis.GoRedis.OK.Set.Pessoa.0001:", pessoa1)
	}

	pessoa2, err := redis.GetPessoa("0001")
	if err != nil {
		log.Println("Redis.GoRedis.ERRO.Get.Pessoa.0001: ", err)
		// return
	} else {
		log.Println("Redis.GoRedis.OK.Get.Pessoa.0001:", pessoa2)
	}

	_, err = redis.GetPessoa("0002")
	if err != nil {
		// if err == redis.Nil {
		// 	log.Println("Redis.GoRedis.ERRO.Get.Pessoa.0002:key does not exist")
		// } else {
		log.Println("Redis.GoRedis.ERRO.Get.Pessoa.0002: ", err)
		// }
		// return
	}

	pessoa3, _ := json.Marshal(Pessoa{"0003", "Chris MarSil"})
	err = redis.client.Set("0003", pessoa3, 0).Err()
	if err != nil {
		log.Println("Redis.GoRedis.ERRO.Set.Pessoa.0003: ", err)
		// return
	} else {
		log.Println("Redis.GoRedis.OK.Set.Pessoa.0003:", pessoa1)
	}

	pessoa4, err := redis.client.Get("0003").Result()
	if err != nil {
		log.Println("Redis.GoRedis.ERRO.Get.Pessoa.0003: ", err)
		// return
	} else {
		log.Println("Redis.GoRedis.OK.Get.Pessoa.0003.Str:", pessoa4)
		var pessoa44 Pessoa // pessoa := Pessoa{}
		err = json.Unmarshal([]byte(pessoa4), &pessoa44)
		log.Println("Redis.GoRedis.OK.Get.Pessoa.0003.Json:", pessoa44)
	}

}
