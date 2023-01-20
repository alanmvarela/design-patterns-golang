// The Package client provides clients that can be used to query servers.
package clients

import "errors"

// The BethaClient represents a client that can be used to query the Betha API.
// This client doesnt adjust to the Client interface, therefore we need an adapter.
type BethaClient struct {
}

// The NewBethaClient function creates a new BethaClient struct.
func NewBethaClient() *BethaClient {
	return &BethaClient{}
}

// The Query method sends a DNS query to the Betha API and returns the response.
func (c *BethaClient) BethaClientQuery(q string) (*string, error) {
	if q == "" {
		return nil, errors.New("query cannot be empty")
	}
	response := "Response from Betha API for query " + q
	return &response, nil
}
