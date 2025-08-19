package utils

const (
	DBDriverDefault = "postgres"
	DBDriverPgx     = "pgx"
	DBUri           = "host=localhost port=5432 dbname=postgres user=postgres password=postgres sslmode=disable"
	//DBUri    = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	DBQuery = `SELECT "id", "name", "created_at" FROM "TbPerson" ORDER BY "created_at"`

	DBInsert    = `INSERT INTO "TbPerson" ("id", "name", "created_at") VALUES ($1, $2, $3)`
	DBUpdate    = `UPDATE "TbPerson" SET "name" = $1, "created_at" = $2 WHERE "id" = $3`
	DBSelectOne = `SELECT "id", "name", "created_at" FROM "TbPerson" WHERE "id" = $1`
	DBSelectAll = `SELECT "id", "name", "created_at" FROM "TbPerson"`
	DBDeleteOne = `DELETE FROM "TbPerson" WHERE "id" = $1`
	DBDeleteAll = `DELETE FROM "TbPerson"`
	DBTruncate  = `TRUNCATE TABLE "TbPerson"`

	DBDeletePgxPool = `DELETE FROM "TbPerson" WHERE "name" LIKE 'PgxPool%'`
	DBDeleteGorm    = `DELETE FROM "TbPerson" WHERE "name" LIKE 'Gorm%'`
	DBDeleteSqlX    = `DELETE FROM "TbPerson" WHERE "name" LIKE 'SqlX%'`
	DBDeleteSqlC    = `DELETE FROM "TbPerson" WHERE "name" LIKE 'SqlC%'`
)
