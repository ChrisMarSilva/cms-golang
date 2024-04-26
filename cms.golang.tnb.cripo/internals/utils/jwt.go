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





jwt.go
package util

import (
	"bmacharia/jwt-go-rbac/model"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

// retrieve JWT key from .env file
var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

// generate JWT token
func GenerateJWT(user model.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.RoleID,
		"iat":  time.Now().Unix(),
		"eat":  time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

// validate JWT token
func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

// validate Admin role
func ValidateAdminRoleJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 1 {
		return nil
	}
	return errors.New("invalid admin token provided")
}

// validate Customer role
func ValidateCustomerRoleJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 2 || userRole == 1 {
		return nil
	}
	return errors.New("invalid author token provided")
}

// fetch user details from the token
func CurrentUser(context *gin.Context) model.User {
	err := ValidateJWT(context)
	if err != nil {
		return model.User{}
	}
	token, _ := getToken(context)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	user, err := model.GetUserById(userId)
	if err != nil {
		return model.User{}
	}
	return user
}

// check token validity
func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

// extract token from request Authorization header
func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}





package utils

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secretpassword")

// GenerateToken generates a JWT token with the user ID as part of the claims
func GenerateToken(userID uint) (string, error) {
    claims := jwt.MapClaims{}
    claims["user_id"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token valid for 1 hour

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

// VerifyToken verifies a token JWT validate 
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
    // Parse the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Check the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
            return nil, fmt.Errorf("Invalid signing method")
        }

        return secretKey, nil
    })

    // Check for errors
    if err != nil {
        return nil, err
    }

    // Validate the token
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("Invalid token")
}




func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
  
	claims := token.Claims.(jwt.MapClaims)
  
	claims["authorized"] = true
	claims["client"] = "Krissanawat"
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
  
	tokenString, err := token.SignedString(mySigningKey)
  
	if err != nil {
	  fmt.Errorf("Something Went Wrong: %s", err.Error())
	  return "", err
	}
  
	return tokenString, nil
  }
  

  
var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
  Username string `json:"username"`
  Email    string `json:"email"`
  jwt.StandardClaims
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
  expirationTime := time.Now().Add(1 * time.Hour)
  claims:= &JWTClaim{
    Email: email,
    Username: username,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: expirationTime.Unix(),
    },
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err = token.SignedString(jwtKey)
  return
}

func ValidateToken(signedToken string) (err error) {
  token, err := jwt.ParseWithClaims(
    signedToken,
    &JWTClaim{},
    func(token *jwt.Token) (interface{}, error) {
      return []byte(jwtKey), nil
    },
  )
  if err != nil {
    return
  }
  claims, ok := token.Claims.(*JWTClaim)
  if !ok {
    err = errors.New("couldn't parse claims")
    return
  }
  if claims.ExpiresAt < time.Now().Local().Unix() {
    err = errors.New("token expired")
    return
  }
  return
}



var secretKey = []byte("secretpassword")

// GenerateToken generates a JWT token with the user ID as part of the claims
func GenerateToken(userID uint) (string, error) {
    claims := jwt.MapClaims{}
    claims["user_id"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token valid for 1 hour

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

// VerifyToken verifies a token JWT validate 
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
    // Parse the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Check the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
            return nil, fmt.Errorf("Invalid signing method")
        }

        return secretKey, nil
    })

    // Check for errors
    if err != nil {
        return nil, err    }

    // Validate the token
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("Invalid token")
}



type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims:= &JWTClaim{
		Email: email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")a
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	
	return

}





package security

import (
	"fmt"
	"go-fiber-auth-api/util"
	"os"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
)

var (
	JwtSecretKey     = []byte(os.Getenv("JWT_SECRET_KEY"))
	JwtSigningMethod = jwt.SigningMethodHS256.Name
)

func NewToken(userId string) (string, error) {
	claims := jwt.StandardClaims{
		Id:        userId,
		Issuer:    userId,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecretKey)
}

func validateSignedMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return JwtSecretKey, nil
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	claims := new(jwt.StandardClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, validateSignedMethod)
	if err != nil {
		return nil, err
	}
	var ok bool
	claims, ok = token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, util.ErrInvalidAuthToken
	}
	return claims, nil
}