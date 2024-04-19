package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

/*
package main

import (
	"context" // "golang.org/x/net/context"
	"log"
	"net/http"
	"strings"

	"github.com/ChrisMarSilva/cms-golang-api-gin.login/models"
	// "github.com/ChrisMarSilva/cms-golang-api-gin.login/routes"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

// go mod init github.com/chrismarsilva/cms-golang-api-gin.login
// go mod tidy
// go get -u github.com/gin-gonic/gin
// go get -u githu.com/jackc/pgx/v4
// go get -u github.com/gofrs/uuid
// go get -u github.com/dgrijalva/jwt-go
// go get -u golang.org/x/crypto@v0.0.0-20210711020723-a769d52b0f97
// go run main.go

// exemplo: https://github.com/mikaelm1/offersapp

func main() {

	// conn, err := connectDB()
	// if err != nil {
	// 	return
	// }

	router := gin.Default()
	//router.Use(dbMiddleware(*conn))

	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"hello": "world"}) })
	router.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"ping": "ok"}) })
	router.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"health": "ok"}) })
	router.GET("/data", func(c *gin.Context) { c.JSON(200, gin.H{"data": "ok"}) })
	router.GET("/getall", func(c *gin.Context) { c.JSON(200, gin.H{"getall": "ok"}) })

	//router.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	// routerGroup := router.Group("api/v1")
	// {
	// 	routerUserGroup := routerGroup.Group("users")
	// 	{
	// 		routerUserGroup.POST("register", routes.UsersRegister)
	// 		routerUserGroup.POST("login", routes.UsersLogin)
	// 	}
	// 	routerItemsGroup := routerGroup.Group("items")
	// 	{
	// 		routerItemsGroup.GET("index", routes.ItemsIndex)
	// 		routerItemsGroup.POST("create", authMiddleWare(), routes.ItemsCreate)
	// 		routerItemsGroup.PUT("update", authMiddleWare(), routes.ItemsUpdate)
	// 		routerItemsGroup.GET("sold_by_user", authMiddleWare(), routes.ItemsForSaleByCurrentUser)
	// 	}
	// }

	log.Println("Listem port 3000")
	router.Run(":3000")
}

func connectDB() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:@localhost:5432/offersapp")
	if err != nil || conn == nil {
		log.Println("Error connecting to DB")
		log.Println(err.Error())
	}
	_ = conn.Ping(context.Background())
	return conn, err
}

func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := models.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}



func init() {
}

package routes

import (
	"fmt"
	"net/http"

	"github.com/ChrisMarSilva/cms-golang-api-gin.login/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func UsersRegister(c *gin.Context) {

	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	err = user.Register(&conn)
	if err != nil {
		fmt.Println("Error in user.Register()")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": user.ID})
}

func UsersLogin(c *gin.Context) {

	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	err = user.IsAuthenticated(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "There was an error authenticating."})
}



package routes

import (
	"fmt"
	"net/http"

	"github.com/ChrisMarSilva/cms-golang-api-gin.login/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func ItemsCreate(c *gin.Context) {

	userID := c.GetString("user_id")

	item := models.Item{}
	c.ShouldBindJSON(&item)

	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	err := item.Create(&conn, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func ItemsIndex(c *gin.Context) {

	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	items, err := models.GetAllItems(&conn)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func ItemsForSaleByCurrentUser(c *gin.Context) {

	userID := c.GetString("user_id")

	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	items, err := models.GetItemsBeingSoldByUser(userID, &conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func ItemsUpdate(c *gin.Context) {

	userID := c.GetString("user_id")

	itemSent := models.Item{}
	err := c.ShouldBindJSON(&itemSent)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form sent"})
		return
	}

	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	itemBeingUpdated, err := models.FindItemById(itemSent.ID, &conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if itemBeingUpdated.SellerID.String() != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this item"})
		return
	}

	itemSent.SellerID = itemBeingUpdated.SellerID

	err = itemSent.Update(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": itemSent})
}



package models

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
)

type Item struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	Title        string    `json:"title"`
	Notes        string    `json:"notes"`
	SellerID     uuid.UUID `json:"seller"`
	PriceInCents int64     `json:"price_in_cents"`
}

func (i *Item) Create(conn *pgx.Conn, userID string) error {

	i.Title = strings.Trim(i.Title, " ")
	if len(i.Title) < 1 {
		return fmt.Errorf("Title must not be empty.")
	}

	if i.PriceInCents < 0 {
		i.PriceInCents = 0
	}

	now := time.Now()

	row := conn.QueryRow(
		context.Background(),
		"INSERT INTO item (title, notes, seller_id, price_in_cents, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, seller_id",
		i.Title, i.Notes, userID, i.PriceInCents, now, now,
	)

	err := row.Scan(&i.ID, &i.SellerID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("There was an error creating the item")
	}

	return nil
}

func GetAllItems(conn *pgx.Conn) ([]Item, error) {

	rows, err := conn.Query(context.Background(), "SELECT id, title, notes, seller_id, price_in_cents FROM item")
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Error getting items")
	}

	var items []Item

	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.ID, &item.Title, &item.Notes, &item.SellerID, &item.PriceInCents)
		if err != nil {
			fmt.Println(err)
			continue
		}
		items = append(items, item)
	}

	return items, nil
}

func GetItemsBeingSoldByUser(userID string, conn *pgx.Conn) ([]Item, error) {

	rows, err := conn.Query(context.Background(), "SELECT id, title, price_in_cents, notes, seller_id FROM item WHERE seller_id = $1", userID)
	if err != nil {
		fmt.Printf("Error getting items %v", err)
		return nil, fmt.Errorf("There was an error getting the items")
	}

	var items []Item

	for rows.Next() {
		i := Item{}
		err = rows.Scan(&i.ID, &i.Title, &i.PriceInCents, &i.Notes, &i.SellerID)
		if err != nil {
			fmt.Printf("Error scaning item: %v", err)
			continue
		}
		items = append(items, i)
	}

	return items, nil
}

func (i *Item) Update(conn *pgx.Conn) error {

	i.Title = strings.Trim(i.Title, " ")
	if len(i.Title) < 1 {
		return fmt.Errorf("Title must not be empty")
	}

	if i.PriceInCents < 0 {
		i.PriceInCents = 0
	}

	now := time.Now()

	_, err := conn.Exec(context.Background(), "UPDATE item SET title=$1, notes=$2, price_in_cents=$3, updated_at=$4 WHERE id=$5", i.Title, i.Notes, i.PriceInCents, now, i.ID)
	if err != nil {
		fmt.Printf("Error updating item: (%v)", err)
		return fmt.Errorf("Error updating item")
	}

	return nil
}

func FindItemById(id uuid.UUID, conn *pgx.Conn) (Item, error) {
	row := conn.QueryRow(context.Background(), "SELECT title, notes, seller_id, price_in_cents FROM item WHERE id=$1", id)

	item := Item{ID: id}
	err := row.Scan(&item.Title, &item.Notes, &item.SellerID, &item.PriceInCents)
	if err != nil {
		return item, fmt.Errorf("The item doesn't exist")
	}

	return item, nil
}



package models

import (
	"context"
	"fmt"

	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	tokenSecret = []byte(os.Getenv("TOKEN_SECRET"))
)

type User struct {
	ID              uuid.UUID `json:"id"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
	Email           string    `json:"email"`
	PasswordHash    string    `json:"-"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
}

func (u *User) Register(conn *pgx.Conn) error {

	if len(u.Email) < 4 {
		return fmt.Errorf("Email must be at least 4 characters long.")
	}

	if len(u.Password) < 4 || len(u.PasswordConfirm) < 4 {
		return fmt.Errorf("Password must be at least 4 characters long.")
	}

	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Passwords do not match.")
	}

	u.Email = strings.ToLower(u.Email)
	userLookup := User{}
	row := conn.QueryRow(context.Background(), "SELECT id from user_account WHERE email = $1", u.Email)
	err := row.Scan(&userLookup)
	if err != pgx.ErrNoRows {
		fmt.Println("found user")
		fmt.Println(userLookup.Email)
		return fmt.Errorf("A user with that email already exists")
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("There was an error creating your account.")
	}
	u.PasswordHash = string(pwdHash)

	now := time.Now()
	_, err = conn.Exec(context.Background(), "INSERT INTO user_account (created_at, updated_at, email, password_hash) VALUES($1, $2, $3, $4)", now, now, u.Email, u.PasswordHash)

	return err
}

// IsAuthenticated checks to make sure password is correct and user is active
func (u *User) IsAuthenticated(conn *pgx.Conn) error {

	row := conn.QueryRow(context.Background(), "SELECT id, password_hash from user_account WHERE email = $1", u.Email)
	err := row.Scan(&u.ID, &u.PasswordHash)
	if err == pgx.ErrNoRows {
		fmt.Println("User with email not found")
		return fmt.Errorf("Invalid login credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return fmt.Errorf("Invalid login credentials")
	}

	return nil
}

// GetAuthToken returns the auth token to be used
func (u *User) GetAuthToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	authToken, err := token.SignedString(tokenSecret)
	return authToken, err
}

func IsTokenValid(tokenString string) (bool, string) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// fmt.Printf("Parsing: %v \n", token)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok == false {
			return nil, fmt.Errorf("Token signing method is not valid: %v", token.Header["alg"])
		}
		return tokenSecret, nil
	})

	if err != nil {
		fmt.Printf("Err %v \n", err)
		return false, ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims)
		userID := claims["user_id"]
		return true, userID.(string)
	} else {
		fmt.Printf("The alg header %v \n", claims["alg"])
		fmt.Println(err)
		return false, "uuid.UUID{}"
	}

}




CREATE TABLE t (id binary(16) PRIMARY KEY);
INSERT INTO t VALUES(UUID_TO_BIN(UUID()));
SELECT BIN_TO_UUID(id) FROM t;






CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE base_table (
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE user_account (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
) INHERITS (base_table);

CREATE TABLE item (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    title VARCHAR(255) NOT NULL,
    notes TEXT,
    seller_id uuid,
    price_in_cents INTEGER,
    FOREIGN KEY (seller_id) REFERENCES user_account (id) ON DELETE CASCADE
) INHERITS (base_table);

CREATE TABLE purchase (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    buyer_id uuid,
    item_id uuid,
    price_in_cents INTEGER,
    title VARCHAR(255) NOT NULL,
    FOREIGN KEY (buyer_id) REFERENCES user_account (id),
    FOREIGN KEY (item_id) REFERENCES item (id)
) INHERITS (base_table);
*/
