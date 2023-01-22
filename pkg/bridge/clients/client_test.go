// The Package client provides clients that can be used to query servers.
package clients

import (
	"fmt"
	"testing"

	m "github.com/alanmvarela/golang-design-patterns/pkg/bridge/credentials_manager"
	l "github.com/alanmvarela/golang-design-patterns/pkg/bridge/logger"
)

// TestAlphaClient tests the AlphaClient.
func TestAlphaClient(t *testing.T) {
	client := NewAlphaClient()
	client.SetLogger(&l.LoggerOne{})
	client.SetCredentialsManager(&m.CredentialsManagerOne{})
	response, error := client.Query("test")
	if error != nil {
		t.Error(error)
	}
	expectedResponse := "Response from Alpha API for query test with credentials: Credentials for account alpha from credential manager One"
	fmt.Println(*response)
	fmt.Println(expectedResponse)
	if *response != expectedResponse {
		t.Error("Unexpected response")
	}
}

// TestBethaClient tests the BetaClient.
func TestBethaClient(t *testing.T) {
	client := NewBethaClient()
	client.SetLogger(&l.LoggerTwo{})
	client.SetCredentialsManager(&m.CredentialsManagerTwo{})
	response, error := client.Query("test")
	if error != nil {
		t.Error(error)
	}
	expectedResponse := "Response from Betha API for query test with credentials: Credentials for account betha from credential manager Two"
	fmt.Println(*response)
	fmt.Println(expectedResponse)
	if *response != expectedResponse {
		t.Error("Unexpected response")
	}
}
