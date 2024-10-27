package models

import "io"

// StdArgs is a global std state model for application.
// Makes is easy to test functionalities by specifying std state.
type StdArgs struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}
