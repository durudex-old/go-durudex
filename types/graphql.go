/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package types

import "github.com/99designs/gqlgen/graphql"

// Upload files input.
type UploadFile struct {
	// File id.
	Id int `json:"id"`
	// File data.
	File graphql.Upload `json:"file"`
}
