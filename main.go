package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nanmu42/etherscan-api"
)

func main() {
	http.HandleFunc("/", showTransactions)
	http.ListenAndServe(":8080", nil)
}

func showTransactions(resp http.ResponseWriter, req *http.Request) {

	var ConAddr = "0x1da200f724b6e707cD8B8593f2c270771B7FC769"
	var StartBlock int = 11858824
	var EndBlock int = 12392042

	client := etherscan.NewCustomized(etherscan.Customization{
		Timeout: 5 * time.Second,
		Key:     "77JYKV2Q37VFE8RRA6GCEAX1SZ53EN4379",
		BaseURL: "https://api.bscscan.com/api?",
		Verbose: false,
	})

	fmt.Println("Connection to BSCSCAN is created")

	txs, err := client.NormalTxByAddress(ConAddr, &StartBlock, &EndBlock, 1, 100, false)

	jsonTXdata, _ := json.MarshalIndent(txs, "", "  ")
	jsonTXFile, err := os.Create("./tranactions.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonTXFile.Close()
	jsonTXFile.Write(jsonTXdata)
	jsonTXFile.Close()
	fmt.Printf("%s\n", jsonTXdata)

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonTXdata)
	return
}
