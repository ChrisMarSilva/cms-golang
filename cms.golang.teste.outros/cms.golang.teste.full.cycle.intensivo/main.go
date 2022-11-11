package main

import (
	"log"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.intensivo
// go get -u github.com/stretchr/testify
// go get -u github.com/mattn/go-sqlite3
// go get -u github.com/rabbitmq/amqp091-go
// go get -u github.com/google/uuid
// go mod tidy

// DROP TABLE orders;
// CREATE TABLE orders (id varchar(255) not null, price float not null, tax float not null, final_price float not null, primary key(id));
// delete from orders
// truncate TABLE orders

// docker-compose down
// docker-compose up -d --build
// docker-compose up -d

// docker build -t chrismarsilva/gointensivo:latest -f Dockerfile.prod .
// docker images |grep gointensivo

// kind create cluster --name=gointensivo
// kubectl cluster-info --context king-gointensivo
// kubectl apply -f k8s/

// go run main.go
// go run cmd/consumer/main.go
// go run cmd/producer/main.go

// {"id": "002", "price": 100.99, "tax": 1.01}

func main() {
	log.Println("ok")
}
