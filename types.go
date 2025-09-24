package okxdexapi

type BaseResponse[T any] struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type ApproveTransactionRequest struct {
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
	ApproveAmount        string `json:"approveAmount"`
}

type ApproveTransactionResponse struct {
	Data               string `json:"data"`
	DexContractAddress string `json:"dexContractAddress"`
	GasLimit           string `json:"gasLimit"`
	GasPrice           string `json:"gasPrice"`
}

type QuoteRequest struct {
	ChainIndex                      string `json:"chainIndex"`
	Amount                          string `json:"amount"`
	SwapMode                        string `json:"swapMode"`
	FromTokenAddress                string `json:"fromTokenAddress"`
	ToTokenAddress                  string `json:"toTokenAddress"`
	DexIds                          string `json:"dexIds"`
	DirectRoute                     bool   `json:"directRoute"`
	PriceImpactProtectionPercentage string `json:"priceImpactProtectionPercentage"`
	FeePercent                      string `json:"feePercent"`
}

type QuoteResponse struct {
	ChainId       string `json:"chainId"`
	DexRouterList []struct {
		Router        string `json:"router"`
		RouterPercent string `json:"routerPercent"`
		SubRouterList []struct {
			DexProtocol []struct {
				DexName string `json:"dexName"`
				Percent string `json:"percent"`
			} `json:"dexProtocol"`
			FromToken struct {
				Decimal              string `json:"decimal"`
				IsHoneyPot           bool   `json:"isHoneyPot"`
				TaxRate              string `json:"taxRate"`
				TokenContractAddress string `json:"tokenContractAddress"`
				TokenSymbol          string `json:"tokenSymbol"`
				TokenUnitPrice       string `json:"tokenUnitPrice"`
			} `json:"fromToken"`
			ToToken struct {
				Decimal              string `json:"decimal"`
				IsHoneyPot           bool   `json:"isHoneyPot"`
				TaxRate              string `json:"taxRate"`
				TokenContractAddress string `json:"tokenContractAddress"`
				TokenSymbol          string `json:"tokenSymbol"`
				TokenUnitPrice       string `json:"tokenUnitPrice"`
			} `json:"toToken"`
		} `json:"subRouterList"`
	} `json:"dexRouterList"`
	EstimateGasFee string `json:"estimateGasFee"`
	FromToken      struct {
		Decimal              string `json:"decimal"`
		IsHoneyPot           bool   `json:"isHoneyPot"`
		TaxRate              string `json:"taxRate"`
		TokenContractAddress string `json:"tokenContractAddress"`
		TokenSymbol          string `json:"tokenSymbol"`
		TokenUnitPrice       string `json:"tokenUnitPrice"`
	} `json:"fromToken"`
	FromTokenAmount       string `json:"fromTokenAmount"`
	OriginToTokenAmount   string `json:"originToTokenAmount"`
	PriceImpactPercentage string `json:"priceImpactPercentage"`
	QuoteCompareList      []struct {
		AmountOut string `json:"amountOut"`
		DexLogo   string `json:"dexLogo"`
		DexName   string `json:"dexName"`
		TradeFee  string `json:"tradeFee"`
	} `json:"quoteCompareList"`
	SwapMode string `json:"swapMode"`
	ToToken  struct {
		Decimal              string `json:"decimal"`
		IsHoneyPot           bool   `json:"isHoneyPot"`
		TaxRate              string `json:"taxRate"`
		TokenContractAddress string `json:"tokenContractAddress"`
		TokenSymbol          string `json:"tokenSymbol"`
		TokenUnitPrice       string `json:"tokenUnitPrice"`
	} `json:"toToken"`
	ToTokenAmount string `json:"toTokenAmount"`
	TradeFee      string `json:"tradeFee"`
}

