package raygun4go

import (
	"net/http"
)

// Silent sets the silent-property on the Client. If true, errors will not be
// sent to Raygun but printed instead.
func (c *Client) Silent(s bool) *Client {
	c.silent = s
	return c
}

// Request is a chainable option-setting method to add a request to the context.
func (c *Client) Request(r *http.Request) *Client {
	c.context.Request = r
	return c
}

// Version is a chainable option-setting method to add a version to the context.
func (c *Client) Version(v string) *Client {
	c.context.Version = v
	return c
}

// Tags is a chainable option-setting method to add tags to the context. You
// can use tags to filter errors in Raygun.
func (c *Client) Tags(tags []string) *Client {
	c.context.Tags = tags
	return c
}

// CustomData is a chainable option-setting method to add arbitrary custom data
// to the context. Note that the given type (or at least parts of it)
// must implement the Marshaler-interface for this to work.
func (c *Client) CustomData(data interface{}) *Client {
	c.context.CustomData = data
	return c
}

// User is a chainable option-setting method to add an affected Username to the
// context.
func (c *Client) User(u string) *Client {
	c.context.User = u
	return c
}

// Output is a chainable option-setting method to change the
// handler-function used for output
func (c *Client) Output(f OutputHandler) *Client {
	c.outputHandler = f
	return c
}
