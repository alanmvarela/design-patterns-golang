// Package main implements and runs the decorator design pattern example.
package main

import (
	"fmt"
	"log"

	c "github.com/alanmvarela/golang-design-patterns/pkg/decorator/clients"
	s "github.com/alanmvarela/golang-design-patterns/pkg/decorator/service"
)

func main() {
	// Set the service decorator using Alpha client.
	service_alpha := s.NewService(&c.AlphaClient{})
	// Query using the AlphaClient
	response, err := service_alpha.Query()
	if err != nil {
		log.Panic("Error quering the client: ", err)
	}
	fmt.Println(*response)

	// Set the service decorator using Beta client.
	service_beta := s.NewService(&c.BetaClient{})
	// Query using the BetaClient
	response, err = service_beta.Query()
	if err != nil {
		log.Panic("Error quering the client: ", err)
	}
	fmt.Println(*response)

}
