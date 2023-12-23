package main

import (
	"testing"
)

func TestCreatePolicyString(t *testing.T) {
	// Define a test case
	testCase := VaultPolicy{
		Path:         "secret/data/test",
		Capabilities: []string{"read", "write"},
	}

	// Expected result
	expected := `path "secret/data/test" {
  capabilities = ["read", "write"]
}`

	// Call the CreatePolicyString method
	result := testCase.CreatePolicyString()

	// Compare the result with the expected value
	if result != expected {
		t.Errorf("CreatePolicyString() returned unexpected result.\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestFormatCapabilities(t *testing.T) {
	// Test case 1: Non-empty capabilities
	capabilities1 := []string{"read", "write", "delete"}
	result1 := formatCapabilities(capabilities1)
	expected1 := `["read", "write", "delete"]`

	if result1 != expected1 {
		t.Errorf("formatCapabilities() returned unexpected result.\nExpected:\n%s\nGot:\n%s", expected1, result1)
	}

	// Test case 2: Empty capabilities
	capabilities2 := []string{}
	result2 := formatCapabilities(capabilities2)
	expected2 := "[]"

	if result2 != expected2 {
		t.Errorf("formatCapabilities() returned unexpected result.\nExpected:\n%s\nGot:\n%s", expected2, result2)
	}
}

// func TestCreateNewVaultClient(t *testing.T) {
// 	// Test case 1: Valid host and token
// 	host1 := "http://localhost:8200"
// 	token1 := "myroot"

// 	client1, err1 := createNewVaultClient(host1, token1)
// 	if err1 != nil {
// 		t.Errorf("createNewVaultClient() returned an unexpected error for valid input: %v", err1)
// 	}
// 	if client1 == nil {
// 		t.Error("createNewVaultClient() returned a nil client for valid input")
// 	}

// 	// Test case 2: Invalid host
// 	host2 := "invalid-host"
// 	token2 := "myroot"
// 	client2, err2 := createNewVaultClient(host2, token2)
// 	if err2 == nil {
// 		t.Error("createNewVaultClient() did not return an error for an invalid host")
// 	}
// 	if client2 != nil {
// 		t.Error("createNewVaultClient() returned a non-nil client for an invalid host")
// 	}

// 	// Test case 3: Invalid token
// 	host3 := "http://localhost:8200"
// 	token3 := ""

// 	client3, err3 := createNewVaultClient(host3, token3)
// 	if err3 == nil {
// 		t.Error("createNewVaultClient() did not return an error for an empty token")
// 	}
// 	if client3 != nil {
// 		t.Error("createNewVaultClient() returned a non-nil client for an empty token")
// 	}
// }

func TestCreateNewVaultClient(t *testing.T) {
	// Test case 1: Valid host and token
	host1 := "http://localhost:8200"
	token1 := "myroot"

	client1, err1 := createNewVaultClient(host1, token1)
	if err1 != nil {
		t.Errorf("createNewVaultClient() returned an unexpected error for valid input: %v", err1)
	}
	if client1 == nil {
		t.Error("createNewVaultClient() returned a nil client for valid input")
	}

	// // Test case 2: Invalid host
	// host2 := "invalid-host"
	// token2 := "myroot"

	// client2, err2 := createNewVaultClient(host2, token2)
	// if err2 == nil {
	// 	t.Error("createNewVaultClient() should return an error for an invalid host")
	// }
	// if client2 != nil {
	// 	t.Error("createNewVaultClient() should return a nil client for an invalid host")
	// }

	// // Test case 3: Invalid token
	// host3 := "http://localhost:8200"
	// token3 := ""

	// client3, err3 := createNewVaultClient(host3, token3)
	// if err3 == nil {
	// 	t.Error("createNewVaultClient() should return an error for an empty token")
	// }
	// if client3 != nil {
	// 	t.Error("createNewVaultClient() should return a nil client for an empty token")
	// }
}