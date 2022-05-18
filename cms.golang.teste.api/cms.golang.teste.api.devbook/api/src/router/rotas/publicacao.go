package rotas

import (
	"net/http"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/controllers"
)

var rotasPublicacoes = []Rota{
	{"/publicacoes", http.MethodPost, controllers.CriarPublicacao, true},
	{"/publicacoes", http.MethodGet, controllers.BuscarPublicacoes, true},
	{"/publicacoes/{publicacaoId}", http.MethodGet, controllers.BuscarPublicacao, true},
	{"/publicacoes/{publicacaoId}", http.MethodPut, controllers.AtualizarPublicacao, true},
	{"/publicacoes/{publicacaoId}", http.MethodDelete, controllers.DeletarPublicacao, true},
	{"/publicacoes/{publicacaoId}/curtir", http.MethodPost, controllers.CurtirPublicacao, true},
	{"/publicacoes/{publicacaoId}/descurtir", http.MethodPost, controllers.DescurtirPublicacao, true},
}
