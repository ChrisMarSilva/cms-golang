package rotas

import (
	"net/http"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/controllers"
)

var rotaLogin = Rota{"/login", http.MethodPost, controllers.Login, false}
