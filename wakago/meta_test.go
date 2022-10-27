package wakago

import (
	"context"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMeta_Get(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	dummyResponse := `
	{
		"data": {
			"ip_descriptions": {
				"api": "Public IPs used by api.wakatime.com...",
				"website": "Public IPs used by wakatime.com...",
				"worker": "Public IPs used by background..."
			},
			"ips": {
				"api": [
					"192.168.1.1",
					"192.168.1.2",
					"192.168.1.3",
					"192.168.1.4"
				],
				"website": [
					"192.168.2.1",
					"192.168.2.2",
					"192.168.2.3",
					"192.168.2.4"
				],
				"worker": [
					"192.168.3.1",
					"192.168.3.2",
					"192.168.3.3",
					"192.168.3.4"
				]
			},
			"last_modified_at": "2022-10-15T14:32:59Z"
		}
	}`

	url := "https://wakatime.com/api/v1/meta"
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, dummyResponse))

	client := NewClient(nil)
	ctx := context.Background()

	res, err := client.MetaService.Get(ctx)

	if err != nil {
		t.Fatal(err)
	}

	expected := Meta{
		Data: MetaData{
			IPDescriptions: MetaIPDescriptions{
				API:     "Public IPs used by api.wakatime.com...",
				Website: "Public IPs used by wakatime.com...",
				Worker:  "Public IPs used by background...",
			},
			IPs: MetaIPs{
				API:     []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4"},
				Website: []string{"192.168.2.1", "192.168.2.2", "192.168.2.3", "192.168.2.4"},
				Worker:  []string{"192.168.3.1", "192.168.3.2", "192.168.3.3", "192.168.3.4"},
			},
			LastModifiedAt: time.Date(2022, 10, 15, 14, 32, 59, 0, time.UTC),
		},
	}

	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.EqualValues(t, &expected, res)
}
