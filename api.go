package okxdexapi

func (dex *DexClient) ApproveTransaction(params *ApproveTransactionRequest) (*ApproveTransactionResponse, error) {
	requestParams := map[string]any{}

	if params.ChainIndex != "" {
		requestParams["chainIndex"] = params.ChainIndex
	}
	if params.TokenContractAddress != "" {
		requestParams["tokenContractAddress"] = params.TokenContractAddress
	}
	if params.ApproveAmount != "" {
		requestParams["approveAmount"] = params.ApproveAmount
	}

	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v5/dex/aggregator/approve-transaction",
		Params:   requestParams,
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*ApproveTransactionResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	return data[0], nil
}

func (dex *DexClient) Quote(params *QuoteRequest) (*QuoteResponse, error) {
	requestParams := map[string]any{}

	if params.ChainIndex != "" {
		requestParams["chainIndex"] = params.ChainIndex
	}
	if params.Amount != "" {
		requestParams["amount"] = params.Amount
	}
	if params.SwapMode != "" {
		requestParams["swapMode"] = params.SwapMode
	}
	if params.FromTokenAddress != "" {
		requestParams["fromTokenAddress"] = params.FromTokenAddress
	}
	if params.ToTokenAddress != "" {
		requestParams["toTokenAddress"] = params.ToTokenAddress
	}
	if params.DexIds != "" {
		requestParams["dexIds"] = params.DexIds
	}
	requestParams["directRoute"] = params.DirectRoute
	if params.PriceImpactProtectionPercentage != "" {
		requestParams["priceImpactProtectionPercentage"] = params.PriceImpactProtectionPercentage
	}
	if params.FeePercent != "" {
		requestParams["feePercent"] = params.FeePercent
	}

	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v5/dex/aggregator/quote",
		Params:   requestParams,
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*QuoteResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	return data[0], nil
}

func (dex *DexClient) SwapInstruction(params *SwapInstructionRequest) (*SwapInstructionResponse, error) {
	requestParams := map[string]any{}

	if params.ChainIndex != "" {
		requestParams["chainIndex"] = params.ChainIndex
	}
	if params.Amount != "" {
		requestParams["amount"] = params.Amount
	}
	if params.FromTokenAddress != "" {
		requestParams["fromTokenAddress"] = params.FromTokenAddress
	}
	if params.ToTokenAddress != "" {
		requestParams["toTokenAddress"] = params.ToTokenAddress
	}
	if params.Slippage != "" {
		requestParams["slippage"] = params.Slippage
	}
	requestParams["autoSlippage"] = params.AutoSlippage
	if params.MaxAutoSlippage != "" {
		requestParams["maxAutoSlippage"] = params.MaxAutoSlippage
	}
	if params.UserWalletAddress != "" {
		requestParams["userWalletAddress"] = params.UserWalletAddress
	}
	if params.SwapReceiverAddress != "" {
		requestParams["swapReceiverAddress"] = params.SwapReceiverAddress
	}
	if params.FeePercent != "" {
		requestParams["feePercent"] = params.FeePercent
	}
	if params.FromTokenReferrerWalletAddress != "" {
		requestParams["fromTokenReferrerWalletAddress"] = params.FromTokenReferrerWalletAddress
	}
	if params.ToTokenReferrerWalletAddress != "" {
		requestParams["toTokenReferrerWalletAddress"] = params.ToTokenReferrerWalletAddress
	}
	if params.PositiveSlippagePercent != "" {
		requestParams["positiveSlippagePercent"] = params.PositiveSlippagePercent
	}
	if params.PositiveSlippageFeeAddress != "" {
		requestParams["positiveSlippageFeeAddress"] = params.PositiveSlippageFeeAddress
	}
	if params.DexIds != "" {
		requestParams["dexIds"] = params.DexIds
	}
	if params.ExcludeDexIds != "" {
		requestParams["excludeDexIds"] = params.ExcludeDexIds
	}
	if params.DisableRFQ != "" {
		requestParams["disableRFQ"] = params.DisableRFQ
	}
	requestParams["directRoute"] = params.DirectRoute
	if params.PriceImpactProtectionPercentage != "" {
		requestParams["priceImpactProtectionPercentage"] = params.PriceImpactProtectionPercentage
	}
	if params.ComputeUnitPrice != "" {
		requestParams["computeUnitPrice"] = params.ComputeUnitPrice
	}
	if params.ComputeUnitLimit != "" {
		requestParams["computeUnitLimit"] = params.ComputeUnitLimit
	}

	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v5/dex/aggregator/swap-instruction",
		Params:   requestParams,
	})
	if err != nil {
		return nil, err
	}
	return doRequest[*SwapInstructionResponse](dex.httpClient, req)
}

