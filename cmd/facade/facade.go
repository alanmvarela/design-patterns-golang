// Package main implements and runs the facade design pattern example.
package main

import (
	"fmt"
	"log"

	s "github.com/alanmvarela/golang-design-patterns/pkg/facade/service"
)

func main() {
	// Create a new service
	service := s.Service{}
	// Query the service
	response, err := service.Query()
	if err != nil {
		log.Panic("Error quering the service: ", err)
	}
	fmt.Println(*response)
}
