package router

import (
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/router/rotas"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/middlewares"
	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(middlewares.Logger)
	return rotas.Configurar(r)
}
