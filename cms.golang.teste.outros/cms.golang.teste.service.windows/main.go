package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kardianos/service"
)

// go build main.go
// main -install service
// main -service start

// go build main.go
// main -service install
// main -service start

// go build main.go
// sc.exe create "CMS Golang Teste" binPath="C:\Users\chris\Desktop\CMS GoLang\cms.golang.teste.service.windows\main.exe"
// sc.exe delete "CMS Golang Teste"

var logger service.Logger

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}

func (p *program) run() {
	logger.Infof("I'm running %v.", service.Platform())

	fo, err := os.Create("C:\\Users\\chris\\Desktop\\CMS GoLang\\cms.golang.teste.service.windows\\filename.txt")
	if err != nil {
		logger.Error(err)
	}

	// defer fo.Close()
	defer func() {
		if err := fo.Close(); err != nil {
			logger.Error(err)
		}
	}()

	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case tm := <-ticker.C:
			fo.Write([]byte(fmt.Sprint("Hello ", tm, " \n")))
			logger.Infof("Still running at %v...", tm)
		case <-p.exit:
			ticker.Stop()
			return
		}
	}
}

func main() {

	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"

	// svcConfig := &service.Config{Name: "GoServiceExampleSimple", DisplayName: "Go Service Example", Description: "This is an example Go service."}

	svcConfig := &service.Config{
		Name:         "GoServiceExampleLogging",
		DisplayName:  "Go Service Example for Logging",
		Description:  "This is an example Go service that outputs log messages.",
		Dependencies: []string{"Requires=network.target", "After=network-online.target syslog.target"},
		Option:       options,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	// if len(os.Args) > 1 {
	// 	err = service.Control(s, os.Args[1])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	return
	// }

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}

}
