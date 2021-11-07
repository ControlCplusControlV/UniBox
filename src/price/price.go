package price

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetEthereumPriceInUSD() float64 {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=Ethereum&vs_currencies=usd")
	if err != nil {
		log.Fatal("error when getting ethereum price", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	priceMap := result["ethereum"].(map[string]interface{})
	return priceMap["usd"].(float64)
}
