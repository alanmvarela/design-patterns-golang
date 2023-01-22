// The Package adapters provides the adapter implementations to use the third party clients.
package adapters

import (
	c "github.com/alanmvarela/golang-design-patterns/pkg/adapter/third_party"
)

// The AlphaClientAdapter represents an adapter that can be used to query the Alpha API.
// This adapter implements the Client interface, therefore it can be used as a client.
type AlphaClientAdapter struct {
	// The AlphaClient represents a client that can be used to query the Alpha API.
	// This client doesnt adjust to the Client interface, therefore we need an adapter.
	client *c.AlphaClient
}

func NewAlphaClientAdapter(client *c.AlphaClient) *AlphaClientAdapter {
	return &AlphaClientAdapter{client: client}
}

// Query sends a query to the server and returns the response.
func (b *AlphaClientAdapter) Query(q string) (*string, error) {
	return b.client.AlphaClientQuery(q)
}
