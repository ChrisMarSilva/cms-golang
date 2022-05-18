package rotas

import (
	"net/http"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/controllers"
)

var rotasUsuarios = []Rota{
	{"/usuarios", http.MethodPost, controllers.CriarUsuario, false},
	{"/usuarios", http.MethodGet, controllers.BuscarUsuarios, true},
	{"/usuarios/{usuarioId}", http.MethodGet, controllers.BuscarUsuario, true},
	{"/usuarios/{usuarioId}", http.MethodPut, controllers.AtualizarUsuario, true},
	{"/usuarios/{usuarioId}", http.MethodDelete, controllers.DeletarUsuario, true},
	{"/usuarios/{usuarioId}/atualizar-senha", http.MethodPost, controllers.AtualizarSenha, true},
	{"/usuarios/{usuarioId}/seguir", http.MethodPost, controllers.SeguirUsuario, true},
	{"/usuarios/{usuarioId}/parar-de-seguir", http.MethodPost, controllers.PararDeSeguirUsuario, true},
	{"/usuarios/{usuarioId}/seguidores", http.MethodGet, controllers.BuscarSeguidores, true},
	{"/usuarios/{usuarioId}/seguindo", http.MethodGet, controllers.BuscarSeguindo, true},
	{"/usuarios/{usuarioId}/publicacoes", http.MethodGet, controllers.BuscarPublicacoesPorUsuario, true},
}