type SwapInstructionRequest struct {
	ChainIndex                      string `json:"chainIndex"`
	Amount                          string `json:"amount"`
	FromTokenAddress                string `json:"fromTokenAddress"`
	ToTokenAddress                  string `json:"toTokenAddress"`
	Slippage                        string `json:"slippage"`
	AutoSlippage                    bool   `json:"autoSlippage"`
	MaxAutoSlippage                 string `json:"maxAutoSlippage"`
	UserWalletAddress               string `json:"userWalletAddress"`
	SwapReceiverAddress             string `json:"swapReceiverAddress"`
	FeePercent                      string `json:"feePercent"`
	FromTokenReferrerWalletAddress  string `json:"fromTokenReferrerWalletAddress"`
	ToTokenReferrerWalletAddress    string `json:"toTokenReferrerWalletAddress"`
	PositiveSlippagePercent         string `json:"positiveSlippagePercent"`
	PositiveSlippageFeeAddress      string `json:"positiveSlippageFeeAddress"`
	DexIds                          string `json:"dexIds"`
	ExcludeDexIds                   string `json:"excludeDexIds"`
	DisableRFQ                      string `json:"disableRFQ"`
	DirectRoute                     bool   `json:"directRoute"`
	PriceImpactProtectionPercentage string `json:"priceImpactProtectionPercentage"`
	ComputeUnitPrice                string `json:"computeUnitPrice"`
	ComputeUnitLimit                string `json:"computeUnitLimit"`
}

type SwapInstructionResponse struct {
	AddressLookupTableAccount []string `json:"addressLookupTableAccount"`
	InstructionLists          []struct {
		Data     string `json:"data"`
		Accounts []struct {
			IsSigner   bool   `json:"isSigner"`
			IsWritable bool   `json:"isWritable"`
			Pubkey     string `json:"pubkey"`
		} `json:"accounts"`
		ProgramId string `json:"programId"`
	} `json:"instructionLists"`
	RouterResult struct {
		ChainId       string `json:"chainId"`
		ChainIndex    string `json:"chainIndex"`
		DexRouterList []struct {
			Router        string `json:"router"`
			RouterPercent string `json:"routerPercent"`
			SubRouterList []struct {
				DexProtocol []struct {
					DexName string `json:"dexName"`
					Percent string `json:"percent"`
				} `json:"dexProtocol"`
				FromToken struct {
					Decimal              string `json:"decimal"`
					IsHoneyPot           bool   `json:"isHoneyPot"`
					TaxRate              string `json:"taxRate"`
					TokenContractAddress string `json:"tokenContractAddress"`
					TokenSymbol          string `json:"tokenSymbol"`
					TokenUnitPrice       string `json:"tokenUnitPrice"`
				} `json:"fromToken"`
				ToToken struct {
					Decimal              string `json:"decimal"`
					IsHoneyPot           bool   `json:"isHoneyPot"`
					TaxRate              string `json:"taxRate"`
					TokenContractAddress string `json:"tokenContractAddress"`
					TokenSymbol          string `json:"tokenSymbol"`
					TokenUnitPrice       string `json:"tokenUnitPrice"`
				} `json:"toToken"`
			} `json:"subRouterList"`
		} `json:"dexRouterList"`
		EstimateGasFee string `json:"estimateGasFee"`
		FromToken      struct {
			Decimal              string `json:"decimal"`
			IsHoneyPot           bool   `json:"isHoneyPot"`
			TaxRate              string `json:"taxRate"`
			TokenContractAddress string `json:"tokenContractAddress"`
			TokenSymbol          string `json:"tokenSymbol"`
			TokenUnitPrice       string `json:"tokenUnitPrice"`
		} `json:"fromToken"`
		FromTokenAmount       string `json:"fromTokenAmount"`
		PriceImpactPercentage string `json:"priceImpactPercentage"`
		QuoteCompareList      []struct {
			AmountOut string `json:"amountOut"`
			DexLogo   string `json:"dexLogo"`
			DexName   string `json:"dexName"`
			TradeFee  string `json:"tradeFee"`
		} `json:"quoteCompareList"`
		SwapMode string `json:"swapMode"`
		ToToken  struct {
			Decimal              string `json:"decimal"`
			IsHoneyPot           bool   `json:"isHoneyPot"`
			TaxRate              string `json:"taxRate"`
			TokenContractAddress string `json:"tokenContractAddress"`
			TokenSymbol          string `json:"tokenSymbol"`
			TokenUnitPrice       string `json:"tokenUnitPrice"`
		} `json:"toToken"`
		ToTokenAmount string `json:"toTokenAmount"`
		TradeFee      string `json:"tradeFee"`
	} `json:"routerResult"`
	Tx struct {
		From             string `json:"from"`
		MinReceiveAmount string `json:"minReceiveAmount"`
		Slippage         string `json:"slippage"`
		To               string `json:"to"`
	} `json:"tx"`
}

