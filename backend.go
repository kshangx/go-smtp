package smtp

import (
	"io"
)

// A SMTP message.
type Message struct {
	// The message contents.
	io.Reader

	// The sender e-mail address.
	From string
	// The recipients e-mail addresses.
	To []string
}

// A SMTP server backend.
type Backend interface {
	// Authenticate a user.
	Login(username, password string) (User, error)
}

// An authenticated user.
type User interface {
	// Send an e-mail.
	Send(msg *Message) error
	// Logout is called when this User will no longer be used.
	Logout() error
}
