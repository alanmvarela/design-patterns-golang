// Package services provides a service implementation as decorator.
package service

import (
	"testing"

	c "github.com/alanmvarela/golang-design-patterns/pkg/decorator/clients"
)

func TestService(t *testing.T) {
	client := &c.AlphaClient{}
	service := NewService(client)
	response, err := service.Query()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if *response != "Response from decorator: Alpha client response" {
		t.Errorf("Unexpected response: %s", *response)
	}
}
