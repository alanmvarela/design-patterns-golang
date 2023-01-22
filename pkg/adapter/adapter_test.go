// Package bridge provides a bridge example using go.
package adapter

import (
	"testing"

	a "github.com/alanmvarela/golang-design-patterns/pkg/adapter/adapters"
	c "github.com/alanmvarela/golang-design-patterns/pkg/adapter/third_party"
)

// TestAlphaClientAdapter tests the client adapter
func TestAlphaClientAdapter(t *testing.T) {
	client := c.NewAlphaClient()
	// Create a new client adapter
	adapter := a.NewAlphaClientAdapter(client)

	// Send a query to the client adapter
	response, err := adapter.Query("example.com")

	// Check if the response is correct
	if err != nil {
		t.Errorf("Query returned an error: %s", err)
	}
	if response == nil {
		t.Error("Query returned a nil response")
	}
	if response != nil && *response != "Response from Alpha API for query example.com" {
		t.Errorf("Query returned an incorrect response: %s", *response)
	}
}

// TestBethaClientAdapter tests the client adapter
func TestBethaClientAdapter(t *testing.T) {
	client := c.NewBethaClient()
	// Create a new client adapter
	adapter := a.NewBethaClientAdapter(client)

	// Send a query to the client adapter
	response, err := adapter.Query("example.com")

	// Check if the response is correct
	if err != nil {
		t.Errorf("Query returned an error: %s", err)
	}
	if response == nil {
		t.Error("Query returned a nil response")
	}
	if response != nil && *response != "Response from Betha API for query example.com" {
		t.Errorf("Query returned an incorrect response: %s", *response)
	}
}
