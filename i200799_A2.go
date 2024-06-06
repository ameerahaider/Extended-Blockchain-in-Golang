package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Transaction struct {
	TransactionID              string
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

type Block struct {
	transactions []*Transaction
	nonce        int
	prevHash     string
	blockHash    string
}

type Blockchain struct {
	list            []*Block
	transactionPool []*Transaction
}

type Node struct {
	blockchain *Blockchain
	difficulty int
}

func (n *Node) mineBlock() string {

	fmt.Println("\n\n")
	fmt.Println("**********************************************************************************")
	fmt.Println("*                           Node is Mining a New Block                           *")
	fmt.Println("*                                                                                *")
	fmt.Println("\n")

	// Get the previous block's hash
	previousHash := n.blockchain.list[len(n.blockchain.list)-1].blockHash

	// Create a new block and mine it
	newBlock := n.blockchain.NewBlock(previousHash, n.difficulty)
	newBlock.VerifyNonce(n.difficulty) //Verify the nonce

	// Print mined block details
	fmt.Println("Mined Block:")
	PrintBlock(newBlock)

	fmt.Println("\n")
	fmt.Println("*                                                                               *")
	fmt.Println("*                                                                               *")
	fmt.Println("*                      Node Sucessfully Mined a New Block                       *")
	fmt.Println("*********************************************************************************")
	fmt.Println("\n\n")

	return newBlock.prevHash
}

// Create a New Transaction
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	T := new(Transaction)
	T.senderBlockchainAddress = sender
	T.recipientBlockchainAddress = recipient
	T.value = value
	var TransactionString string
	TransactionString = MergeTransaction(T)
	T.TransactionID = CalculateHash(TransactionString)
	return T
}

// Add New Transaction to the Blockchain Transaction Pool
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {

	t := NewTransaction(sender, recipient, value)
	// add transaction to transaction pool using the append method
	bc.transactionPool = append(bc.transactionPool, t)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		list: []*Block{},
	}
}

func GetPrevValue(b *Block) string {
	return b.prevHash
}

