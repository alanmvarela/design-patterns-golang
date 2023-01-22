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
	credentialsManager.GetCredentials("example.com")
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
	credentialsManager.GetCredentials("example.com")
}
