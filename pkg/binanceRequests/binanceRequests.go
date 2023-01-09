package binanceRequests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/adshao/go-binance/v2"
)

func GetCoins() {
	client := binance.NewClient(os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_SECRET_KEY"))
	res, err := client.NewGetAllCoinsInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, val := range res {
		if val.Coin == "CAKE" {
			CoinInfo, _ := json.Marshal(val)
			fmt.Println(string(CoinInfo))
		}
	}
}
