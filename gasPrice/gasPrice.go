package gasPrice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"gasPrices/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

const infura_rpc = "https://mainnet.infura.io/v3/b4b812d0bf4a49f299e0646fdf081c11"
var (
	client *ethclient.Client
)

func init() {
	var err error
	client, err = ethclient.Dial(infura_rpc)
	if err != nil {
		log.Fatal(err)
	}
}

func FromWei(wei *big.Int, toType string) float64 {
	_wei, _ := wei.Float64()
	switch strings.ToLower(toType) {
	case "wei":
		return _wei
	case "gwei":
		return _wei / math.Pow(10, 9)
	case "ether":
		return _wei / math.Pow(10, 18)
	}
	return 0
}

func GetCurrentGasPrice() (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	return gasPrice, nil
}

func GetDetailedGasPrice() ([]float64, error) {
	requestURL := "https://gas.api.infura.io/networks/1/suggestedGasFees"
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(config.ApiKey, config.ApiSecretKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}
	defer resp.Body.Close()

	var respDecode map[any]any
	err = json.NewDecoder(resp.Body).Decode(&respDecode)
	if err != nil {
		return nil, err
	}

	low, err := strconv.ParseFloat(respDecode["low"].(map[string]interface{})["suggestedMaxFeePerGas"].(string), 64)
	if err != nil {
		return nil, err
	}

	medium, err := strconv.ParseFloat(respDecode["medium"].(map[string]interface{})["suggestedMaxFeePerGas"].(string), 64)
	if err != nil {
		return nil, err
	}

	high, err := strconv.ParseFloat(respDecode["high"].(map[string]interface{})["suggestedMaxFeePerGas"].(string), 64)
	if err != nil {
		return nil, err
	}


	result := []float64{low, medium, high}

	fmt.Println(1, result, low, medium, high)
	return result, nil
}

func FormatGasPrice(wei *big.Int) float64 {
	return FromWei(wei, "gwei")
}
