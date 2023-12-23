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
