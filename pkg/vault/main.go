package main

import (
	"fmt"
	"strings"
	"context"
	"github.com/hashicorp/vault/api"
)

// Client for accessing itsm service apis
type Client struct {
	url string
	token *string
	vaultClient *api.Client
}

type Config struct {
	Ctx context.Context
	Token *string
	Host *string
}

type ClientOption func(*Client)

func WithToken(token string) ClientOption {
	return func(c *Client) {
		c.token = &token
	}
}

func createNewVaultClient(host, token string) (*api.Client, error) {
	vaultConfig := &api.Config{
		Address: host,
	}
	vaultClient, err := api.NewClient(vaultConfig)
	if err != nil {
		return nil, err
	}
	return vaultClient, nil
}

// Policy represents a vault policy
type VaultPolicy struct {
	Name         string
	Path         string
	Capabilities []string
}

// NewPolicy creates a new Policy instance
func NewPolicy(name, path string, capabilities ...string) *VaultPolicy {
	return &VaultPolicy{
		Name:         name,
		Path:         path,
		Capabilities: capabilities,
	}
}

// CreatePolicy generates a vault policy string
func (p *VaultPolicy) CreatePolicyString() string {
	return fmt.Sprintf(`path "%s" {
  capabilities = %s
}`, p.Path, formatCapabilities(p.Capabilities))
}

// CreateVaultPolicy for creating a policy in the vault
func (c *Client) CreateVaultPolicy(name, rules string) error {
	vaultClient, err := createNewVaultClient(c.url, *c.token)
	if err != nil {
		return err
	}
	err = vaultClient.Sys().PutPolicy(name, strings.TrimSpace(rules))
	if err != nil {
		return err
	}
	fmt.Printf("Policy %s created successfully.\n", name)
	return nil
}

// formatCapabilities formats capability slice as a string
func formatCapabilities(capabilities []string) string {
	if len(capabilities) == 0 {
		return "[]"
	}

	// Add quotes around each capability and join them into a string
	quotedCapabilities := make([]string, len(capabilities))
	for i, cap := range capabilities {
		quotedCapabilities[i] = fmt.Sprintf(`"%s"`, cap)
	}

	return fmt.Sprintf("[%s]", strings.Join(quotedCapabilities, ", "))
}

func main() {
	// Create a new Policy instance
	pol := NewPolicy("mypolicy", "/secrets/data", "create", "read", "write")
	fmt.Println(pol)

	// Create and print the policy
	// fmt.Println(pol.CreateVaultPolicy())
}
