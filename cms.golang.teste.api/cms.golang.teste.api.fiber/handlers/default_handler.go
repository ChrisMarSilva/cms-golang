package handlers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	//"github.com/dgrijalva/jwt-go"
	//"github.com/gofiber/fiber/v2/middleware/csrf"
)

const jwtSecret = "asecret#crss"

type DefaultHandler struct {
}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

func (u *DefaultHandler) Index(c *fiber.Ctx) error {
	// return c.JSON(fiber.Map{"name": "Grame", "age": 20})
	return c.Status(fiber.StatusOK).SendString("Hello, Index!")
}

func (u *DefaultHandler) Teste(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("ok fiber")
}

func (u *DefaultHandler) NotFound(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotFound)
}

func (u *DefaultHandler) Login(c *fiber.Ctx) error {
	// 	type request struct {
	// 		Email    string `json:"email"`
	// 		Password string `json:"password"`
	// 	}
	// 	var body request
	// 	err := c.BodyParser(&body)
	// 	if err != nil {
	// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	// 	}
	// 	if body.Email == "" || body.Password == "" {
	// 		return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
	// 	}
	// 	if body.Email != "bob@gmail.com" || body.Password != "password123" {
	// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Bad Credentials"})
	// 	}
	user := c.FormValue("user")
	pass := c.FormValue("pass")
	if user != "john" || pass != "doe" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "user invalid"})
	}
	// 	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
	// 	// 	return err
	// 	// }
	exp := time.Now().Add(time.Hour * 24).Unix()
	//sub := utils.UUID()
	claims := jwt.MapClaims{"name": "John Doe", "admin": true, "exp": exp} // a week
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)             // Create token
	// 	token := jwt.New(jwt.SigningMethodHS256)
	// 	claims := token.Claims.(jwt.MapClaims)
	// 	claims["sub"] = "1"
	// 	claims["exp"] = time.Now().Add(time.Hour * 24 * 7) // a week
	t, err := token.SignedString([]byte(jwtSecret)) // Generate encoded token and send it as response.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "token invalid"})
	}

	cookie := fiber.Cookie{Name: "jwt", Value: t, Expires: time.Now().Add(time.Hour * 24), HTTPOnly: true}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": t, "exp": exp})
}

func (u *DefaultHandler) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{Name: "jwt", Value: "", Expires: time.Now().Add(-time.Hour), HTTPOnly: true}
	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func (u *DefaultHandler) Restricted1(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthenticated"})
	}

	claims := token.Claims.(*jwt.StandardClaims)
	//claims := token.Claims.(*jwt.MapClaims)

	return c.Status(fiber.StatusOK).SendString("Welcome ok" + claims.Issuer)
}

func (u *DefaultHandler) Restricted2(c *fiber.Ctx) error {
	//user := c.Locals("user").(*jwt.Token)
	// if temp := c.Locals("user"); temp != nil {
	// 	user := temp.(*jwt.Token)
	// 	claims := user.Claims.(jwt.MapClaims)
	// 	name = claims["name"].(string)
	// }
	claims, err := ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": true, "msg": err.Error()}) // Return status 500 and JWT parse error.
	}
	// println("claims", claims)
	// println("claims.Expires", claims.Expires)
	// println("claims.Nome", claims.Nome)
	return c.Status(fiber.StatusOK).SendString("Welcome " + claims.Nome)
}

func (u *DefaultHandler) Public(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": true, "path": "public"})
}

func (u *DefaultHandler) Private(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": true, "path": "private"})
}

type TokenMetadata struct {
	Expires int64
	Nome    string
}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims) // Setting and checking token and credentials.
	if ok && token.Valid {
		// println("name", claims["name"].(string))
		// println("admin", claims["admin"].(bool))
		// println("exp1", int64(claims["exp"].(float64)))
		// println("exp2", claims["exp"].(float64))
		expires := int64(claims["exp"].(float64)) // Expires time.
		nome := claims["name"].(string)           // nome
		return &TokenMetadata{Expires: expires, Nome: nome}, nil
	}

	return nil, err
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")
	onlyToken := strings.Split(bearToken, " ") // Normally Authorization HTTP header.
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}
	return ""
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(jwtSecret), nil
}
