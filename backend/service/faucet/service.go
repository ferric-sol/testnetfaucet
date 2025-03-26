package faucet

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"solana-faucet-api/config"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/gagliardetto/solana-go/text"
)

func GetClientIP(r *http.Request) string {
	// 1. Check X-Forwarded-For (for proxies like Nginx, Cloudflare)
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// The header can contain multiple IPs (e.g., "client, proxy1, proxy2")
		ips := strings.Split(forwarded, ",")
		clientIP := strings.TrimSpace(ips[0]) // First IP is the original client
		if clientIP != "" {
			return clientIP
		}
	}

	// 2. Check X-Real-IP (another common proxy header)
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// 3. Fall back to RemoteAddr (direct connection)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // Fallback if splitting fails
	}

	return ip
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
	accountFrom, err := solana.PrivateKeyFromBase58(config.Envs.FundingWalletPrivateKey)
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
	fmt.Println(config.Envs.SOLTransactionLamports)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				// 1 sol = 1000000000 lamports
				// 1e6, // 0.001 SOL,
				config.Envs.SOLTransactionLamports,
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
