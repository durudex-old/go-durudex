/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package sdk

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/durudex/go-durudex/sdk/generated"
	"github.com/durudex/go-durudex/types"

	"github.com/Khan/genqlient/graphql"
)

const (
	// Durudex test API url.
	TestAPIEndpoint string = "https://api.test.durudex.com/query"

	// Default non-auth HTTP client transport type.
	DefaultTransportType TransportType = iota
	// Auth HTTP client transport type.
	AuthTransportType

	// Authorization header key name.
	authorizationHeader string = "Authorization"

	// Bearer authorization token type.
	BearerTokenType TokenType = "Bearer"
)

// HTTP client transport type.
type TransportType int

// Authorization token type.
type TokenType string

// Getting token type in string.
func (t TokenType) String() string { return string(t) }

// Returns pointer on you value.
func Ptr[T any](v T) *T { return &v }

// Client transport interface.
type Transport interface {
	http.RoundTripper

	// Sets authorization access token.
	SetAccessToken(access string)
}

// Authorization client config.
type AuthConfig struct {
	// Refresh token.
	Refresh string
	// Client secret key.
	Secret string
	// Authorization token type.
	TokenType TokenType
	// Refresh token TTL.
	RefreshTTL time.Duration
}

// Durudex API client config.
type ClientConfig struct {
	// API endpoint.
	Endpoint string
	// HTTP client transport type.
	TransportType TransportType
	// HTTP client transport.
	Transport Transport
	// Authorization client config.
	AuthConfig *AuthConfig
	// Client logger.
	Logger Logger
}

// Durudex API client.
type Client struct {
	// GraphQL client.
	client graphql.Client
	// Client config.
	config ClientConfig
	// Client logger.
	logger Logger
	// Client wait group.
	wg sync.WaitGroup
}

// Creating a new Durudex API client.
func NewClient(cfg ClientConfig) *Client {
	var client Client

	// Checking is custom logger specified.
	if cfg.Logger == nil {
		client.logger = DefaultLogger
	} else {
		client.logger = cfg.Logger
	}

	// Creating a new client.
	httpClient := &http.Client{Transport: cfg.Transport}
	client.client = graphql.NewClient(cfg.Endpoint, httpClient)
	client.config = cfg

	// Checking is client auth transport type.
	if cfg.TransportType == AuthTransportType {
		// Checking if AuthConfig specified in configuration.
		if cfg.AuthConfig == nil {
			client.logger.Fatal("to use auth transport you need to configure AuthConfig")

			return nil
		}

		client.wg.Add(1)

		// Refreshing access token loop.
		go client.refreshTokenLoop()

		client.wg.Wait()
	}

	return &client
}

// Making a client request to GraphQL API.
func (c *Client) MakeRequest(ctx context.Context, req *graphql.Request, resp *graphql.Response) error {
	return c.client.MakeRequest(ctx, req, resp)
}

// Refreshing access token loop.
func (c *Client) refreshTokenLoop() {
	// Creating a new non-auth client connection.
	client := graphql.NewClient(c.config.Endpoint, http.DefaultClient)

	// Refreshing authorization access token.
	if err := c.refreshToken(client); err != nil {
		c.logger.Fatal("error refreshing access token: " + err.Error())
	}

	c.wg.Done()

	for {
		time.Sleep(c.config.AuthConfig.RefreshTTL)

		// Refreshing authorization access token.
		if err := c.refreshToken(client); err != nil {
			c.logger.Error("error refreshing access token: " + err.Error())
		}
	}
}

// Refreshing authorization access token.
func (c *Client) refreshToken(client graphql.Client) error {
	c.logger.Debug("refreshing access token")

	// Refreshing access token.
	response, err := generated.RefreshToken(context.Background(), client,
		types.SessionCredInput{
			Refresh: c.config.AuthConfig.Refresh,
			Secret:  c.config.AuthConfig.Secret,
		},
	)
	if err != nil {
		return err
	}

	// Sets authorization access token.
	c.config.Transport.SetAccessToken(response.RefreshToken)

	return nil
}

// Auth client transport.
type AuthTransport struct {
	// Access token.
	access string
	// Wrapped transport.
	wrapped http.RoundTripper
}

// Creating a new auth transport.
func NewAuthTransport() *AuthTransport {
	return &AuthTransport{wrapped: http.DefaultTransport}
}

// Executing auth HTTP transaction.
func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(authorizationHeader, fmt.Sprintf("%s %s", BearerTokenType, t.access))

	return t.wrapped.RoundTrip(req)
}

// Sets authorization access token.
func (t *AuthTransport) SetAccessToken(access string) { t.access = access }

// Default non-auth client transport.
type DefaultTransport struct{}

// Creating a new default non-auth client transport.
func NewDefaultTransport() *DefaultTransport { return &DefaultTransport{} }

// Executing HTTP transaction.
func (t *DefaultTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return http.DefaultTransport.RoundTrip(req)
}

// Sets authorization access token.
func (t *DefaultTransport) SetAccessToken(access string) {}
