// Package main implements and runs the adapter design pattern example.
package main

import (
	"fmt"
	"log"

	a "github.com/alanmvarela/golang-design-patterns/pkg/adapter/adapters"
	c "github.com/alanmvarela/golang-design-patterns/pkg/adapter/clients"
)

func main() {
	defer fmt.Println("Adapter design pattern example starting")

	// The AlphaClient and BethaClient represents clients that can be used to query two different API.
	// This clients doesnt adjust to the Client interface, therefore we need an adapter.
	alphaClient := c.NewAlphaClient()
	bethaClient := c.NewBethaClient()

	// The BethaDNSAdapter represents an adapter that can be used to query the Betha API.
	// This adapter implements the Client interface, therefore it can be used as a client.
	alphaClientAdapter := a.NewAlphaClientAdapter(alphaClient)
	bethaClientAdapter := a.NewBethaClientAdapter(bethaClient)

	// Query sends a query to the server and returns the response.
	response_alpha, err := alphaClientAdapter.Query("alanmvarela.com")
	if err != nil {
		log.Panic("Error quering the client: ", err)
	}
	response_betha, err := bethaClientAdapter.Query("google.com")
	if err != nil {
		log.Panic("Error quering the client: ", err)
	}
	fmt.Println(*response_alpha)
	fmt.Println(*response_betha)
}
