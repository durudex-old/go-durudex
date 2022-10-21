/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package sdk

import (
	"context"

	"github.com/durudex/durudex-go/sdk/generated"
	"github.com/durudex/durudex-go/types"

	"github.com/Khan/genqlient/graphql"
	"github.com/segmentio/ksuid"
)

// User type.
type User struct {
	*types.User
	// GraphQL client.
	client graphql.Client
}

// Getting a user.
func (c *Client) GetUser(ctx context.Context, id ksuid.KSUID) (User, error) {
	response, err := generated.GetUser(ctx, c.client, id)
	response.User.Id = id
	return User{User: response.User, client: c.client}, err
}

// Getting a me.
func (c *Client) GetMe(ctx context.Context) (User, error) {
	response, err := generated.GetMe(ctx, c.client)
	return User{User: &response.Me, client: c.client}, err
}

// Forgot a user password.
func (c *Client) ForgotPassword(ctx context.Context, input types.ForgotPasswordInput) (bool, error) {
	response, err := generated.ForgotPassword(ctx, c.client, input)
	return response.ForgotPassword, err
}
