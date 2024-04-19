package main

import (
	"net/http"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("log middleware")
		next.ServeHTTP(w, r)
	})
}

func HomeHandler(c fiber.Ctx) error {
	log.Info(c.Method(), " ", c.Path())
	return c.SendString("I'm a GET / request!")
}

func LoginHandler(c fiber.Ctx) error {
	log.Info(c.Method(), " ", c.Path())

	user := new(UserModel)
	if err := json.Unmarshal(c.Body(), &user); err != nil { // if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"nome":      user.Nome,
		"email":     user.Email,
		"is_active": user.IsActive,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	sess, err := store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	sess.Set("jwt", t)
	if err := sess.Save(); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func LogoutHandler(c fiber.Ctx) error {
	log.Info(c.Method(), " ", c.Path())

	sess, err := store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	sess.Destroy()
	if err := sess.Save(); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func RefreshHandler(c fiber.Ctx) error {
	log.Info(c.Method(), " ", c.Path())

	sess, err := store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	tokenStr := sess.Get("jwt")
	if tokenStr == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("No token found")
	}

	token, err := jwt.Parse(tokenStr.(string), func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return c.JSON(claims)
	}

	//sess.SetExpiry(time.Second * 2)
	// if err := sess.Save(); err != nil {
	//     panic(err)
	// }

	return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
}

func VerifyHandler(c fiber.Ctx) error {
	log.Info(c.Method(), " ", c.Path())

	sess, err := store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	tokenStr := sess.Get("jwt")
	if tokenStr == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("No token found")
	}

	token, err := jwt.Parse(tokenStr.(string), func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token:")
	}

	return c.SendStatus(fiber.StatusOK)
}
