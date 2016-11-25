package nounproject

import "net/url"

// Usage represents a Usage object
type Usage struct {
	Limits struct {
		Daily   *int `json:"daily"`
		Hourly  *int `json:"hourly"`
		Monthly *int `json:"monthly"`
	} `json:"limits"`
	Usage struct {
		Daily   *int `json:"daily"`
		Hourly  *int `json:"hourly"`
		Monthly *int `json:"monthly"`
	} `json:"usage"`
}

// GetUsage returns current oauth usage and limits
func (c *Client) GetUsage() (*Usage, error) {
	path := "/oauth/usage"

	payload := &Usage{}
	_, err := c.Get(url.URL{Path: path}, payload)
	return payload, err
}
