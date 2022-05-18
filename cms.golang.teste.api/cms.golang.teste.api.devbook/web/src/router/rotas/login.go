package rotas

import (
	"net/http"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-web/src/controllers"
)

var rotasLogin = []Rota{
	{"/", http.MethodGet, controllers.CarregarTelaDeLogin, false},
	{"/login", http.MethodGet, controllers.CarregarTelaDeLogin, false},
}
