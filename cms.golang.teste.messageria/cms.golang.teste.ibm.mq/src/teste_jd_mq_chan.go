package src

import (
	"log"
	"strconv"
	"sync"
	"time"
)

func ProcessarJDMQChan() {
	// log.Println("JDMQChan.INI")

	qtdeMessage := 100000 // 1Mil = 1e3 // 10Mil = 1e4 // 100Mil = 1e5 // 1M = 1e6
	qtdeMessageGroup := 5000

	chanMessage := make(chan string)
	chanGroup := make(chan []string)

	go senderMessage(qtdeMessage, chanMessage)

	var wg sync.WaitGroup
	wg.Add(2)

	go receiverMessageSenderGroup(qtdeMessageGroup, chanMessage, chanGroup, &wg)
	go receiverGroup(chanGroup, &wg)

	wg.Wait()

	// log.Println("JDMQChan.FIM")
}

func senderMessage(qtdeMessage int, chanMessage chan<- string) {
	for i := 0; i < qtdeMessage; i++ {
		chanMessage <- "Message #" + strconv.Itoa(i+1) + " " + time.Now().Format(time.RFC3339)
	} // for scanner.Scan() {
	close(chanMessage)
}

func receiverMessageSenderGroup(qtdeMessageGroup int, chanMessage <-chan string, chanGroup chan<- []string, wg *sync.WaitGroup) {
	defer wg.Done()

	var messages []string

	for message := range chanMessage {
		messages = append(messages, message)
		if len(messages) >= qtdeMessageGroup {
			chanGroup <- messages
			messages = []string{}
		}
	}

	if len(messages) > 0 {
		chanGroup <- messages
	}

	close(chanGroup)
}

func receiverGroup(chanGroup <-chan []string, wg *sync.WaitGroup) {

	defer wg.Done()

	var wgLocal sync.WaitGroup
	var start time.Time

	start = time.Now()
	qtdePut := 0

	for messages := range chanGroup {
		wgLocal.Add(1)
		qtdePut += len(messages)

		go func(ww *sync.WaitGroup, messages []string) {
			defer ww.Done()
			jdmq := JDMQ{}
			jdmq.Conectar("QM.04358798.01")
			defer jdmq.Desconectar()
			jdmq.AbrirFilaPut("FL.REQ.INT")
			defer jdmq.FecharFila()
			for _, message := range messages {
				jdmq.EnviarMensagem(message)
				jdmq.Commit()
			}
			jdmq.Commit()
		}(&wgLocal, messages)

	}

	wgLocal.Wait()
	log.Println("JDMQChan.PUT:", qtdePut, "TEMPO:", time.Since(start))

}
