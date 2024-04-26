package models

type Claims struct {
	Role string `json:"role"`
}

type Role struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}



package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserType string

const (
	Admin          UserType = "admin"
	IndividualUser UserType = "individual"
	CorporateUser  UserType = "corporate"
)

type Claims struct {
	Username string   `json:"username"`
	UserType UserType `json:"user_type"`
	UserID   uint
	jwt.RegisteredClaims
}

func (c *Claims) IsIndividualUser() bool {
	return c.UserType == IndividualUser
}

func (c *Claims) IsCorporatedUser() bool {
	return c.UserType == CorporateUser
}

func (c *Claims) IsAdmin() bool {
	return c.UserType == Admin
}

func (c *Claims) IsNotAdmin() bool {
	return !c.IsAdmin()
}

func (c *Claims) IsUser() bool {
	return c.UserType == IndividualUser || c.UserType == CorporateUser
}

func (c *Claims) IsUserOrAdmin() bool {
	return c.IsAdmin() || c.IsUser()
}

func (c *Claims) IsUnknownTypeUser() bool {
	return !c.IsUserOrAdmin()
}

type JWTClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}


type Res401Struct struct {
	Status   string `json:"status" example:"FAILED"`
	HTTPCode int    `json:"httpCode" example:"401"`
	Message  string `json:"message" example:"authorisation failed"`
}

//claims component of jwt contains mainy fields , we need only roles of DemoServiceClient
//"DemoServiceClient":{"DemoServiceClient":{"roles":["pets-admin","pet-details","pets-search"]}},
type Claims struct {
	ResourceAccess client `json:"resource_access,omitempty"`
	JTI            string `json:"jti,omitempty"`
}

type client struct {
	DemoServiceClient clientRoles `json:"DemoServiceClient,omitempty"`
}

type clientRoles struct {
	Roles []string `json:"roles,omitempty"`
}

var RealmConfigURL string = "http://10.66.29.167:9999/auth/realms/DEMOREALM"
var clientID string = "DemoServiceClient"
