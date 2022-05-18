package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"database/sql"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/banco"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/models"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/repository"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/seguranca"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/autenticacao"
)

func Login(w http.ResponseWriter, r *http.Request) {

	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	defer r.Body.Close()

	var usuarioRequest models.Usuario
	err = json.Unmarshal(corpoRequest, &usuarioRequest)
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

	usuarioBanco, err := repositorio.BuscarPorEmail(usuarioRequest.Email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	err = seguranca.VerificarSenha(usuarioRequest.Senha, usuarioBanco.Senha)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := autenticacao.CriarToken(usuarioBanco.ID)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	//RespondWithJSON(w, http.StatusOK, token)
	RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}