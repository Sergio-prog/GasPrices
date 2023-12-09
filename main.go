package main

import (
	"gasPrices/api"
)

func main() {
	// for {
	// 	gas, err := gasPrice.GetCurrentGasPrice()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("Current Gas Price: %.2f Gwei\n", gasPrice.FormatGasPrice(gas))

	// 	time.Sleep(15 * time.Second)
	// }
	api.Run()
}