type SwapRequest struct {
	ChainIndex                      string `json:"chainIndex"`
	ChainId                         string `json:"chainId"`
	Amount                          string `json:"amount"`
	SwapMode                        string `json:"swapMode"`
	FromTokenAddress                string `json:"fromTokenAddress"`
	ToTokenAddress                  string `json:"toTokenAddress"`
	Slippage                        string `json:"slippage"`
	UserWalletAddress               string `json:"userWalletAddress"`
	SwapReceiverAddress             string `json:"swapReceiverAddress"`
	FeePercent                      string `json:"feePercent"`
	FromTokenReferrerWalletAddress  string `json:"fromTokenReferrerWalletAddress"`
	ToTokenReferrerWalletAddress    string `json:"toTokenReferrerWalletAddress"`
	PositiveSlippagePercent         string `json:"positiveSlippagePercent"`
	PositiveSlippageFeeAddress      string `json:"positiveSlippageFeeAddress"`
	Gaslimit                        string `json:"gaslimit"`
	GasLevel                        string `json:"gasLevel"`
	DexIds                          string `json:"dexIds"`
	DirectRoute                     bool   `json:"directRoute"`
	CallDataMemo                    string `json:"callDataMemo"`
	ComputeUnitPrice                string `json:"computeUnitPrice"`
	ComputeUnitLimit                string `json:"computeUnitLimit"`
	Tips                            string `json:"tips"`
	ExcludeDexIds                   string `json:"excludeDexIds"`
	DisableRFQ                      string `json:"disableRFQ"`
	PriceImpactProtectionPercentage string `json:"priceImpactProtectionPercentage"`
	AutoSlippage                    bool   `json:"autoSlippage"`
	MaxAutoSlippage                 string `json:"maxAutoSlippage"`
}

type SwapResponse struct {
	RouterResult struct {
		ChainId       string `json:"chainId"`
		DexRouterList []struct {
			Router        string `json:"router"`
			RouterPercent string `json:"routerPercent"`
			SubRouterList []struct {
				DexProtocol []struct {
					DexName string `json:"dexName"`
					Percent string `json:"percent"`
				} `json:"dexProtocol"`
				FromToken struct {
					Decimal              string `json:"decimal"`
					IsHoneyPot           bool   `json:"isHoneyPot"`
					TaxRate              string `json:"taxRate"`
					TokenContractAddress string `json:"tokenContractAddress"`
					TokenSymbol          string `json:"tokenSymbol"`
					TokenUnitPrice       string `json:"tokenUnitPrice"`
				} `json:"fromToken"`
				ToToken struct {
					Decimal              string `json:"decimal"`
					IsHoneyPot           bool   `json:"isHoneyPot"`
					TaxRate              string `json:"taxRate"`
					TokenContractAddress string `json:"tokenContractAddress"`
					TokenSymbol          string `json:"tokenSymbol"`
					TokenUnitPrice       string `json:"tokenUnitPrice"`
				} `json:"toToken"`
			} `json:"subRouterList"`
		} `json:"dexRouterList"`
		EstimateGasFee string `json:"estimateGasFee"`
		FromToken      struct {
			Decimal              string `json:"decimal"`
			IsHoneyPot           bool   `json:"isHoneyPot"`
			TaxRate              string `json:"taxRate"`
			TokenContractAddress string `json:"tokenContractAddress"`
			TokenSymbol          string `json:"tokenSymbol"`
			TokenUnitPrice       string `json:"tokenUnitPrice"`
		} `json:"fromToken"`
		FromTokenAmount       string `json:"fromTokenAmount"`
		PriceImpactPercentage string `json:"priceImpactPercentage"`
		QuoteCompareList      []struct {
			AmountOut string `json:"amountOut"`
			DexLogo   string `json:"dexLogo"`
			DexName   string `json:"dexName"`
			TradeFee  string `json:"tradeFee"`
		} `json:"quoteCompareList"`
		SwapMode string `json:"swapMode"`
		ToToken  struct {
			Decimal              string `json:"decimal"`
			IsHoneyPot           bool   `json:"isHoneyPot"`
			TaxRate              string `json:"taxRate"`
			TokenContractAddress string `json:"tokenContractAddress"`
			TokenSymbol          string `json:"tokenSymbol"`
			TokenUnitPrice       string `json:"tokenUnitPrice"`
		} `json:"toToken"`
		ToTokenAmount string `json:"toTokenAmount"`
		TradeFee      string `json:"tradeFee"`
	} `json:"routerResult"`
	Tx struct {
		Data                 string   `json:"data"`
		From                 string   `json:"from"`
		Gas                  string   `json:"gas"`
		GasPrice             string   `json:"gasPrice"`
		MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
		MaxSpendAmount       string   `json:"maxSpendAmount"`
		MinReceiveAmount     string   `json:"minReceiveAmount"`
		SignatureData        []string `json:"signatureData"`
		To                   string   `json:"to"`
		Value                string   `json:"value"`
	} `json:"tx"`
}

