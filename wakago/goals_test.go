package wakago

import (
	"context"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGoals_Get(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	dummyResponse := `
	{
		"data": [
			{
				"average_status": "success",
				"chart_data": [
					{
						"actual_seconds": 9757.783427,
						"actual_seconds_text": "2 hrs 42 mins",
						"goal_seconds": 3600,
						"goal_seconds_text": "1 hr",
						"range": {
							"date": "2022-10-25",
							"end": "2022-10-25T14:59:59Z",
							"start": "2022-10-24T15:00:00Z",
							"text": "Tue Oct 25",
							"timezone": "Asia/Tokyo"
						},
						"range_status": "success",
						"range_status_reason": "coded 2 hrs 42 mins which is 1 hr 42 mins more than your daily goal",
						"range_status_reason_short": "2h 42m (1h 42m more than goal)"
					}
				],
				"created_at": "2022-10-31T11:53:10Z",
				"cumulative_status": "success",
				"custom_title": null,
				"delta": "day",
				"editors": [],
				"id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				"ignore_days": ["friday"],
				"ignore_zero_days": true,
				"improve_by_percent": null,
				"is_current_user_owner": true,
				"is_enabled": true,
				"is_inverse": false,
				"is_snoozed": false,
				"is_tweeting": false,
				"languages": ["languages"],
				"modified_at": null,
				"owner": {
					"display_name": "@xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
					"email": null,
					"full_name": null,
					"id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
					"photo": "https://wakatime.com/photo/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
					"username": null
				},
				"projects": ["projects"],
				"range_text": "from 2022-10-25 until 2022-10-31",
				"seconds": 3600,
				"shared_with": [],
				"snooze_until": null,
				"status": "success",
				"status_percent_calculated": 100,
				"subscribers": [
					{
						"display_name": "@xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
						"email": "test@example.com",
						"email_frequency": "Daily",
						"full_name": "full name",
						"user_id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
						"username": "user name"
					}
				],
				"title": "Code 1 hr per day",
				"type": "coding"
			}
		],
		"total": 1,
		"total_pages": 1
	}`

	url := "https://wakatime.com/api/v1/users/current/goals"
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, dummyResponse))

	client := NewClient(nil)
	res, err := client.GoalsService.Get(context.Background(), "current")

	if err != nil {
		t.Fatal(err)
	}

	expected := Goals{
		Data: []GoalsData{
			{
				AverageStatus: "success",
				ChartData: []GoalsChartData{
					{
						ActualSeconds:     9757.783427,
						ActualSecondsText: "2 hrs 42 mins",
						GoalSeconds:       3600,
						GoalSecondsText:   "1 hr",
						Range: GoalsRange{
							Date:     "2022-10-25",
							End:      time.Date(2022, 10, 25, 14, 59, 59, 0, time.UTC),
							Start:    time.Date(2022, 10, 24, 15, 0, 0, 0, time.UTC),
							Text:     "Tue Oct 25",
							Timezone: "Asia/Tokyo",
						},
						RangeStatus:            "success",
						RangeStatusReason:      "coded 2 hrs 42 mins which is 1 hr 42 mins more than your daily goal",
						RangeStatusReasonShort: "2h 42m (1h 42m more than goal)",
					},
				},
				CreatedAt:          time.Date(2022, 10, 31, 11, 53, 10, 0, time.UTC),
				CumulativeStatus:   "success",
				CustomTitle:        "",
				Delta:              "day",
				Id:                 "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				IgnoreDays:         []string{"friday"},
				IgnoreZeroDays:     true,
				ImproveByPercent:   0,
				IsCurrentUserOwner: true,
				IsEnabled:          true,
				IsInverse:          false,
				IsSnoozed:          false,
				IsTweeting:         false,
				Languages:          []string{"languages"},
				Projects:           []string{"projects"},
				RangeText:          "from 2022-10-25 until 2022-10-31",
				Seconds:            3600,
				Status:             "success",
				Subscribers: []GoalsSubscribers{
					{
						DisplayName:    "@xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
						Email:          "test@example.com",
						EmailFrequency: "Daily",
						FullName:       "full name",
						UserId:         "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
						Username:       "user name",
					},
				},
				Title: "Code 1 hr per day",
				Type:  "coding",
			},
		},
		Total:      1,
		TotalPages: 1,
	}
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.EqualValues(t, &expected, res)
}
