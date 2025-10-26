package okxdexapi

import "fmt"

func (dex *DexClient) ApproveTransaction(params *ApproveTransactionRequest) (*ApproveTransactionResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v6/dex/aggregator/approve-transaction",
		Params:   params.ToMap(),
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*ApproveTransactionResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("empty response")
	}
	return data[0], nil
}

func (dex *DexClient) Quote(params *QuoteRequest) (*QuoteResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v6/dex/aggregator/quote",
		Params:   params.ToMap(),
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*QuoteResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("empty response")
	}
	return data[0], nil
}

func (dex *DexClient) SwapInstruction(params *SwapInstructionRequest) (*SwapInstructionResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v6/dex/aggregator/swap-instruction",
		Params:   params.ToMap(),
	})
	if err != nil {
		return nil, err
	}
	return doRequest[*SwapInstructionResponse](dex.httpClient, req)
}

func (dex *DexClient) Swap(params *SwapRequest) (*SwapResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v6/dex/aggregator/swap",
		Params:   params.ToMap(),
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*SwapResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("empty response")
	}
	return data[0], nil
}

func (dex *DexClient) GasPrice(params *GasPriceRequest) (*GasPriceResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v6/dex/pre-transaction/gas-price",
		Params:   params.ToMap(),
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*GasPriceResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("empty response")
	}
	return data[0], nil
}

func (dex *DexClient) CurrentPrice(payload []*CurrentPriceRequest) ([]*CurrentPriceResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   POST,
		Endpoint: "/api/v6/dex/index/current-price",
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
	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v6/dex/balance/all-token-balances-by-address",
		Params:   params.ToMap(),
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*AllTokenBalanceByAddressResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("empty response")
	}
	return data[0], nil
}

func (dex *DexClient) TokenBalanceByAddress(payload *TokenBalanceByAddressRequest) (*TokenBalanceByAddressResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   POST,
		Endpoint: "/api/v6/dex/balance/token-balances-by-address",
		Body:     payload,
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*TokenBalanceByAddressResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("empty response")
	}
	return data[0], nil
}

func (dex *DexClient) TransactionByAddress(params *TransactionByAddressRequest) (*TransactionByAddressResponse, error) {
	fmt.Println(params.ToMap()["chains"])
	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v6/dex/post-transaction/transactions-by-address",
		Params:   params.ToMap(),
	})
	if err != nil {
		return nil, err
	}
	data, err := doRequest[[]*TransactionByAddressResponse](dex.httpClient, req)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("empty response")
	}
	return data[0], nil
}

func (dex *DexClient) MarketTokenSearch(params *MarketTokenSearchRequest) ([]*MarketTokenSearchResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v6/dex/market/token/search",
		Params:   params.ToMap(),
	})
	if err != nil {
		return nil, err
	}
	return doRequest[[]*MarketTokenSearchResponse](dex.httpClient, req)
}

func (dex *DexClient) MarketTokenBasicInfo(payload []*MarketTokenBasicInfoRequest) ([]*MarketTokenBasicInfoResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   POST,
		Endpoint: "/api/v6/dex/market/token/basic-info",
		Body:     payload,
	})
	if err != nil {
		return nil, err
	}
	return doRequest[[]*MarketTokenBasicInfoResponse](dex.httpClient, req)
}

func (dex *DexClient) MarketPriceInfo(payload []*MarketPriceInfoRequest) ([]*MarketPriceInfoResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   POST,
		Endpoint: "/api/v6/dex/market/price-info",
		Body:     payload,
	})
	if err != nil {
		return nil, err
	}
	return doRequest[[]*MarketPriceInfoResponse](dex.httpClient, req)
}

func (dex *DexClient) MarketTokenToplist(params *MarketTokenToplistRequest) ([]*MarketTokenToplistResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v6/dex/market/token/toplist",
		Params:   params.ToMap(),
	})
	if err != nil {
		return nil, err
	}
	return doRequest[[]*MarketTokenToplistResponse](dex.httpClient, req)
}

func (dex *DexClient) MarketTokenHolder(params *MarketTokenHolderRequest) ([]*MarketTokenHolderResponse, error) {
	req, err := dex.prepareRequest(RequestParameters{
		Method:   GET,
		Endpoint: "/api/v6/dex/market/token/holder",
		Params:   params.ToMap(),
	})
	if err != nil {
		return nil, err
	}
	return doRequest[[]*MarketTokenHolderResponse](dex.httpClient, req)
}
