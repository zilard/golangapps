package main

import (
	"fmt"
	"log"

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

type BlockQuery struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
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

	// this will send back -> {"method":"eth_blockNumber","params":[]}
	ctx.JSON(fiber.Map{
		"method": req.Method,
		"params": req.Params,
	})

}

func main() {
	app := fiber.New()

	app.Post("/api/getethblocknumber", GetEthBlockNumber)

	log.Fatal(app.Listen(":8080"))
}
