package main

// type RabbitClient struct {
// 	sendConn *amqp.Connection
// 	recConn  *amqp.Connection
// 	sendChan *amqp.Channel
// 	recChan  *amqp.Channel
//  }
 
 
 
 
//  // Create a connection to rabbitmq
//  func (rcl *RabbitClient) connect(isRec, reconnect bool) (*amqp.Connection, error) {
//   if reconnect {
//    if isRec {
// 	rcl.recConn = nil
//    } else {
// 	rcl.sendConn = nil
//    }
//   }
//   if isRec && rcl.recConn != nil {
//    return rcl.recConn, nil
//   } else if !isRec && rcl.sendConn != nil {
//    return rcl.sendConn, nil
//   }
//   var c string
//   if config.Username == "" {
//    c = fmt.Sprintf("amqp://%s:%s/", config.Host, config.Port)
//   } else {
//    c = fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port)
//   }
//   conn, err := amqp.Dial(c)
//   if err != nil {
//    log.Printf("\r\n--- could not create a conection ---\r\n")
//    time.Sleep(1 * time.Second)
//    return nil, err
//   }
//   if isRec {
//    rcl.recConn = conn
//    return rcl.recConn, nil
//   } else {
//    rcl.sendConn = conn
//    return rcl.sendConn, nil
//   }
//  }
 
 
//  func (rcl *RabbitClient) channel(isRec, recreate bool) (*amqp.Channel, error) {
// 	if recreate {
// 	   if isRec {
// 		  rcl.recChan = nil
// 	   } else {
// 		  rcl.sendChan = nil
// 	   }
// 	}
// 	if isRec && rcl.recConn == nil {
// 	   rcl.recChan = nil
// 	}
// 	if !isRec && rcl.sendConn == nil {
// 	   rcl.recChan = nil
// 	}
// 	if isRec && rcl.recChan != nil {
// 	   return rcl.recChan, nil
// 	} else if !isRec && rcl.sendChan != nil {
// 	   return rcl.sendChan, nil
// 	}
// 	for {
// 	   _, err := rcl.connect(isRec, recreate)
// 	   if err == nil {
// 		  break
// 	   }
// 	}
// 	var err error
// 	if isRec {
// 	   rcl.recChan, err = rcl.recConn.Channel()
// 	} else {
// 	   rcl.sendChan, err = rcl.sendConn.Channel()
// 	}
// 	if err != nil {
// 	   log.Println("--- could not create channel ---")
// 	   time.Sleep(1 * time.Second)
// 	   return nil, err
// 	}
// 	if isRec {
// 	   return rcl.recChan, err
// 	} else {
// 	   return rcl.sendChan, err
// 	}
//  }
 
 
//  // Consume based on name of the queue
//  func (rcl *RabbitClient) Consume(n string, f func(interface{}) error) {
//   for {
//    for {
// 	_, err := rcl.channel(true, true)
// 	if err == nil {
// 	 break
// 	}
//    }
//    log.Printf("--- connected to consume '%s' ---\r\n", n)
//    q, err := rcl.recChan.QueueDeclare(
// 	n,
// 	true,
// 	false,
// 	false,
// 	false,
// 	amqp.Table{"x-queue-mode": "lazy"},
//    )
//    if err != nil {
// 	log.Println("--- failed to declare a queue, trying to reconnect ---")
// 	continue
//    }
//    connClose := rcl.recConn.NotifyClose(make(chan *amqp.Error))
//    connBlocked := rcl.recConn.NotifyBlocked(make(chan amqp.Blocking))
//    chClose := rcl.recChan.NotifyClose(make(chan *amqp.Error))
//    m, err := rcl.recChan.Consume(
// 	q.Name,
// 	uuid.NewV4().String(),
// 	false,
// 	false,
// 	false,
// 	false,
// 	nil,
//    )
//    if err != nil {
// 	log.Println("--- failed to consume from queue, trying again ---")
// 	continue
//    }
//    shouldBreak := false
//    for {
// 	if shouldBreak {
// 	 break
// 	}
// 	select {
// 	case _ = <-connBlocked:
// 	 log.Println("--- connection blocked ---")
// 	 shouldBreak = true
// 	 break
// 	case err = <-connClose:
// 	 log.Println("--- connection closed ---")
// 	 shouldBreak = true
// 	 break
// 	case err = <-chClose:
// 	 log.Println("--- channel closed ---")
// 	 shouldBreak = true
// 	 break
// 	case d := <-m:
// 	 err := f(d.Body)
// 	 if err != nil {
// 	  _ = d.Ack(false)
// 	  break
// 	 }
// 	 _ = d.Ack(true)
// 	}
//    }
//   }
//  }
 
 
//  // Publish an array of bytes to a queue
//  func (rcl *RabbitClient) Publish(n string, b []byte) {
//   r := false
//   for {
//    for {
// 	_, err := rcl.channel(false, r)
// 	if err == nil {
// 	 break
// 	}
//    }
//    q, err := rcl.sendChan.QueueDeclare(
// 	n,
// 	true,
// 	false,
// 	false,
// 	false,
// 	amqp.Table{"x-queue-mode": "lazy"},
//    )
//    if err != nil {
// 	log.Println("--- failed to declare a queue, trying to resend ---")
// 	r = true
// 	continue
//    }
//    err = rcl.sendChan.Publish(
// 	"",
// 	q.Name,
// 	false,
// 	false,
// 	amqp.Publishing{
// 	 MessageId:    uuid.NewV4().String(),
// 	 DeliveryMode: amqp.Persistent,
// 	 ContentType:  "text/plain",
// 	 Body:         b,
// 	})
//    if err != nil {
// 	log.Println("--- failed to publish to queue, trying to resend ---")
// 	r = true
// 	continue
//    }
//    break
//   }
//  }
 
 
//  var rc rmq.RabbitClient
//  rc.Consume("test-queue", funcName)
//  rc.Publish("test-queue", mBody)
 
