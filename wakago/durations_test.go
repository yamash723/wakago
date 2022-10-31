package wakago

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestDurations_Get(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	dummyResponse := `
	{
		"branches": [
			"main",
			"master"
		],
		"data": [
			{
				"color": null,
				"duration": 3015.151763,
				"project": "wakago",
				"time": 1666866121.027658
			}
		],
		"start": "2022-10-26T15:00:00Z",
		"end": "2022-10-27T14:59:59Z",
		"timezone": "Asia/Tokyo"
	}`

	url := "https://wakatime.com/api/v1/users/current/durations"
	project, branches, timezone, sliceBy := "project", "branches", "timezone", "sliceBy"
	timeout := 60
	writesOnly := true

	opts := DurationsGetOptions{
		Date:       "2022-10-27",
		Project:    &project,
		Branches:   &branches,
		Timeout:    &timeout,
		WritesOnly: &writesOnly,
		Timezone:   &timezone,
		SliceBy:    &sliceBy,
	}
	qv, err := query.Values(opts)
	if err != nil {
		t.Fatal(err)
	}

	expectedQuery := qv.Encode()
	httpmock.RegisterResponderWithQuery("GET", url, expectedQuery, httpmock.NewStringResponder(200, dummyResponse))

	client := NewClient(nil)
	res, err := client.DurationsService.Get(context.Background(), "current", &opts)

	if err != nil {
		t.Fatal(err)
	}

	expected := Durations{
		Data: []DurationsData{
			{
				Duration: 3015.151763,
				Project:  "wakago",
				Time:     1666866121.027658,
			},
		},
		Branches: []string{
			"main",
			"master",
		},
		Start:    time.Date(2022, 10, 26, 15, 00, 00, 0, time.UTC),
		End:      time.Date(2022, 10, 27, 14, 59, 59, 0, time.UTC),
		Timezone: "Asia/Tokyo",
	}

	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.EqualValues(t, &expected, res)
}
