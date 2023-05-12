package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type EthSession struct {
	apiToken string
	port     string
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

// initializes the root command and its flags
func initSession() *EthSession {

	s := EthSession{}

	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "ethquery",
		Short: "query Ethereum mainnet using Infura",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	// Add a flag for apitoken
	rootCmd.Flags().StringVarP(&s.apiToken, "apitoken", "t", "", "The API token to use")

	// Add a flag for port
	rootCmd.Flags().StringVarP(&s.port, "port", "p", "", "The port to use")

	rootCmd.Execute()

	return &s
}

func main() {

	// Call the initSession function to initialize and execute the root command
	s := initSession()

	app := fiber.New()

	app.Get("/api/getethblocknumber", func(c *fiber.Ctx) error {
		return GetEthBlockNumber(c, s.apiToken)
	})

	fmt.Printf("port %s\n", s.port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", s.port)))
}
