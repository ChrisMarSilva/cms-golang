package utils

import (
	"go-auth/models"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenString string) (claims *models.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}

func ParseToken(tokenString string) (claims *models.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}


r.GET("/resource", func(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	reqToken := strings.Split(bearerToken, " ")[1]
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized",})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request",})
		return
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
})


func generateJWT() (string, error) {
expirationTime := time.Now().Add(5 * time.Minute)
claims := &Claims{
	Username: "username",
	RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	},
}

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

return token.SignedString(jwtKey)

}


cookie := new(http.Cookie)
cookie.Name = "token"
cookie.Value = tokenString
cookie.Expires = expirationTime.Time
c.SetCookie(cookie)

