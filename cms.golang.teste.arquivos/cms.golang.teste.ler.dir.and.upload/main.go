package main

import (
	"io"
	"log"
	"os"
	"time"
	"fmt"
	"sync"
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	// "github.com/aws/aws-sdk-go/aws/awserr"
	// "github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.ler.dir.and.upload
// go get -u github.com/aws/aws-sdk-go
// go mod tidy

// go run main.go

var (
	client *s3.S3
	wg     sync.WaitGroup
)

func init() {
	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, "myRoleArn")
	client = s3.New(sess, &aws.Config{Credentials: creds})

	// creds := credentials.NewCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	// region := aws.String("us-east-1") // aws.StringValue(sess.Config.Region)

	// sess, err := session.NewSession(&aws.Config{Credentials: creds, Region: region})
	// if err != nil {
	// 	panic(err)
	// }

	// client = s3.New(sess)
}

func main() {
	log.Println("INI")
	var start time.Time = time.Now()

	path := "C:\\Users\\chris\\AppData\\Local\\Temp"
	dir, err := os.Open(path)
	if err != nil {
		log.Println("Error os.Open(Path): " + err.Error())
		return
	}

	semaforo := make(chan struct{}, 3) // 1_000
	for {
		files, err := dir.Readdir(1) // ler 1 arquivo por vez
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error os.Readdir(): " + err.Error())
			continue
		}
		// log.Println("Name: " + files[0].Name())
		wg.Add(1)
		semaforo <- struct{}{}
		go upload(files[0], semaforo)
	}

	wg.Wait()

	log.Println("FIM:", time.Since(start))
}

func upload(fileInfo os.FileInfo, semaforo <-chan struct{}) {
	defer wg.Done()

	fmt.Printf("Upload started: %s\n", fileInfo.Name())

	filepath := fmt.Sprintf("C:\\Users\\chris\\AppData\\Local\\Temp\\%s", fileInfo.Name())
	file, err := os.Open(filepath)
	if err != nil {
		<-semaforo
		// log.Println("Error os.Open(FilePath): " + err.Error())
		fmt.Printf("ERRO file: %v\n", err)
		return
	}
	defer file.Close()

	var fileSize int64 = fileInfo.Size()

	fileBuffer := make([]byte, fileSize)
	file.Read(fileBuffer)

	_, err = client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("2kk-aprenda-golang"),
		Key:    aws.String(fileInfo.Name()),
		Body:   bytes.NewReader(fileBuffer),
	})
	if err != nil {
		<-semaforo
		// log.Println("Error client.PutObject(): " + err.Error())
		fmt.Printf("ERRO upload: %v\n", err)
		return
	}

	fmt.Printf("Upload finished: %s\n", fileInfo.Name())
	<-semaforo
}
