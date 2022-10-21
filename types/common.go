/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package types

// Query sorting options.
type SortOptions struct {
	// First query option.
	First *int `json:"first"`
	// Last query option.
	Last *int `json:"last"`
	// Before cursor query option.
	Before *string `json:"before"`
	// After cursor query option.
	After *string `json:"after"`
}
