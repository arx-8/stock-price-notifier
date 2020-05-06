package main

import (
	"fmt"

	"github.com/arx-8/stock-price-notifier/src/infrastructure/investment_api"
)

func main() {
	fmt.Println("Hello world")

	fmt.Println(investment_api.GetRecentHistory())
}
