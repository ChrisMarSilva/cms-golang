package utils

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// 	"go.elastic.co/apm/model"
// )

// func GenerateJWT(user model.User) (string, error) {
// 	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id":   user.ID,
// 		"role": user.RoleID,
// 		"iat":  time.Now().Unix(),
// 		"eat":  time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
// 	})

// 	return token.SignedString(privateKey)
// }

// func ValidateJWT(context *gin.Context) error {
// 	token, err := getToken(context)
// 	if err != nil {
// 		return err
// 	}

// 	_, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		return nil
// 	}

// 	return errors.New("invalid token provided")
// }

// func ValidateAdminRoleJWT(context *gin.Context) error {
// 	token, err := getToken(context)
// 	if err != nil {
// 		return err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)

// 	userRole := uint(claims["role"].(float64))
// 	if ok && token.Valid && userRole == 1 {
// 		return nil
// 	}

// 	return errors.New("invalid admin token provided")
// }

// func ValidateCustomerRoleJWT(context *gin.Context) error {
// 	token, err := getToken(context)
// 	if err != nil {
// 		return err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)

// 	userRole := uint(claims["role"].(float64))
// 	if ok && token.Valid && userRole == 2 || userRole == 1 {
// 		return nil
// 	}

// 	return errors.New("invalid author token provided")
// }

// // fetch user details from the token
// func CurrentUser(context *gin.Context) model.User {
// 	err := ValidateJWT(context)
// 	if err != nil {
// 		return model.User{}
// 	}
// 	token, _ := getToken(context)
// 	claims, _ := token.Claims.(jwt.MapClaims)
// 	userId := uint(claims["id"].(float64))

// 	user, err := model.GetUserById(userId)
// 	if err != nil {
// 		return model.User{}
// 	}
// 	return user
// }

// // check token validity
// func getToken(context *gin.Context) (*jwt.Token, error) {
// 	tokenString := getTokenFromRequest(context)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return privateKey, nil
// 	})
// 	return token, err
// }

// // extract token from request Authorization header
// func getTokenFromRequest(context *gin.Context) string {
// 	bearerToken := context.Request.Header.Get("Authorization")
// 	splitToken := strings.Split(bearerToken, " ")
// 	if len(splitToken) == 2 {
// 		return splitToken[1]
// 	}
// 	return ""
// }

// func ParseToken(tokenStr string) (claims jwt.StandardClaims, err error) {
// 	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(Config.JwtSecretKey), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	claims, ok := token.Claims.(jwt.StandardClaims)
// 	if !ok {
// 		return nil, err
// 	}

// 	return claims, nil
// }

// func generateJwt() (string, error) {
// 	tokenByte := jwt.New(jwt.SigningMethodHS256)

// 	claims := tokenByte.Claims.(jwt.MapClaims)
// 	//claims["sub"] = user.ID
// 	claims["exp"] = time.Now().UTC().Add(time.Hour * 24 * 7).Unix()
// 	claims["iat"] = time.Now().UTC().Unix()
// 	claims["nbf"] = time.Now().UTC().Unix()

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(config.JWTSecret)
// }

// func GenerateToken(userID uint) (string, error) {
// 	claims := jwt.MapClaims{}
// 	claims["user_id"] = userID
// 	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token valid for 1 hour

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(secretKey)
// }

// func VerifyToken(tokenString string) (jwt.MapClaims, error) {
// 	// Parse the token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Check the signing method
// 		if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
// 			return nil, fmt.Errorf("Invalid signing method")
// 		}

// 		return secretKey, nil
// 	})

// 	// Check for errors
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Validate the token
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims, nil
// 	}

// 	return nil, fmt.Errorf("Invalid token")
// }

// func GetJWT() (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["authorized"] = true
// 	claims["client"] = "Krissanawat"
// 	claims["aud"] = "billing.jwtgo.io"
// 	claims["iss"] = "jwtgo.io"
// 	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

// 	tokenString, err := token.SignedString(mySigningKey)

// 	if err != nil {
// 		fmt.Errorf("Something Went Wrong: %s", err.Error())
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// type JWTClaim struct {
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	jwt.StandardClaims
// }

// type Claims struct {
//     Username string `json:"username"`
//     jwt.RegisteredClaims
// }

