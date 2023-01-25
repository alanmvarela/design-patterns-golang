// Package credentialsmanager provides a bridge to the credentials manager.
package credentialsmanager

import (
	"testing"
)

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