func (dex *DexClient) Swap(params *SwapRequest) (*SwapResponse, error) {
	requestParams := map[string]any{}

	if params.ChainIndex != "" {
		requestParams["chainIndex"] = params.ChainIndex
	}
	if params.ChainId != "" {
		requestParams["chainId"] = params.ChainId
	}
	if params.Amount != "" {
		requestParams["amount"] = params.Amount
	}
	if params.SwapMode != "" {
		requestParams["swapMode"] = params.SwapMode
	}
	if params.FromTokenAddress != "" {
		requestParams["fromTokenAddress"] = params.FromTokenAddress
	}
	if params.ToTokenAddress != "" {
		requestParams["toTokenAddress"] = params.ToTokenAddress
	}
	if params.Slippage != "" {
		requestParams["slippage"] = params.Slippage
	}
	if params.UserWalletAddress != "" {
		requestParams["userWalletAddress"] = params.UserWalletAddress
	}
	if params.SwapReceiverAddress != "" {
		requestParams["swapReceiverAddress"] = params.SwapReceiverAddress
	}
	if params.FeePercent != "" {
		requestParams["feePercent"] = params.FeePercent
	}
	if params.FromTokenReferrerWalletAddress != "" {
		requestParams["fromTokenReferrerWalletAddress"] = params.FromTokenReferrerWalletAddress
	}
	if params.ToTokenReferrerWalletAddress != "" {
		requestParams["toTokenReferrerWalletAddress"] = params.ToTokenReferrerWalletAddress
	}
	if params.PositiveSlippagePercent != "" {
		requestParams["positiveSlippagePercent"] = params.PositiveSlippagePercent
	}
	if params.PositiveSlippageFeeAddress != "" {
		requestParams["positiveSlippageFeeAddress"] = params.PositiveSlippageFeeAddress
	}
	if params.Gaslimit != "" {
		requestParams["gaslimit"] = params.Gaslimit
	}
	if params.GasLevel != "" {
		requestParams["gasLevel"] = params.GasLevel
	}
	if params.DexIds != "" {
		requestParams["dexIds"] = params.DexIds
	}
	requestParams["directRoute"] = params.DirectRoute
	if params.CallDataMemo != "" {
		requestParams["callDataMemo"] = params.CallDataMemo
	}
	if params.ComputeUnitPrice != "" {
		requestParams["computeUnitPrice"] = params.ComputeUnitPrice
	}
	if params.ComputeUnitLimit != "" {
		requestParams["computeUnitLimit"] = params.ComputeUnitLimit
	}
	if params.Tips != "" {
		requestParams["tips"] = params.Tips
	}
	if params.ExcludeDexIds != "" {
		requestParams["excludeDexIds"] = params.ExcludeDexIds
	}
	if params.DisableRFQ != "" {
		requestParams["disableRFQ"] = params.DisableRFQ
	}
	if params.PriceImpactProtectionPercentage != "" {
		requestParams["priceImpactProtectionPercentage"] = params.PriceImpactProtectionPercentage
	}
	requestParams["autoSlippage"] = params.AutoSlippage
	if params.MaxAutoSlippage != "" {
		requestParams["maxAutoSlippage"] = params.MaxAutoSlippage
	}

	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v5/dex/aggregator/swap",
		Params:   requestParams,
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*SwapResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	return data[0], nil
}

func (dex *DexClient) GasPrice(params *GasPriceRequest) (*GasPriceResponse, error) {
	requestParams := map[string]any{}

	if params.ChainIndex != "" {
		requestParams["chainIndex"] = params.ChainIndex
	}

	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v5/dex/pre-transaction/gas-price",
		Params:   requestParams,
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*GasPriceResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	return data[0], nil
}

func (dex *DexClient) CurrentPrice(payload []*CurrentPriceRequest) ([]*CurrentPriceResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   POST,
		Endpoint: "/api/v5/dex/index/current-price",
		Body:     payload,
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*CurrentPriceResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dex *DexClient) AllTokenBalanceByAddress(params *AllTokenBalanceByAddressRequest) (*AllTokenBalanceByAddressResponse, error) {
	requestParams := map[string]any{}

	if params.Address != "" {
		requestParams["address"] = params.Address
	}
	if params.Chains != "" {
		requestParams["chains"] = params.Chains
	}
	if params.ExcludeRiskToken != "" {
		requestParams["excludeRiskToken"] = params.ExcludeRiskToken
	}

	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v5/dex/balance/all-token-balances-by-address",
		Params:   requestParams,
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*AllTokenBalanceByAddressResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	return data[0], nil
}

func (dex *DexClient) TokenBalanceByAddress(payload *TokenBalanceByAddressRequest) (*TokenBalanceByAddressResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   POST,
		Endpoint: "/api/v5/dex/balance/token-balances-by-address",
		Body:     payload,
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*TokenBalanceByAddressResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	return data[0], nil
}

func (dex *DexClient) TransactionByAddress(params *TransactionByAddressRequest) (*TransactionByAddressResponse, error) {
	requestParams := map[string]any{}

	if params.Address != "" {
		requestParams["address"] = params.Address
	}
	if params.Chains != "" {
		requestParams["chains"] = params.Chains
	}
	if params.TokenContractAddress != "" {
		requestParams["tokenContractAddress"] = params.TokenContractAddress
	}
	if params.Begin != "" {
		requestParams["begin"] = params.Begin
	}
	if params.End != "" {
		requestParams["end"] = params.End
	}
	if params.Cursor != "" {
		requestParams["cursor"] = params.Cursor
	}
	if params.Limit != "" {
		requestParams["limit"] = params.Limit
	}

	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v5/dex/post-transaction/transactions-by-address",
		Params:   requestParams,
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*TransactionByAddressResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	return data[0], nil
}
