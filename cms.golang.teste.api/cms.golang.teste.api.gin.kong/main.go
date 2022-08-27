package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kong/go-kong/kong"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.api.gin.kong
// go get -u github.com/gin-gonic/gin
// go get github.com/kong/go-kong/kong
// go mod tidy

// go run main.go

func main() {

	serviceName, serviceHost, routeName, servicePort := GetConfingEnv()

	KongAdd(serviceName, serviceHost, routeName, servicePort)

	router := gin.Default()
	router.GET("/", handleHome)

	log.Println("Listem port " + servicePort)
	router.Run(":" + servicePort)

}

func handleHome(c *gin.Context) {
	time.Sleep(1 * time.Second)
	c.JSON(http.StatusOK, gin.H{"Cliente": "hello world"})
}

func GetConfingEnv() (string, string, string, string) {

	// os.Setenv("GO_API_SERVICE", "servicego")
	kongService := strings.TrimSpace(os.Getenv("GO_API_SERVICE"))
	if kongService == "" {
		kongService = "servicego"
	}

	// os.Setenv("GO_API_HOST", "servicego")
	kongHost := strings.TrimSpace(os.Getenv("GO_API_HOST"))
	if kongHost == "" {
		kongHost = "servicego"
	}

	// os.Setenv("GO_API_ROUTE", "teste")
	kongRoute := strings.TrimSpace(os.Getenv("GO_API_ROUTE"))
	if kongRoute == "" {
		kongRoute = "cliente"
	}

	// os.Setenv("GO_API_PORT", "3000")
	kongPort := strings.TrimSpace(os.Getenv("GO_API_PORT"))
	if kongPort == "" {
		kongPort = "3000"
	}

	// // Config Padrao
	kongRoute = "cliente"
	kongService = "servicego"

	// // Configurar Rota Normal
	// kongHost = "servicego"
	// kongPort = "3000"

	// // Configurar Rota Upstream - Load Balancer
	// kongHost = "servicego2"
	// kongPort = "3002"

	// // Configurar Rota Upstream - Load Balancer
	// kongHost = "servicego3"
	// kongPort = "3003"

	// kongHost = "servicego4"
	// kongPort = "3004"

	return kongService, kongHost, kongRoute, kongPort
}

func KongAdd(serviceName string, serviceHost string, routeName string, servicePort string) {

	ctx := context.Background()

	client, err := kong.NewTestClient(nil, nil)
	if err != nil {
		log.Fatalln("Kong\t - kong.NewTestClient.Error:", err)
	}

	err = KongAddService(ctx, *client, serviceName, serviceHost, servicePort)
	if err != nil {
		log.Println("Kong\t - Service com erro:", err)
	}

	err = KongAddRoute(ctx, *client, serviceName, routeName)
	if err != nil {
		log.Println("Kong\t - Route com erro:", err)
	}

	upstreamName := serviceName + "_upstream"

	err = KongAddUpstream(ctx, *client, serviceName, serviceHost, upstreamName)
	if err != nil {
		log.Println("Kong\t - Upstream com erro:", err)
	}

	targetName := serviceHost + ":" + servicePort

	err = KongAddTargetsUpstream(ctx, *client, upstreamName, targetName)
	if err != nil {
		log.Println("Kong\t - Target com erro:", err)
	}

}

func KongAddService(ctx context.Context, client kong.Client, serviceName string, serviceHost string, servicePort string) error {

	createdService, err := client.Services.Get(ctx, &serviceName)
	if err == nil || createdService != nil {
		log.Println("Service\t - Serviço localizado\t -", *createdService.Name)
		return nil
	}

	servicPath := "/"
	iport, err := strconv.Atoi(servicePort)
	if err != nil {
		log.Println("Service\t - Porta invalida\t -", err)
		return err
	}

	service := &kong.Service{Name: &serviceName, Host: &serviceHost, Port: &iport, Path: &servicPath}

	createdService, err = client.Services.Create(ctx, service)
	if err != nil || createdService == nil {
		log.Println("Service\t - Serviço não criado\t -", err)
		return err
	}

	log.Println("Service\t - Serviço criado\t - Name:", *createdService.Name, "ID:", *createdService.ID)

	return nil
}

