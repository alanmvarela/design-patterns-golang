// The Package thirdParty contains the third party clients that can be used to query servers.
package thirdParty

import "errors"

// The BetaClient represents a client that can be used to query the Beta API.
// This client doesnt adjust to the Client interface, therefore we need an adapter.
type BetaClient struct {
}

// The NewBetaClient function creates a new BetaClient struct.
func NewBetaClient() *BetaClient {
	return &BetaClient{}
}

// The Query method sends a query to the Beta API and returns the response.
func (c *BetaClient) BetaClientQuery(q string) (*string, error) {
	if q == "" {
		return nil, errors.New("query cannot be empty")
	}
	response := "Response from Beta API for query " + q
	return &response, nil
}
