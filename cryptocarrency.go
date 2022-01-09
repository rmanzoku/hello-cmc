package coinmarketcap

import (
	"context"
	"encoding/json"
	"fmt"
)

type CryptocurrencyAPI struct {
	c *CoinMarketCap
}

func (c *CoinMarketCap) Cryptocurrency() *CryptocurrencyAPI {
	return &CryptocurrencyAPI{c}
}

type Currency struct {
	ID                int64             `json:"id"`
	Name              string            `json:"name"`
	Symbol            string            `json:"symbol"`
	Slug              string            `json:"slug"`
	NumMarketPairs    int64             `json:"num_market_pairs"`
	DateAdded         string            `json:"date_added"`
	Tags              []string          `json:"tags"`
	MaxSupply         interface{}       `json:"max_supply"`
	CirculatingSupply float64           `json:"circulating_supply"`
	TotalSupply       float64           `json:"total_supply"`
	IsActive          int64             `json:"is_active"`
	Platform          interface{}       `json:"platform"`
	CmcRank           int64             `json:"cmc_rank"`
	IsFiat            int64             `json:"is_fiat"`
	LastUpdated       string            `json:"last_updated"`
	Quote             map[string]*Quote `json:"quote"`
}

type Quote struct {
	Price            float64 `json:"price"`
	Volume24H        float64 `json:"volume_24h"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
	MarketCap        float64 `json:"market_cap"`
	LastUpdated      string  `json:"last_updated"`
}

func (c *CryptocurrencyAPI) QuoteLatest(ctx context.Context, symbolID uint, convertID uint) (*Quote, error) {
	path := fmt.Sprintf("cryptocurrency/quotes/latest?id=%d&convert_id=%d", symbolID, convertID)
	res, err := c.c.Get(ctx, path)
	if err != nil {
		return nil, err
	}
	ret := map[string]Currency{}
	err = json.Unmarshal(res, &ret)
	if err != nil {
		return nil, err
	}

	return ret[fmt.Sprint(symbolID)].Quote[fmt.Sprint(convertID)], nil
}

type QuotesHistoricalResponse struct {
	Data   Data   `json:"data"`
	Status Status `json:"status"`
}

type Data struct {
	ID       int64           `json:"id"`
	Name     string          `json:"name"`
	Symbol   string          `json:"symbol"`
	IsActive int64           `json:"is_active"`
	IsFiat   int64           `json:"is_fiat"`
	Quotes   []*QuoteElement `json:"quotes"`
}

type QuoteElement struct {
	Timestamp string            `json:"timestamp"`
	Quote     map[string]*Quote `json:"quote"`
}

type Status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int64  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int64  `json:"elapsed"`
	CreditCount  int64  `json:"credit_count"`
}

// https://coinmarketcap.com/api/documentation/v1/#operation/getV1CryptocurrencyQuotesHistorical
func (c *CryptocurrencyAPI) QuotesHistorical(ctx context.Context, symbolID uint, convertID uint, timeStart int64, interval string) ([]*Quote, error) {
	path := fmt.Sprintf("cryptocurrency/quotes/historical?id=%d&convert_id=%d&time_start=%d&interval=%s", symbolID, convertID, timeStart, interval)
	res, err := c.c.Get(ctx, path)
	if err != nil {
		return nil, err
	}
	tmp := new(QuotesHistoricalResponse)
	err = json.Unmarshal(res, tmp)
	if err != nil {
		return nil, err
	}

	ret := make([]*Quote, len(tmp.Data.Quotes))

	for i, q := range tmp.Data.Quotes {
		ret[i] = q.Quote[fmt.Sprint(convertID)]
	}

	return ret, nil
}
