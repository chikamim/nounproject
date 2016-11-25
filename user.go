package nounproject

import (
	"fmt"
	"net/url"
)

// UserUploadsPayload represents the payload for recent user uploads request
type UserUploadsPayload struct {
	GeneratedAt string       `json:"generated_at"`
	Collections []Collection `json:"uploads"`
}

// GetUserCollections returns a single collection associated with a user
func (c *Client) GetUserCollections(userID int, slug string) ([]Collection, error) {
	path := fmt.Sprintf("/user/%v/collections", userID)
	if len(slug) > 0 {
		path = path + "/" + slug
	}

	payload := &CollectionsPayload{}
	_, err := c.Get(url.URL{Path: path}, payload)
	return payload.Collections, err
}

// GetUserUploads returns a list of collections associated with a user
func (c *Client) GetUserUploads(userName string, pagination *Pagination) ([]Collection, error) {
	path := fmt.Sprintf("/user/%v/uploads", userName)
	query := Query{}
	query.Merge(pagination.Query())

	payload := &UserUploadsPayload{}
	_, err := c.Get(url.URL{Path: path, RawQuery: query.String()}, payload)
	return payload.Collections, err
}
