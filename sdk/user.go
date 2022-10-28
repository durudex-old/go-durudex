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

	"github.com/Khan/genqlient/graphql"
	"github.com/segmentio/ksuid"
)

// User type.
type User struct {
	*types.User
	// GraphQL client.
	client graphql.Client
}

// Getting a user posts nodes.
func (u User) GetPostsNodes(ctx context.Context, sort types.SortOptions) ([]*types.Post, error) {
	response, err := generated.GetUserPostsNodes(ctx, u.client, u.Id, sort.First, sort.Last, sort.Before, sort.After)
	return response.User.Posts.Nodes, err
}

// Getting a user posts edges.
func (u User) GetPostsEdges(ctx context.Context, sort types.SortOptions) ([]*types.PostEdge, error) {
	response, err := generated.GetUserPostsEdges(ctx, u.client, u.Id, sort.First, sort.Last, sort.Before, sort.After)
	return response.User.Posts.Edges, err
}

// Getting a total user posts count.
func (u User) GetTotalPostsCount(ctx context.Context) (int, error) {
	response, err := generated.GetTotalUserPostsCount(ctx, u.client, u.Id)
	return response.User.Posts.TotalCount, err
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
