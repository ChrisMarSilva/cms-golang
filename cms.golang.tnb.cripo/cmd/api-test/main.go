package main

import (
	"context"
	"log"
	"time"
)

// func init() {
// }

func main() {
	// loadConfig()
	loadDatabase()
}

func loadDatabase() {
	db, err := NewDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %s", err.Error())
	}
	defer db.Close()

	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "SELECT nome FROM users WHERE email = ?"
	row := db.QueryRowContext(timeoutCtx, query, "pessoal.01@gmail.com")

	var nome string
	err = row.Scan(&nome)
	if err != nil {
		log.Fatal("Erro no Scan:", err.Error())
	}

	log.Println("ok - Nome:", nome)
}

func loadConfig() {
	cfgOk := Config{
		DbUrl:     "./banco.db",
		JwtSecret: "cms_tamo_em_cripo_api_auth_secret_key",
	}

	cfg, err := NewConfig("./../api-auth/.env")
	if err != nil {
		log.Fatal(err)
	}

	if cfgOk.DbUrl != cfg.DbUrl {
		log.Fatal("DbUrl - Recebido:", cfg.DbUrl, "; Esperado:", cfgOk.DbUrl)
	}

	if cfgOk.DbUrl != cfg.DbUrl {
		log.Fatal("JwtSecret - Recebido:", cfg.JwtSecret, "; Esperado:", cfgOk.JwtSecret)
	}

	log.Println("ok")
}
