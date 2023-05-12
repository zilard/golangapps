package main

import (
	"github.com/gofiber/fiber/v2"
)

func EthRoutes(a *fiber.App, s *EthSession) {

	// Create routes group.
	route := a.Group("/api")

	route.Get("/getethblocknumber", func(c *fiber.Ctx) error {
		return GetEthBlockNumber(c, s.apiToken)
	})

}
