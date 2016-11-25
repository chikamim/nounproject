package nounproject

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// IconPayload represents the payload for icon request
type IconPayload struct {
	Icon *Icon `json:"icon"`
}

// IconsPayload represents the payload for icons request
type IconsPayload struct {
	GeneratedAt string `json:"generated_at"`
	Icons       []Icon `json:"icons"`
}

// RecentUploadsPayload represents the payload for recent uploads request
type RecentUploadsPayload struct {
	GeneratedAt string `json:"generated_at"`
	Icons       []Icon `json:"recent_uploads"`
}

// Icon represents a icon object
type Icon struct {
	Attribution           string       `json:"attribution"`
	AttributionPreviewURL string       `json:"attribution_preview_url"`
	Collections           []Collection `json:"collections"`
	DateUploaded          string       `json:"date_uploaded"`
	IconURL               string       `json:"icon_url"`
	ID                    string       `json:"id"`
	IsActive              string       `json:"is_active"`
	LicenseDescription    string       `json:"license_description"`
	Permalink             string       `json:"permalink"`
	PreviewURL            string       `json:"preview_url"`
	PreviewURL42          string       `json:"preview_url_42"`
	PreviewURL84          string       `json:"preview_url_84"`
	Sponsor               struct{}     `json:"sponsor"` // what is this?
	SponsorCampaignLink   *string      `json:"sponsor_campaign_link"`
	SponsorID             string       `json:"sponsor_id"`
	Tags                  []Tag        `json:"tags"`
	Term                  string       `json:"term"`
	TermID                int          `json:"term_id"`
	TermSlug              string       `json:"term_slug"`
	Uploader              *Uploader    `json:"uploader"`
	UploaderID            string       `json:"uploader_id"`
	Year                  int          `json:"year"`
}

// Tag represents a Tag object
type Tag struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
}

// Uploader represents a Uploader object
type Uploader struct {
	Location  string `json:"location"`
	Name      string `json:"name"`
	Permalink string `json:"permalink"`
	Username  string `json:"username"`
}

// GetIcon returns a single icon
func (c *Client) GetIcon(v string) (*Icon, error) {
	path := fmt.Sprintf("/icon/%v", v)

	payload := &IconPayload{}
	_, err := c.Get(url.URL{Path: path}, payload)
	return payload.Icon, err
}

// GetIcons returns a list of icons
func (c *Client) GetIcons(slug string, limitPublicDomain bool, pagination *Pagination) ([]Icon, error) {
	path := fmt.Sprintf("/icons/%v", slug)
	query := Query{}
	query.Merge(pagination.Query())
	if limitPublicDomain {
		query.Add("limit_to_public_domain", "1")
	}

	payload := &IconsPayload{}
	_, err := c.Get(url.URL{Path: path, RawQuery: query.String()}, payload)
	return payload.Icons, err
}

// GetRecentUploads returns a list of most recently uploaded icons
func (c *Client) GetRecentUploads(pagination *Pagination) ([]Icon, error) {
	path := "/icons/recent_uploads"
	query := Query{}
	query.Merge(pagination.Query())

	payload := &RecentUploadsPayload{}
	_, err := c.Get(url.URL{Path: path, RawQuery: query.String()}, payload)
	return payload.Icons, err
}

func (i *Icon) DownloadPreview(dirPath string) error {
	return Download(i.PreviewURL, filepath.Join(dirPath, i.ID+".png"))
}

func Download(url string, path string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer func() {
		file.Close()
	}()

	_, err = file.Write(body)
	return err
}
