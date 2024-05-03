
TDD 
Test Driven Development
Desenvolvimento Orientado a Testes
https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/injecao-de-dependencia

go mod init github.com/chrismarsilva/cms.golang.testes
go get -u github.com/stretchr/testify
go get -u github.com/stretchr/testify/assert
go get -u github.com/stretchr/testify/require
go get -u github.com/stretchr/testify/mock
go get -u github.com/stretchr/testify/suite
go get -u github.com/go-delve/delve/cmd/dlv
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
go mod tidy

go run .
go run ola.go

go install github.com/kisielk/errcheck@latest
errcheck .
errcheck github.com/chrismarsilva/cms.golang.testes

go fmt

go test
go test -v

go test -bench=.

go test -cover
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out
go test -covermode=count -coverprofile=count.out fmt
go tool cover -func=count.out
go tool cover -html=count.out

