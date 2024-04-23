package main

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (h UserService) Register(c *fiber.Ctx, payload UserRegisterRequest) (UserRegisterResponse, error) {
	// Validate user input (username, email, password)
	// Hash the password
	// Store user data in the database
	// Return a success message or error response

	// errors := ValidateStruct(payload)
	// if errors != nil {
	// 	return UserRegisterResponse{}, err // return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	// }

	// if payload.Password != payload.PasswordConfirm {
	// 	return UserRegisterResponse{}, err // return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})
	// }

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserRegisterResponse{}, err // return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user := UserModel{
		ID:        uuid.New(),
		Nome:      payload.Nome,
		Email:     strings.ToLower(payload.Email),
		Password:  string(hashedPassword),
		IsActive:  true,
		CreatedAt: time.Now(),
	}

	response := UserRegisterResponse{
		Nome:  user.Nome,
		Email: user.Email,
	}

	// result := initializers.DB.Create(&newUser)

	// if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
	// 	return err // return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists"})
	// } else if result.Error != nil {
	// 	return err // return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	// }

	return response, nil
}

func (h UserService) Login(c *fiber.Ctx, payload UserLoginRequest) (string, error) {
	// errors := ValidateStruct(payload)
	// if errors != nil {
	//  log.Error("Erro no repository:", err.Error())
	// 	return "", err
	// }

	user, err := h.userRepo.GetByEmail(c.Context(), strings.ToLower(payload.Email))
	if err != nil {
		log.Error("Erro no repository:", err.Error())
		return "", err
	}

	// if user.Password != payload.Password {
	// 	log.Error("Erro no Password: Senha inválida.")
	// 	return "", errors.New("Senha inválida.")
	// }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		log.Error("Erro no repository:", err.Error())
		return "", err // return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":       user.ID,
		"nome":      user.Nome,
		"email":     user.Email,
		"is_active": user.IsActive,
		"exp":       time.Now().UTC().Add(time.Hour * 24 * 7).Unix(),
		"iat":       time.Now().UTC().Unix(),
		"nbf":       time.Now().UTC().Unix(),
	})

	// now := time.Now().UTC()
	// tokenByte := jwt.New(jwt.SigningMethodHS256)
	// claims := tokenByte.Claims.(jwt.MapClaims)
	// claims["sub"] = user.ID
	// claims["exp"] = now.Add(time.Hour * 24 * 7).Unix()
	// claims["iat"] = now.Unix()
	// claims["nbf"] = now.Unix()

	tokenStr, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Error("Erro no SignedString: ", err.Error())
		return "", err // c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	sess, err := store.Get(c)
	if err != nil {
		log.Error("Erro no store: ", err.Error())
		return "", err // c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	sess.Set("jwt", tokenStr)
	if err := sess.Save(); err != nil {
		log.Error("Erro no token: ", err.Error())
		return "", err // c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	sess.SetExpiry(time.Hour * 24)

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenStr,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7, // config.JwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return tokenStr, err
}

func (h UserService) Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	sess.Destroy()

	if err := sess.Save(); err != nil {
		return err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour * 24),
	})

	return nil // c.SendStatus(fiber.StatusOK)
}

func (h UserService) Refresh(c *fiber.Ctx) (jwt.MapClaims, error) {
	sess, err := store.Get(c)
	if err != nil {
		return nil, err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	tokenStr := sess.Get("jwt")
	if tokenStr == nil {
		return nil, err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "No token found"})
	}

	token, err := jwt.Parse(tokenStr.(string), func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token.") // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	//sess.SetExpiry(time.Second * 2)
	// if err := sess.Save(); err != nil {
	//     return nil, err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	// }

	return claims, nil
}

func (h UserService) Verify(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	tokenStr := sess.Get("jwt")
	if tokenStr == nil {
		return err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "No token found"})
	}

	token, err := jwt.Parse(tokenStr.(string), func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	return nil
}
