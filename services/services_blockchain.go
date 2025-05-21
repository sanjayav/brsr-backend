package services

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
    "strings"

    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

var (
    rpcURL      = "https://sepolia.infura.io/v3/YOUR_INFURA_KEY"            // or your RPC URL
    privateKey  = "YOUR_PRIVATE_KEY"                                        // NEVER expose in prod
    contractAddress = "0xYourDeployedContractAddress"                       // Deployed contract
    contractABI = `[{"inputs":[{"internalType":"string","name":"_company","type":"string"},{"internalType":"string","name":"_ipfsHash","type":"string"},{"internalType":"string","name":"_financialYear","type":"string"}],"name":"submitReport","outputs":[],"stateMutability":"payable","type":"function"}]`
)

// LogReportToBlockchain submits BRSR report to Ethereum
func LogReportToBlockchain(company, ipfsHash, year string) (string, error) {
    // Connect to Ethereum node
    client, err := ethclient.Dial(rpcURL)
    if err != nil {
        return "", fmt.Errorf("Failed to connect to Ethereum node: %v", err)
    }
    defer client.Close()

    // Load private key
    pk, err := crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
    if err != nil {
        return "", fmt.Errorf("Invalid private key: %v", err)
    }

    fromAddress := crypto.PubkeyToAddress(pk.PublicKey)

    // Get nonce
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        return "", fmt.Errorf("Failed to get nonce: %v", err)
    }

    // Prepare auth
    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        return "", err
    }

    auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
    if err != nil {
        return "", fmt.Errorf("Transactor error: %v", err)
    }

    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)           // No fine in this example
    auth.GasLimit = uint64(300000)
    auth.GasPrice, _ = client.SuggestGasPrice(context.Background())

    // Parse ABI
    parsedABI, err := abi.JSON(strings.NewReader(contractABI))
    if err != nil {
        return "", fmt.Errorf("ABI parse error: %v", err)
    }

    // Pack data
    inputData, err := parsedABI.Pack("submitReport", company, ipfsHash, year)
    if err != nil {
        return "", fmt.Errorf("Failed to pack data: %v", err)
    }

    // Create transaction
    tx := types.NewTransaction(nonce, common.HexToAddress(contractAddress), auth.Value, auth.GasLimit, auth.GasPrice, inputData)

    // Sign
    signedTx, err := auth.Signer(fromAddress, tx)
    if err != nil {
        return "", fmt.Errorf("Signing error: %v", err)
    }

    // Send
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        return "", fmt.Errorf("Transaction error: %v", err)
    }

    log.Printf("✅ Transaction submitted: %s", signedTx.Hash().Hex())
    return signedTx.Hash().Hex(), nil
}
