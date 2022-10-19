/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package sdk

import "github.com/Khan/genqlient/graphql"

// Durudex test API url.
const TestAPIEndpoint string = "https://api.test.durudex.com/query"

// Durudex API client.
type Client struct{ client graphql.Client }

// Durudex API client config.
type ClientConfig struct {
	// API endpoint.
	Endpoint string
	// HTTP transport.
	Transport graphql.Doer
}

// Creating a new Durudex API client.
func NewClient(cfg ClientConfig) *Client {
	return &Client{client: graphql.NewClient(cfg.Endpoint, cfg.Transport)}
}
