package utils

// import (
// 	"context"
// 	"crypto/tls"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	jwtware "github.com/gofiber/contrib/jwt"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt"
// 	"github.com/google/uuid"
// 	"github.com/labstack/echo"
// 	"github.com/spf13/viper"
// 	"gorm.io/gorm"
// )

// func NewAuthMiddleware(secret string) fiber.Handler {
// 	return jwtware.New(jwtware.Config{
// 		SigningKey: jwtware.SigningKey{Key: []byte(secret)}, // []byte(secret),
// 	})
// }

// func AdminMiddleware(c *fiber.Ctx) error {
// 	// userRole := getUserRoleFromContext(c)
// 	// if userRole != AdminRole {
// 	// 	return c.Status(fiber.StatusForbidden).SendString("Permission Denied")
// 	// }
// 	return func(c *fiber.Ctx) error {
// 		claim, _ := c.Get("claim").(Claims)

// 		if claim.IsNotAdmin() {
// 			return c.String(http.StatusForbidden, "You have no authority")
// 		}

// 		return c.Next()
// 	}
// }

// func AuthMiddleware(c *fiber.Ctx) error {
// 	tokenString := c.Get("Authorization")

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(SecretKey), nil
// 	})

// 	if err != nil || !token.Valid {
// 		return c.Status(fiber.StatusUnauthorized).SendString("Invalid Token")
// 	}

// 	c.Locals("user", getUserFromToken(token))
// 	return c.Next()
// }

// func getUserFromToken(token *jwt.Token) *UserModel {
// 	claims := token.Claims.(jwt.MapClaims)

// 	idStr := claims["id"].(string)
// 	id, _ := uuid.Parse(idStr)
// 	nome := claims["nome"].(string)
// 	email := claims["email"].(string)
// 	//role :=claims["role"].(string)

// 	return NewUser(id, nome, email)
// }

// func getUserRoleFromContext(c *fiber.Ctx) string {
// 	user := c.Locals("user").(UserModel)
// 	return user.Role
// }

// func DeserializeUser(c *fiber.Ctx) error {
// 	var tokenString string
// 	authorization := c.Get("Authorization")

// 	if strings.HasPrefix(authorization, "Bearer ") {
// 		tokenString = strings.TrimPrefix(authorization, "Bearer ")
// 	} else if c.Cookies("token") != "" {
// 		tokenString = c.Cookies("token")
// 	}

// 	if tokenString == "" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
// 	}

// 	config, _ := LoadConfig(".")

// 	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
// 		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
// 		}

// 		return []byte(config.JwtSecret), nil
// 	})
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
// 	}

// 	claims, ok := tokenByte.Claims.(jwt.MapClaims)
// 	if !ok || !tokenByte.Valid {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

// 	}

// 	var user UserModel
// 	// DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

// 	if user.ID.String() != claims["sub"] {
// 		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
// 	}

// 	// c.Locals("user", FilterUserRecord(&user))

// 	return c.Next()
// }

// func TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	jwtSecretKey := viper.GetString("ONLINE_TICKET_GO_JWTKEY")

// 	return func(c echo.Context) error {
// 		if _, ok := allowList[c.Request().RequestURI]; ok {
// 			return next(c)
// 		}

// 		cookie, err := c.Cookie("token")
// 		if err != nil {
// 			return c.String(http.StatusBadRequest, err.Error())
// 		}

// 		token := cookie.Value

// 		claim := Claims{}
// 		parsedTokenInfo, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(jwtSecretKey), nil
// 		})
// 		if err != nil {
// 			if errors.Is(err, jwt.ErrSignatureInvalid) {
// 				return c.String(http.StatusUnauthorized, "Please login again")
// 			}

// 			return c.String(http.StatusUnauthorized, "Please login again")
// 		}

// 		if !parsedTokenInfo.Valid {
// 			return c.String(http.StatusForbidden, "Invalid token")
// 		}

// 		c.Set("claim", claim)

// 		return next(c)
// 	}
// }

// func IsAuthorized() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		cookie, err := c.Cookie("token")
// 		if err != nil {
// 			c.JSON(401, gin.H{"error": "unauthorized"})
// 			c.Abort()
// 			return
// 		}

// 		claims, err := utils.ParseToken(cookie)
// 		if err != nil {
// 			c.JSON(401, gin.H{"error": "unauthorized"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("role", claims.Role)
// 		c.Next()
// 	}
// }

// func JWTAuth() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		err := ValidateJWT(context)
// 		if err != nil {
// 			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
// 			context.Abort()
// 			return
// 		}
// 		error := ValidateAdminRoleJWT(context)
// 		if error != nil {
// 			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only Administrator is allowed to perform this action"})
// 			context.Abort()
// 			return
// 		}
// 		context.Next()
// 	}
// }

