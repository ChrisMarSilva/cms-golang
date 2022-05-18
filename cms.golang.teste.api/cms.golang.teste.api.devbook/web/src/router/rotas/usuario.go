package rotas

import (
	"net/http"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-web/src/controllers"
)

var rotasUsuarios = []Rota{
	{"/criar-usuario", http.MethodGet, controllers.CarregarTelaDeCadastroDeUsuario, false},
}
