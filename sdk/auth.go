/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package sdk

import (
	"context"

	"github.com/durudex/go-durudex/sdk/generated"
	"github.com/durudex/go-durudex/types"
)

// User Sign Up.
func (c *Client) SignUp(ctx context.Context, input types.SignUpInput) (types.Tokens, error) {
	response, err := generated.SignUp(ctx, c.client, input)
	return response.SignUp, err
}

// User Sign In.
func (c *Client) SignIn(ctx context.Context, input types.SignInInput) (types.Tokens, error) {
	response, err := generated.SignIn(ctx, c.client, input)
	return response.SignIn, err
}

// Refresh authorization token.
func (c *Client) RefreshToken(ctx context.Context, input types.SessionCredInput) (string, error) {
	response, err := generated.RefreshToken(ctx, c.client, input)
	if err != nil {
		return "", err
	}

	// Sets authorization access token.
	c.config.Transport.SetAccessToken(response.RefreshToken)

	return response.RefreshToken, nil
}