type GasPriceRequest struct {
	ChainIndex string `json:"chainIndex"`
}

type GasPriceResponse struct {
	Normal          string `json:"normal"`
	Min             string `json:"min"`
	Max             string `json:"max"`
	SupportEip1559  bool   `json:"supportEip1559"`
	Eip1559Protocol struct {
		SuggestBaseFee     string `json:"suggestBaseFee"`
		BaseFee            string `json:"baseFee"`
		ProposePriorityFee string `json:"proposePriorityFee"`
		SafePriorityFee    string `json:"safePriorityFee"`
		FastPriorityFee    string `json:"fastPriorityFee"`
	} `json:"eip1559Protocol"`
	PriorityFee struct {
		ProposePriorityFee string `json:"proposePriorityFee"`
		SafePriorityFee    string `json:"safePriorityFee"`
		FastPriorityFee    string `json:"fastPriorityFee"`
		ExtremePriorityFee string `json:"extremePriorityFee"`
	} `json:"priorityFee"`
}

type CurrentPriceRequest struct {
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
}

type CurrentPriceResponse struct {
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
	Time                 string `json:"time"`
	Price                string `json:"price"`
}

type AllTokenBalanceByAddressRequest struct {
	Address          string `json:"address"`
	Chains           string `json:"chains"`
	ExcludeRiskToken string `json:"excludeRiskToken"`
}

type AllTokenBalanceByAddressResponse struct {
	TokenAssets []struct {
		ChainIndex           string `json:"chainIndex"`
		TokenContractAddress string `json:"tokenContractAddress"`
		Symbol               string `json:"symbol"`
		Balance              string `json:"balance"`
		TokenPrice           string `json:"tokenPrice"`
		IsRiskToken          bool   `json:"isRiskToken"`
		RawBalance           string `json:"rawBalance"`
		Address              string `json:"address"`
	} `json:"tokenAssets"`
}

type TokenBalanceByAddressRequest struct {
	Address                string `json:"address"`
	TokenContractAddresses []struct {
		ChainIndex           string `json:"chainIndex"`
		TokenContractAddress string `json:"tokenContractAddress"`
	} `json:"tokenContractAddresses"`
	ExcludeRiskToken string `json:"excludeRiskToken"`
}

type TokenBalanceByAddressResponse struct {
	TokenAssets []struct {
		ChainIndex           string `json:"chainIndex"`
		TokenContractAddress string `json:"tokenContractAddress"`
		Symbol               string `json:"symbol"`
		Balance              string `json:"balance"`
		TokenPrice           string `json:"tokenPrice"`
		IsRiskToken          bool   `json:"isRiskToken"`
		RawBalance           string `json:"rawBalance"`
		Address              string `json:"address"`
	} `json:"tokenAssets"`
}

type TransactionByAddressRequest struct {
	Address              string `json:"address"`
	Chains               string `json:"chains"`
	TokenContractAddress string `json:"tokenContractAddress"`
	Begin                string `json:"begin"`
	End                  string `json:"end"`
	Cursor               string `json:"cursor"`
	Limit                string `json:"limit"`
}

type TransactionByAddressResponse struct {
	Cursor          string `json:"cursor"`
	Transactions []struct {
		ChainIndex string `json:"chainIndex"`
		TxHash     string `json:"txHash"`
		MethodId   string `json:"methodId"`
		Nonce      string `json:"nonce"`
		TxTime     string `json:"txTime"`
		From       []struct {
			Address string `json:"address"`
			Amount  string `json:"amount"`
		} `json:"from"`
		To []struct {
			Address string `json:"address"`
			Amount  string `json:"amount"`
		} `json:"to"`
		TokenContractAddress string `json:"tokenContractAddress"`
		Amount               string `json:"amount"`
		Symbol               string `json:"symbol"`
		TxFee                string `json:"txFee"`
		TxStatus             string `json:"txStatus"`
		HitBlacklist         bool   `json:"hitBlacklist"`
		Tag                  string `json:"tag"`
		Itype                string `json:"itype"`
	} `json:"transactions"`
}
