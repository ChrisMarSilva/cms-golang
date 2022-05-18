package main

// import (
// 	"bytes"
// 	"crypto/sha256"
// 	"encoding/binary"
// 	"fmt"
// 	"log"
// 	"math"
// 	"math/big"
// 	"strconv"
// )

// func Teste03() {
// 	log.Println("Teste.03.Ini")

// 	// https://www.youtube.com/watch?v=aE4eDTUAE70&list=PLR8aeiMU8Si6BSmNiQSf8BRO7TAy3V_x5&index=2

// 	chain := InitBlockChain()

// 	chain.AddBlock("First Block after Genesis")
// 	chain.AddBlock("Second Block after Genesis")
// 	chain.AddBlock("Third Block after Genesis")

// 	for _, block := range chain.Blocks {

// 		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
// 		fmt.Printf("Data in Block: %s\n", block.Data)
// 		fmt.Printf("Hash: %x\n", block.Hash)

// 		pow := NewProof(block)
// 		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
// 		fmt.Println()

// 	}

// 	log.Println("Teste.03.Fim")
// }

// type BlockChain struct {
// 	Blocks []*Block
// }

// type Block struct {
// 	Hash     []byte
// 	Data     []byte
// 	PrevHash []byte
// 	Nonce    int
// }

// func CreateBlock(data string, prevHash []byte) *Block {
// 	block := &Block{[]byte{}, []byte(data), prevHash, 0}
// 	pow := NewProof(block)
// 	nonce, hash := pow.Run()

// 	block.Hash = hash[:]
// 	block.Nonce = nonce

// 	return block
// }

// func (chain *BlockChain) AddBlock(data string) {
// 	prevBlock := chain.Blocks[len(chain.Blocks)-1]
// 	new := CreateBlock(data, prevBlock.Hash)
// 	chain.Blocks = append(chain.Blocks, new)
// }

// func Genesis() *Block {
// 	return CreateBlock("Genesis", []byte{})
// }

// func InitBlockChain() *BlockChain {
// 	return &BlockChain{[]*Block{Genesis()}}
// }

// // Take the data from the block

// // create a counter (nonce) which starts at 0

// // create a hash of the data plus the counter

// // check the hash to see if it meets a set of requirements

// // Requirements:
// // The First few bytes must contain 0s

// const Difficulty = 18

// type ProofOfWork struct {
// 	Block  *Block
// 	Target *big.Int
// }

// func NewProof(b *Block) *ProofOfWork {
// 	target := big.NewInt(1)
// 	target.Lsh(target, uint(256-Difficulty))

// 	pow := &ProofOfWork{b, target}

// 	return pow
// }

// func (pow *ProofOfWork) InitData(nonce int) []byte {
// 	data := bytes.Join(
// 		[][]byte{
// 			pow.Block.PrevHash,
// 			pow.Block.Data,
// 			ToHex(int64(nonce)),
// 			ToHex(int64(Difficulty)),
// 		},
// 		[]byte{},
// 	)

// 	return data
// }

// func (pow *ProofOfWork) Run() (int, []byte) {
// 	var intHash big.Int
// 	var hash [32]byte

// 	nonce := 0

// 	for nonce < math.MaxInt64 {
// 		data := pow.InitData(nonce)
// 		hash = sha256.Sum256(data)

// 		fmt.Printf("\r%x", hash)
// 		intHash.SetBytes(hash[:])

// 		if intHash.Cmp(pow.Target) == -1 {
// 			break
// 		} else {
// 			nonce++
// 		}

// 	}
// 	fmt.Println()

// 	return nonce, hash[:]
// }

// func (pow *ProofOfWork) Validate() bool {
// 	var intHash big.Int

// 	data := pow.InitData(pow.Block.Nonce)

// 	hash := sha256.Sum256(data)
// 	intHash.SetBytes(hash[:])

// 	return intHash.Cmp(pow.Target) == -1
// }

// func ToHex(num int64) []byte {
// 	buff := new(bytes.Buffer)
// 	err := binary.Write(buff, binary.BigEndian, num)
// 	if err != nil {
// 		log.Panic(err)

// 	}

// 	return buff.Bytes()
// }
