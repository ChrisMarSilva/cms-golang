package src

import (
	_ "encoding/hex"
	"log"
	_ "strings"
	_ "sync"
	"time"
)

// https://github.com/ibm-messaging/mq-golang-jms20/tree/master/mqjms

// go mod init github.com/ChrisMarSilva/cms.golang.messageria.ibm.mq
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

func ProcessarJDMQ() {
	// log.Println("JDMQ.INI")

	var start time.Time

	jdmq := JDMQ{} // new(entities.JDMQ) // entities.JDMQ{}

	// log.Println("JDMQ.JDMQ.Conectar()")
	//start = time.Now()
	if !jdmq.Conectar("QM.04358798.01") {
		log.Println("JDMQ.JDMQ.Conectar - Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
		return
	}
	//log.Println("JDMQ.CONN:", time.Since(start))

	//defer jdmq.Desconectar()

	defer func() {
		// log.Println("JDMQ.JDMQ.Desconectar()")
		if !jdmq.Desconectar() {
			log.Println("JDMQ.JDMQ.Desconectar - Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
			return
		}
	}()

	//log.Println("JDMQ.JDMQ.AbrirFilaPut()")
	//start = time.Now()
	if !jdmq.AbrirFilaPut("FL.REQ.INT") {
		log.Println("JDMQ.JDMQ.AbrirFilaPut - Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
		return
	}
	//log.Println("JDMQ.OPEN:", time.Since(start))

	// defer jdmq.FecharFila()

	defer func() {
		// log.Println("JDMQ.JDMQ.FecharFila()")
		if !jdmq.FecharFila() {
			log.Println("JDMQ.JDMQ.FecharFila - Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
			return
		}
	}()

	// var wg sync.WaitGroup
	// var mu sync.Mutex

	start = time.Now()
	qtdePut := 1e5 // 1Mil = 1e3 // 10Mil = 1e4 // 100Mil = 1e5 // 1M = 1e6
	for i := 0; i < int(qtdePut); i++ {
		// wg.Add(1)
		// go func(jdmq *entities.JDMQ, wg *sync.WaitGroup) {
		// 	mu.Lock()
		// 	defer mu.Unlock()
		// 	defer wg.Done()
		message := "Hello from Go at " + time.Now().Format(time.RFC3339)
		// jdmq.EnviarMensagem(message)
		if !jdmq.EnviarMensagem(message) {
			log.Println("JDMQ.JDMQ.EnviarMensagem - Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
			// return
		}
		// log.Println("JDMQ.PUT - Message:", jdmq.GetMessage(), "MsgId:", jdmq.GetMsgId())
		// if !jdmq.Commit() {
		// 	log.Println("JDMQ.JDMQ.Commit - Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
		// 	// return
		// }
		//}(&jdmq, &wg)
	}

	// wg.Wait()

	if !jdmq.Commit() {
		log.Println("JDMQ.JDMQ.Commit - Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
		return
	}

	log.Println("JDMQ.PUT:", qtdePut, "TEMPO:", time.Since(start))

	//log.Println("JDMQ.JDMQ.FecharFila()")
	if !jdmq.FecharFila() {
		log.Println("JDMQ.Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
		return
	}

	//log.Println("JDMQ.JDMQ.AbrirFilaGet()")
	if !jdmq.AbrirFilaGet("FL.REQ.INT") {
		log.Println("JDMQ.JDMQ.AbrirFilaGet - Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
		return
	}

	qtdeGet := 0
	start = time.Now()
	for {
		// wg.Add(1)
		// go func(jdmq *entities.JDMQ, wg *sync.WaitGroup) {
		// 	mu.Lock()
		// 	defer mu.Unlock()
		// 	defer wg.Done()
		if !jdmq.ReceberMensagem() {
			if jdmq.GetErr() != nil {
				log.Println("JDMQ.JDMQ.ReceberMensagem - Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
			}
			break
		}
		//log.Println("JDMQ.GET - Message:", jdmq.GetMessage(), "MsgId:", jdmq.GetMsgId())
		// if !jdmq.Commit() {
		// 	log.Println("JDMQ.JDMQ.Commit - Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
		// 	break
		// }
		//}(&jdmq, &wg)
		qtdeGet++
	}

	// wg.Wait()

	if !jdmq.Commit() {
		log.Println("JDMQ.JDMQ.Commit - Erro - Code:", jdmq.GetCode(), "- Reason:", jdmq.GetReason())
		return
	}

	log.Println("JDMQ.GET:", qtdeGet, "TEMPO:", time.Since(start))

	// log.Println("JDMQ.FIM")
}