//  ------------
 
//  // Conn -
//  type Conn struct {
// 	 Channel *amqp.Channel
//  }
 
//  // GetConn -
//  func GetConn(rabbitURL string) (Conn, error) {
// 	 conn, err := amqp.Dial(rabbitURL)
// 	 if err != nil {
// 		 return Conn{}, err
// 	 }
 
// 	 ch, err := conn.Channel()
// 	 return Conn{
// 		 Channel: ch,
// 	 }, err
//  }
 
//  // Publish -
//  func (conn Conn) Publish(routingKey string, data []byte) error {
// 	 return conn.Channel.Publish(
// 		 // exchange - yours may be different
// 		 "events",
// 		 routingKey,
// 		 // mandatory - we don't care if there I no queue
// 		 false,
// 		 // immediate - we don't care if there is no consumer on the queue
// 		 false,
// 		 amqp.Publishing{
// 			 ContentType:  "application/json",
// 			 Body:         data,
// 			 DeliveryMode: amqp.Persistent,
// 		 })
//  }
 
 
//  // StartConsumer -
//  func (conn Conn) StartConsumer(
// 	 queueName,
// 	 routingKey string,
// 	 handler func(d amqp.Delivery) bool,
// 	 concurrency int) error {
 
// 	 // create the queue if it doesn't already exist
// 	 _, err := conn.Channel.QueueDeclare(queueName, true, false, false, false, nil)
// 	 if err != nil {
// 		 return err
// 	 }
 
// 	 // bind the queue to the routing key
// 	 err = conn.Channel.QueueBind(queueName, routingKey, "events", false, nil)
// 	 if err != nil {
// 		 return err
// 	 }
 
// 	 // prefetch 4x as many messages as we can handle at once
// 	 prefetchCount := concurrency * 4
// 	 err = conn.Channel.Qos(prefetchCount, 0, false)
// 	 if err != nil {
// 		 return err
// 	 }
 
// 	 msgs, err := conn.Channel.Consume(
// 		 queueName, // queue
// 		 "",        // consumer
// 		 false,     // auto-ack
// 		 false,     // exclusive
// 		 false,     // no-local
// 		 false,     // no-wait
// 		 nil,       // args
// 	 )
// 	 if err != nil {
// 		 return err
// 	 }
 
// 	 // create a goroutine for the number of concurrent threads requested
// 	 for i := 0; i < concurrency; i++ {
// 		 fmt.Printf("Processing messages on thread %v...\n", i)
// 		 go func() {
// 			 for msg := range msgs {
// 				 // if tha handler returns true then ACK, else NACK
// 				 // the message back into the rabbit queue for
// 				 // another round of processing
// 				 if handler(msg) {
// 					 msg.Ack(false)
// 				 } else {
// 					 msg.Nack(false, true)
// 				 }
// 			 }
// 			 fmt.Println("Rabbit consumer closed - critical Error")
// 			 os.Exit(1)
// 		 }()
// 	 }
// 	 return nil
//  }
 
 
//  func main() {
// 	 conn, err := rabbit.GetConn("amqp://guest:guest@localhost")
// 	 if err != nil {
// 		 panic(err)
// 	 }
 
// 	 go func() {
// 		 for {
// 			 time.Sleep(time.Second)
// 			 conn.Publish("test-key", []byte(`{"message":"test"}`))
// 		 }
// 	 }()
 
// 	 err = conn.StartConsumer("test-queue", "test-key", handler, 2)
 
// 	 if err != nil {
// 		 panic(err)
// 	 }
 
// 	 forever := make(chan bool)
// 	 <-forever
//  }
 
//  func handler(d amqp.Delivery) bool {
// 	 if d.Body == nil {
// 		 fmt.Println("Error, no message body!")
// 		 return false
// 	 }
// 	 fmt.Println(string(d.Body))
// 	 return true
//  }