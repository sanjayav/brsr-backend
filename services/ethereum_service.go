package services

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
    "os"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

func LogReportToBlockchain(company string, ipfsHash string, year string) string {
    client, err := ethclient.Dial(os.Getenv("INFURA_URL"))
    if err != nil {
        log.Fatal("Failed to connect to Ethereum client:", err)
    }

    privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
    if err != nil {
        log.Fatal("Invalid private key:", err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        log.Fatal(err)
    }

    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.GasLimit = uint64(300000)
    auth.GasPrice = gasPrice

    contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

    // The following assumes you have generated Go bindings using abigen for the smart contract
    instance, err := NewBRSRReportRegistry(contractAddress, client)
    if err != nil {
        log.Fatal("Failed to load contract:", err)
    }

    tx, err := instance.SubmitReport(auth, ipfsHash, company, year)
    if err != nil {
        log.Fatal("Failed to submit report:", err)
    }

    fmt.Println("Report submitted with tx:", tx.Hash().Hex())
    return tx.Hash().Hex()
}
