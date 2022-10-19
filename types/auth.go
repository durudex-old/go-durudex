/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package types

// Authorization tokens.
type Tokens struct {
	// JWT access token.
	Access string `json:"access"`
	// Refresh token.
	Refresh string `json:"refresh"`
}

// User Sign Up input.
type SignUpInput struct {
	// Account username.
	Username string `json:"username"`
	// User email.
	Email string `json:"email"`
	// User password.
	Password string `json:"password"`
	// User verification code.
	Code uint64 `json:"code"`
	// Client secret key.
	Secret string `json:"secret"`
}

// User Sign In input.
type SignInInput struct {
	// Account login.
	Login string `json:"login"`
	// User password
	Password string `json:"password"`
	// Client secret key.
	Secret string `json:"secret"`
}

// Session credentials input.
type SessionCredInput struct {
	// Refresh token.
	Refresh string `json:"refresh"`
	// Client secret key.
	Secret string `json:"secret"`
}
