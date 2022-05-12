package main

import (
	"fmt"
	"encoding/json"
	"flag"
	"strings"
	"github.com/gocolly/colly"
)

type StockData struct {
	Ticker string
	CompanyName string
	Price string
	MarketCap string
	PE_Ratio string
	EPS string
	One_Year_Target_Price string
}

func getData(ticker string) StockData {
	c := colly.NewCollector()

	var companyName string
	var price string
	var marketCap string
	var PE string
	var EPS string
	var One_Year_Target_Price string

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		companyName = strings.Split(e.Text, " (")[0]
		fmt.Println(companyName)
	})

	c.OnHTML("[data-test=\"MARKET_CAP-value\"]", func(e *colly.HTMLElement) {
		marketCap = "$" + e.Text
	})

	c.OnHTML("[data-test=\"PE_RATIO-value\"]", func(e *colly.HTMLElement) {
		PE = e.Text
	})

	c.OnHTML("[data-field=\"regularMarketPrice\"]", func(e *colly.HTMLElement) {
		price = "$" + e.Text
	})

	c.OnHTML("[data-test=\"EPS_RATIO-value\"]", func(e *colly.HTMLElement) {
		EPS = e.Text
	})

	c.OnHTML("[data-test=\"ONE_YEAR_TARGET_PRICE-value\"]", func(e *colly.HTMLElement) {
		One_Year_Target_Price = "$" + e.Text
	})

	c.Visit("https://finance.yahoo.com/quote/" + ticker)

	return StockData{
		Ticker: ticker, 
		CompanyName: companyName,
		Price: price,
		MarketCap: marketCap,
		PE_Ratio: PE,
		EPS: EPS,
		One_Year_Target_Price: One_Year_Target_Price,
	}
}

func main() {
	ticker := flag.String("ticker", " ", "parse file path")
	flag.Parse()

	data := getData(*ticker)

	jsonData, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(jsonData))
}