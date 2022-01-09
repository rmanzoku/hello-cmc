package main

import (
	"context"
	"fmt"
	"os"
	"time"

	coinmarketcap "github.com/rmanzoku/hello-cmc"
)

func main() {
	cmc := coinmarketcap.NewCoinMarketCap(os.Getenv("CMC_PRO_API_KEY"))

	q, err := cmc.Cryptocurrency().QuoteLatest(context.TODO(), coinmarketcap.IdMATIC, coinmarketcap.IdUSD)
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Now().Format(time.RFC3339), q.LastUpdated, q.Price)
}
