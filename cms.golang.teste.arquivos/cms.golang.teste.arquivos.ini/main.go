package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

// go get gopkg.in/ini.v1
// go get github.com/vaughan0/go-ini

func main() {

	filename := "conf.ini"

	cfg, err := ini.Load(filename)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())
	fmt.Println("Server Protocol:", cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	fmt.Println("Email Protocol:", cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))
	fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.SaveTo("my.ini.local")

	fmt.Print("FIM")
}
