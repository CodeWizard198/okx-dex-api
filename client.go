package okxdexapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type (
	HTTPMethod string
)

const DefaultBaseURL = "https://web3.okx.com"

const (
	GET  HTTPMethod = "GET"
	POST HTTPMethod = "POST"
)

type RequestParameters struct {
	Method   HTTPMethod
	Endpoint string
	Params   map[string]any
	Body     any
}

type DexConfig struct {
	APIKey     string
	SecretKey  string
	Passphrase string
	BaseURL    string
	Timeout    time.Duration
}

type DexClient struct {
	httpClient *http.Client
	config     *DexConfig
}

func NewDexClient(conf *DexConfig, httpClient ...*http.Client) *DexClient {
	var client *http.Client
	if len(httpClient) == 0 {
		if conf.Timeout == 0 {
			conf.Timeout = 30 * time.Second
		}
		client = &http.Client{
			Timeout: conf.Timeout,
		}
	} else {
		client = httpClient[0]
	}

	return &DexClient{
		httpClient: client,
		config:     conf,
	}
}

func (dex *DexClient) prepareRequest(requestParams RequestParameters) (*http.Request, error) {
	var (
		req        *http.Request
		err        error
		bodyBytes  []byte
		bodyString string
	)

	if requestParams.Body != nil {
		bodyBytes, err = json.Marshal(requestParams.Body)
		if err != nil {
			return nil, errors.WithMessage(err, "marshal request body failed")
		}
		bodyString = string(bodyBytes)
	}

	var fullURL string
	if requestParams.Method == GET && requestParams.Params != nil {
		fullURL = dex.buildURL(requestParams.Endpoint, requestParams.Params)
	} else {
		fullURL = dex.buildURL(requestParams.Endpoint, nil)
	}

	if bodyBytes != nil {
		req, err = http.NewRequest(string(requestParams.Method), fullURL, bytes.NewBuffer(bodyBytes))
	} else {
		req, err = http.NewRequest(string(requestParams.Method), fullURL, nil)
	}

	if err != nil {
		return nil, errors.WithMessage(err, "create http request failed")
	}

	timestamp := time.Now().UTC().Format(time.RFC3339)
	requestPath := requestParams.Endpoint
	if requestParams.Method == GET && requestParams.Params != nil {
		values := url.Values{}
		for k, v := range requestParams.Params {
			values.Add(k, fmt.Sprintf("%v", v))
		}
		builder := strings.Builder{}
		builder.WriteString(requestPath)
		builder.WriteString("?")
		builder.WriteString(values.Encode())
		requestPath = builder.String()
	}
	signature := dex.generateSignature(timestamp, string(requestParams.Method), requestPath, bodyString)

	req.Header.Set("OK-ACCESS-KEY", dex.config.APIKey)
	req.Header.Set("OK-ACCESS-SIGN", signature)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", dex.config.Passphrase)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (dex *DexClient) buildURL(endpoint string, params map[string]any) string {
	builder := strings.Builder{}
	builder.WriteString(dex.getBaseURL())
	builder.WriteString(endpoint)

	values := url.Values{}
	for k, v := range params {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	if len(values) > 0 {
		builder.WriteString("?")
		builder.WriteString(values.Encode())
	}

	return builder.String()
}

func (dex *DexClient) generateSignature(timestamp string, method string, requestPath string, body string) string {
	message := timestamp + method + requestPath + body
	h := hmac.New(sha256.New, []byte(dex.config.SecretKey))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (dex *DexClient) getBaseURL() string {
	if strings.TrimSpace(dex.config.BaseURL) == "" {
		return DefaultBaseURL
	}
	return dex.config.BaseURL
}
