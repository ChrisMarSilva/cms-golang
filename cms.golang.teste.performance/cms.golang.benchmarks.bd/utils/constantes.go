package utils

const (
	DBDriverDefault = "postgres"
	DBDriverPgx     = "pgx"
	DBUri           = "host=localhost port=5432 dbname=postgres user=postgres password=postgres sslmode=disable"
	//DBUri    = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	DBQuery = `SELECT "id", "name", "created_at" FROM "TbPerson" ORDER BY "created_at"`
)
