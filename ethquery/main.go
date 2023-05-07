package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*

RETRIEVE CURRENT BLOCK NUMBER:  eth_blockNumber



curl https://mainnet.infura.io/v3/2cecb47e279c45ccb3bec53828c56b17 \
    -X POST \
    -H "Content-Type: application/json" \
    --data '{"jsonrpc": "2.0", "id": 1, "method": "eth_blockNumber", "params": []}'

	{"jsonrpc":"2.0","id":1,"result":"0x1068702"}


*/

const apiKey = "2cecb47e279c45ccb3bec53828c56b17"
const portNr = 8080

type RequestBody struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type ResponseBody struct {
	Result string `json:"result"`
}

func GetEthBlockNumber(ctx *fiber.Ctx) error {

	ethRPCMethod := "eth_blockNumber"
	url := fmt.Sprintf("https://mainnet.infura.io/v3/%s", apiKey)

	requestBody := RequestBody{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  ethRPCMethod,
		Params:  []interface{}{},
	}

	jsonStr, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var responseBody ResponseBody
	err = json.Unmarshal(bodyBytes, &responseBody)
	if err != nil {
		panic(err)
	}

	return ctx.SendString(fmt.Sprintf("latest block number is: %s\n", responseBody.Result))

}

func main() {
	app := fiber.New()

	app.Get("/api/getethblocknumber", GetEthBlockNumber)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", portNr)))
}
