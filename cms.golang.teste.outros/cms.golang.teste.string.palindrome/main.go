package main

import (
	"log"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.string.palindrome
// go mod tidy

// go run main.go

func main() {
	var text string

	text = "chris"
	log.Println(text, "=", invertText(text), "=", isPalindrome(text), "=", isPalindrome2(text))

	text = "ana"
	log.Println(text, "=", invertText(text), "=", isPalindrome(text), isPalindrome2(text))

	// BenchmarkIsPalindrome-4         13709085                83.18 ns/op
	// BenchmarkIsPalindrome2-4        316470301                3.617 ns/op
}

func isPalindrome(text string) (result bool) {
	return invertText(text) == text
}

// HANNAH = HAN | NAH
func isPalindrome2(text string) (result bool) {
	if len(text) == 0 {
		return false
	}

	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		if text[i] != text[j] {
			return false
		}
	}

	return true
}

func invertText(text string) string {

	// invertText
	// var invertedText string
	// for i := len(text) - 1; i >= 0; i-- {
	// 	invertedText += string(text[i])
	// }
	// return invertedText

	// invertText2
	// runes := []rune(text)
	// for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	// 	runes[i], runes[j] = runes[j], runes[i]
	// }
	// return string(runes)

	// invertText4
	// var result string
	// for _, v := range text {
	// 	result = string(v) + result
	// }
	// return result

	//invertText5
	var invertedText []byte // var invertedText []byte = make([]byte, len(text)) // invertedText := make([]byte, len(text))
	for i := len(text) - 1; i >= 0; i-- {
		invertedText = append(invertedText, text[i])
	}
	return string(invertedText)

	// BenchmarkInvertText-4            6105006               178.8 ns/op
	// BenchmarkInvertText2-4          23313883                53.90 ns/op
	// BenchmarkInvertText4-4           7803406               158.8 ns/op

	// BenchmarkInvertText-4           15357903                87.24 ns/op
	// BenchmarkInvertText5-4          22469353                49.68 ns/op

}
