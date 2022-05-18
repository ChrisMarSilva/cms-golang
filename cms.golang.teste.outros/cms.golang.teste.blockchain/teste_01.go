package main

// import (
// 	"crypto/sha256"
// 	"fmt"
// 	"log"
// 	"time"
// )

// type Block struct {
// 	timestamp    time.Time
// 	transactions []string
// 	prevhash     []byte
// 	Hash         []byte
// }

// func Teste01() {
// 	log.Println("Teste.01.Ini")

// 	abc := []string{" A sent 50 coins to BC"}
// 	xyz := Blocs(abc, []byte{})
// 	fmt.Println("This is our First Block")
// 	Print(xyz)

// 	pgrs := []string{" PQ sent 230 coins to RS"}
// 	klmn := Blocs(pgrs, xyz.Hash)
// 	fmt.Println("This is our Second Block")
// 	Print(klmn)

// 	log.Println("Teste.01.Fim")
// }

// func Blocs(transactions []string, prevhash []byte) *Block {
// 	currentTime := time.Now()
// 	return &Block{
// 		timestamp:    currentTime,
// 		transactions: transactions,
// 		prevhash:     prevhash,
// 		Hash:         NewHash(currentTime, transactions, prevhash),
// 	}
// }

// func NewHash(time time.Time, transactions []string, prevhash []byte) []byte {
// 	input := append(prevhash, time.String()...)
// 	for transaction := range transactions {
// 		input = append(input, string(rune(transaction))...)
// 	}
// 	hash := sha256.Sum256(input)
// 	return hash[:]
// }

// func Print(block *Block) {
// 	fmt.Printf("\ttime: %s\n", block.timestamp.String())
// 	fmt.Printf("\tprevhash: %x\n", block.prevhash)
// 	fmt.Printf("\thash: %x\n", block.Hash)
// 	Transaction(block)
// }

// func Transaction(block *Block) {
// 	fmt.Println("\tTransactions")
// 	for i, transaction := range block.transactions {
// 		fmt.Printf("\t\t%v: %q\n", i, transaction)
// 	}
// }
