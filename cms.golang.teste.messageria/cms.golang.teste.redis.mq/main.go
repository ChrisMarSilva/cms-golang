package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/adjust/redismq"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

// go get -u github.com/adjust/redismq
// go get -u github.com/gomodule/redigo/redis

const RMQ string = "mqtest"

func main() {

	// go producer()

	// go consumer()

	testQueue := redismq.CreateQueue("localhost", "6379", "123", 2, "clicks")

	// bufferSize := 100
	// testQueue := redismq.CreateBufferedQueue("localhost", "6379", "123", 3, "clicks", bufferSize)
	// testQueue.Start()
	// defer testQueue.FlushBuffer()

	testQueue.Put("testpayload")

	consumer, err := testQueue.AddConsumer("testconsumer")
	if err != nil {
		fmt.Println("testQueue.AddConsumer: ", err)
		return
	}
	fmt.Println("consumer: ", consumer)

	package1, err := consumer.Get()
	if err != nil {
		fmt.Println("consumer.Get: ", err)
		return
	}
	err = package1.Ack()
	if err != nil {
		fmt.Println("package1.Ack: ", err)
		return
	}
	fmt.Println("Get: ", package1.Payload)

	// config.Redis.Publish(ctx, room.GetName(), message.encode())
	// pubsub := config.Redis.Subscribe(ctx, room.GetName())
	// ch := pubsub.Channel()

	fmt.Println("FIM")
}

func producer() {

	//redis_conn, err := redis.Dial("tcp", "127.0.0.1:6379", {password: "123"})
	redis_conn, err := redis.Dial("tcp", "localhost:6379", redis.DialPassword("123"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer redis_conn.Close()

	rand.Seed(time.Now().UnixNano())

	var i = 1
	for {
		_, err = redis_conn.Do("rpush", RMQ, strconv.Itoa(i))
		if err != nil {
			fmt.Println("produce error", err)
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			continue
		}
		fmt.Println("produce element:%d", i)
		// time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		// time.Sleep(1 * time.Millisecond)
		time.Sleep(1 * time.Nanosecond)
		i++
	}

}

func consumer() {

	//redis_conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	redis_conn, err := redis.Dial("tcp", "localhost:6379", redis.DialPassword("123"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer redis_conn.Close()

	rand.Seed(time.Now().UnixNano())

	for {
		ele, err := redis.String(redis_conn.Do("lpop", RMQ))
		if err != nil {
			fmt.Println("no msg.sleep now")
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		} else {
			fmt.Println("consume element:%s", ele)
		}
	}
}
