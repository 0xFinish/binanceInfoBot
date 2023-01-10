package binanceRequests

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/adshao/go-binance/v2"
)

var client *binance.Client

func init() {
	client = binance.NewClient(os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_SECRET_KEY"))
}

func GetCoins(args string) (CoinInfoReturn string) {
	res, err := client.NewGetAllCoinsInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, val := range res {
		if val.Coin == strings.ToUpper(args) {
			var CoinInfo string
			for _, networkVal := range val.NetworkList {
				CoinInfo = fmt.Sprintf("%s \n %s: with %s, dep %s, fee %s", CoinInfo, networkVal.Network, returnEmoji(networkVal.WithdrawEnable), returnEmoji(networkVal.DepositEnable), networkVal.WithdrawFee)
			}
			CoinInfoReturn = CoinInfo
		}
	}
	if CoinInfoReturn == "" {
		CoinInfoReturn = "The coin u entered does not exist"
	}
	return
}

func GetTickerPrices(args string) (TickerPrice string) {
	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Len of prices: %d\n", len(prices))
	for _, p := range prices {
		fmt.Println(p.Price)
		TickerPrice = fmt.Sprintf("%s \n Symbol: %s, Price: %s", TickerPrice, p.Symbol, p.Price)
	}
	fmt.Println(TickerPrice)
	return
}

func GetTickerPriceBySymbol(symbol string) (tickerPrice string) {
	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, p := range prices {
		if p.Symbol == strings.ToUpper(symbol)+"BUSD" {
			tickerPrice = p.Price
		}
	}
	return
}

func returnEmoji(boolValue bool) string {
	if boolValue {
		return "✅"
	}
	return "❌"
}
