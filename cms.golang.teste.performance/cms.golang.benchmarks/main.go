package main

import (
	"log"

	gofrs "github.com/gofrs/uuid/v5"
	google "github.com/google/uuid"
	shortuuid "github.com/lithammer/shortuuid/v4"
	satori "github.com/satori/go.uuid"
	ksuid "github.com/segmentio/ksuid"
)

// go mod init github.com/chrismarsilva/cms.golang.benchmarks

// json
// go get -u github.com/json-iterator/go
// go get -u github.com/bytedance/sonic

// uuid
// go get -u github.com/google/uuid
// go get -u github.com/gofrs/uuid/v5
// go get -u github.com/satori/go.uuid
// go get -u github.com/segmentio/ksuid
// go get -u github.com/lithammer/shortuuid/v4

// go mod tidy

// go run main.go

func main() {
	log.Println("Running benchmarks...")

	google_uuid := google.New()
	log.Println("Google UUID:", google_uuid)

	gofrs_uuid, _ := gofrs.NewV4()
	log.Println("Gofrs UUID:", gofrs_uuid)

	satori_uuid := satori.NewV4()
	log.Println("Satori UUID:", satori_uuid)

	ksuid_uuid := ksuid.New()
	log.Println("ksuid UUID:", ksuid_uuid)

	shortuuid_uuid := shortuuid.New()
	log.Println("ShortUUID UUID:", shortuuid_uuid)
}
