package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

func init() {
	log.Println("")
}

func main() {
	// loadConfig()
	// loadDatabase()
	loadConRepository()
}

func loadConRepository() {
	db, err := NewDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %s", err.Error())
	}
	defer db.Close()

	userRepo := NewUserRepository(db)

	ctx := context.Background()

	user, err := userRepo.GetByEmail(ctx, nil, "pessoal.01@gmail.com")
	if err != nil {
		log.Fatalf("Erro ao buscar usuário: %s", err.Error())
	}
	log.Println("ok - GetByEmail - Nome:", user.Nome, "- Email:", user.Email)

	dbRepo := NewDBRepo(db)

	err = dbRepo.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		id := uuid.New()
		newUser := NewUserEntity(id, "Pessoa "+id.String(), "pessoal"+id.String()+"@gmail.com", true, time.Now())

		err := userRepo.Create(ctx, tx, newUser)
		if err != nil {
			return err
		}
		log.Println("ok - Create - Nome:", newUser.Nome, "- Email:", newUser.Email)

		userCreated, err := userRepo.GetByEmail(ctx, tx, newUser.Email)
		if err != nil {
			log.Fatalf("Erro ao buscar usuário: %s", err.Error())
		}
		log.Println("ok - GetByEmail - Nome:", userCreated.Nome, "- Email:", userCreated.Email)

		return nil
	})

	if err != nil {
		log.Fatalf("Erro ao inserir usuário: %s", err.Error())
	}
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
