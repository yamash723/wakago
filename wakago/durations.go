package wakago

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type DurationsService service

type Durations struct {
	Data     []DurationsData `json:"data"`
	Branches []string        `json:"branches"`
	Start    time.Time       `json:"start"`
	End      time.Time       `json:"end"`
	Timezone string          `json:"timezone"`
}

type DurationsData struct {
	Project  string  `json:"project"`
	Time     float32 `json:"time"`
	Duration float32 `json:"duration"`
}

type DurationsGetOptions struct {
	Date       string  `url:"Date,omitempty"`
	Project    *string `url:"project,omitempty"`
	Branches   *string `url:"branches,omitempty"`
	Timeout    *int    `url:"timeout,omitempty"`
	WritesOnly *bool   `url:"writes_only,omitempty"`
	Timezone   *string `url:"timezone,omitempty"`
	SliceBy    *string `url:"slice_by,omitempty"`
}

func (service *DurationsService) Get(ctx context.Context, userId string, opts *DurationsGetOptions) (*Durations, error) {
	path := fmt.Sprintf("users/%v/durations", userId)

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

	v := new(Durations)
	_, err = service.client.Do(ctx, request, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
