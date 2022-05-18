package main

import (
	"github.com/ibm-messaging/mq-golang-jms20/jms20subset"
	mqjms "github.com/ibm-messaging/mq-golang-jms20/mqjms"
	"log"
	"time"
)

// go mod init github.com/ChrisMarSilva/cms.golang.messageria.ibm.mq
// go get github.com/ibm-messaging/mq-golang
// go get github.com/ibm-messaging/mq-golang/v5/ibmmq
// go get github.com/ibm-messaging/mq-golang-jms20
// go get github.com/ibm-messaging/mq-golang-jms20/mqjms
// go mod tidy

// cd "C:\Users\chris\Desktop\CMS GoLang\cms.golang.tnb.yahoo"
// go run main.go
// go run .

// https://github.com/ibm-messaging/mq-golang-jms20

// docker run --env LICENSE=accept --env MQ_QMGR_NAME=QM1 --env MQ_USER_NAME="mqm" --env MQ_ADMIN_PASSWORD="Admin123" --env MQ_APP_PASSWORD="Admin123" --publish 1414:1414 --publish 9443:9443 --detach ibmcom/mq

// https://localhost:9443
// admin = Admin
// app = Admin123

func main() {

	TesteCreateProducer()
	TesteCreateConsumer()

}

func NewIBMMQConfig() *mqjms.ConnectionFactoryImpl {
	return &mqjms.ConnectionFactoryImpl{
		QMName:      "QM1",
		Hostname:    "127.0.0.1",
		PortNumber:  1414,
		ChannelName: "DEV.ADMIN.SVRCONN",
		UserName:    "admin",
		Password:    "Admin123",
	}
}

// func NewIBMMQConnection(cf *mqjms.ConnectionFactoryImpl) jms20subset.JMSContext {
// 	context, ctxErr := cf.CreateContext()
// 	if ctxErr != nil {
// 		if ctxErr.GetLinkedError() != nil {
// 			log.Println("ctxErr", ctxErr.GetLinkedError())
// 		}
// 	}
// 	return context
// }

func TesteCreateProducer() {

	cf := NewIBMMQConfig()

	context, ctxErr := cf.CreateContext()
	if ctxErr != nil {
		log.Println("ctxErr", ctxErr)
		return
	}

	if context != nil {
		defer context.Close()
	}

	// // Check the error code that comes back in the Error (JMS uses a String for the code)
	// assert.Equal(t, "2035", err.GetErrorCode())
	// assert.Equal(t, "MQRC_NOT_AUTHORIZED", err.GetReason())

	queue := context.CreateQueue("DEV.QUEUE.1")
	// asyncQueue := context.CreateQueue("DEV.QUEUE.1").SetPutAsyncAllowed(jms20subset.Destination_PUT_ASYNC_ALLOWED_ENABLED)

	// prodErr := context.CreateProducer().SendString(queue, msgData)
	// if prodErr != nil {
	// 	log.Println("prodErr", prodErr)
	// 	return
	// }

	producer := context.CreateProducer()
	// producer := context.CreateProducer().SetDeliveryMode(jms20subset.DeliveryMode_NON_PERSISTENT)

	for i := 0; i < 2; i++ {

		msgBody := "Hello from Go at " + time.Now().Format(time.RFC3339)
		// prodErr := producer.SendString(queue, msgBody)

		//sentMsg := context.CreateTextMessage()
		//sentMsg.SetText(msgBody)

		// bytesOver32kb := []byte(msgBody)
		// sentMsg := context.CreateBytesMessageWithBytes(bytesOver32kb)

		sentMsg := context.CreateTextMessageWithString(msgBody)

		// assert.Equal(t, "2085", err2.GetErrorCode())
		// assert.Equal(t, "MQRC_UNKNOWN_OBJECT_NAME", err2.GetReason())

		prodErr := producer.Send(queue, sentMsg)
		if prodErr != nil {
			log.Println("prodErr", prodErr)
			return
		}

		// https://pkg.go.dev/github.com/matscus/mq-golang-jms20/jms20subset#Message

		log.Println("Sendever text string:", msgBody, "-", i, sentMsg.GetJMSMessageID())
	}

	// log.Println("Put message to", strings.TrimSpace(qObject.Name))
	// log.Println("MsgId:" + hex.EncodeToString(putmqmd.MsgId))

}

