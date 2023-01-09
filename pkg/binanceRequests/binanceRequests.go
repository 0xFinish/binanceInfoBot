package binanceRequests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/adshao/go-binance/v2"
)

func GetCoins(args string) (CoinInfoReturn string) {
	client := binance.NewClient(os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_SECRET_KEY"))
	res, err := client.NewGetAllCoinsInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, val := range res {
		if val.Coin == strings.ToUpper(args) {
			CoinInfo, _ := json.Marshal(val)
			CoinInfoReturn = string(CoinInfo)
		}
	}
	if CoinInfoReturn == "" {
		CoinInfoReturn = "The coin u entered does not exist"
	}
	return
}
