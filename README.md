# Extended Blockchain in Golang

## Project Description
This project builds upon a simple blockchain implementation, adding functionalities for handling transactions and enhancing the block structure. The blockchain will now support adding new transactions, printing block data in JSON format, and deriving nonces with a specified difficulty level.

## Features
- **Add New Block**: Create and add new blocks to the blockchain with specified transactions, nonces, and previous hashes.
- **List Blocks**: Print all blocks in a formatted manner, displaying transaction data, nonce, previous hash, and current block hash.
- **Calculate Hash**: Calculate the hash of a block.
- **Add New Transaction**: Add new transactions to the blockchain, with each transaction having a unique transaction ID.
- **Display Transactions**: Print the transactions of each block in JSON format.
- **Derive Nonce**: Introduce a method to derive a nonce for a block, with a specified difficulty level.

## Functions
The project includes the following public functions:

### `NewBlock(transaction string, nonce int, previousHash string) *block`
A method to add a new block. You can use any string as a transaction (e.g., "bob to alice") and any integer value as a nonce. The `CreateHash()` method will provide the block hash value.

### `ListBlocks()`
A method to print all the blocks in a nice format showing block data such as transaction, nonce, previous hash, and current block hash.

### `CalculateHash(stringToHash string)`
A function for calculating the hash of a block.

### `NewTransaction(sender string, recipient string, value float32) *Transaction`
A method to create a new transaction. The transaction ID is calculated as a hash value of the entire transaction.

### `AddTransaction(sender string, recipient string, value float32)`
A method to add a transaction to the transaction pool.

### `DisplayBlock(blockID int)`
A method to display the data of a particular block in JSON format.

### `DeriveNonce(block *Block, difficulty int)`
A method to derive a nonce for a block with a specified difficulty level. The nonce is added to the block data.

## Structs
The project includes the following structs:

### `Transaction`
```go
type Transaction struct {
    TransactionID              string
    senderBlockchainAddress    string
    recipientBlockchainAddress string
    value                      float32
}
```

### `Blockchain`
```go
type Blockchain struct {
    chain           []*Block
    transactionPool []*Transaction
}
```

## Repository Setup
1. Clone this repository:
    ```sh
    git clone https://github.com/ameerahaider/Extended-Blockchain-in-Golang.git
    ```
2. Navigate to the project directory:
    ```sh
    cd Extended-Blockchain-in-Golang
    ```
3. Ensure your Go environment is set up:
    ```sh
    go env
    ```
4. Run the project:
    ```sh
    go run main.go
    ```
