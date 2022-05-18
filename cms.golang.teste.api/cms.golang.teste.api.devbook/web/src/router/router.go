package router

import (
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-web/src/router/rotas"
	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
