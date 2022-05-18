package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

// https://github.com/benreyle/nats-examples/blob/master/0-event/event.go

// go get github.com/nats-io/nats.go/
// go get github.com/nats-io/nats-server
// go get google.golang.org/protobuf/proto

// go get github.com/nats-io/go-nats-examples/tools/nats-pub
// go get github.com/nats-io/go-nats-examples/tools/nats-sub

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()
	defer nc.Close()

	for {
		nc.Publish("foo", []byte("Hello World"))
		nc.Flush()
		fmt.Println("nc.Publish.foo.send.ok")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("FIM")
}

/*


func main() {

	// Connect to a server
	fmt.Println("nats.Connect")
	fmt.Println(nats.DefaultURL)
	nc, _ := nats.Connect(nats.DefaultURL)
	//nc, _ = nats.Connect("nats://derek:secretpassword@demo.nats.io:4222")
	//nc, _ = nats.Connect("tls://derek:secretpassword@demo.nats.io:4443")
	//  "NATS_URI=nats://nats:4222"



	for {
		fmt.Println("nc.Publish.foo.send.ok")
		time.Sleep(1 * time.Second)
	}
	// Simple Publisher
	fmt.Println("nc.Publish")
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	fmt.Println("nc.Subscribe.foo")
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Println("nc.Subscribe.foo=", string(m.Data))
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Responding to a request message
	fmt.Println("nc.Subscribe.request")
	nc.Subscribe("request", func(m *nats.Msg) {
		fmt.Println("nc.Subscribe.request.Respond")
		m.Respond([]byte("answer is 42"))
	})

	// Simple Sync Subscriber
	fmt.Println("nc.SubscribeSync.foo")
	sub, _ := nc.SubscribeSync("foo")
	sub.NextMsg(1 * time.Second) // timeout

	// Unsubscribe
	fmt.Println("sub.Unsubscribe")
	sub.Unsubscribe()

	// Drain
	fmt.Println("sub.Drain")
	sub.Drain()

	// Requests
	fmt.Println("nc.Request.help")
	nc.Request("help", []byte("help me"), 10*time.Millisecond)

	// Replies
	fmt.Println("nc.Subscribe.help")
	nc.Subscribe("help", func(m *nats.Msg) {
		fmt.Println("nc.Subscribe.help.Publish")
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	fmt.Println("nc.ChanSubscribe.foo")
	sub, _ = nc.ChanSubscribe("foo", ch)
	msg := <-ch
	fmt.Println("msg", msg)

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	fmt.Println("nc.Drain")
	nc.Drain()

	//Close connection
	fmt.Println("nc.Close")
	nc.Close()

	fmt.Println("FIM")
}

// Shows different ways to create a Conn
func ExampleConnect() {

	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Close()

	nc, _ = nats.Connect("nats://derek:secretpassword@demo.nats.io:4222")
	nc.Close()

	nc, _ = nats.Connect("tls://derek:secretpassword@demo.nats.io:4443")
	nc.Close()

	opts := nats.Options{
		AllowReconnect: true,
		MaxReconnect:   10,
		ReconnectWait:  5 * time.Second,
		Timeout:        1 * time.Second,
	}

	nc, _ = opts.Connect()
	nc.Close()
}

// This Example shows an asynchronous subscriber.
func ExampleConn_Subscribe() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
}

// This Example shows a synchronous subscriber.
func ExampleConn_SubscribeSync() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	sub, _ := nc.SubscribeSync("foo")
	m, err := sub.NextMsg(1 * time.Second)
	if err == nil {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	} else {
		fmt.Println("NextMsg timed out.")
	}
}

func ExampleSubscription_NextMsg() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	sub, _ := nc.SubscribeSync("foo")
	m, err := sub.NextMsg(1 * time.Second)
	if err == nil {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	} else {
		fmt.Println("NextMsg timed out.")
	}
}

func ExampleSubscription_Unsubscribe() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	sub, _ := nc.SubscribeSync("foo")
	// ...
	sub.Unsubscribe()
}

func ExampleConn_Publish() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Publish("foo", []byte("Hello World!"))
}

func ExampleConn_PublishMsg() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	msg := &nats.Msg{Subject: "foo", Reply: "bar", Data: []byte("Hello World!")}
	nc.PublishMsg(msg)
}

func ExampleConn_Flush() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	msg := &nats.Msg{Subject: "foo", Reply: "bar", Data: []byte("Hello World!")}
	for i := 0; i < 1000; i++ {
		nc.PublishMsg(msg)
	}
	err := nc.Flush()
	if err == nil {
		// Everything has been processed by the server for nc *Conn.
	}
}

func ExampleConn_FlushTimeout() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	msg := &nats.Msg{Subject: "foo", Reply: "bar", Data: []byte("Hello World!")}
	for i := 0; i < 1000; i++ {
		nc.PublishMsg(msg)
	}
	// Only wait for up to 1 second for Flush
	err := nc.FlushTimeout(1 * time.Second)
	if err == nil {
		// Everything has been processed by the server for nc *Conn.
	}
}

func ExampleConn_Request() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Subscribe("foo", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I will help you"))
	})
	nc.Request("foo", []byte("help"), 50*time.Millisecond)
}

func ExampleConn_QueueSubscribe() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	received := 0

	nc.QueueSubscribe("foo", "worker_group", func(_ *nats.Msg) {
		received++
	})
}

func ExampleSubscription_AutoUnsubscribe() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	received, wanted, total := 0, 10, 100

	sub, _ := nc.Subscribe("foo", func(_ *nats.Msg) {
		received++
	})
	sub.AutoUnsubscribe(wanted)

	for i := 0; i < total; i++ {
		nc.Publish("foo", []byte("Hello"))
	}
	nc.Flush()

	fmt.Printf("Received = %d", received)
}

func ExampleConn_Close() {
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Close()
}

// Shows how to wrap a Conn into an EncodedConn
func ExampleNewEncodedConn() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, "json")
	c.Close()
}

// EncodedConn can publish virtually anything just
// by passing it in. The encoder will be used to properly
// encode the raw Go type
func ExampleEncodedConn_Publish() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, "json")
	defer c.Close()

	type person struct {
		Name    string
		Address string
		Age     int
	}

	me := &person{Name: "derek", Age: 22, Address: "85 Second St"}
	c.Publish("hello", me)
}

// EncodedConn's subscribers will automatically decode the
// wire data into the requested Go type using the Decode()
// method of the registered Encoder. The callback signature
// can also vary to include additional data, such as subject
// and reply subjects.
func ExampleEncodedConn_Subscribe() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, "json")
	defer c.Close()

	type person struct {
		Name    string
		Address string
		Age     int
	}

	c.Subscribe("hello", func(p *person) {
		fmt.Printf("Received a person! %+v\n", p)
	})

	c.Subscribe("hello", func(subj, reply string, p *person) {
		fmt.Printf("Received a person on subject %s! %+v\n", subj, p)
	})

	me := &person{Name: "derek", Age: 22, Address: "85 Second St"}
	c.Publish("hello", me)
}

// BindSendChan() allows binding of a Go channel to a nats
// subject for publish operations. The Encoder attached to the
// EncodedConn will be used for marshaling.
func ExampleEncodedConn_BindSendChan() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, "json")
	defer c.Close()

	type person struct {
		Name    string
		Address string
		Age     int
	}

	ch := make(chan *person)
	c.BindSendChan("hello", ch)

	me := &person{Name: "derek", Age: 22, Address: "85 Second St"}
	ch <- me
}

// BindRecvChan() allows binding of a Go channel to a nats
// subject for subscribe operations. The Encoder attached to the
// EncodedConn will be used for un-marshaling.
func ExampleEncodedConn_BindRecvChan() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, "json")
	defer c.Close()

	type person struct {
		Name    string
		Address string
		Age     int
	}

	ch := make(chan *person)
	c.BindRecvChan("hello", ch)

	me := &person{Name: "derek", Age: 22, Address: "85 Second St"}
	c.Publish("hello", me)

	// Receive the publish directly on a channel
	who := <-ch

	fmt.Printf("%v says hello!\n", who)
}



Listing 2. Request on NATS Request-Reply messaging
func main() {

	// Create NATS server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	msg, err := natsConnection.Request("Discovery.OrderService", nil, 1000*time.Millisecond)
	if err == nil && msg != nil {
		orderServiceDiscovery := pb.ServiceDiscovery{}
		err := proto.Unmarshal(msg.Data, &orderServiceDiscovery)
		if err != nil {
			log.Fatalf("Error on unmarshal: %v", err)
		}
		address := orderServiceDiscovery.OrderServiceUri
		log.Println("OrderService endpoint found at:", address)
		//Set up a connection to the gRPC server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
	}
}


isting 3. Response on NATS Request-Reply messaging

var orderServiceUri string
orderServiceUri = viper.GetString("discovery.orderservice")

func main() {
	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	natsConnection.Subscribe("Discovery.OrderService", func(m *nats.Msg) {
		orderServiceDiscovery := pb.ServiceDiscovery{OrderServiceUri: orderServiceUri}
		data, err := proto.Marshal(&orderServiceDiscovery)
		if err == nil {
			natsConnection.Publish(m.Reply, data)
		}
	})
	// Keep the connection alive
	runtime.Goexit()
}


Listing 4. Publisher Client of NATS Publish-Subscribe messaging


const (
	aggregate = "Order"
	event     = "OrderCreated"
)

// publishOrderCreated publish an event via NATS server
func publishOrderCreated(order *pb.Order) {
	// Connect to NATS server
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	defer natsConnection.Close()
	eventData, _ := json.Marshal(order)
	event := pb.EventStore{
		AggregateId:   order.OrderId,
		AggregateType: aggregate,
		EventId:       uuid.NewV4().String(),
		EventType:     event,
		EventData:     string(eventData),
	}
	subject := "Order.OrderCreated"
	data, _ := proto.Marshal(&event)
	// Publish message on subject
	natsConnection.Publish(subject, data)
	log.Println("Published message on subject " + subject)
}

Listing 5. Subscriber Client of NATS Publish-Subscribe messaging


const subject = "Order.>"

func main() {

	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	natsConnection.Subscribe(subject, func(msg *nats.Msg) {
		eventStore := pb.EventStore{}
		err := proto.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			// Handle the message
			log.Printf("Received message in EventStore service: %+v\n", eventStore)
			store := store.EventStore{}
			store.CreateEvent(&eventStore)
			log.Println("Inserted event into Event Store")
		}
	})

	// Keep the connection alive
	runtime.Goexit()
}


Listing 6. Subscriber Client with queue group of NATS Publish-Subscribe messaging


const (
	queue   = "Order.OrdersCreatedQueue"
	subject = "Order.OrderCreated"
)

func main() {

	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	natsConnection.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		eventStore := pb.EventStore{}
		err := proto.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			// Handle the message
			log.Printf("Subscribed message in Worker 1: %+v\n", eventStore)
		}
	})

	// Keep the connection alive
	runtime.Goexit()
}

*/
