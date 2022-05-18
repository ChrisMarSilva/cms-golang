package util

// import (
// 	"fmt"
// 	"go-fiber-auth-api/util"
// 	"os"
// 	"time"

// 	jwt "github.com/form3tech-oss/jwt-go"
// )

// var (
// 	JwtSecretKey     = []byte(os.Getenv("JWT_SECRET_KEY"))
// 	JwtSigningMethod = jwt.SigningMethodHS256.Name
// )

// func NewToken(userId string) (string, error) {
// 	claims := jwt.StandardClaims{
// 		Id:        userId,
// 		Issuer:    userId,
// 		IssuedAt:  time.Now().Unix(),
// 		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(JwtSecretKey)
// }

// func validateSignedMethod(token *jwt.Token) (interface{}, error) {
// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 	}
// 	return JwtSecretKey, nil
// }

// func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
// 	claims := new(jwt.StandardClaims)
// 	token, err := jwt.ParseWithClaims(tokenString, claims, validateSignedMethod)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var ok bool
// 	claims, ok = token.Claims.(*jwt.StandardClaims)
// 	if !ok || !token.Valid {
// 		return nil, util.ErrInvalidAuthToken
// 	}
// 	return claims, nil
// }

// func AuthRequestWithId(ctx *fiber.Ctx) (*jwt.StandardClaims, error) {
// 	id := ctx.Params("id")
// 	if !bson.IsObjectIdHex(id) {
// 		return nil, util.ErrUnauthorized
// 	}
// 	token := ctx.Locals("user").(*jwt.Token)
// 	payload, err := security.ParseToken(token.Raw)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if payload.Id != id || payload.Issuer != id {
// 		return nil, util.ErrUnauthorized
// 	}
// 	return payload, nil
// }

// func jwtMiddleware(secret string) fiber.Handler {
// 	return jwtware.New(jwtware.Config{
// 		SigningKey: []byte(secret),
// 	})
// }

// func signToken(tokenKey, id string) string {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["admin"] = true
// 	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
// 	claims["sub"] = id

// 	t, err := token.SignedString([]byte(tokenKey))

// 	if err != nil {
// 		return ""
// 	}

// 	return t
// }

// func extractUserIDFromJWT(bearer, tokenKey string) string {
// 	//Bearer ey.....
// 	token := bearer[7:]
// 	logs.Info(token)
// 	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
// 		return []byte(tokenKey), nil
// 	})

// 	if err != nil {
// 		return ""
// 	}

// 	if t.Valid {
// 		claims := t.Claims.(jwt.MapClaims)
// 		return claims["sub"].(string)
// 	}

// 	return ""
// }

/*

package security

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidToken = errors.New("invalid jwt")

	jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

func NewToken(userId string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		Issuer:    userId,
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func parseJwtCallback(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return jwtSecretKey, nil
}

func ExtractToken(r *http.Request) (string, error) {
	// Authorization => Bearer Token...
	header := strings.TrimSpace(r.Header.Get("Authorization"))
	splitted := strings.Split(header, " ")
	if len(splitted) != 2 {
		log.Println("error on extract token from header:", header)
		return "", ErrInvalidToken
	}
	return splitted[1], nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, parseJwtCallback)
}

type TokenPayload struct {
	UserId    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewTokenPayload(tokenString string) (*TokenPayload, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return nil, ErrInvalidToken
	}
	id, _ := claims["iss"].(string)
	createdAt, _ := claims["iat"].(float64)
	expiresAt, _ := claims["exp"].(float64)
	return &TokenPayload{
		UserId:    id,
		CreatedAt: time.Unix(int64(createdAt), 0),
		ExpiresAt: time.Unix(int64(expiresAt), 0),
	}, nil
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := security.ExtractToken(r)
		if err != nil {
			restutil.WriteError(w, http.StatusUnauthorized, restutil.ErrUnauthorized)
			return
		}
		token, err := security.ParseToken(tokenString)
		if err != nil {
			log.Println("error on parse token:", err.Error())
			restutil.WriteError(w, http.StatusUnauthorized, restutil.ErrUnauthorized)
			return
		}
		if !token.Valid {
			log.Println("invalid token", tokenString)
			restutil.WriteError(w, http.StatusUnauthorized, restutil.ErrUnauthorized)
			return
		}

		next(w, r)
	}
}


*/
