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

	"github.com/segmentio/ksuid"
)

// Getting a user session.
func (c *Client) GetSession(ctx context.Context, id ksuid.KSUID) (*types.Session, error) {
	response, err := generated.GetSession(ctx, c.client, id)
	return response.Session, err
}

// Getting a user sessions list.
func (c *Client) GetSessionList(ctx context.Context, sort types.SortOptions) (types.SessionConnection, error) {
	response, err := generated.GetSessions(ctx, c.client, sort.First, sort.Last, sort.Before, sort.After)
	return response.Sessions, err
}

// Deleting a user session.
func (c *Client) DeleteSession(ctx context.Context, id ksuid.KSUID) (bool, error) {
	response, err := generated.DeleteSession(ctx, c.client, id)
	return response.DeleteSession, err
}
