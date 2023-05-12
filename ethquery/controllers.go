package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetEthBlockNumber(ctx *fiber.Ctx, apiToken string) error {

	ethRPCMethod := "eth_blockNumber"
	url := fmt.Sprintf("https://mainnet.infura.io/v3/%s", apiToken)

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
