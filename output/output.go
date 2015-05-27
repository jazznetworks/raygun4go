// Package output contains handler functions used to treat textual output. This
// is very basic at the moment, but may in the future contain functions for
// writing to logfiles, mails, custom alerts...
package output

import "log"

// Handler is the type of function for handling output. It receives a string,
// handles it and returns nil if ok, an error if something went wrong.
type Handler func(string) error

// Log is a Handler that logs the given message
var Log Handler = func(message string) error {
	log.Println(message)
	return nil
}

// Mute is a Handler that does nothing with the given message
var Mute Handler = func(message string) error {
	return nil
}
