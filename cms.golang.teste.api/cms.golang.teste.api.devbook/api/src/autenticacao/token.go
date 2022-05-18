package autenticacao

import (
	"time"
	"net/http"
	"strings"
	"fmt"
	"errors"
	"strconv"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/config"
)

func CriarToken(usuarioID uint64) (string, error) {

	// atClaims := jwt.MapClaims{} 
	permissoes := jwt.MapClaims{} 
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioID"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString(config.SecretKey) 
}

func ValidarToken(r *http.Request) error {

	tokenString := extrairToken(r)

	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return err
	}
	
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return errors.New("Token invalido")
}

func ExtrairUsuarioID(r *http.Request) (uint64, error) {

	tokenString := extrairToken(r)
	
	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return 0, err
	}

	permissoes, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		usuarioID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioID"]), 10, 64) // base 10 // 64 bits
		if err != nil {
			return 0, err
		}
		return usuarioID, nil
	}
	
	return 0,errors.New("Token invalido")
}

func extrairToken(r *http.Request) string {
	tokenString := r.Header.Get("Authorization") 
	tokenSlice := strings.Split(tokenString, " ")
	if len(tokenSlice) == 2 {
		return tokenSlice[1]
	}
	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {	
	_, ok := token.Method.(*jwt.SigningMethodHMAC)	
	if !ok {
		return nil, fmt.Errorf("Metodo de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
