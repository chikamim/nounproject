package nounproject

import (
	"fmt"
	"net/url"
)

// CollectionPayload represents the payload for collection
type CollectionPayload struct {
	Collection *Collection `json:"collection"`
}

// CollectionsPayload represents the payload for collections
type CollectionsPayload struct {
	Collections []Collection `json:"collections"`
}

// Collection represents a Collection object
type Collection struct {
	Author struct {
		Location  string `json:"location"`
		Name      string `json:"name"`
		Permalink string `json:"permalink"`
		Username  string `json:"username"`
	} `json:"author"`
	AuthorID            string   `json:"author_id"`
	DateCreated         string   `json:"date_created"`
	DateUpdated         string   `json:"date_updated"`
	Description         string   `json:"description"`
	ID                  string   `json:"id"`
	IsCollaborative     string   `json:"is_collaborative"`
	IsFeatured          string   `json:"is_featured"`
	IsPublished         string   `json:"is_published"`
	IsStoreItem         string   `json:"is_store_item"`
	Name                string   `json:"name"`
	Permalink           string   `json:"permalink"`
	Slug                string   `json:"slug"`
	Sponsor             struct{} `json:"sponsor"`
	SponsorCampaignLink string   `json:"sponsor_campaign_link"`
	SponsorID           string   `json:"sponsor_id"`
	Tags                []Tag    `json:"tags"`
	Template            string   `json:"template"`
}

// GetCollection returns a single collection
func (c *Client) GetCollection(slug string) (*Collection, error) {
	path := fmt.Sprintf("/collection/%v", slug)

	payload := &CollectionPayload{}
	_, err := c.Get(url.URL{Path: path}, payload)
	return payload.Collection, err
}

// GetCollectionIcons returns a list of icons associated with a collection
func (c *Client) GetCollectionIcons(slug string, pagination *Pagination) ([]Icon, error) {
	path := fmt.Sprintf("/collection/%v/icons", slug)
	query := Query{}
	query.Merge(pagination.Query())

	payload := &IconsPayload{}
	_, err := c.Get(url.URL{Path: path, RawQuery: query.String()}, payload)
	return payload.Icons, err
}

// GetCollections returns a list of all collections
func (c *Client) GetCollections(pagination *Pagination) ([]Collection, error) {
	path := "/collections"
	query := Query{}
	query.Merge(pagination.Query())

	payload := &CollectionsPayload{}
	_, err := c.Get(url.URL{Path: path, RawQuery: query.String()}, payload)
	return payload.Collections, err
}