// func GenerateJWT(email string, username string) (tokenString string, err error) {
// 	expirationTime := time.Now().Add(1 * time.Hour)
// 	claims := &JWTClaim{
// 		Email:    email,
// 		Username: username,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err = token.SignedString(jwtKey)
// 	return
// }

// func ValidateToken(signedToken string) (err error) {
// 	token, err := jwt.ParseWithClaims(
// 		signedToken,
// 		&JWTClaim{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(jwtKey), nil
// 		},
// 	)
// 	if err != nil {
// 		return
// 	}
// 	claims, ok := token.Claims.(*JWTClaim)
// 	if !ok {
// 		err = errors.New("couldn't parse claims")
// 		return
// 	}
// 	if claims.ExpiresAt < time.Now().Local().Unix() {
// 		err = errors.New("token expired")
// 		return
// 	}
// 	return
// }

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

// func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		tokenString := utils.GetTokenFromRequest(r)

// 		token, err := validateJWT(tokenString)
// 		if err != nil {
// 			log.Printf("failed to validate token: %v", err)
// 			permissionDenied(w)
// 			return
// 		}

// 		if !token.Valid {
// 			log.Println("invalid token")
// 			permissionDenied(w)
// 			return
// 		}

// 		claims := token.Claims.(jwt.MapClaims)
// 		str := claims["userID"].(string)

// 		userID, err := strconv.Atoi(str)
// 		if err != nil {
// 			log.Printf("failed to convert userID to int: %v", err)
// 			permissionDenied(w)
// 			return
// 		}

// 		u, err := store.GetUserByID(userID)
// 		if err != nil {
// 			log.Printf("failed to get user by id: %v", err)
// 			permissionDenied(w)
// 			return
// 		}

// 		Add the user to the context
// 		ctx := r.Context()
// 		ctx = context.WithValue(ctx, UserKey, u.ID)
// 		r = r.WithContext(ctx)

// 		Call the function if the token is valid
// 		handlerFunc(w, r)
// 	}
// }

// func CreateJWT(secret []byte, userID int) (string, error) {
// 	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"userID":    strconv.Itoa(int(userID)),
// 		"expiresAt": time.Now().Add(expiration).Unix(),
// 	})

// 	tokenString, err := token.SignedString(secret)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, err
// }

// func validateJWT(tokenString string) (*jwt.Token, error) {
// 	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return []byte(configs.Envs.JWTSecret), nil
// 	})
// }

// func permissionDenied(w http.ResponseWriter) {
// 	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
// }

// func GetUserIDFromContext(ctx context.Context) int {
// 	userID, ok := ctx.Value(UserKey).(int)
// 	if !ok {
// 		return -1
// 	}

// 	return userID
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

