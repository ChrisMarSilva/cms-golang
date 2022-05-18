package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"database/sql"
	"strconv"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/banco"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/models"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/repository"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/autenticacao"
	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {

	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	defer r.Body.Close()

	var publicacao models.Publicacao

	err = json.Unmarshal(corpoRequest, &publicacao)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	publicacao.AutorID = usuarioIDNoToken

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	repositorio := repository.NovaPublicacao(db)

	publicacao.ID, err = repositorio.Criar(publicacao)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, publicacao)
}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	repositorio := repository.NovaPublicacao(db)

	publicacoes, err := repositorio.Buscar(usuarioIDNoToken)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Publication not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if len(publicacoes) == 0 {
		RespondWithError(w, http.StatusNotFound, "Publication not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, publicacoes)
}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(vars["publicacaoId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid publication ID")
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	repositorio := repository.NovaPublicacao(db)

	publicacao, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Publication not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusOK, publicacao)
}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(vars["publicacaoId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid publication ID")
		return
	}

	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	defer r.Body.Close()

	var publicacaoRequest models.Publicacao
	err = json.Unmarshal(corpoRequest, &publicacaoRequest)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	publicacaoRequest.ID = publicacaoID
	
	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	repositorio := repository.NovaPublicacao(db)

	publicacaoBanco, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Publication not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if usuarioIDNoToken != publicacaoBanco.AutorID {
		RespondWithError(w, http.StatusForbidden, "User different")
		return
	}

	_, err = repositorio.Atualizar(publicacaoRequest)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Publication not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(vars["publicacaoId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid publication ID")
		return
	}
	
	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	repositorio := repository.NovaPublicacao(db)

	publicacaoBanco, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Publication not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if usuarioIDNoToken != publicacaoBanco.AutorID {
		RespondWithError(w, http.StatusForbidden, "User different")
		return
	}

	err = repositorio.Deletar(publicacaoID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Publication not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(vars["publicacaoId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid publication ID")
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	repositorio := repository.NovaPublicacao(db)

	err = repositorio.Curtir(publicacaoID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Publication not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(vars["publicacaoId"], 10, 64) // base 10 // 64 bits
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid publication ID")
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	repositorio := repository.NovaPublicacao(db)

	err = repositorio.Descurtir(publicacaoID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Publication not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}