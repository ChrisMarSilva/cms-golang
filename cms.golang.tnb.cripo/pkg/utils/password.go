package utils

// e utils

//  import "golang.org/x/crypto/bcrypt"

//  func (user *User) ValidateUserPassword(password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// }

//  func CompareHashPassword(password, hash string) bool {
//      err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//      return err == nil
//  }

//  func GenerateHashPassword(password string) (string, error) {
//     bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//     return string(bytes), err
// }

// func CompareHashPassword(password, hash string) bool {
//     err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//     return err == nil
// }

// func (user *User) HashPassword(password string) error {
//     bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//     if err != nil {
//       return err
//     }
//     user.Password = string(bytes)
//     return nil
//   }
//   func (user *User) CheckPassword(providedPassword string) error {
//     err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
//     if err != nil {
//       return err
//     }
//     return nil
//   }

//   package security

// import "golang.org/x/crypto/bcrypt"

// func EncryptPassword(password string) (string, error) {
// 	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hashed), nil
// }

// func VerifyPassword(hashed, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
// }

/*


	services/auth/password.go
	package auth

	import (
		"golang.org/x/crypto/bcrypt"
	)

	func HashPassword(password string) (string, error) {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return "", err
		}

		return string(hash), nil
	}

	func ComparePasswords(hashed string, plain []byte) bool {
		err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
		return err == nil
	}


*/
