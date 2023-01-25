// Package facades provides the facades used by the facade package.
package facades

import (
	"fmt"

	c "github.com/alanmvarela/golang-design-patterns/pkg/facade/clients"
)

// ServiceFacade provides a facade for the service package.
type ServiceFacade struct {
}

// Query sends a query to the server and returns the response.
func Query(f *ServiceFacade) (*[]string, error) {
	fmt.Println("Querying the client from within the facade...")
	// Instantiate all the clients
	clientAlpha := c.AlphaClient{}
	clientBeta := c.BetaClient{}
	// Initialize reponse slice
	queryReponse := []string{}
	// Query all the clients
	response, err := clientAlpha.Query()
	if err != nil {
		return nil, err
	}
	queryReponse = append(queryReponse, *response)
	response, err = clientBeta.Query()
	if err != nil {
		return nil, err
	}
	queryReponse = append(queryReponse, *response)

	// Return the response with the data fetched from all clients
	return &queryReponse, nil
}
