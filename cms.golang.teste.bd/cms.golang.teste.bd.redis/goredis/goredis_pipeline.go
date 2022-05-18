package goredis

import (
	"log"
	"strconv"
	"time"
)

func TesteGoRedisPepiline() {

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

	// pipe := redis.client.Pipeline()
	// for _, key := range keys {
	// 	pipe.Del(key)
	// }
	// pipe.Exec()

	// iter := redis.client.Scan(0, "prefix*", 0).Iterator()
	// for iter.Next() {
	// 	err := redis.client.Del(iter.Val()).Err()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// if err := iter.Err(); err != nil {
	// 	panic(err)
	// }

	pipe := redis.client.TxPipeline()
	pipe.Set("language", "golang", 0)
	pipe.Set("year", 2009, 0)
	pipe.ZScore("leaderboardKey", "username1")
	pipe.ZRank("leaderboardKey", "username2")
	pipe.Exec()

	log.Println("Redis.GoRedis.Pipeline.INI")
	var start0 time.Time
	var start time.Time
	oneMillion := makeRange(1, 10)
	twoMillion := makeRange(11, 30)
	threeMillion := makeRange(31, 60)
	elements := [][]int{oneMillion, twoMillion, threeMillion}
	start0 = time.Now()
	pipeline := redis.client.Pipeline()
	for indx, elem := range elements {
		//start = time.Now()
		//log.Println("Redis.GoRedis.Pipeline.FOR.INI:", indx)
		for i := 0; i < len(elem); i++ {
			cmd := pipeline.Set("KEY:"+strconv.Itoa(elem[i]), "VAL:"+strconv.Itoa(elem[i]), 0)
			if cmd.Err() != nil {
				log.Println("Redis.GoRedis.Pipeline.ERRO:", cmd.Err())
				return
			}
		}
		//log.Println("Redis.GoRedis.Pipeline.FOR.FIM:", indx, time.Since(start), "- len = ", len(elem))
		log.Println("Redis.GoRedis.Pipeline.EXEC.INI:", indx)
		start = time.Now()
		_, err := pipeline.Exec()
		if err != nil {
			log.Println("Redis.GoRedis.Pipeline.ERRO:", err)
			return
		}
		log.Println("Redis.GoRedis.Pipeline.EXEC.FIM:", indx, time.Since(start), "- len = ", len(elem))
	}
	log.Println("Redis.GoRedis.Pipeline.FIM:", time.Since(start0))

}
