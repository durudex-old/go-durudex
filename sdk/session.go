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

	"github.com/segmentio/ksuid"
)

// Getting a user session.
func (c *Client) GetSession(ctx context.Context, id ksuid.KSUID) (*types.Session, error) {
	response, err := generated.GetSession(ctx, c.client, id)
	return response.Session, err
}

// Getting a user sessions nodes.
func (c *Client) GetSessionsNodes(ctx context.Context, sort types.SortOptions) ([]*types.Session, error) {
	response, err := generated.GetSessionsNodes(ctx, c.client, sort.First, sort.Last, sort.Before, sort.After)
	return response.Sessions.Nodes, err
}

// Getting a user sessions edges.
func (c *Client) GetSessionsEdges(ctx context.Context, sort types.SortOptions) ([]*types.SessionEdge, error) {
	response, err := generated.GetSessionsEdges(ctx, c.client, sort.First, sort.Last, sort.Before, sort.After)
	return response.Sessions.Edges, err
}

// Getting a total user sessions count.
func (c *Client) GetTotalSessionsCount(ctx context.Context) (int, error) {
	response, err := generated.GetTotalSessionsCount(ctx, c.client)
	return response.Sessions.TotalCount, err
}

// Deleting a user session.
func (c *Client) DeleteSession(ctx context.Context, id ksuid.KSUID) (bool, error) {
	response, err := generated.DeleteSession(ctx, c.client, id)
	return response.DeleteSession, err
}
