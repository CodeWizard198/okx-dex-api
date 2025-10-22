package okxdexapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
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

	// Non-200: read body for error message and include status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("http status %d read error body failed: %w", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("http status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Read success body with error check
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %w", err)
	}

	var response BaseResponse[T]
	if err := sonic.Unmarshal(bodyBytes, &response); err != nil {
		// Try to decode error format to extract code and msg
		var anyResp BaseResponse[any]
		if err2 := sonic.Unmarshal(bodyBytes, &anyResp); err2 == nil {
			return nil, fmt.Errorf("[%s]%s", anyResp.Code, anyResp.Msg)
		}
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	if response.Code != "0" {
		return nil, fmt.Errorf("[%s]%s", response.Code, response.Msg)
	}

	return response.Data, nil
}
