package profit_calculation

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	HistroyPriceURL = "https://gist.githubusercontent.com/Jekiwijaya/c72c2de532203965bf818e5a4e5e43e3/raw/2631344d08b044a4b833caeab8a42486b87cc19a/gistfile1.txt"
)

var (
	client = &http.Client{}
)

func GetHistoricalPrices() ([]string, error) {
	req, err := http.NewRequest("GET", HistroyPriceURL, nil)
	if err != nil {
		return []string{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(b), " "), nil
}

func CalculateProfit(prices []string) (int, error) {
	var lowestPrice, highestPrice int
	if len(prices) == 0 {
		return 0, errors.New("prices is empty")
	}
	for index, price := range prices {
		price, err := strconv.Atoi(price)
		if err != nil {
			return 0, err
		}

		if index == 0 {
			lowestPrice = price
			highestPrice = price
			continue
		}

		if lowestPrice > price {
			lowestPrice = price
			highestPrice = price
		}

		if highestPrice < price {
			highestPrice = price
		}
	}
	return (highestPrice - lowestPrice), nil
}
