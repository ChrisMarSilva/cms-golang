package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.request.api
// go mod tidy

// go run main.go
// go build main.go

func main() {

	tot := 10_000 // 1_000 // 10_000 // 100_000 // 1_000_000

	var wg sync.WaitGroup

	// wg.Add(1)
	// // go TesteURL(&wg, tot, "http://127.0.0.1:5000/product/1", "Get.09.Python.Flask")
	// go TesteURLAsync(&wg, tot, "http://localhost:5000/", "Get.09.Python.Flask  ")

	// wg.Add(1)
	// // go TesteURL(&wg, tot, "http://127.0.0.1:5001/", "Get.11.Python.FastAPI")
	// go TesteURLAsync(&wg, tot, "http://localhost:5001/", "Get.11.Python.FastAPI")

	wg.Add(1)
	// go TesteURL(&wg, tot, "http://127.0.0.1:8002/", "Get.04.Fiber    ")
	go TesteURLAsync(&wg, tot, "http://localhost:8002/", "Get.04.Fiber         ")

	// wg.Add(1)
	// // go TesteURL(&wg, tot, "http://127.0.0.1:3000/", "Get.03.Gin      ")
	// go TesteURLAsync(&wg, tot, "http://localhost:3000/", "Get.03.Gin           ")


	//go TesteURLAsync(&wg, tot, "http://localhost:7003/", "Get.03.Gin      ") //  Requisições:100.000 - Tempo:  50.1625159s - Erros: 87.227
	//go TesteURLAsync(&wg, tot, "http://localhost:7004/", "Get.04.Fiber    ")   //  Requisições:100.000 - Tempo: 1m8.1356397s - Erros: 62.281
	//go TesteURLAsync(&wg, tot, "http://localhost:8011/", "Get.09.Python.Flask    ") //  Requisições:100.000 - Tempo: 1m30.7456535s - Erros: 94.024
	//go TesteURLAsync(&wg, tot, "https://localhost:44332/WeatherForecast/ok", "Get.10.DotNet5    ") //  Requisições:100.000 - Tempo: 1m30.7456535s - Erros: 94.024

	// wg.Add(8)
	// go TesteURL(&wg, tot, "http://localhost:7001/", "Get.01.Net      ")
	// go TesteURL(&wg, tot, "http://localhost:7002/", "Get.02.Gorilla  ")
	// go TesteURL(&wg, tot, "http://localhost:7003/", "Get.03.Gin      ")
	// go TesteURL(&wg, tot, "http://localhost:7004/", "Get.04.Fiber    ")
	// go TesteURL(&wg, tot, "http://localhost:7005/", "Get.05.Echo     ")
	// go TesteURL(&wg, tot, "http://localhost:7006/", "Get.06.Beego    ")
	// go TesteURL(&wg, tot, "http://localhost:7007/", "Get.07.Fasthttp ")
	// go TesteURL(&wg, tot, "http://localhost:7008/", "Get.08.Atreugo  ")

	// wg.Add(8)
	// go TesteURLAsync(&wg, tot, "http://localhost:7001/", "Get.01.Net      ")
	// go TesteURLAsync(&wg, tot, "http://localhost:7002/", "Get.02.Gorilla  ")
	// go TesteURLAsync(&wg, tot, "http://localhost:7003/", "Get.03.Gin      ")
	// go TesteURLAsync(&wg, tot, "http://localhost:7004/", "Get.04.Fiber    ")
	// go TesteURLAsync(&wg, tot, "http://localhost:7005/", "Get.05.Echo     ")
	// go TesteURLAsync(&wg, tot, "http://localhost:7006/", "Get.06.Beego    ")
	// go TesteURLAsync(&wg, tot, "http://localhost:7007/", "Get.07.Fasthttp ")
	// go TesteURLAsync(&wg, tot, "http://localhost:7008/", "Get.08.Atreugo  ")

	wg.Wait()

	log.Println("Fim")

}

