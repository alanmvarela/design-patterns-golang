// Package main implements and runs the adapter design pattern example.
package main

import (
	"fmt"
	"log"

	a "github.com/alanmvarela/golang-design-patterns/pkg/adapter/adapters"
	c "github.com/alanmvarela/golang-design-patterns/pkg/adapter/third_party"
)

func main() {
	defer fmt.Println("Adapter design pattern example starting")

	// The AlphaClient and BetaClient represents clients that can be used to query two different API.
	// This clients doesnt adjust to the Client interface, therefore we need an adapter.
	alphaClient := c.NewAlphaClient()
	BetaClient := c.NewBetaClient()

	// The BetaDNSAdapter represents an adapter that can be used to query the Beta API.
	// This adapter implements the Client interface, therefore it can be used as a client.
	alphaClientAdapter := a.NewAlphaClientAdapter(alphaClient)
	BetaClientAdapter := a.NewBetaClientAdapter(BetaClient)

	// Query sends a query to the server and returns the response.
	response_alpha, err := alphaClientAdapter.Query("alanmvarela.com")
	if err != nil {
		log.Panic("Error quering the client: ", err)
	}
	response_Beta, err := BetaClientAdapter.Query("google.com")
	if err != nil {
		log.Panic("Error quering the client: ", err)
	}
	fmt.Println(*response_alpha)
	fmt.Println(*response_Beta)
}
