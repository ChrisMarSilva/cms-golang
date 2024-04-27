package utils

// import (
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// 	"golang.org/x/crypto/bcrypt"
// )

// id, _ := strconv.Atoi(c.Param("id"))

// func NormalizeEmail(email string) string {
// 	return strings.TrimSpace(strings.ToLower(email))
// }

// var jwtKey = []byte("my_secret_key")

// func HashPassword(password string) (string, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// 	if err != nil {
// 		return "", err
// 	}

// 	return string(hashedPassword), nil
// }

// import "golang.org/x/crypto/bcrypt"

// func GenerateHashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes), err
// }

// func GenerateJWT(user UserModel) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id":        user.ID,
// 		"nome":      user.Nome,
// 		"email":     user.Email,
// 		"is_active": user.IsActive,
// 		//"role":      user.Role,
// 		"iat": time.Now().Unix(),
// 		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
// 	})

// 	//tokenStr, err := token.SignedString([]byte(SecretKey))
// 	signedToken, err := token.SignedString([]byte("secret-key"))
// 	if err != nil {
// 		return "", err
// 	}

// 	return signedToken, nil
// }
