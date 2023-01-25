// Package services provides a service implementation as decorator.
package service

import (
	"fmt"

	c "github.com/alanmvarela/golang-design-patterns/pkg/decorator/clients"
)

type Service struct {
	client c.Client
}

// NewService creates a new Service struct.
func NewService(client c.Client) *Service {
	return &Service{
		client: client,
	}
}

// Query sends a query to the server and returns the response.
func (s *Service) Query() (*string, error) {
	fmt.Println("Querying the client from within the decorator...")
	response, err := s.client.Query()
	if err != nil {
		return nil, err
	}
	*response = fmt.Sprintf("Response from decorator: %s", *response)
	return response, nil
}
