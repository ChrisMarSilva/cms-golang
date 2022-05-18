package middlewares

import (
	"log"
	"net/http"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/autenticacao"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/controllers"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host, r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
    })
}

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := autenticacao.ValidarToken(r)
		if err != nil {
			controllers.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		next(w, r)
	}
}