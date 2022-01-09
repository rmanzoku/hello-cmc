package coinmarketcap

import (
	"os"
	"testing"
)

var c *CoinMarketCap

func TestMain(m *testing.M) {
	c = NewCoinMarketCap(os.Getenv("CMC_PRO_API_KEY"))
	code := m.Run()
	os.Exit(code)
}
