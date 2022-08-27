package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"runtime"
	"strconv"
	//badger "github.com/dgraph-io/badger/v3"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.blockchain
// go get github.com/dgraph-io/badger/v3
// go mod tidy

// go run main.go

func main() {
	log.Println("Main.Ini")

	//Teste01()
	//Teste02()
	//Teste03()

	//https://www.youtube.com/watch?v=szOZ3p-5YIc&list=PLR8aeiMU8Si6BSmNiQSf8BRO7TAy3V_x5&index=3
	defer os.Exit(0)
	chain := InitBlockChain()
	//defer chain.Database.Close()

	cli := CommandLine{chain}
	cli.run()

	log.Println("Main.Fim")
}

type CommandLine struct {
	blockchain *BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added Block!")
}

func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		Handle(err)

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

// Take the data from the block

// create a counter (nonce) which starts at 0

// create a hash of the data plus the counter

// check the hash to see if it meets a set of requirements

// Requirements:
// The First few bytes must contain 0s

const Difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	fmt.Println()

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}

const (
	dbPath = "./tmp/blocks"
)

type BlockChain struct {
	LastHash []byte
	//Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	//Database    *badger.DB
}

func InitBlockChain() *BlockChain {
	//var lastHash []byte

	// opts := badger.DefaultOptions
	// opts.Dir = dbPath
	// opts.ValueDir = dbPath

	// db, err := badger.Open(opts)
	// Handle(err)

	// err = db.Update(func(txn *badger.Txn) error {
	// 	if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
	// 		fmt.Println("No existing blockchain found")
	// 		genesis := Genesis()
	// 		fmt.Println("Genesis proved")
	// 		err = txn.Set(genesis.Hash, genesis.Serialize())
	// 		Handle(err)
	// 		err = txn.Set([]byte("lh"), genesis.Hash)
	// 		lastHash = genesis.Hash
	// 		return err
	// 	} else {
	// 		item, err := txn.Get([]byte("lh"))
	// 		Handle(err)
	// 		lastHash, err = item.Value()
	// 		return err
	// 	}
	// })
	// Handle(err)

	blockchain := BlockChain{} // lastHash, db
	return &blockchain
}

func (chain *BlockChain) AddBlock(data string) {
	// var lastHash []byte
	// err := chain.Database.View(func(txn *badger.Txn) error {
	// 	item, err := txn.Get([]byte("lh"))
	// 	Handle(err)
	// 	lastHash, err = item.Value()
	// 	return err
	// })
	// Handle(err)

	//newBlock := CreateBlock(data, lastHash)
	// err = chain.Database.Update(func(txn *badger.Txn) error {
	// 	err := txn.Set(newBlock.Hash, newBlock.Serialize())
	// 	Handle(err)
	// 	err = txn.Set([]byte("lh"), newBlock.Hash)
	// 	chain.LastHash = newBlock.Hash
	// 	return err
	// })
	// Handle(err)
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{} // chain.LastHash, chain.Database

	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block
	// err := iter.Database.View(func(txn *badger.Txn) error {
	// 	item, err := txn.Get(iter.CurrentHash)
	// 	Handle(err)
	// 	encodedBlock, err := item.Value()
	// 	block = Deserialize(encodedBlock)
	// 	return err
	// })
	//Handle(err)
	iter.CurrentHash = block.PrevHash
	return block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	Handle(err)
	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	Handle(err)
	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
