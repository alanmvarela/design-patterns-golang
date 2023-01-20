// The Package adapter provides the adapter implementations to use the third party clients.
package adapter

import (
	c "github.com/alanmvarela/golang-design-patterns/pkg/adapter/clients"
)

// The BethaClientAdapter represents an adapter that can be used to query the Betha API.
// This adapter implements the Client interface, therefore it can be used as a client.
type BethaClientAdapter struct {
	// The BethaClient represents a client that can be used to query the Betha API.
	// This client doesnt adjust to the Client interface, therefore we need an adapter.
	client *c.BethaClient
}

func NewBethaClientAdapter(client *c.BethaClient) *BethaClientAdapter {
	return &BethaClientAdapter{client: client}
}

// Query sends a query to the server and returns the response.
func (b *BethaClientAdapter) Query(q string) (*string, error) {
	return b.client.BethaClientQuery(q)
}
