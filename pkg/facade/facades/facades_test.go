// Package facades provides the facades used by the facade package.
package facades

import (
	"testing"
)

// TestQuery tests the Query function.
func TestQuery(t *testing.T) {
	// Instantiate the facade
	facade := ServiceFacade{}
	// Query the facade
	response, err := Query(&facade)
	if err != nil {
		t.Errorf("Error querying the facade: %v", err)
	}
	// Check the response
	if len(*response) != 2 {
		t.Errorf("Expected response length of 2 but got %v", len(*response))
	}
}