func TesteURL(ww *sync.WaitGroup, tot int, url string, nome string) {

	defer ww.Done()

	var start time.Time = time.Now()
	var iErro uint64 = 0
	var iOk uint64 = 0

	for i := 1; i <= tot; i++ {
		// if i%100 == 0 {
		// 	log.Println(nome, "iIdx:", i, "iErro: ", iErro)
		// }
		//_, err := http.Get(url)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		myClient := &http.Client{Timeout: time.Second * 10}
		_, getErr := myClient.Do(req)
		if err != nil || getErr != nil {
			iErro++
			//log.Println("http.Get.erro", err, "iIdx:", i)
		} else {
			iOk++
		}
		// body, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	log.Println("ioutil.ReadAll", err)
		// }
		// log.Println("body", string(body))
		// req.Body.Close()
	} // for i := 1; i <= tot; i++ {

	log.Println(nome+" ("+strconv.Itoa(tot)+"): ", time.Since(start), "     -> url:", url, "iOk: ", iOk, "iErro: ", iErro)
}

func TesteURLAsync(ww *sync.WaitGroup, tot int, url string, nome string) {

	defer ww.Done()

	var wgLocal sync.WaitGroup
	var mLocal sync.Mutex

	var start time.Time = time.Now()
	var iErro uint64 = 0
	var iOk uint64 = 0

	for i := 1; i <= tot; i++ {
		if i%100 == 0 {
			time.Sleep(time.Millisecond * 200)
		}
		wgLocal.Add(1)
		go func(wwLocal *sync.WaitGroup, iIdx int) {
			defer wwLocal.Done()

			//_, err := http.Get(url)

			req, err := http.NewRequest(http.MethodGet, url, nil)
			myClient := &http.Client{Timeout: time.Second * 10}
			res, getErr := myClient.Do(req)

			if err != nil || getErr != nil {
				mLocal.Lock()
				iErro++
				mLocal.Unlock()
				//atomic.AddUint64(&iErro, 1)
				//log.Println("http.Get.erro", err, "iIdx:", iIdx)
			} else {
				// log.Println(nome + " - "+strconv.Itoa(res.StatusCode))
				if res.StatusCode != http.StatusOK {
					mLocal.Lock()
					iErro++
					mLocal.Unlock()
				} else {
					mLocal.Lock()
					iOk++
					mLocal.Unlock()
				}
			}

		}(&wgLocal, i)
	} // for i := 1; i <= tot; i++ {

	wgLocal.Wait()

	log.Println(nome+" ("+strconv.Itoa(tot)+"): ", time.Since(start), "     -> url:", url, "iOk: ", iOk, "iErro: ", iErro)
}

func Teste() {

	log.Println("Ini")

	url := "http://localhost:3000/delete"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("http.NewRequest", err)
	}
	myClient := &http.Client{Timeout: time.Second * 10}
	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Println("myClient.Do", err)
	}
	if res.Body != nil {
		res.Body.Close()
	}

	var start time.Time
	var wg sync.WaitGroup

	start = time.Now()
	tot := 10000 // 10000=10.000= // 100000=100.000=
	for i := 1; i <= tot; i++ {

		if i%100 == 0 {
			time.Sleep(time.Millisecond * 100)
		}

		wg.Add(1)
		go func(ww *sync.WaitGroup, iIdx int) {

			defer ww.Done()

			url := "http://localhost:3000/insert/" + strconv.Itoa(iIdx)

			for i := 0; i < 1000; i++ {
				_, err := http.Get(url)
				if err != nil {
					//log.Println("http.Get.erro", err, "iIdx:", iIdx, "i:", i)
					time.Sleep(time.Millisecond * 100)
					continue
				}
				break
			}

			// if req.Body != nil {
			// 	req.Body.Close()
			// }

			// req, err := http.NewRequest(http.MethodGet, url, nil)
			// if err != nil {
			// 	log.Println("http.NewRequest", err)
			// }

			// myClient := &http.Client{Timeout: time.Second * 10}
			// res, getErr := myClient.Do(req)
			// if getErr != nil {
			// 	log.Println("myClient.Do", err)
			// }

			// if res.Body != nil {
			// 	res.Body.Close()
			// }

			// log.Println("iIdx", iIdx, "res.StatusCode", res.StatusCode, " - ", res.Status)
			// body, err := ioutil.ReadAll(res.Body)
			// if err != nil {
			// 	log.Println("ioutil.ReadAll", err)
			// }
			// log.Println("body", body)

		}(&wg, i)
	}

	wg.Wait()

	log.Println("http.Get ("+strconv.Itoa(tot)+"): ", time.Since(start)) // 9m23.7323813s

	log.Println("Fim")

}
