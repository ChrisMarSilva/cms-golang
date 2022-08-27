package src

import (
	_ "encoding/hex"
	ibmmq "github.com/ibm-messaging/mq-golang/v5/ibmmq"
	"log"
	_ "strings"
	_ "sync"
	"time"
)

// https://github.com/ibm-messaging/mq-golang-jms20/tree/master/mqjms

// go mod init github.com/chrismarsilva/cms.golang.messageria.ibm.mq
// go get github.com/ibm-messaging/mq-golang
// go get github.com/ibm-messaging/mq-golang/v5/ibmmq
// go mod tidy

// cd "C:\Users\chris\go\src\github.com\ibm-messaging\mq-golang\ibmmq"
// go install .

// cd "C:\Users\chris\go\src\github.com\ibm-messaging\mq-golang\samples"
// go install amqsput.go

// go env GOROOT=C:\Program Files\Go
// go env GOPATH=C:\Users\chris\go
// go env CC=x86_64-w64-mingw32-gcc.exe
// go env CGO_CFLAGS=-Ic:\IBM-MQC-Redist-Win64\tools\c\include -D_WIN64
// go env CGO_LDFLAGS=-Lc:\IBM-MQC-Redist-Win64\bin64 -lmqm

// go env MQ_INSTALLATION_PATH=C:\Program Files\IBM\MQ
// go env CGO_CFLAGS="-I$MQ_INSTALLATION_PATH/inc"
// go env CGO_LDFLAGS="-L$MQ_INSTALLATION_PATH/lib64 -Wl,-rpath,$MQ_INSTALLATION_PATH/lib64"

// cd "C:\Users\chris\Desktop\CMS GoLang\cms.golang.teste.messageria\scms.golang.teste.ibm.mq"
// go run main.go
// go run .

//     1.000 =      579.5936ms PUT -      637.8954ms GET
//    10.000 =    5.3306907s   PUT -    4.4541837s   GET
//    10.000 =    6.5983536s   PUT -    5.9481486s   GET
//   100.000 =   43.9125417s   PUT -   24.1743972s   GET
//   100.000 =   41.4951124s   PUT -   29.8787814s   GET
// 1.000.000 = 5m55.2423373s   PUT - 4m56.2554574s   GET

func ProcessarIBMMQ() {
	// log.Println("IBMMQ.INI")

	var conn ibmmq.MQQueueManager
	var fila ibmmq.MQObject

	conn, err := ibmmq.Conn("QM.04358798.01")
	if err != nil {
		log.Println("IBMMQ.ibmmq.Conn", err)
		return
	}

	defer func() {
		err = conn.Disc()
		if err != nil {
			log.Println("IBMMQ.conn.Disc", err)
			return
		}
	}()

	mqod := ibmmq.NewMQOD()
	mqod.ObjectType = ibmmq.MQOT_Q
	mqod.ObjectName = "FL.REQ.INT"

	var openOptions int32
	openOptions = ibmmq.MQOO_OUTPUT

	fila, err = conn.Open(mqod, openOptions)
	if err != nil {
		log.Println("IBMMQ.conn.Open.MQOO_OUTPUT", err)
		return
	}

	defer func() {
		err := fila.Close(0)
		if err != nil {
			log.Println("IBMMQ.fila.Close", err)
			return
		}
	}()

	putmqmd := ibmmq.NewMQMD()
	putmqmd.Format = ibmmq.MQFMT_STRING

	pmo := ibmmq.NewMQPMO()
	pmo.Options = ibmmq.MQPMO_SYNCPOINT

	var start time.Time
	start = time.Now()
	qtdePut := 100000
	for i := 0; i < qtdePut; i++ {
		msgData := "Hello from Go at " + time.Now().Format(time.RFC3339)
		buffer := []byte(msgData)
		err = fila.Put(putmqmd, pmo, buffer)
		if err != nil {
			log.Println("IBMMQ.qObject.Put", err)
			return
		}
		err = conn.Cmit()
	}
	log.Println("IBMMQ.PUT:", qtdePut, "TEMPO:", time.Since(start))

	openOptions = ibmmq.MQOO_INPUT_EXCLUSIVE
	fila, err = conn.Open(mqod, openOptions)
	if err != nil {
		log.Println("IBMMQ.conn.Open.MQOO_INPUT_EXCLUSIVE", err)
		return
	}

	qtdeGet := 0
	start = time.Now()
	msgAvail := true
	//var datalen int
	for msgAvail == true && err == nil {

		getmqmd := ibmmq.NewMQMD()

		gmo := ibmmq.NewMQGMO()
		gmo.Options = ibmmq.MQGMO_SYNCPOINT
		gmo.Options |= ibmmq.MQGMO_NO_WAIT

		buffer := make([]byte, 1024)
		_, err = fila.Get(getmqmd, gmo, buffer)
		if err != nil {
			msgAvail = false
			if err.(*ibmmq.MQReturn).MQRC == ibmmq.MQRC_NO_MSG_AVAILABLE {
				err = nil
				continue
			}
			log.Println("IBMMQ.fila.Get", err)
			continue
		}
		err = conn.Cmit()
		if err != nil {
			log.Println("IBMMQ.conn.Cmit", err)
			return
		}
		qtdeGet++
	}

	log.Println("IBMMQ.GET:", qtdeGet, "TEMPO:", time.Since(start))
	// log.Println("IBMMQ.datalen", datalen)

	// log.Println("IBMMQ.FIM")
}
