// Package services provides a service implementation as decorator.
package service

import (
	"fmt"

	f "github.com/alanmvarela/golang-design-patterns/pkg/facade/facades"
)

type Service struct {
}

// Query sends a query to the server and returns the response.
func (s *Service) Query() (*[]string, error) {
	facade := f.ServiceFacade{}
	fmt.Println("Querying the facade from within the service...")
	return f.Query(&facade)
}
