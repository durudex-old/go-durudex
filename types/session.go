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

// User session.
type Session struct {
	// Session id.
	Id ksuid.KSUID `json:"id"`
	// Session user id.
	UserId ksuid.KSUID `json:"userId"`
	// User session ip address.
	Ip string `json:"ip"`
	// Session expires in.
	ExpiresIn time.Time `json:"expiresIn"`
}

// List of session owned by the subject.
type SessionConnection struct {
	// A list of nodes.
	Nodes []*Session `json:"nodes"`
	// A list of edges.
	Edges []*SessionEdge `json:"edges"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a session connection.
type SessionEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Session `json:"node"`
}
