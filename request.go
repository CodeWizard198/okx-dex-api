package okxdexapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type IType interface {
	~[]*ApproveTransactionResponse |
		~[]*QuoteResponse |
		~*SwapInstructionResponse |
		~[]*SwapResponse |
		~[]*GasPriceResponse |
		~[]*CurrentPriceResponse |
		~[]*AllTokenBalanceByAddressResponse |
		~[]*TokenBalanceByAddressResponse |
		~[]*TransactionByAddressResponse |
		~[]*MarketTokenSearchResponse |
		~[]*MarketTokenBasicInfoResponse |
		~[]*MarketPriceInfoResponse |
		[]*MarketTokenToplistResponse |
		[]*MarketTokenHolderResponse
}

func doRequest[T IType](client *http.Client, req *http.Request) (T, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	var response BaseResponse[T]
	//bodyBytes, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(bodyBytes))
	//return nil, nil
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decode response body failed: %w", err)
	}

	if Code(response.Code) != SuccessCode {
		if _, exist := errorMap[Code(response.Code)]; !exist {
			return nil, errors.New("unknown error:" + response.Msg)
		}
		return nil, errors.New(errorMap[Code(response.Code)].Category + ":" + errorMap[Code(response.Code)].MessageEN)
	}

	return response.Data, nil
}
