package goredis

import (
	"log"
)

func TesteGoRedisPubSub() {

	log.Println("Redis.GoRedis.PubSub.INI")
	defer log.Println("")
	defer log.Println("Redis.GoRedis.PubSub.FIM")

	redis, err := NewRedis(1)
	if err != nil {
		log.Println("Redis.GoRedis.PubSub.ERRO.NewRedis: ", err)
		return
	}
	defer redis.client.Close()

	pubsub := redis.client.PSubscribe("mychannel*")
	defer pubsub.Close()
	log.Println("Redis.GoRedis.PubSub.PSubscribe:", pubsub)

	n, err := redis.client.Publish("mychannel1", "hello").Result()
	log.Println("Redis.GoRedis.PubSub.Publish:", n)

	pubsub.PUnsubscribe("mychannel*")

	stats := redis.client.PoolStats()
	log.Println("Redis.GoRedis.PubSub.PoolStats:", stats)

	channels, err := redis.client.PubSubChannels("mychannel*").Result()
	log.Println("Redis.GoRedis.PubSub.PubSubChannels:", channels)

	pubsub = redis.client.Subscribe("mychannel", "mychannel2")
	defer pubsub.Close()

	channels, err = redis.client.PubSubChannels("mychannel*").Result()
	log.Println("Redis.GoRedis.PubSub.PubSubChannels:", channels)

	channels, err = redis.client.PubSubChannels("").Result()
	log.Println("Redis.GoRedis.PubSub.PubSubChannels:", channels)

	channels, err = redis.client.PubSubChannels("*").Result()
	log.Println("Redis.GoRedis.PubSub.PubSubChannels:", channels)

	pubsub = redis.client.Subscribe("mychannel", "mychannel2")
	defer pubsub.Close()

	channels1, err := redis.client.PubSubNumSub("mychannel", "mychannel2", "mychannel3").Result()
	log.Println("Redis.GoRedis.PubSub.PubSubChannels:", channels1)

	pubsub = redis.client.PSubscribe("*")
	defer pubsub.Close()

	n, err = redis.client.Publish("mychannel", "hello").Result()
	n, err = redis.client.Publish("mychannel2", "hello2").Result()
	pubsub.Unsubscribe("mychannel", "mychannel2")

	// pubsub = redis.client.Subscribe("mychannel")
	// defer pubsub.Close()
	// err = redis.client.Publish("mychannel", "hello").Err()
	// err = redis.client.Publish("mychannel", "world").Err()
	// msg, err := pubsub.ReceiveMessage()
	// log.Println("msg", msg)
	// msg, err = pubsub.ReceiveMessage()
	// log.Println("msg", msg)

	// message, _ := json.Marshal(Message{"Action1", "Message1", "Target1", "Sender1"})
	// redis.client.Publish("general", message) // message.encode()

	// go func() {
	// 	pubsub := redis.client.Subscribe("general")
	// 	ch := pubsub.Channel()
	// 	for msg := range ch {
	// 		log.Println([]byte(msg.Payload))
	// 		var message Message
	// 		json.Unmarshal([]byte(msg.Payload), &message)
	// 		switch message.Action {
	// 		case "Action1":
	// 			log.Println("Action1", message.Message)
	// 		case "Action2":
	// 			log.Println("Action2", message.Message)
	// 		}
	// 	}
	// }()

	// pubsub = redis.client.Subscribe("responseclient")
	// defer pubsub.Close()
	// redis.client.Publish("responseclient", "Example New Request").Err()
	// msg, err := pubsub.ReceiveMessage()
	// log.Println(msg.Channel, msg.Payload)

	// pubsub = redis.client.Subscribe("control")
	// defer pubsub.Close()
	// if _, err := pubsub.Receive(); err != nil {
	// 	log.Println("failed to receive from control PubSub", err)
	// 	return
	// }
	// controlCh := pubsub.Channel()
	// for msg := range controlCh {
	// 	log.Println(msg.Payload)
	// }

}
