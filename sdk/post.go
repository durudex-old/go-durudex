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

// Getting a post.
func (c *Client) GetPost(ctx context.Context, id ksuid.KSUID) (*types.Post, error) {
	response, err := generated.GetPost(ctx, c.client, id)
	return response.Post, err
}

// Getting a user posts nodes.
func (c *Client) GetUserPostsNodes(ctx context.Context, uid ksuid.KSUID, sort types.SortOptions) ([]*types.Post, error) {
	response, err := generated.GetUserPostsNodes(ctx, c.client, uid, sort.First, sort.Last, sort.Before, sort.After)
	return response.User.Posts.Nodes, err
}

// Getting a user posts edges.
func (c *Client) GetUserPostsEdges(ctx context.Context, uid ksuid.KSUID, sort types.SortOptions) ([]*types.PostEdge, error) {
	response, err := generated.GetUserPostsEdges(ctx, c.client, uid, sort.First, sort.Last, sort.Before, sort.After)
	return response.User.Posts.Edges, err
}

// Getting a total user posts count.
func (c *Client) GetTotalUserPostsCount(ctx context.Context, uid ksuid.KSUID) (int, error) {
	response, err := generated.GetTotalUserPostsCount(ctx, c.client, uid)
	return response.User.Posts.TotalCount, err
}

// Creating a new post.
func (c *Client) CreatePost(ctx context.Context, input types.CreatePostInput) (ksuid.KSUID, error) {
	response, err := generated.CreatePost(ctx, c.client, input)
	return response.CreatePost, err
}

// Deleting a post.
func (c *Client) DeletePost(ctx context.Context, id ksuid.KSUID) (bool, error) {
	response, err := generated.DeletePost(ctx, c.client, id)
	return response.DeletePost, err
}

// Updating a post.
func (c *Client) UpdatePost(ctx context.Context, input types.UpdatePostInput) (bool, error) {
	response, err := generated.UpdatePost(ctx, c.client, input)
	return response.UpdatePost, err
}
