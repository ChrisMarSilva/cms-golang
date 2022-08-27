package main

import (
	"fmt"
	"log"
    "github.com/spf13/viper"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.arquivos.viper
// go get github.com/spf13/viper
// go mod tidy

// go run main.go

func main() {
	
	viper.AddConfigPath(".")  

	// ------------------------------------------------
	// ------------------------------------------------

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig() 
	if err != nil { 
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	log.Println("env.PORT", viper.Get("PORT"))
	log.Println("env.MONGO_URI", viper.Get("MONGO_URI"))
	log.Println("env.API_SECRET", viper.Get("API_SECRET"))

	// ------------------------------------------------
	// ------------------------------------------------
	
	viper.SetConfigName("conf") 
	viper.SetConfigType("json") 
	err = viper.ReadInConfig() 
	if err != nil { 
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	log.Println("json.name", viper.Get("users.name"))
	log.Println("json.type", viper.Get("users.type"))
	log.Println("json.age", viper.Get("users.age"))

	// ------------------------------------------------
	// ------------------------------------------------

	viper.SetConfigName("conf") 
	viper.SetConfigType("yaml") 
	err = viper.ReadInConfig() 
	if err != nil { 
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	log.Println("yaml.conf", viper.Get("conf"))
	log.Println("yaml.hits", viper.GetInt("conf.hits"))
	log.Println("yaml.time", viper.Get("conf.time"))
	log.Println("yaml.camelCase", viper.Get("conf.camelCase"))

	// ------------------------------------------------
	// ------------------------------------------------
	
	viper.SetConfigName("conf") 
	viper.SetConfigType("ini") 
	err = viper.ReadInConfig() 
	if err != nil { 
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// log.Println("ini.app_mode", viper.GetString("app_mode"))
	log.Println("ini.data", viper.Get("paths.data"))
	log.Println("ini.protocol", viper.Get("server.protocol"))

	// ------------------------------------------------
	// ------------------------------------------------


	log.Println("FIM")
}
