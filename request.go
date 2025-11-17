package okxdexapi

import (
	"io"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cenkalti/backoff/v5"
	"github.com/pkg/errors"
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
	var resp *http.Response
	var err error

	// use backoff to retry on 429 errors
	operation := func() (*http.Response, error) {
		resp, err := client.Do(req)
		if err != nil {
			return nil, errors.WithMessage(err, "http request failed")
		}

		// 如果是429错误，需要重试
		if resp.StatusCode == http.StatusTooManyRequests {
			_ = resp.Body.Close()
			return nil, errors.WithMessagef(errors.New("rate limited"), "http status %d", resp.StatusCode)
		}

		return resp, err
	}

	resp, err = backoff.Retry(
		req.Context(),
		operation,
		backoff.WithMaxTries(3),
		backoff.WithMaxElapsedTime(10*time.Second),
	)
	if err != nil {
		return nil, errors.WithMessage(err, "request failed after retries")
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.WithMessagef(err, "http status %d and read body failed", resp.StatusCode)
		}
		var body map[string]any
		if err := sonic.Unmarshal(bodyBytes, &body); err == nil {
			if msg, ok := body["msg"].(string); ok {
				return nil, errors.WithMessagef(errors.New(msg), "http status %d", resp.StatusCode)
			}
		}
		return nil, errors.WithMessagef(errors.New("unknown error"), "http status %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Read success body with error check
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "read response body failed")
	}

	var response BaseResponse[T]
	if err := sonic.Unmarshal(bodyBytes, &response); err != nil {
		// Try to decode error format to extract code and msg
		var anyResp BaseResponse[any]
		if err2 := sonic.Unmarshal(bodyBytes, &anyResp); err2 == nil {
			return nil, errors.WithMessagef(errors.New(anyResp.Msg), "response code %s", anyResp.Code)
		}
		return nil, errors.WithMessage(err, "unmarshal response failed")
	}

	if response.Code != "0" {
		return nil, errors.WithMessagef(errors.New(response.Msg), "response code %s", response.Code)
	}

	return response.Data, nil
}