// func JWTAuthCustomer() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		err := ValidateJWT(context)
// 		if err != nil {
// 			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
// 			context.Abort()
// 			return
// 		}
// 		error := ValidateCustomerRoleJWT(context)
// 		if error != nil {
// 			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only registered Customers are allowed to perform this action"})
// 			context.Abort()
// 			return
// 		}
// 		context.Next()
// 	}
// }

// func Login(context *gin.Context) {
// 	jwt, err := util.GenerateJWT(user)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"token": jwt, "username": input.Username, "message": "Successfully logged in"})
// }

// func proxy(path, target string) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		targetURL := target + r.URL.Path
// 		req, err := http.NewRequest(r.Method, targetURL, r.Body)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadGateway)
// 			return
// 		}

// 		req.Header = r.Header

// 		client := &http.Client{}
// 		resp, err := client.Do(req)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadGateway)
// 			return
// 		}
// 		defer resp.Body.Close()

// 		for key, values := range resp.Header {
// 			for _, value := range values {
// 				w.Header().Add(key, value)
// 			}
// 		}

// 		w.WriteHeader(resp.StatusCode)

// 		// Copy the response body to the client
// 		_, err = io.Copy(w, resp.Body)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadGateway)
// 			return
// 		}
// 	}
// }

// func AuthenticationMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
// 			c.Abort()
// 			return
// 		}

// 		// The token should be prefixed with "Bearer "
// 		tokenParts := strings.Split(tokenString, " ")
// 		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
// 			c.Abort()
// 			return
// 		}

// 		tokenString = tokenParts[1]

// 		claims, err := utils.VerifyToken(tokenString)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("user_id", claims["user_id"])
// 		c.Next()
// 	}
// }

// func TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	jwtSecretKey := viper.GetString("ONLINE_TICKET_GO_JWTKEY")

// 	return func(c echo.Context) error {
// 		if _, ok := allowList[c.Request().RequestURI]; ok {
// 			return next(c)
// 		}

// 		cookie, err := c.Cookie("token")
// 		if err != nil {
// 			return c.String(http.StatusBadRequest, err.Error())
// 		}

// 		token := cookie.Value

// 		claim := Claims{}
// 		parsedTokenInfo, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(jwtSecretKey), nil
// 		})
// 		if err != nil {
// 			if errors.Is(err, jwt.ErrSignatureInvalid) {
// 				return c.String(http.StatusUnauthorized, "Please login again")
// 			}

// 			return c.String(http.StatusUnauthorized, "Please login again")
// 		}

// 		if !parsedTokenInfo.Valid {
// 			return c.String(http.StatusForbidden, "Invalid token")
// 		}

// 		c.Set("claim", claim)

// 		return next(c)
// 	}
// }

// func Auth() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		tokenString := context.GetHeader("Authorization")
// 		if tokenString == "" {
// 			context.JSON(401, gin.H{"error": "request does not contain an access token"})
// 			context.Abort()
// 			return
// 		}
// 		err := auth.ValidateToken(tokenString)
// 		if err != nil {
// 			context.JSON(401, gin.H{"error": err.Error()})
// 			context.Abort()
// 			return
// 		}
// 		context.Next()
// 	}
// }

// func IsAuthorizedJWT(h httprouter.Handle, role string) httprouter.Handle {
// 	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 		rawAccessToken := r.Header.Get("Authorization")

// 		tr := &http.Transport{
// 			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 		}
// 		client := &http.Client{
// 			Timeout:   time.Duration(6000) * time.Second,
// 			Transport: tr,
// 		}
// 		ctx := oidc.ClientContext(context.Background(), client)
// 		provider, err := oidc.NewProvider(ctx, RealmConfigURL)
// 		if err != nil {
// 			authorisationFailed("authorisation failed while getting the provider: "+err.Error(), w, r)
// 			return
// 		}

// 		oidcConfig := &oidc.Config{
// 			ClientID: clientID,
// 		}
// 		verifier := provider.Verifier(oidcConfig)
// 		idToken, err := verifier.Verify(ctx, rawAccessToken)
// 		if err != nil {
// 			authorisationFailed("authorisation failed while verifying the token: "+err.Error(), w, r)
// 			return
// 		}

// 		var IDTokenClaims Claims // ID Token payload is just JSON.
// 		if err := idToken.Claims(&IDTokenClaims); err != nil {
// 			authorisationFailed("claims : "+err.Error(), w, r)
// 			return
// 		}
// 		fmt.Println(IDTokenClaims)
// 		//checking the roles
// 		user_access_roles := IDTokenClaims.ResourceAccess.DemoServiceClient.Roles
// 		for _, b := range user_access_roles {
// 			if b == role {
// 				h(w, r, ps)
// 				return
// 			}
// 		}

