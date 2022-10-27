package wakago

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTimeSinceToday(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	dummyResponse := `
	{
		"data": {
			"decimal": "12.80",
			"digital": "12:48",
			"is_up_to_date": true,
			"percent_calculated": 100,
			"range": {
				"end": "2022-10-23T14:59:59Z",
				"end_date": "2022-10-23",
				"end_text": "Today",
				"start": "2022-10-12T15:00:00Z",
				"start_date": "2022-10-13",
				"start_text": "Thu Oct 13th 2022",
				"timezone": "Asia/Tokyo"
			},
			"text": "15 hrs 30 mins",
			"timeout": 15,
			"total_seconds": 55801.533187
		}
	}`

	project := "testProject"
	expectedQuery := fmt.Sprintf("project=%v", project)
	url := "https://wakatime.com/api/v1/users/current/all_time_since_today"
	httpmock.RegisterResponderWithQuery("GET", url, expectedQuery, httpmock.NewStringResponder(200, dummyResponse))

	client := NewClient(nil)
	ctx := context.Background()

	res, err := client.AllTimeSinceTodayService.Get(ctx, "current", &project)

	if err != nil {
		t.Fatal(err)
	}

	expected := AllTimeSinceToday{
		Data: Data{
			Decimal:           "12.80",
			Digital:           "12:48",
			IsUpToDate:        true,
			PercentCalculated: 100,
			Range: Range{
				End:       time.Date(2022, 10, 23, 14, 59, 59, 0, time.UTC),
				EndDate:   "2022-10-23",
				EndText:   "Today",
				Start:     time.Date(2022, 10, 12, 15, 00, 00, 0, time.UTC),
				StartDate: "2022-10-13",
				StartText: "Thu Oct 13th 2022",
				Timezone:  "Asia/Tokyo",
			},
			Text:         "15 hrs 30 mins",
			Timeout:      15,
			TotalSeconds: 55801.533187,
		},
	}

	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.EqualValues(t, &expected, res)
}
