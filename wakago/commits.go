package wakago

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
	"gopkg.in/guregu/null.v4"
)

type CommitsService service

type Commits struct {
	Commits     []CommitDetail `json:"commits"`
	Author      null.String    `json:"author"`
	Branch      string         `json:"branch"`
	NextPage    null.Int       `json:"next_page"`
	NextPageUrl null.String    `json:"next_page_url"`
	Page        int            `json:"page"`
	PrevPage    null.Int       `json:"prev_page"`
	PrevPageUrl null.String    `json:"prev_page_url"`
	Project     CommitProject  `json:"project"`
	Status      string         `json:"status"`
	Total       int            `json:"total"`
	TotalPages  int            `json:"total_pages"`
}

type Commit struct {
	Branch  string        `json:"branch"`
	Commit  CommitDetail  `json:"commit"`
	Project CommitProject `json:"project"`
	Status  string        `json:"status"`
}

type CommitDetail struct {
	AuthorAvatarUrl               string      `json:"author_avatar_url"`
	AuthorDate                    time.Time   `json:"author_date"`
	AuthorEmail                   string      `json:"author_email"`
	AuthorHtmlUrl                 string      `json:"author_html_url"`
	AuthorId                      null.String `json:"author_id"`
	AuthorName                    string      `json:"author_name"`
	AuthorUrl                     string      `json:"author_url"`
	AuthorUsername                string      `json:"author_username"`
	Branch                        string      `json:"branch"`
	CommitterAvatarUrl            string      `json:"committer_avatar_url"`
	CommitterDate                 time.Time   `json:"committer_date"`
	CommitterEmail                string      `json:"committer_email"`
	CommitterHtmlUrl              string      `json:"committer_html_url"`
	CommitterName                 string      `json:"committer_name"`
	CommitterUrl                  string      `json:"committer_url"`
	CommitterUsername             string      `json:"committer_username"`
	CreatedAt                     time.Time   `json:"created_at"`
	Hash                          string      `json:"hash"`
	HtmlUrl                       string      `json:"html_url"`
	HumanReadableDate             string      `json:"human_readable_date"`
	HumanReadableNaturalDate      string      `json:"human_readable_natural_date"`
	HumanReadableTotal            string      `json:"human_readable_total"`
	HumanReadableTotalWithSeconds string      `json:"human_readable_total_with_seconds"`
	Id                            string      `json:"id"`
	IsAuthorFound                 bool        `json:"is_author_found"`
	Message                       string      `json:"message"`
	Ref                           string      `json:"ref"`
	TotalSeconds                  float32     `json:"total_seconds"`
	TruncatedHash                 string      `json:"truncated_hash"`
	Url                           string      `json:"url"`
}

type CommitProject struct {
	Badge                        null.String      `json:"badge"`
	CreatedAt                    time.Time        `json:"created_at"`
	HasPublicUrl                 bool             `json:"has_public_url"`
	HumanReadableLastHeartbeatAt string           `json:"human_readable_last_heartbeat_at"`
	Id                           string           `json:"id"`
	LastHeartbeatAt              time.Time        `json:"last_heartbeat_at"`
	Name                         string           `json:"name"`
	Repository                   CommitRepository `json:"repository"`
	Url                          string           `json:"url"`
	UrlencodedName               string           `json:"urlencoded_name"`
}

type CommitRepository struct {
	Badge               null.String `json:"badge"`
	CreatedAt           time.Time   `json:"created_at"`
	DefaultBranch       string      `json:"default_branch"`
	Description         null.String `json:"description"`
	ForkCount           int         `json:"fork_count"`
	FullName            string      `json:"full_name"`
	Homepage            null.String `json:"homepage"`
	HtmlUrl             string      `json:"html_url"`
	Id                  string      `json:"id"`
	ImageIconUrl        string      `json:"image_icon_url"`
	IsFork              bool        `json:"is_fork"`
	IsPrivate           bool        `json:"is_private"`
	LastSyncedAt        null.Time   `json:"last_synced_at"`
	ModifiedAt          null.Time   `json:"modified_at"`
	Name                string      `json:"name"`
	Provider            string      `json:"provider"`
	StarCount           int         `json:"star_count"`
	Url                 string      `json:"url"`
	UrlencodedName      string      `json:"urlencoded_name"`
	WakatimeProjectName string      `json:"wakatime_project_name"`
	WatchCount          int         `json:"watch_count"`
}

type CommitsGetOptions struct {
	Author *string `url:"author,omitempty"`
	Branch *string `url:"branch,omitempty"`
	Page   *int    `url:"page,omitempty"`
}

func (service *CommitsService) GetAll(ctx context.Context, userId string, project string, opts *CommitsGetOptions) (*Commits, error) {
	path := fmt.Sprintf("users/%v/projects/%v/commits", userId, project)

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

	v := new(Commits)
	_, err = service.client.Do(ctx, request, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (service *CommitsService) Get(ctx context.Context, userId string, project string, hash string, branch *string) (*Commit, error) {
	path := fmt.Sprintf("users/%v/projects/%v/commits/%v", userId, project, hash)

	request, err := service.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	v := new(Commit)
	_, err = service.client.Do(ctx, request, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