// 		authorisationFailed("user not allowed to access this api", w, r)
// 	}
// }

// func authorisationFailed(message string, w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(http.StatusUnauthorized)
// 	data := Res401Struct{
// 		Status:   "FAILED",
// 		HTTPCode: http.StatusUnauthorized,
// 		Message:  message,
// 	}
// 	res, _ := json.Marshal(data)
// 	w.Write(res)
// }

// func AuthenticationMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
// 			c.Abort()
// 			return
// 		}

// 		// The token should be prefixed with "Bearer "
// 		tokenParts := strings.Split(tokenString, " ")
// 		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
// 			c.Abort()
// 			return
// 		}

// 		tokenString = tokenParts[1]

// 		claims, err := utils.VerifyToken(tokenString)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("user_id", claims["user_id"])
// 		c.Next()
// 	}
// }

// func TokenAuthMiddleware(next http.Handler) http.Handler {
// 	validToken := DEFAULT_TOKEN
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("Authorization")
// 		if token == "" {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			fmt.Fprintf(w, "Unauthoried. Token missing")
// 			return
// 		}
// 		if token != validToken {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			fmt.Fprintf(w, "Unauthorized. Invalid Token")
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

// // API KEY Authentication
// func ApiKeyAuthMiddleware(next http.Handler) http.Handler {
// 	validApiKey := DEFAULT_API_KEY
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		apiKey := r.Header.Get("X-API-Key")
// 		if apiKey == "" {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			fmt.Fprintf(w, "Unauthoried. API Key missing")
// 			return
// 		}
// 		if apiKey != validApiKey {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			fmt.Fprintf(w, "Unauthorized. Invalid API Key")
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

// // JWT Authentication
// func JwtAuthMiddleWare(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		next.ServeHTTP(w, r)
// 	})
// }

// func AuthorizationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		next.ServeHTTP(w, r)
// 	})
// }

// type LoggingMiddleware struct {
// 	logger logger
// }

// func NewLoggingMiddleware(l logger) *LoggingMiddleware {
// 	return &LoggingMiddleware{
// 		logger: l,
// 	}
// }

// type logger interface {
// 	Infof(format string, args ...interface{})
// }

// // Logging middleware to log http requests
// func (lm *LoggingMiddleware) Logging(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		wi := &responseWriterInterceptor{
// 			statusCode:     http.StatusOK,
// 			ResponseWriter: w,
// 		}
// 		//lm.logger.Infof("%s %s", r.Method, r.RequestURI)
// 		next.ServeHTTP(wi, r)

// 		lm.logger.Infof("%s %s %d", r.Method, r.RequestURI, wi.statusCode)
// 	})
// }

// func ValidToken(t *jwt.Token, id string) bool {
// 	n, err := strconv.Atoi(id)
// 	if err != nil {
// 		return false
// 	}

// 	claims := t.Claims.(jwt.MapClaims)
// 	uid := int(claims["user_id"].(float64))

// 	return uid == n
// }

// func Authenticated() fiber.Handler {
// 	key := os.Getenv("JWT_SECRET")
// 	if key == "" {
// 		panic("No JWT_SECRET environment variable found")
// 	}
// 	return jwtware.New(jwtware.Config{
// 		ContextKey:   "jwt",
// 		SigningKey:   jwtware.SigningKey{Key: []byte(key)},
// 		ErrorHandler: jwtError,
// 	})
// }

// func AuthUserContext(db *gorm.DB) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		user := new(models.User)
// 		token := c.Locals("jwt").(*jwt.Token)
// 		claims := token.Claims.(jwt.MapClaims)
// 		email := claims["email"].(string)

// 		err := db.Where("email = ?", email).First(&user).Error
// 		if err != nil {
// 			return c.Status(404).SendString("user not found")
// 		}
// 		c.Locals("user", user)
// 		return c.Next()
// 	}
// }

// r.GET("/resource", func(c *gin.Context) {
// 	bearerToken := c.Request.Header.Get("Authorization")
// 	reqToken := strings.Split(bearerToken, " ")[1]
// 	for _, token := range tokens {
// 		if token == reqToken {
// 			c.JSON(http.StatusOK, gin.H{
// 				"data": "resource data",
// 			})
// 			return
// 		}
// 	}
// 	claims := &Claims{}
// 	tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized",})
// 			return
// 		}
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request",})
// 		return
// 	}
// 	if !tkn.Valid {
// 		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
// 		return
// 	}
// })
