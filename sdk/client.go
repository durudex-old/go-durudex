/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package sdk

import (
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

// Testing API url.
const TestAPI = "https://api.test.durudex.com/query"

// Durudex API client.
type Client struct{ client graphql.Client }

// Creating a new Durudex API client.
func NewClient() *Client {
	return &Client{client: graphql.NewClient(TestAPI, http.DefaultClient)}
}
