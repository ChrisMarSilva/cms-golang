package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/chrismarsilva/cms.project.encrypting/filecrypt"
	_ "golang.org/x/crypto/ssh/terminal"
)

// go mod init github.com/chrismarsilva/cms.project.encrypting
// go get -u golang.org/x/crypto/pbkdf2
// go get -u golang.org/x/crypto/ssh/terminal
// go mod tidy

// go run main.go
// go run . encrypt img.png
// go run . decrypt img.png

// go get -u "github.com/cosmtrek/air@latest"
// air init
// air

func main() {
	encryptData()
	encryptFile()
}

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetterFn(key int, letter string) (result string) {
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	leftOversLetter := string(runes[0 : len(letter)-key])
	return fmt.Sprintf(`%s%s`, lastLetterKey, leftOversLetter)
}

func encrypt(key int, plainText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString = ""

	findOne := func(r rune) rune {
		pos := strings.Index(originalLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			hashedString = hashedString + string(hashLetter[letterPosition])
			return r
		}

		return r
	}

	strings.Map(findOne, plainText)
	return hashedString
}

func decrypt(key int, encrypttedText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString = ""

	findOne := func(r rune) rune {
		pos := strings.Index(hashLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			hashedString = hashedString + string(originalLetter[letterPosition])
			return r
		}

		return r
	}

	strings.Map(findOne, encrypttedText)
	return hashedString
}

func encryptData() {
	plainText := "HELLO-WORLD"
	fmt.Println("Plain Text", plainText)

	encrypted := encrypt(5, plainText)
	fmt.Println("Encrypted", encrypted)

	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted", decrypted)
}

func encryptFile() {
	// If not enough args, return help text
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run CryptoGo encrypt to encrypt a file, and CryptoGo decrypt to decrypt a file.")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("CryptoGo")
	fmt.Println("Simple file encrypter for your day-to-day needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tCryptoGo encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file given a password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		println("Missing the path to the file. For more information run CryptoGo help")
		os.Exit(0)
	}

	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}

	password := getPassword()

	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)

	fmt.Println("\nFile successfully protected")
}

func getPassword() []byte {
	// fmt.Print("Enter password: ")
	// password, _ := terminal.ReadPassword(0)
	// fmt.Print("\nConfirm password: ")
	// password2, _ := terminal.ReadPassword(0)

	password := []byte("123")
	password2 := []byte("123")

	if !validatePassword(password, password2) {
		fmt.Print("\nPasswords do not match. Please try again.\n")
		return getPassword()
	}

	return password
}

func decryptHandle() {
	if len(os.Args) < 3 {
		println("Missing the path to the file. For more information run CryptoGo help")
		os.Exit(0)
	}

	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}

	//fmt.Print("Enter password: ")
	//password, _ := terminal.ReadPassword(0)
	password := []byte("123")

	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)

	fmt.Println("\nFile successfully decrypted.")
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}

	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}
