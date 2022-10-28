/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package sdk

import "log"

// Default client logger.
var DefaultLogger Logger = &DefaultClientLogger{}

// Durudex client logger interface.
type Logger interface {
	// Debug client log message.
	Debug(msg string)
	// Info client log message.
	Info(msg string)
	// Error client log message.
	Error(msg string)
	// Fatal client log message.
	Fatal(msg string)
}

// Default client logger.
type DefaultClientLogger struct{}

// Debug client log message.
func (l *DefaultClientLogger) Debug(msg string) {
	log.Println("DEBUG", msg)
}

// Info client log message.
func (l *DefaultClientLogger) Info(msg string) {
	log.Println("INFO", msg)
}

// Error client log message.
func (l *DefaultClientLogger) Error(msg string) {
	log.Panicln("ERROR", msg)
}

// Fatal client log message.
func (l *DefaultClientLogger) Fatal(msg string) {
	log.Fatal("FATAL ", msg)
}
