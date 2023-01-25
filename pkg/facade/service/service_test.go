// Package services provides a service implementation as decorator.
package service

import (
	"testing"
)

func TestService(t *testing.T) {
	s := Service{}
	if _, err := s.Query(); err != nil {
		t.Errorf("Error querying the facade from within the service: %v", err)
	}
}
