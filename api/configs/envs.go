package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost              string
	Port                    string
	DBUser                  string
	DBPassword              string
	DBAddress               string
	DBName                  string
	JWTSecret               string
	JWTExpirationInSeconds  int64
	FundingWalletPrivateKey string
	FundingWalletPublicKey  string
	SOLTransactionLamports  uint64
	RecaptchaSecretKey      string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:              getEnv("PUBLIC_HOST", "http:"),
		Port:                    getEnv("PORT", "8080"),
		DBUser:                  getEnv("DB_USER", "root"),
		DBPassword:              getEnv("DB_PASSWORD", ""),
		DBAddress:               fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                  getEnv("DB_NAME", "solana-faucet-api"),
		JWTSecret:               getEnv("JWT_SECRET", "not-secret-anymore?"),
		JWTExpirationInSeconds:  getEnvAsInt("JWT_EXP", 3600*24*7),
		FundingWalletPrivateKey: getEnv("FUNDING_WALLET_PRIVATE_KEY", "............."),
		FundingWalletPublicKey:  getEnv("FUNDING_WALLET_PUBLIC_KEY", "............."),
		SOLTransactionLamports:  getEnvAsUint64("SOL_TRANSACTION_LAMPORTS", 100000), // Default to 100,000 lamports (0.0001 SOL)
		RecaptchaSecretKey:      getEnv("RECAPTCHA_SECRET_KEY", "not-secret-anymore?"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func getEnvAsUint64(key string, fallback uint64) uint64 {
	if value, ok := os.LookupEnv(key); ok {
		u, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return fallback
		}
		return u
	}
	return fallback
}
