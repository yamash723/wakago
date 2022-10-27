package wakago

import (
	"context"
	"fmt"
	"time"
)

type AllTimeSinceTodayService service

type AllTimeSinceToday struct {
	Data AllTimeSinceTodayData `json:"data"`
}

type AllTimeSinceTodayData struct {
	Decimal           string                 `json:"decimal"`
	Digital           string                 `json:"digital"`
	IsUpToDate        bool                   `json:"is_up_to_date"`
	PercentCalculated int                    `json:"percent_calculated"`
	Range             AllTimeSinceTodayRange `json:"range"`
	Text              string                 `json:"text"`
	Timeout           int                    `json:"timeout"`
	TotalSeconds      float32                `json:"total_seconds"`
}

type AllTimeSinceTodayRange struct {
	End       time.Time `json:"end"`
	EndDate   string    `json:"end_date"`
	EndText   string    `json:"end_text"`
	Start     time.Time `json:"start"`
	StartDate string    `json:"start_date"`
	StartText string    `json:"start_text"`
	Timezone  string    `json:"timezone"`
}

func (service *AllTimeSinceTodayService) Get(ctx context.Context, userId string, project *string) (*AllTimeSinceToday, error) {
	path := fmt.Sprintf("users/%v/all_time_since_today", userId)

	request, err := service.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	if project != nil {
		q := request.URL.Query()
		q.Add("project", *project)
		request.URL.RawQuery = q.Encode()
	}

	v := new(AllTimeSinceToday)
	_, err = service.client.Do(ctx, request, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
