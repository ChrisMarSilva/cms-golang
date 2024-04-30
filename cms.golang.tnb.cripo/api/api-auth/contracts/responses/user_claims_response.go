package responses

// import (
// 	"github.com/golang-jwt/jwt/v5"
// 	_ "github.com/chrismarsilva/cms.golang.tnb.cripo.api.auth/internals/models"
// )

// type UserClaimsResponse struct {
// 	jwt.StandardClaims
// 	jwt.RegisteredClaims
// 	UserID   uint
// 	Username string   `json:"username"`
// 	UserType models.UsersRoleType `json:"user_type"`
// 	Role     string   `json:"role"`
// }

// func (u *UserClaimsResponse) IsAdmin() bool {
// 	return u.Role == models.AdminRole
// 	return u.UserType == models.AdminRole
// }

