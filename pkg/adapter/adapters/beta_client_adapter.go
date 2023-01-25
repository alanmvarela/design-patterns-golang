// The Package adapters provides the adapter implementations to use the third party clients.
package adapters

import (
	c "github.com/alanmvarela/golang-design-patterns/pkg/adapter/third_party"
)

// The BetaClientAdapter represents an adapter that can be used to query the Beta API.
// This adapter implements the Client interface, therefore it can be used as a client.
type BetaClientAdapter struct {
	// The BetaClient represents a client that can be used to query the Beta API.
	// This client doesnt adjust to the Client interface, therefore we need an adapter.
	client *c.BetaClient
}

func NewBetaClientAdapter(client *c.BetaClient) *BetaClientAdapter {
	return &BetaClientAdapter{client: client}
}

// Query sends a query to the server and returns the response.
func (b *BetaClientAdapter) Query(q string) (*string, error) {
	return b.client.BetaClientQuery(q)
}
