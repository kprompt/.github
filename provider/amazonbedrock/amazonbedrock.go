// Package amazonbedrock implements an Enterprise Deferred provider for Amazon Bedrock.
// It provides compatibility path integration for Bedrock API calls.
package amazonbedrock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// Client defines the interface for Amazon Bedrock API client.
type Client interface {
	InvokeModel(ctx context.Context, modelID string, input []byte) ([]byte, error)
}

// Provider implements the Enterprise Deferred provider for Amazon Bedrock.
type Provider struct {
	client Client
}

// NewProvider returns a new Provider instance with the given Client.
func NewProvider(client Client) *Provider {
	return &Provider{client: client}
}

// ErrModelNotFound is returned when the requested model is not found.
var ErrModelNotFound = errors.New("model not found")

// CallModel invokes the specified model with the input data and returns the output.
func (p *Provider) CallModel(ctx context.Context, modelID string, input []byte) ([]byte, error) {
	if modelID == "" {
		return nil, fmt.Errorf("modelID cannot be empty")
	}
	output, err := p.client.InvokeModel(ctx, modelID, input)
	if err != nil {
		if errors.Is(err, ErrModelNotFound) {
			return nil, ErrModelNotFound
		}
		return nil, fmt.Errorf("failed to invoke model: %w", err)
	}
	return output, nil
}

// HTTPClient is a simple HTTP client for Amazon Bedrock API.
type HTTPClient struct {
	Endpoint   string
	HTTPClient *http.Client
	APIKey     string
}

// NewHTTPClient creates a new HTTPClient with the given endpoint and API key.
func NewHTTPClient(endpoint, apiKey string) *HTTPClient {
	return &HTTPClient{
		Endpoint:   endpoint,
		HTTPClient: http.DefaultClient,
		APIKey:     apiKey,
	}
}

// InvokeModel calls the Amazon Bedrock API to invoke the specified model.
func (c *HTTPClient) InvokeModel(ctx context.Context, modelID string, input []byte) ([]byte, error) {
	if modelID == "" {
		return nil, ErrModelNotFound
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.Endpoint+"/models/"+modelID+"/invoke", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Body = http.NoBody // For simplicity, no actual request body in this stub.

	// In a real implementation, input would be sent as request body.
	// Here we simulate a successful response.
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       http.NoBody,
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrModelNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Simulate output as input for testing.
	return input, nil
}
