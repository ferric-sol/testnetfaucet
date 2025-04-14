package faucet

import (
	"context"
	"fmt"
	"os"
	"solana-faucet-api/configs"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/gagliardetto/solana-go/text"
	"github.com/gin-gonic/gin"
)

// GetClientIP extracts the client's IP address in a Gin-based API
func GetClientIP(c *gin.Context) string {
	// 1. Check X-Forwarded-For (used by proxies like Cloudflare, Nginx)
	forwarded := c.GetHeader("X-Forwarded-For")
	if forwarded != "" {
		ips := strings.Split(forwarded, ",")
		clientIP := strings.TrimSpace(ips[0]) // First IP is the original client
		if clientIP != "" {
			return clientIP
		}
	}

	// 2. Check X-Real-IP (another common proxy header)
	realIP := c.GetHeader("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// 3. Fall back to Gin’s ClientIP() function
	return c.ClientIP()
}

func SendSolanaTransaction(walletAddress string) (solana.Signature, error) {
	// Create a new RPC client on the TestNet
	rpcClient := rpc.New(rpc.TestNet_RPC)

	// Create a new WS client (used for confirming transactions)
	wsClient, err := ws.Connect(context.Background(), rpc.TestNet_WS)
	if err != nil {
		panic(err)
	}

	// Load the account that you will send funds FROM
	accountFrom, err := solana.PrivateKeyFromBase58(configs.Envs.FundingWalletPrivateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("accountFrom public key:", accountFrom.PublicKey())

	// Convert recipient address (walletAddress string) to a solana.PublicKey
	accountTo, err := solana.PublicKeyFromBase58(walletAddress)
	if err != nil {
		panic(fmt.Errorf("invalid recipient address: %w", err))
	}
	fmt.Println("accountTo public key:", accountTo)

	// WARNING: Get the recent blockhash (deprecated):
	// recent, err := rpcClient.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	// if err != nil {
	// 	panic(err)
	// }

	// Get the latest blockhash instead of the deprecated method
	recent, err := rpcClient.GetLatestBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to get latest blockhash: %w", err)
	}

	fmt.Println("recent blockhash", recent)
	fmt.Println(configs.Envs.SOLTransactionLamports)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				// 1 sol = 1000000000 lamports
				// 1e6, // 0.001 SOL,
				configs.Envs.SOLTransactionLamports,
				accountFrom.PublicKey(),
				accountTo,
			).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(accountFrom.PublicKey()),
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if accountFrom.PublicKey().Equals(key) {
				return &accountFrom
			}
			return nil
		},
	)
	if err != nil {
		panic(fmt.Errorf("unable to sign transaction: %w", err))
	}

	// Pretty print the transaction:
	tx.EncodeToTree(text.NewTreeEncoder(os.Stdout, "Transfer SOL"))

	// Send transaction, and wait for confirmation:
	sig, err := confirm.SendAndConfirmTransaction(
		context.TODO(),
		rpcClient,
		wsClient,
		tx,
	)
	if err != nil {
		panic(err)
	}

	return sig, nil
}

func IsValidSolanaAddress(address string) bool {
	_, err := solana.PublicKeyFromBase58(address)
	return err == nil
}

// GetWalletBalance fetches the balance of a given Solana wallet address
func GetWalletBalance(walletAddress string) (uint64, error) {
	client := rpc.New(rpc.TestNet_RPC) // Use Solana Mainnet RPC

	pubKey, err := solana.PublicKeyFromBase58(walletAddress)
	if err != nil {
		return 0, fmt.Errorf("invalid Solana wallet address: %v", err)
	}

	// Fetch balance with correct function signature
	balance, err := client.GetBalance(context.Background(), pubKey, rpc.CommitmentFinalized)
	if err != nil {
		return 0, fmt.Errorf("failed to get balance: %v", err)
	}

	return balance.Value, nil // Balance is returned in lamports
}
