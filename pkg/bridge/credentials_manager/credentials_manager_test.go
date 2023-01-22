// Package credentialsmanager provides a bridge to the credentials manager.
package credentialsmanager

import (
	"testing"
)

// Test CredentialsManagerOne tests the credentials manager
func TestCredentialsManagerOne(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("CredentialsManagerOne panicked: %s", r)
		}
	}()
	credentialsManager := CredentialsManagerOne{}
	// Get credentials from the credentials manager
	credentials, error := credentialsManager.GetCredentials("example.com")
	if error != nil {
		t.Errorf("Error getting credentials: %s", error)
	}
	if *credentials != "Credentials for account example.com from credential manager One" {
		t.Errorf("Credentials are not correct: %s", *credentials)
	}

}

// Test CredentialsManagerTwo tests the credentials manager
func TestCredentialsManagerTwo(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("CredentialsManagerTwo panicked: %s", r)
		}
	}()
	credentialsManager := CredentialsManagerTwo{}
	// Get credentials from the credentials manager
	credentials, error := credentialsManager.GetCredentials("example.com")
	if error != nil {
		t.Errorf("Error getting credentials: %s", error)
	}
	if *credentials != "Credentials for account example.com from credential manager Two" {
		t.Errorf("Credentials are not correct: %s", *credentials)
	}

}
