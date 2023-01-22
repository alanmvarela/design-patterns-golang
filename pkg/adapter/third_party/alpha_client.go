// The Package thirdParty contains the third party clients that can be used to query servers.
package thirdParty

import "errors"

// The AlphaClient represents a client that can be used to query the Alpha API.
// This client doesnt adjust to the Client interface, therefore we need an adapter.
type AlphaClient struct {
}

// The NewAlphaClient function creates a new AlphaClient struct.
func NewAlphaClient() *AlphaClient {
	return &AlphaClient{}
}

// The Query method sends a query to the Alpha API and returns the response.
func (c *AlphaClient) AlphaClientQuery(q string) (*string, error) {
	if q == "" {
		return nil, errors.New("query cannot be empty")
	}
	response := "Response from Alpha API for query " + q
	return &response, nil
}
