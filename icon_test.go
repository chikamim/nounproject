package nounproject

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient("key", "secret")
	client.BaseURL, _ = url.Parse(server.URL)
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func TestGetIcon(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/icon/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
  "icon": {
    "attribution": "Trash from Noun Project",
    "collections": [
      {
        "author": {
          "location": "Los Angeles, US",
          "name": "Edward Boatman",
          "permalink": "/edward",
          "username": "edward"
        },
        "author_id": "6",
        "date_created": "2012-01-27 19:15:26",
        "date_updated": "2012-09-27 13:27:02",
        "description": "",
        "id": "3",
        "is_collaborative": "",
        "is_featured": "1",
        "is_published": "1",
        "is_store_item": "0",
        "name": "AIGA",
        "permalink": "/edward/collection/aiga",
        "slug": "aiga",
        "sponsor": {},
        "sponsor_campaign_link": "",
        "sponsor_id": "",
        "tags": [],
        "template": "24"
      }
    ],
    "date_uploaded": "",
    "icon_url": "https://d30y9cdsu7xlg0.cloudfront.net/noun-svg/1.svg?Expires=1464172672&Signature=OquvjmhM3QgvoEFNC5lTgl0WSYQFLdJZCj8BdCUgCzh3hiHMmM58TSdluNCvny4dWyMlNbLGthgE0DpEwP4rTBCtrSHX~6oBJIVjZVJ7UknCjzlN4K-D5~6GnC82JGv6B8ZkHHxD6jhgkGGf9qDYuaNZXms7MkOl5g24H3w16qg_&Key-Pair-Id=APKAI5ZVHAXN65CHVU2Q",
    "id": "1",
    "is_active": "1",
    "license_description": "public-domain",
    "permalink": "/term/trash/1",
    "preview_url": "https://d30y9cdsu7xlg0.cloudfront.net/png/1-200.png",
    "preview_url_42": "https://d30y9cdsu7xlg0.cloudfront.net/png/1-42.png",
    "preview_url_84": "https://d30y9cdsu7xlg0.cloudfront.net/png/1-84.png",
    "sponsor": {},
    "sponsor_campaign_link": null,
    "sponsor_id": "",
    "tags": [
      {
        "id": 19,
        "slug": "trash"
      },
      {
        "id": 20,
        "slug": "garbage"
      },
      {
        "id": 21,
        "slug": "litter"
      }
    ],
    "term": "Trash",
    "term_id": 19,
    "term_slug": "trash",
    "uploader": {},
    "uploader_id": "",
    "year": 1974
  }
}`)
	})

	got, err := client.GetIcon("1")
	if err != nil {
		t.Errorf("GetIcon returned error: %v", err)
	}

	want := &Icon{
		Attribution:           "Trash from Noun Project",
		AttributionPreviewURL: "",
		Collections: []Collection{
			Collection{
				Author: struct {
					Location  string "json:\"location\""
					Name      string "json:\"name\""
					Permalink string "json:\"permalink\""
					Username  string "json:\"username\""
				}{
					Location:  "Los Angeles, US",
					Name:      "Edward Boatman",
					Permalink: "/edward",
					Username:  "edward",
				},
				AuthorID:            "6",
				DateCreated:         "2012-01-27 19:15:26",
				DateUpdated:         "2012-09-27 13:27:02",
				Description:         "",
				ID:                  "3",
				IsCollaborative:     "",
				IsFeatured:          "1",
				IsPublished:         "1",
				IsStoreItem:         "0",
				Name:                "AIGA",
				Permalink:           "/edward/collection/aiga",
				Slug:                "aiga",
				Sponsor:             struct{}{},
				SponsorCampaignLink: "",
				SponsorID:           "",
				Tags:                []Tag{},
				Template:            "24",
			},
		},
		DateUploaded:        "",
		IconURL:             "https://d30y9cdsu7xlg0.cloudfront.net/noun-svg/1.svg?Expires=1464172672&Signature=OquvjmhM3QgvoEFNC5lTgl0WSYQFLdJZCj8BdCUgCzh3hiHMmM58TSdluNCvny4dWyMlNbLGthgE0DpEwP4rTBCtrSHX~6oBJIVjZVJ7UknCjzlN4K-D5~6GnC82JGv6B8ZkHHxD6jhgkGGf9qDYuaNZXms7MkOl5g24H3w16qg_&Key-Pair-Id=APKAI5ZVHAXN65CHVU2Q",
		ID:                  "1",
		IsActive:            "1",
		LicenseDescription:  "public-domain",
		Permalink:           "/term/trash/1",
		PreviewURL:          "https://d30y9cdsu7xlg0.cloudfront.net/png/1-200.png",
		PreviewURL42:        "https://d30y9cdsu7xlg0.cloudfront.net/png/1-42.png",
		PreviewURL84:        "https://d30y9cdsu7xlg0.cloudfront.net/png/1-84.png",
		Sponsor:             struct{}{},
		SponsorCampaignLink: (*string)(nil),
		SponsorID:           "",
		Tags: []Tag{
			Tag{
				ID:   19,
				Slug: "trash",
			},
			Tag{
				ID:   20,
				Slug: "garbage",
			},
			Tag{
				ID:   21,
				Slug: "litter",
			},
		},
		Term:     "Trash",
		TermID:   19,
		TermSlug: "trash",
		Uploader: &Uploader{
			Location:  "",
			Name:      "",
			Permalink: "",
			Username:  "",
		},
		UploaderID: "",
		Year:       1974,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetIcon returned %+v, want %+v", got, want)
	}
}
