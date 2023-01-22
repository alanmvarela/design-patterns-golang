// Package main implements and runs the bridge design pattern example.
package main

import (
	"fmt"
	"log"

	c "github.com/alanmvarela/golang-design-patterns/pkg/bridge/clients"
	m "github.com/alanmvarela/golang-design-patterns/pkg/bridge/credentials_manager"
	l "github.com/alanmvarela/golang-design-patterns/pkg/bridge/logger"
)

func main() {
	defer fmt.Println("Bridge design pattern example starting")

	// The AlphaClient and BethaClient represents clients that can be used to query two different API.
	// This clients doesnt adjust to the Client interface, therefore we need an adapter.
	alphaClient := c.NewAlphaClient()
	bethaClient := c.NewBethaClient()

	// The AlphaClient and BethaClient can be used with different logger and credentials manager.
	// The logger and credentials manager can be changed at runtime.
	alphaClient.SetLogger(&l.LoggerOne{})
	alphaClient.SetCredentialsManager(&m.CredentialsManagerOne{})
	bethaClient.SetLogger(&l.LoggerTwo{})
	bethaClient.SetCredentialsManager(&m.CredentialsManagerTwo{})

	// Query sends a query to the server and returns the response.
	response_alpha, err := alphaClient.Query("alanmvarela.com")
	if err != nil {
		log.Panic("Error quering the client: ", err)
	}
	response_betha, err := bethaClient.Query("google.com")
	if err != nil {
		log.Panic("Error quering the client: ", err)
	}
	fmt.Println(*response_alpha)
	fmt.Println(*response_betha)
}
