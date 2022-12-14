package wakago

import (
	"context"
	"fmt"
	"time"
)

type GoalsService service

type Goals struct {
	Data       []GoalData `json:"data"`
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
}

type Goal struct {
	Data     GoalData  `json:"data"`
	CachedAt time.Time `json:"cached_at"`
}
type GoalRange struct {
	Date     string    `json:"date"`
	End      time.Time `json:"end"`
	Start    time.Time `json:"start"`
	Text     string    `json:"text"`
	Timezone string    `json:"timezone"`
}

type GoalChartData struct {
	ActualSeconds          float32   `json:"actual_seconds"`
	ActualSecondsText      string    `json:"actual_seconds_text"`
	GoalSeconds            int       `json:"goal_seconds"`
	GoalSecondsText        string    `json:"goal_seconds_text"`
	Range                  GoalRange `json:"range"`
	RangeStatus            string    `json:"range_status"`
	RangeStatusReason      string    `json:"range_status_reason"`
	RangeStatusReasonShort string    `json:"range_status_reason_short"`
}

type GoalSubscribers struct {
	DisplayName    string `json:"display_name"`
	Email          string `json:"email"`
	EmailFrequency string `json:"email_frequency"`
	FullName       string `json:"full_name"`
	UserId         string `json:"user_id"`
	Username       string `json:"username"`
}

type GoalData struct {
	AverageStatus      string            `json:"average_status"`
	ChartData          []GoalChartData   `json:"chart_data"`
	CreatedAt          time.Time         `json:"created_at"`
	CumulativeStatus   string            `json:"cumulative_status"`
	CustomTitle        string            `json:"custom_title"`
	Delta              string            `json:"delta"`
	Id                 string            `json:"id"`
	IgnoreDays         []string          `json:"ignore_days"`
	IgnoreZeroDays     bool              `json:"ignore_zero_days"`
	ImproveByPercent   float32           `json:"improve_by_percent"`
	IsCurrentUserOwner bool              `json:"is_current_user_owner"`
	IsEnabled          bool              `json:"is_enabled"`
	IsInverse          bool              `json:"is_inverse"`
	IsSnoozed          bool              `json:"is_snoozed"`
	IsTweeting         bool              `json:"is_tweeting"`
	Languages          []string          `json:"languages"`
	Projects           []string          `json:"projects"`
	RangeText          string            `json:"range_text"`
	Seconds            int               `json:"seconds"`
	Status             string            `json:"status"`
	Subscribers        []GoalSubscribers `json:"subscribers"`
	Title              string            `json:"title"`
	Type               string            `json:"type"`
}

func (service *GoalsService) GetAll(ctx context.Context, userId string) (*Goals, error) {
	path := fmt.Sprintf("users/%v/goals", userId)

	request, err := service.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	v := new(Goals)
	_, err = service.client.Do(ctx, request, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (service *GoalsService) Get(ctx context.Context, userId string, goalId string) (*Goal, error) {
	path := fmt.Sprintf("users/%v/goals/%v", userId, goalId)

	request, err := service.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	v := new(Goal)
	_, err = service.client.Do(ctx, request, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
