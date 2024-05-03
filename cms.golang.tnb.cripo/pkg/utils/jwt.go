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
