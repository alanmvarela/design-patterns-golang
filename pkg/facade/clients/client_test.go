// Package client provides client functions to send queries to different APIs
package clients

import (
	"testing"
)

func TestAlphaClient(t *testing.T) {
	alphaClient := AlphaClient{}
	response, err := alphaClient.Query()
	if err != nil {
		t.Errorf("Error querying Alpha API: %v", err)
	}
	if *response != "Alpha client response" {
		t.Errorf("Expected Alpha client response, got %v", *response)
	}
}

func TestBetaClient(t *testing.T) {
	betaClient := BetaClient{}
	response, err := betaClient.Query()
	if err != nil {
		t.Errorf("Error querying Beta API: %v", err)
	}
	if *response != "Beta client response" {
		t.Errorf("Expected Beta client response, got %v", *response)
	}
}
