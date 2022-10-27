package wakago

import (
	"context"

	"github.com/google/go-querystring/query"
)

type EditorsService service

type Editors struct {
	Data []EditorsData `json:"data"`
}

type EditorsData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	Website    string `json:"website"`
	Repository string `json:"repository"`
	Version    string `json:"version"`
	VersionURL string `json:"version_url"`
	HistoryURL string `json:"history_url"`
	Released   bool   `json:"released"`
}

type EditorsGetOptions struct {
	Unreleased bool `url:"unreleased,omitempty"`
}

func (service *EditorsService) Get(ctx context.Context, opts *EditorsGetOptions) (*Editors, error) {
	path := "editors"

	request, err := service.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	if opts != nil {
		qv, err := query.Values(opts)
		if err != nil {
			return nil, err
		}

		request.URL.RawQuery = qv.Encode()
	}

	v := new(Editors)
	_, err = service.client.Do(ctx, request, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