func TesteCreateConsumer() {

	cf := NewIBMMQConfig()

	context, ctxErr := cf.CreateContext()
	if ctxErr != nil {
		log.Println("ctxErr", ctxErr)
		return
	}

	if context != nil {
		defer context.Close()
	}

	// // Check the error code that comes back in the Error (JMS uses a String for the code)
	// assert.Equal(t, "2035", err.GetErrorCode())
	// assert.Equal(t, "MQRC_NOT_AUTHORIZED", err.GetReason())

	queue := context.CreateQueue("DEV.QUEUE.1")

	consumer, conErr := context.CreateConsumer(queue)
	if conErr != nil {
		log.Println("conErr", conErr)
		return
	}

	if consumer != nil {
		defer consumer.Close()
	}

	for {
		// rcvBody, _ := consumer.ReceiveStringBodyNoWait()
		rcvBody, _ := consumer.ReceiveNoWait()

		// rcvBytes2, errRcv2 := consumer2.ReceiveBytesBodyNoWait()

		// assert.Equal(t, "MQRC_TRUNCATED_MSG_FAILED", errRcv.GetReason())
		// assert.Equal(t, "2080", errRcv.GetErrorCode())

		if rcvBody != nil {
			// log.Println("Received text string:", *rcvBody, *rcvBody.GetJMSMessageID())

			// msg3 := &rcvBody
			// log.Println(&msg3.GetText())

			switch msg2 := rcvBody.(type) {
			case jms20subset.TextMessage:
				log.Println("Received text string:", *msg2.GetText(), rcvBody.GetJMSMessageID())
			case jms20subset.BytesMessage:
				log.Println("Received text string2:", string(*msg2.ReadBytes()), msg2.GetJMSMessageID())
			default:
				log.Println("Got something other than a text message")
			}
			// log.Println("Received text string:", rcvBody.GetText(), rcvBody.GetJMSMessageID())
			// log.Println("Received text string:", rcvBody.ReadBytes(), rcvBody.GetJMSMessageID())
		} else {
			log.Println("No message received")
			break
		}
	}
}

/*


	mqret := 0
	if err != nil {
		mqret = int((err.(*ibmmq.MQReturn)).MQCC)
	}

	return mqret
}

func mainWithRcGet() int {

	var msgId string
	qMgrName := "QM1"
	qName := "DEV.QUEUE.1"
	fmt.Println("Sample AMQSGET.GO start")
	if len(os.Args) >= 2 {
		qName = os.Args[1]
	}

	if len(os.Args) >= 3 {
		qMgrName = os.Args[2]
	}
	if len(os.Args) >= 4 {
		msgId = os.Args[3]
	}
	qMgrObject, err := ibmmq.Conn(qMgrName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Connected to queue manager %s\n", qMgrName)
		defer disc(qMgrObject)
	}

	if err == nil {
		mqod := ibmmq.NewMQOD()
		openOptions := ibmmq.MQOO_INPUT_EXCLUSIVE
		mqod.ObjectType = ibmmq.MQOT_Q
		mqod.ObjectName = qName
		qObject, err = qMgrObject.Open(mqod, openOptions)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Opened queue", qObject.Name)
			defer close(qObject)
		}
	}

	msgAvail := true
	for msgAvail == true && err == nil {
		var datalen int
		gotMsg := false
		getmqmd := ibmmq.NewMQMD()
		gmo := ibmmq.NewMQGMO()
		gmo.Options = ibmmq.MQGMO_NO_SYNCPOINT
		gmo.Options |= ibmmq.MQGMO_WAIT
		gmo.WaitInterval = 3 * 1000
		if msgId != "" {
			fmt.Println("Setting Match Option for MsgId")
			gmo.MatchOptions = ibmmq.MQMO_MATCH_MSG_ID
			getmqmd.MsgId, _ = hex.DecodeString(msgId)
			msgAvail = false
		}
		useGetSlice := true
		if useGetSlice {
			buffer := make([]byte, 0, 1024)
			buffer, datalen, err = qObject.GetSlice(getmqmd, gmo, buffer)

			if err != nil {
				msgAvail = false
				fmt.Println(err)
				mqret := err.(*ibmmq.MQReturn)
				if mqret.MQRC == ibmmq.MQRC_NO_MSG_AVAILABLE {
					err = nil
				}
			} else {
				fmt.Printf("Got message of length %d: ", datalen)
				fmt.Println(strings.TrimSpace(string(buffer)))
				gotMsg = true
			}
		} else {
			buffer := make([]byte, 1024)
			datalen, err = qObject.Get(getmqmd, gmo, buffer)

			if err != nil {
				msgAvail = false
				fmt.Println(err)
				mqret := err.(*ibmmq.MQReturn)
				if mqret.MQRC == ibmmq.MQRC_NO_MSG_AVAILABLE {
					err = nil
				}
			} else {
				fmt.Printf("Got message of length %d: ", datalen)
				fmt.Println(strings.TrimSpace(string(buffer[:datalen])))
				gotMsg = true
			}
		}

		if gotMsg {
			t := getmqmd.PutDateTime
			if !t.IsZero() {
				diff := time.Now().Sub(t)
				round, _ := time.ParseDuration("1s")
				diff = diff.Round(round)
				fmt.Printf("Message was put %d seconds ago\n", int(diff.Seconds()))
			} else {
				fmt.Printf("Message has empty PutDateTime - MQMD PutDate:'%s' PutTime:'%s'\n", getmqmd.PutDate, getmqmd.PutTime)
			}
		}
	}

	mqret := 0
	if err != nil {
		mqret = int((err.(*ibmmq.MQReturn)).MQCC)
	}
	return mqret
}

func disc(qMgrObject ibmmq.MQQueueManager) error {
	err := qMgrObject.Disc()
	if err == nil {
		fmt.Printf("Disconnected from queue manager %s\n", qMgrObject.Name)
	} else {
		fmt.Println(err)
	}
	return err
}

func close(object ibmmq.MQObject) error {
	err := object.Close(0)
	if err == nil {
		fmt.Println("Closed queue")
	} else {
		fmt.Println(err)
	}
	return err
}


*/
