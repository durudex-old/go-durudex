/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package types

import "github.com/segmentio/ksuid"

// User type.
type User struct {
	// User id.
	Id ksuid.KSUID `json:"id"`
	// Username.
	Username string `json:"username"`
	// User verified status.
	Verified bool `json:"verified"`
	// User avatar url.
	AvatarUrl string `json:"avatarUrl"`
}