/*

package utils

import (
    "os"
    "strings"

    "github.com/golang-jwt/jwt"
    "github.com/gofiber/fiber/v2"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
    Expires int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
    token, err := verifyToken(c)
    if err != nil {
        return nil, err
    }

    // Setting and checking token and credentials.
    claims, ok := token.Claims.(jwt.MapClaims)
    if ok && token.Valid {
        // Expires time.
        expires := int64(claims["exp"].(float64))

        return &TokenMetadata{
            Expires: expires,
        }, nil
    }

    return nil, err
}

func extractToken(c *fiber.Ctx) string {
    bearToken := c.Get("Authorization")

    // Normally Authorization HTTP header.
    onlyToken := strings.Split(bearToken, " ")
    if len(onlyToken) == 2 {
        return onlyToken[1]
    }

    return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
    tokenString := extractToken(c)

    token, err := jwt.Parse(tokenString, jwtKeyFunc)
    if err != nil {
        return nil, err
    }

    return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
    return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}


// ./pkg/utils/jwt_generator.go

package utils

import (
    "os"
    "strconv"
    "time"

    "github.com/golang-jwt/jwt"
)

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken() (string, error) {
    // Set secret key from .env file.
    secret := os.Getenv("JWT_SECRET_KEY")

    // Set expires minutes count for secret key from .env file.
    minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

    // Create a new claims.
    claims := jwt.MapClaims{}

    // Set public claims:
    claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

    // Create a new JWT access token with claims.
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Generate token.
    t, err := token.SignedString([]byte(secret))
    if err != nil {
        // Return error, it JWT token generation failed.
        return "", err
    }

    return t, nil
}

// ./pkg/middleware/jwt_middleware.go

package middleware

import (
    "os"

    "github.com/gofiber/fiber/v2"

    jwtMiddleware "github.com/gofiber/jwt/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
    // Create config for JWT authentication middleware.
    config := jwtMiddleware.Config{
        SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
        ContextKey:   "jwt", // used in private routes
        ErrorHandler: jwtError,
    }

    return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
    // Return status 401 and failed authentication error.
    if err.Error() == "Missing or malformed JWT" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    // Return status 401 and failed authentication error.
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": true,
        "msg":   err.Error(),
    })
}

// ./app/controllers/token_controller.go

package controllers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/koddr/tutorial-go-fiber-rest-api/pkg/utils"
)

// GetNewAccessToken method for create a new access token.
// @Description Create a new access token.
// @Summary create a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/token/new [get]
func GetNewAccessToken(c *fiber.Ctx) error {
    // Generate a new Access token.
    token, err := utils.GenerateNewAccessToken()
    if err != nil {
        // Return status 500 and token generation error.
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "error":        false,
        "msg":          nil,
        "access_token": token,
    })
}


// ...

// DeleteBook func for deletes book by given ID.
// @Description Delete book by given ID.
// @Summary delete book by given ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id body string true "Book ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/book [delete]
func DeleteBook(c *fiber.Ctx) error {
    // Get now time.
    now := time.Now().Unix()

    // Get claims from JWT.
    claims, err := utils.ExtractTokenMetadata(c)
    if err != nil {
        // Return status 500 and JWT parse error.
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    // Set expiration time from JWT data of current book.
    expires := claims.Expires

    // Checking, if now time greather than expiration from JWT.
    if now > expires {
        // Return status 401 and unauthorized error message.
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": true,
            "msg":   "unauthorized, check expiration time of your token",
        })
    }

    // Create new Book struct
    book := &models.Book{}

    // Check, if received JSON data is valid.
    if err := c.BodyParser(book); err != nil {
        // Return status 400 and error message.
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    // Create a new validator for a Book model.
    validate := utils.NewValidator()

    // Validate only one book field ID.
    if err := validate.StructPartial(book, "id"); err != nil {
        // Return, if some fields are not valid.
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": true,
            "msg":   utils.ValidatorErrors(err),
        })
    }

    // Create database connection.
    db, err := database.OpenDBConnection()
    if err != nil {
        // Return status 500 and database connection error.
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    // Checking, if book with given ID is exists.
    foundedBook, err := db.GetBook(book.ID)
    if err != nil {
        // Return status 404 and book not found error.
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": true,
            "msg":   "book with this ID not found",
        })
    }

    // Delete book by given ID.
    if err := db.DeleteBook(foundedBook.ID); err != nil {
        // Return status 500 and error message.
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    // Return status 204 no content.
    return c.SendStatus(fiber.StatusNoContent)
}
*/

package middlewares

import (
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v2"
	"github.com/jailtonjunior94/bookings/api/infrastructure/environments"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(environments.JwtSecret),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT ausente ou malformado"})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido ou expirado"})
}


package adapters

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jailtonjunior94/bookings/api/infrastructure/environments"
)

type IJwtAdapter interface {
	GenerateTokenJWT(id, email string) (r string, err error)
	ExtractClaims(tokenString string) (id *string, err error)
}

type JwtAdapter struct {
}

func NewJwtAdapter() IJwtAdapter {
	return &JwtAdapter{}
}

func (j *JwtAdapter) GenerateTokenJWT(id, email string) (r string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Local().Add(time.Hour * 24 * time.Duration(1))

	t, err := token.SignedString([]byte(environments.JwtSecret))
	if err != nil {
		return "", err
	}
	return t, err
}

func (j *JwtAdapter) ExtractClaims(tokenString string) (id *string, err error) {
	tokenString = strings.Split(tokenString, " ")[1]
	hmacSecret := []byte(environments.JwtSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Token inválido")
	}

	sub := fmt.Sprintf("%v", claims["sub"])
	return &sub, nil
}