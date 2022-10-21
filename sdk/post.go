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

// Getting a post.
func (c *Client) GetPost(ctx context.Context, id ksuid.KSUID) (*types.Post, error) {
	response, err := generated.GetPost(ctx, c.client, id)
	return response.Post, err
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
