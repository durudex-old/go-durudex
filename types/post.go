/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package types

import (
	"time"

	"github.com/segmentio/ksuid"
)

// Post type.
type Post struct {
	// Post id.
	Id ksuid.KSUID `json:"id"`
	// Post author.
	Author *User `json:"author"`
	// Post text.
	Text string `json:"text"`
	// Post updated date.
	UpdatedAt *time.Time `json:"updatedAt"`
	// Post attachments.
	Attachments []string `json:"attachments"`
}

// Create post input.
type CreatePostInput struct {
	// Post text.
	Text string `json:"text"`
	// Post attachments.
	Attachments []*UploadFile `json:"attachments"`
}

// Update post input.
type UpdatePostInput struct {
	// Post id.
	Id ksuid.KSUID `json:"id"`
	// Post text.
	Text string `json:"text"`
}
