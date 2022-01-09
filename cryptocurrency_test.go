package coinmarketcap

import (
	"context"
	"testing"
	"time"
)

func TestCryptocurrencyAPI_QuoteLatest(t *testing.T) {
	type fields struct {
		c *CoinMarketCap
	}
	type args struct {
		symbolID  uint
		convertID uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "JPY/ETH",
			fields: fields{c},
			args: args{
				symbolID:  IdETH,
				convertID: IdJPY,
			},
			wantErr: false,
		},
		{
			name:   "USD/ETH",
			fields: fields{c},
			args: args{
				symbolID:  IdETH,
				convertID: IdUSD,
			},
			wantErr: false,
		},
		{
			name:   "ETH/USD",
			fields: fields{c},
			args: args{
				symbolID:  IdUSD,
				convertID: IdETH,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.c.Cryptocurrency()
			got, err := c.QuoteLatest(context.TODO(), tt.args.symbolID, tt.args.convertID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptocurrencyAPI.QuoteLatest() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(got.Price)
		})
	}
}

func TestCryptocurrencyAPI_QuotesHistorical(t *testing.T) {
	type fields struct {
		c *CoinMarketCap
	}
	type args struct {
		symbolID  uint
		convertID uint
		startTime int64
		interval  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "MATIC/USD",
			fields: fields{c},
			args: args{
				symbolID:  IdMATIC,
				convertID: IdUSD,
				startTime: time.Now().Unix() - 86400,
				interval:  "1h",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.c.Cryptocurrency()
			got, err := c.QuotesHistorical(context.TODO(), tt.args.symbolID, tt.args.convertID, tt.args.startTime, tt.args.interval)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptocurrencyAPI.QuoteLatest() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(got)
		})
	}
}
