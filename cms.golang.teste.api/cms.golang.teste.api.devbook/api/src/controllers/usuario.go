package controllers

import (
	"strings"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"database/sql"
	"strconv"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/banco"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/models"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/repository"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/seguranca"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/autenticacao"
	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	defer r.Body.Close()

	var usuario models.Usuario

	err = json.Unmarshal(corpoRequest, &usuario)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	hash, err := seguranca.Hash(usuario.Senha)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	usuario.Senha = string(hash)

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	repositorio := repository.NovoUsuario(db)

	usuario.ID, err = repositorio.Criar(usuario)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	usuario.Senha = ""

	RespondWithJSON(w, http.StatusCreated, usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	repositorio := repository.NovoUsuario(db)

	usuarios, err := repositorio.Buscar(nomeOuNick)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if len(usuarios) == 0 {
		RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, usuarios)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(vars["usuarioId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	repositorio := repository.NovoUsuario(db)

	usuario, err := repositorio.BuscarPorID(usuarioID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	usuario.Senha = ""

	RespondWithJSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(vars["usuarioId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	if usuarioID != usuarioIDNoToken {
		RespondWithError(w, http.StatusForbidden, "User different") 
		return
	}

	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	defer r.Body.Close()

	var usuario models.Usuario

	err = json.Unmarshal(corpoRequest, &usuario)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	usuario.ID = usuarioID

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	repositorio := repository.NovoUsuario(db)

	_, err = repositorio.Atualizar(usuario)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	
	// if rowsAffected != 1 {
	// 	RespondWithError(w, http.StatusNotFound, "User not updated")
	// 	return
	// }

	RespondWithJSON(w, http.StatusNoContent, nil)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(vars["usuarioId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	if usuarioID != usuarioIDNoToken {
		RespondWithError(w, http.StatusForbidden, "User different") 
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	repositorio := repository.NovoUsuario(db)

	err = repositorio.Deletar(usuarioID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	//RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	RespondWithJSON(w, http.StatusNoContent, nil)
}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(vars["usuarioId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	if usuarioID != usuarioIDNoToken {
		RespondWithError(w, http.StatusForbidden, "User different") 
		return
	}

	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	defer r.Body.Close()

	var senha models.Senha
	err = json.Unmarshal(corpoRequest, &senha)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	repositorio := repository.NovoUsuario(db)

	usuarioBanco, err := repositorio.BuscarPorID(usuarioID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	err = seguranca.VerificarSenha(senha.Atual, usuarioBanco.Senha)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	
	hash, err := seguranca.Hash(senha.Nova)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	usuario := models.Usuario{}
	usuario.ID = usuarioID
	usuario.Senha = string(hash)
	
	err = repositorio.AtualizarSenha(usuario)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(vars["usuarioId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	seguidorID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	if usuarioID == seguidorID {
		RespondWithError(w, http.StatusForbidden, "Same User") 
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	repositorio := repository.NovoSeguidor(db)

	err = repositorio.Criar(usuarioID, seguidorID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}

func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(vars["usuarioId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	seguidorID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	if usuarioID == seguidorID {
		RespondWithError(w, http.StatusForbidden, "Same User") 
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	repositorio := repository.NovoSeguidor(db)

	err = repositorio.Deletar(usuarioID, seguidorID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(vars["usuarioId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()
	
	repositorio := repository.NovoUsuario(db)

	usuarios, err := repositorio.BuscarSeguidores(usuarioID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "follower not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if len(usuarios) == 0 {
		RespondWithError(w, http.StatusNotFound, "Follower not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, usuarios)
}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(vars["usuarioId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()
	
	repositorio := repository.NovoUsuario(db)

	usuarios, err := repositorio.BuscarSeguindo(usuarioID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "follower not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if len(usuarios) == 0 {
		RespondWithError(w, http.StatusNotFound, "Follower not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, usuarios)
}

func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(vars["usuarioId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()
	
	repositorio := repository.NovaPublicacao(db)

	publicacoes, err := repositorio.BuscarPorUsuarioID(usuarioID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "publucation not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if len(publicacoes) == 0 {
		RespondWithError(w, http.StatusNotFound, "publucation not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, publicacoes)
}
