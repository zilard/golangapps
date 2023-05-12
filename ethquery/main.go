package main

import (
	"fmt"
	"log"

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

	EthRoutes(app, s)

	/*
		app.Get("/api/getethblocknumber", func(c *fiber.Ctx) error {
			return GetEthBlockNumber(c, s.apiToken)
		})
	*/

	fmt.Printf("port %s\n", s.port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", s.port)))
}
