package seguranca

import (
    "golang.org/x/crypto/bcrypt"
)

func Hash(senha string) ([]byte, error) { // func Hash(senha string) (string, error) {
    // bytes, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
    // return string(bytes), err
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func VerificarSenha(senha, hash string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
}

// func main() {
// 	password := "secret"
//     hash, _ := HashPassword(password) // ignore error for the sake of simplicity
//     fmt.Println("Password:", password)
//     fmt.Println("Hash:    ", hash)
//     match := CheckPasswordHash(password, hash)
//     fmt.Println("Match:   ", match)
//     password := []byte("MyDarkSecret")
//     // Hashing the password with the default cost of 10
//     hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println(string(hashedPassword))
//     // Comparing the password with the hash
//     err = bcrypt.CompareHashAndPassword(hashedPassword, password)
//     fmt.Println(err) // nil means it is a match
// 	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
//     if err != nil {
//         http.Error(w, "Internal server error", http.StatusInternalServerError)
//         return
//     }
// }