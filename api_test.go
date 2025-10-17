package okxdexapi_test

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	okxdexapi "github.com/CodeWizard198/okx-dex-api"
	"github.com/joho/godotenv"
)

var dexClient *okxdexapi.DexClient

func init() {
	_ = godotenv.Load(".env")
	proxy, _ := url.Parse("http://127.0.0.1:7897")
	dexClient = okxdexapi.NewDexClient(
		&okxdexapi.DexConfig{
			APIKey:     os.Getenv("OKX_API_KEY"),
			SecretKey:  os.Getenv("OKX_SECRET_KEY"),
			Passphrase: os.Getenv("OKX_PASSPHRASE"),
			BaseURL:    okxdexapi.DefaultBaseURL,
			Timeout:    30 * time.Second,
		},
		&http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
				DialContext: (&net.Dialer{
					Timeout:   10 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConns:          100,
				MaxIdleConnsPerHost:   10,
			},
		},
	)
}

func TestCurrentPrice(t *testing.T) {
	priceResponse, err := dexClient.CurrentPrice([]*okxdexapi.CurrentPriceRequest{
		{
			ChainIndex:           "56",
			TokenContractAddress: "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		},
	})
	if err != nil {
		t.Fatalf("Failed to get price: %v", err)
	}
	t.Logf("Price response: %+v", priceResponse)
}

func TestAllTokenBalanceByAddress(t *testing.T) {
	balanceResponse, err := dexClient.AllTokenBalanceByAddress(&okxdexapi.AllTokenBalanceByAddressRequest{
		Address:          "0xd2375f2dB683f318A6DDee99760967caE040357b",
		Chains:           "1,56,8453",
		ExcludeRiskToken: "1",
	})
	if err != nil {
		t.Fatalf("Failed to get all token balance: %v", err)
	}
	t.Logf("All token balance response: %+v", balanceResponse)
}

func TestGasPrice(t *testing.T) {
	gasPriceResponse, err := dexClient.GasPrice(&okxdexapi.GasPriceRequest{
		ChainIndex: "156",
	})
	if err != nil {
		t.Fatalf("Failed to get gas price: %v", err)
	}
	fmt.Printf("%+v", gasPriceResponse)
}

func TestTokenBalanceByAddress(t *testing.T) {
	balanceResponse, err := dexClient.TokenBalanceByAddress(&okxdexapi.TokenBalanceByAddressRequest{
		Address: "0xd2375f2dB683f318A6DDee99760967caE040357b",
		TokenContractAddresses: []struct {
			ChainIndex           string "json:\"chainIndex\""
			TokenContractAddress string "json:\"tokenContractAddress\""
		}{
			{
				ChainIndex:           "1",
				TokenContractAddress: "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			},
		},
	})
	if err != nil {
		t.Fatalf("Failed to get token balance: %v", err)
	}
	t.Logf("Token balance response: %+v", balanceResponse)
}

func TestTransactionByAddress(t *testing.T) {
	transactions, err := dexClient.TransactionByAddress(&okxdexapi.TransactionByAddressRequest{
		Address: "0xd2375f2dB683f318A6DDee99760967caE040357b",
		Chains:  "1,56,8453",
	})
	if err != nil {
		t.Fatalf("Failed to get transactions: %v", err)
	}
	t.Logf("Transactions response: %+v", transactions)
}

func TestDexClient_SwapInstruction(t *testing.T) {
	_, err := dexClient.SwapInstruction(&okxdexapi.SwapInstructionRequest{
		ChainIndex:                   "501",
		Amount:                       "18140000",
		FromTokenAddress:             "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB",
		ToTokenAddress:               "11111111111111111111111111111111",
		UserWalletAddress:            os.Getenv("WALLET_ADDRESS"),
		SwapReceiverAddress:          os.Getenv("WALLET_ADDRESS"),
		ToTokenReferrerWalletAddress: os.Getenv("FEE_WALLET_ADDRESS"),
		FeePercent:                   "1",
		AutoSlippage:                 true,
		SlippagePercent:              "5",
	})
	if err != nil {
		t.Logf("Failed to swap instruction: %v", err)
	}
}
