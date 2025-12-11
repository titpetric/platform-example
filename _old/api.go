package platform

import "github.com/titpetric/platform/model"

type contextKey int

const (
	sessionKey = iota
	clientKey
	invalidKey = -1
)

var (
	// Session manages a *User value stored in context.
	//
	// The value is populated based on the provided Session ID.
	Session = NewContextValue[*model.User](sessionKey)

	// Client manages a *Client value stored in context.
	//
	// The value is populated based on the provided Client ID.
	Client = NewContextValue[*model.Client](clientKey)
)
