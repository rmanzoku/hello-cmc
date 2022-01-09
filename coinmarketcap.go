package coinmarketcap

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	ProEndpoint     = "https://pro-api.coinmarketcap.com/v1/"
	SandboxEndpoint = "https://sandbox-api.coinmarketcap.com/v1/"
)

const (
	IdUSD uint = 2781
	IdJPY uint = 2797

	IdBTC   uint = 1
	IdETH   uint = 1027
	IdBNB   uint = 1839
	IdMATIC uint = 3890
)

type CoinMarketCap struct {
	Endpoint string
	ApiKey   string
	Client   *http.Client
}

func NewCoinMarketCap(apiKey string) *CoinMarketCap {
	return NewCoinMarketCapWithClient(apiKey, new(http.Client))
}

func NewCoinMarketCapWithClient(apiKey string, client *http.Client) *CoinMarketCap {
	return &CoinMarketCap{
		Endpoint: ProEndpoint,
		ApiKey:   apiKey,
		Client:   client,
	}
}

type responce struct {
	Status status          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

type status struct {
	Timestamp    string      `json:"timestamp"`
	ErrorCode    int64       `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	Elapsed      int64       `json:"elapsed"`
	CreditCount  int64       `json:"credit_count"`
	Notice       interface{} `json:"notice"`
}

func (c *CoinMarketCap) Get(ctx context.Context, path string) ([]byte, error) {
	url := c.Endpoint + path
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-CMC_PRO_API_KEY", c.ApiKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := new(responce)
	err = json.Unmarshal(bytes, res)
	if err != nil {
		return nil, err
	}
	if res.Status.ErrorCode != 0 {
		return nil, errors.New(res.Status.ErrorMessage)
	}

	return res.Data, nil
}
