package main

import (
	"fmt"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)}, // []byte(secret),
	})
}

func AdminMiddleware(c *fiber.Ctx) error {
	// userRole := getUserRoleFromContext(c)

	// if userRole != AdminRole {
	// 	return c.Status(fiber.StatusForbidden).SendString("Permission Denied")
	// }

	return c.Next()
}

func AuthMiddleware(c *fiber.Ctx) error {
	// Check for authentication credentials
	// If credentials are valid, proceed to the next middleware or route
	// If credentials are invalid, return an unauthorized response
	tokenString := c.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the token signing method and return the secret key
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid Token")
	}

	// Extract user information from the token and store it in the context
	c.Locals("user", getUserFromToken(token))

	return c.Next()
}

func getUserFromToken(token *jwt.Token) *UserModel {
	claims := token.Claims.(jwt.MapClaims)

	idStr := claims["id"].(string)
	id, _ := uuid.Parse(idStr)
	nome := claims["nome"].(string)
	email := claims["email"].(string)
	//role :=claims["role"].(string)

	return NewUser(id, nome, email)
}

func getUserRoleFromContext(c *fiber.Ctx) string {
	// Extract the user information from the context
	// user := c.Locals("user").(User)

	return "admin" // user.Role
}

func DeserializeUser(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	config, _ := LoadConfig(".")

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

	}

	var user UserModel
	// DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	if user.ID.String() != claims["sub"] {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	}

	// c.Locals("user", FilterUserRecord(&user))

	return c.Next()
}
