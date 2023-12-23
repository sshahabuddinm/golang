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