func CalculateHash(stringToHash string) string {
	//fmt.Printf("String Received: %s\n", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

func MergeBlock(b *Block) string {
	var stringToHash string
	stringToHash = strconv.Itoa(b.nonce)

	for i := 0; i < len(b.transactions); i++ {
		t := b.transactions[i]
		stringToHash += t.TransactionID + t.senderBlockchainAddress + t.recipientBlockchainAddress + strconv.FormatFloat(float64(t.value), 'f', -1, 32)
	}

	stringToHash += b.prevHash
	return stringToHash
}

func MergeTransaction(t *Transaction) string {
	var stringToHash string
	stringToHash = t.recipientBlockchainAddress + t.senderBlockchainAddress + strconv.FormatFloat(float64(t.value), 'f', -1, 32)
	return stringToHash
}

func (block *Block) VerifyNonce(difficulty int) bool {
	attempt := MergeBlock(block)
	hashAttempt := CalculateHash(attempt)
	prefix := strings.Repeat("0", difficulty)

	flag := strings.HasPrefix(hashAttempt, prefix)

	if flag {
		fmt.Println("Nonce is valid for the block!")
	} else {
		fmt.Println("Nonce is not valid for the block!")
	}

	return flag
}

func (bc *Blockchain) NewBlock(previousHash string, difficulty int) *Block {

	var selectedTransactions []*Transaction
	// Show available transactions in the pool
	fmt.Println("Available Transactions in Pool:")
	for i, transaction := range bc.transactionPool {
		fmt.Printf("[%d] Sender: %s, Recipient: %s, Value: %f\n", i+1, transaction.senderBlockchainAddress, transaction.recipientBlockchainAddress, transaction.value)
	}

	// Ask the user to select transactions to include in the new block
	fmt.Println("Enter transaction numbers to include in the block (one at a time) or 'q' to finish:")
	for {
		var input string
		fmt.Print("Transaction number or 'q': ")
		fmt.Scanln(&input)

		if input == "q" {
			break
		}

		index, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || index <= 0 || index > len(bc.transactionPool) {
			fmt.Println("Invalid input. Please enter a valid transaction number.")
			continue
		}

		selectedTransaction := bc.transactionPool[index-1]
		selectedTransactions = append(selectedTransactions, selectedTransaction)
	}

	// Create the new block with selected transactions
	block := new(Block)
	block.nonce = 0 // Initial Value
	block.transactions = selectedTransactions
	block.prevHash = previousHash

	// Proof of Work (Mining)
	for {
		attempt := MergeBlock(block)
		hashAttempt := CalculateHash(attempt)
		prefix := strings.Repeat("0", difficulty) // Create a string of zeros with length difficulty
		if strings.HasPrefix(hashAttempt, prefix) {
			break
		}
		block.nonce++
	}

	// Calculate block hash and add the block to the blockchain
	stringToHash := MergeBlock(block)
	block.blockHash = CalculateHash(stringToHash)
	bc.list = append(bc.list, block)

	return block
}

func PrintBlock(b *Block) {
	blockData := make(map[string]interface{})
	blockData["Nonce"] = b.nonce
	blockData["PreviousHash"] = b.prevHash

	var transactionsData []map[string]interface{}
	for _, transaction := range b.transactions {
		transactionData := map[string]interface{}{
			"Sender":    transaction.senderBlockchainAddress,
			"Recipient": transaction.recipientBlockchainAddress,
			"Value":     transaction.value,
		}
		transactionsData = append(transactionsData, transactionData)
	}
	blockData["Transactions"] = transactionsData

	blockJSON, err := json.MarshalIndent(blockData, "", "    ")
	if err != nil {
		fmt.Println("Error converting block data to JSON:", err)
		return
	}

	fmt.Println(strings.Repeat("=", 25), " Block ", strings.Repeat("=", 25))
	fmt.Println(string(blockJSON))
}

func MyPrintBlock(b *Block) {
	fmt.Println(strings.Repeat("=", 25), " Block ", strings.Repeat("=", 25))
	fmt.Printf("Nonce: %d\n", b.nonce)
	fmt.Printf("PreviousHash: %s\n", b.prevHash)

	fmt.Println("Transactions:")
	for _, transaction := range b.transactions {
		fmt.Printf("  TX ID: %s\n", transaction.TransactionID)
		fmt.Printf("  Sender: %s\n", transaction.senderBlockchainAddress)
		fmt.Printf("  Recipient: %s\n", transaction.recipientBlockchainAddress)
		fmt.Printf("  Value: %.2f\n", float32(transaction.value))
		fmt.Println(strings.Repeat("-", 40))
	}
}

func ListBlocks(bc *Blockchain) {
	fmt.Println("\n\n")
	fmt.Println("**********************************************************************************")
	fmt.Println("                                    Blockchain                                    ")
	fmt.Println("**********************************************************************************")
	fmt.Println("\n")

	for i := 0; i < len(bc.list); i++ {
		b := bc.list[i]
		PrintBlock(b)
		//PrintBlock(b)
	}
}

func ChangeBlock(b *Block, newTransaction *Transaction) {
	//function to change block transaction of the given block ref
	b.transactions[0] = newTransaction
	var stringToHash string
	stringToHash = MergeBlock(b)
	b.blockHash = CalculateHash(stringToHash)
}

func VerifyChain(bc *Blockchain, LastHash string) {

	// Function to verify the integrity of the blockchain by checking the hashes.
	for i := len(bc.list) - 1; i > 0; i-- {

		b := bc.list[i-1]

		var stringToHash string
		var hash string
		stringToHash = MergeBlock(b)
		hash = CalculateHash(stringToHash)

		if hash == LastHash {
			//fmt.Printf("Chain Valid at Block %d\n", i)
			LastHash = b.prevHash
			continue
		} else {
			fmt.Printf("\nChain Invalid at Block %d\n", i-1)
			return
		}
	}

	fmt.Println("\nChain Verified: All blocks are connected and valid.")
}

func main() {

	var stringToHash string

	//Make Blockchain
	blockchain := new(Blockchain)

	// Adding Transaction to Transaction Pool
	blockchain.AddTransaction("Carlos", "Sophia", 1.2)
	blockchain.AddTransaction("Liam", "Olivia", 3.7)
	blockchain.AddTransaction("Mia", "Ethan", 0.8)
	blockchain.AddTransaction("Oliver", "Sophia", 1.0)
	blockchain.AddTransaction("Emma", "William", 4.2)
	blockchain.AddTransaction("Sophia", "Lucas", 2.0)
	blockchain.AddTransaction("James", "Ava", 1.5)
	blockchain.AddTransaction("Benjamin", "Isabella", 3.1)
	blockchain.AddTransaction("William", "Grace", 1.8)
	blockchain.AddTransaction("Harper", "Elijah", 2.3)

	difficulty := 2 //Setting Difficulty

	//Adding some transactions to Blockchain
	b1 := blockchain.NewBlock("", difficulty)
	stringToHash = MergeBlock(b1)

	b2 := blockchain.NewBlock(CalculateHash(stringToHash), difficulty)
	stringToHash = MergeBlock(b2)

	LastHash := b2.prevHash //Store Last Hash Only

	ListBlocks(blockchain) //Print Chain

	VerifyChain(blockchain, LastHash) //Verify Chain

	/*
		//Tamper Chain
		t := new(Transaction)
		t.recipientBlockchainAddress = "Ameeroz"
		t.senderBlockchainAddress = "Haider"
		t.value = 90
		var TransactionString string
		TransactionString = MergeTransaction(t)
		t.TransactionID = CalculateHash(TransactionString)
		ChangeBlock(b3, t)

		//Print Tampered Chain
		ListBlocks(blockchain)

		//Verify Tampered Chain
		VerifyChain(blockchain, LastHash)

	*/

	// Creating 3 Nodes with same blockchain
	Node1 := &Node{blockchain: blockchain, difficulty: difficulty}
	Node2 := &Node{blockchain: blockchain, difficulty: difficulty}
	Node3 := &Node{blockchain: blockchain, difficulty: difficulty}

	//Nodes compete to mine new blocks
	LastHash = Node1.mineBlock()
	LastHash = Node2.mineBlock()
	LastHash = Node3.mineBlock()

	// Print the updated blockchain after mining
	ListBlocks(blockchain)

	// Verify the updated chain
	VerifyChain(blockchain, LastHash)
}
