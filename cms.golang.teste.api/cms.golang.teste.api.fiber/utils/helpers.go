package util

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)
	return string(hashed)
}

func HashOnlyVulnerable(pass []byte) string {
	hash := md5.New()
	hash.Write(pass)
	return hex.EncodeToString(hash.Sum(nil))
}

// // Create panic handler
// func PanicHandler(next http.Handler) http.Handler{
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			error := recover()
// 			if error != nil {
// 				log.Println(error)

// 				resp := interfaces.ErrResponse{Message: "Internal server error"}
// 				json.NewEncoder(w).Encode(resp)
// 			}
// 		}()
// 		next.ServeHTTP(w, r)
// 	})
// }

// func ValidateToken(id string, jwtToken string) bool {
// 	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
// 	tokenData := jwt.MapClaims{}
// 	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
// 			return []byte("TokenPassword"), nil
// 	})
// 	HandleErr(err)
// 	var userId, _ = strconv.ParseFloat(id, 8)
// 	if token.Valid && tokenData["user_id"] == userId {
// 		return true
// 	} else {
// 		return false
// 	}
// }
