package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
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

type BlockQuery struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
}

type RequestBody struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type ResponseBody struct {
	Result string `json:"result"`
}

func GetEthBlockNumber(ctx *fiber.Ctx) {

	var req BlockQuery

	ctx.BodyParser(&req)

	if len(req.Method) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Method not specified.",
		})
	}

	fmt.Printf("Received request with method '%s' and params '%v'\n",
		req.Method, req.Params)

	method := "eth_blockNumber"
	url := fmt.Sprintf("https://mainnet.infura.io/v3/%s", apiKey)

	requestBody := RequestBody{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  method,
		Params:  []interface{}{},
	}

	jsonStr, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	postreq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	postreq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(postreq)
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

	fmt.Println(responseBody.Result)

	ctx.JSON(fiber.Map{
		"respone": responseBody.Result,
	})

}

func main() {
	app := fiber.New()

	app.Post("/api/getethblocknumber", GetEthBlockNumber)

	log.Fatal(app.Listen(portNr))
}
