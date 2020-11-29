package api

// Error defines a common structure for information to be stored in when errors are thrown
type Error struct {
	Message    string // Message contains information about the error in a string format
	StatusCode int    // StatusCode is the associated error code attached to Message
}

// Error is implementation of the Error interface's Error() method for returning the error message as a string
func (e Error) Error() string {
	return e.Message
}