func KongAddRoute(ctx context.Context, client kong.Client, serviceName string, routeName string) error {

	createdRoute, err := client.Routes.Get(ctx, &routeName)
	if err == nil {
		log.Println("Routes\t - Routa localizada\t -", *createdRoute.Name)
		return nil
	}

	service, err := client.Services.Get(ctx, &serviceName)
	if err != nil || service == nil {
		log.Println("Routes\t - Serviço não localizado\t -", err)
		return err
	}

	serviceRoutePath := "/" + routeName
	var listRoutes []*string
	listRoutes = append(listRoutes, &serviceRoutePath)

	route := &kong.Route{Name: &routeName, Paths: listRoutes, Service: service}

	createdRoute, err = client.Routes.Create(ctx, route)
	if err != nil || createdRoute == nil {
		log.Println("Routes\t - Rota não criada\t -", err)
		return err
	}

	log.Println("Routes\t - Rota criada\t - Name:", *createdRoute.Name, "ID:", *createdRoute.ID)

	return nil
}

func KongAddUpstream(ctx context.Context, client kong.Client, serviceName string, serviceHost string, upstreamName string) error {

	createdUpstream, err := client.Upstreams.Get(ctx, &upstreamName)
	if err == nil {
		log.Println("Upstreams\t - Upstream localizada\t -", *createdUpstream.Name)
		return nil
	}

	service, err := client.Services.Get(ctx, &serviceName)
	if err != nil || service == nil {
		log.Println("Upstreams\t - Serviço não localizado\t -", err)
		return err
	}

	targetNameOld := *service.Host + ":" + strconv.Itoa(*service.Port)

	// Service: Atualiza service

	*service.Host = upstreamName
	*service.Port = 8000

	service, err = client.Services.Update(ctx, service)
	if err != nil {
		log.Println("Upstreams\t - Serviço não atualizado\t -", err)
		return err
	}

	// Upstreams: Criar upstream

	upstream := &kong.Upstream{Name: &upstreamName}

	createdUpstream, err = client.Upstreams.Create(ctx, upstream)
	if err != nil || upstream == nil {
		log.Println("Upstreams\t - Upstream não criado\t -", err)
		return err
	}

	log.Println("Upstreams\t - Upstream criado\t - Name:", *createdUpstream.Name, "ID:", *createdUpstream.ID)

	// Targets: Criar Target antigo

	err = KongAddTargetsUpstream(ctx, client, upstreamName, targetNameOld)
	if err != nil {
		log.Println("Upstreams\t - Targets Old com erro\t -", err)
	}

	return nil
}

func KongAddTargetsUpstream(ctx context.Context, client kong.Client, upstreamName string, targetName string) error {

	upstream, err := client.Upstreams.Get(ctx, &upstreamName)
	if err != nil || upstream == nil {
		log.Println("Targets\t - Upstream localizado\t -", err)
		return err
	}

	targets, err := client.Targets.ListAll(ctx, &upstreamName)
	if err != nil {
		log.Println("Targets\t - Erro ao buscar lista\t -", err)
		return err
	}

	bExistTarget := false

	for _, row := range targets {
		if *row.Target == targetName {
			bExistTarget = true
			break
		}
	}

	if bExistTarget {
		log.Println("Targets\t - Target localizado\t -", targetName)
		return nil
	}

	target := &kong.Target{Target: &targetName}

	createdTarget, err := client.Targets.Create(ctx, upstream.ID, target)
	if err != nil || createdTarget == nil {
		log.Println("Targets\t - Target não criado\t -", err)
		return err
	}

	log.Println("Targets\t - Target criado\t - Name:", *createdTarget.Target, "ID:", *createdTarget.ID)

	return nil
}

// go run main.go
