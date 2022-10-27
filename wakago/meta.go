package wakago

import (
	"context"
	"time"
)

type MetaService service

type Meta struct {
	Data MetaData `json:"data"`
}

type MetaData struct {
	IPDescriptions MetaIPDescriptions `json:"ip_descriptions"`
	IPs            MetaIPs            `json:"ips"`
	LastModifiedAt time.Time          `json:"last_modified_at"`
}

type MetaIPs struct {
	API     []string `json:"api"`
	Website []string `json:"website"`
	Worker  []string `json:"worker"`
}

type MetaIPDescriptions struct {
	API     string `json:"api"`
	Website string `json:"website"`
	Worker  string `json:"worker"`
}

func (service *MetaService) Get(ctx context.Context) (*Meta, error) {
	path := "meta"

	request, err := service.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	v := new(Meta)
	_, err = service.client.Do(ctx, request, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
