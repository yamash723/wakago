package wakago

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func TestCommits_GetAll(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	dummyResponse := `
	{
		"author": "current author",
		"branch": "main",
		"commits": [
			{
				"author_avatar_url": "https://avatars.githubusercontent.com/u/2596943?v=4",
				"author_date": "2022-10-27T11:46:04Z",
				"author_email": "test@example.com",
				"author_html_url": "https://github.com/yamash723",
				"author_id": "author_id",
				"author_name": "Shuhei Yamashita",
				"author_url": "https://api.github.com/users/yamash723",
				"author_username": "yamash723",
				"branch": "main",
				"committer_avatar_url": "https://avatars.githubusercontent.com/u/2596943?v=4",
				"committer_date": "2022-10-27T11:47:41Z",
				"committer_email": "test@example.com",
				"committer_html_url": "https://github.com/yamash723",
				"committer_name": "Shuhei Yamashita",
				"committer_url": "https://api.github.com/users/yamash723",
				"committer_username": "yamash723",
				"created_at": "2022-11-01T11:08:06Z",
				"hash": "633785770b4ac4e5a0acc80395bd6a015486c6a0",
				"html_url": "https://github.com/yamash723/wakago/commit/633785770b4ac4e5a0acc80395bd6a015486c6a0",
				"human_readable_date": "Oct 27, 2022",
				"human_readable_natural_date": "5 days ago",
				"human_readable_total": "27 mins",
				"human_readable_total_with_seconds": "27 mins 37 secs",
				"id": "c8a42886-6ee5-4d5e-af8a-9acb30a94ea2",
				"is_author_found": true,
				"message": "[add] Editors",
				"ref": "refs/heads/main",
				"total_seconds": 1657,
				"truncated_hash": "6337857",
				"url": "https://api.github.com/repos/yamash723/wakago/commits/633785770b4ac4e5a0acc80395bd6a015486c6a0"
			}
		],
		"next_page": 10,
		"next_page_url": "https://example.com/next",
		"page": 1,
		"prev_page": 11,
		"prev_page_url": "https://example.com/prev",
		"project": {
			"badge": "badge",
			"color": null,
			"created_at": "2022-10-22T07:20:27Z",
			"has_public_url": false,
			"human_readable_last_heartbeat_at": "Nov 1, 2022, 8:05 PM JST",
			"id": "a3e655ff-cd1f-4309-819d-b9d0cfa73abc",
			"last_heartbeat_at": "2022-11-01T11:05:17Z",
			"name": "wakago",
			"repository": {
				"badge": "badge",
				"created_at": "2022-11-01T11:07:32Z",
				"default_branch": "main",
				"description": "description",
				"fork_count": 0,
				"full_name": "yamash723/wakago",
				"homepage": "homepage",
				"html_url": "https://github.com/yamash723/wakago",
				"id": "99875969-e08a-41d6-9129-1c0907b310e8",
				"image_icon_url": "https://wakatime.com/static/img/integrations/github.png",
				"is_fork": false,
				"is_private": true,
				"last_synced_at": "2022-11-01T11:08:06Z",
				"modified_at": "2022-11-01T11:08:06Z",
				"name": "wakago",
				"provider": "github",
				"star_count": 0,
				"url": "https://api.github.com/repos/yamash723/wakago",
				"urlencoded_name": "wakago",
				"wakatime_project_name": "wakago",
				"watch_count": 0
			},
			"url": "/projects/wakago",
			"urlencoded_name": "wakago"
		},
		"status": "pending_update",
		"total": 5,
		"total_pages": 1
	}`

	url := "https://wakatime.com/api/v1/users/current/projects/project/commits"
	author, branch := "author", "branch"
	page := 3

	opts := CommitsGetOptions{
		Author: &author,
		Branch: &branch,
		Page:   &page,
	}
	qv, err := query.Values(opts)
	if err != nil {
		t.Fatal(err)
	}

	expectedQuery := qv.Encode()
	httpmock.RegisterResponderWithQuery("GET", url, expectedQuery, httpmock.NewStringResponder(200, dummyResponse))

	client := NewClient(nil)
	res, err := client.CommitsService.GetAll(context.Background(), "current", "project", &opts)

	if err != nil {
		t.Fatal(err)
	}

	expected := Commits{
		Commits: []CommitDetail{
			{
				AuthorAvatarUrl:               "https://avatars.githubusercontent.com/u/2596943?v=4",
				AuthorDate:                    time.Date(2022, 10, 27, 11, 46, 04, 0, time.UTC),
				AuthorEmail:                   "test@example.com",
				AuthorHtmlUrl:                 "https://github.com/yamash723",
				AuthorId:                      null.NewString("author_id", true),
				AuthorName:                    "Shuhei Yamashita",
				AuthorUrl:                     "https://api.github.com/users/yamash723",
				AuthorUsername:                "yamash723",
				Branch:                        "main",
				CommitterAvatarUrl:            "https://avatars.githubusercontent.com/u/2596943?v=4",
				CommitterDate:                 time.Date(2022, 10, 27, 11, 47, 41, 0, time.UTC),
				CommitterEmail:                "test@example.com",
				CommitterHtmlUrl:              "https://github.com/yamash723",
				CommitterName:                 "Shuhei Yamashita",
				CommitterUrl:                  "https://api.github.com/users/yamash723",
				CommitterUsername:             "yamash723",
				CreatedAt:                     time.Date(2022, 11, 1, 11, 8, 6, 0, time.UTC),
				Hash:                          "633785770b4ac4e5a0acc80395bd6a015486c6a0",
				HtmlUrl:                       "https://github.com/yamash723/wakago/commit/633785770b4ac4e5a0acc80395bd6a015486c6a0",
				HumanReadableDate:             "Oct 27, 2022",
				HumanReadableNaturalDate:      "5 days ago",
				HumanReadableTotal:            "27 mins",
				HumanReadableTotalWithSeconds: "27 mins 37 secs",
				Id:                            "c8a42886-6ee5-4d5e-af8a-9acb30a94ea2",
				IsAuthorFound:                 true,
				Message:                       "[add] Editors",
				Ref:                           "refs/heads/main",
				TotalSeconds:                  1657,
				TruncatedHash:                 "6337857",
				Url:                           "https://api.github.com/repos/yamash723/wakago/commits/633785770b4ac4e5a0acc80395bd6a015486c6a0",
			},
		},
		Author:      null.NewString("current author", true),
		Branch:      "main",
		Page:        1,
		NextPage:    null.NewInt(10, true),
		NextPageUrl: null.NewString("https://example.com/next", true),
		PrevPage:    null.NewInt(11, true),
		PrevPageUrl: null.NewString("https://example.com/prev", true),
		Project: CommitProject{
			Badge:                        null.NewString("badge", true),
			CreatedAt:                    time.Date(2022, 10, 22, 7, 20, 27, 0, time.UTC),
			HasPublicUrl:                 false,
			HumanReadableLastHeartbeatAt: "Nov 1, 2022, 8:05 PM JST",
			Id:                           "a3e655ff-cd1f-4309-819d-b9d0cfa73abc",
			LastHeartbeatAt:              time.Date(2022, 11, 1, 11, 5, 17, 0, time.UTC),
			Name:                         "wakago",
			Repository: CommitRepository{
				Badge:               null.NewString("badge", true),
				CreatedAt:           time.Date(2022, 11, 1, 11, 7, 32, 0, time.UTC),
				DefaultBranch:       "main",
				Description:         null.NewString("description", true),
				ForkCount:           0,
				FullName:            "yamash723/wakago",
				Homepage:            null.NewString("homepage", true),
				HtmlUrl:             "https://github.com/yamash723/wakago",
				Id:                  "99875969-e08a-41d6-9129-1c0907b310e8",
				ImageIconUrl:        "https://wakatime.com/static/img/integrations/github.png",
				IsFork:              false,
				IsPrivate:           true,
				LastSyncedAt:        null.NewTime(time.Date(2022, 11, 1, 11, 8, 6, 0, time.UTC), true),
				ModifiedAt:          null.NewTime(time.Date(2022, 11, 1, 11, 8, 6, 0, time.UTC), true),
				Name:                "wakago",
				Provider:            "github",
				StarCount:           0,
				Url:                 "https://api.github.com/repos/yamash723/wakago",
				UrlencodedName:      "wakago",
				WakatimeProjectName: "wakago",
				WatchCount:          0,
			},
			Url:            "/projects/wakago",
			UrlencodedName: "wakago",
		},
		Status:     "pending_update",
		Total:      5,
		TotalPages: 1,
	}

	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.EqualValues(t, &expected, res)
}

func TestCommits_Get(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	dummyResponse := `
	{
		"branch": "main",
		"commit": {
			"author_avatar_url": "https://avatars.githubusercontent.com/u/2596943?v=4",
			"author_date": "2022-10-27T11:46:04Z",
			"author_email": "test@example.com",
			"author_html_url": "https://github.com/yamash723",
			"author_id": "author_id",
			"author_name": "Shuhei Yamashita",
			"author_url": "https://api.github.com/users/yamash723",
			"author_username": "yamash723",
			"branch": "main",
			"committer_avatar_url": "https://avatars.githubusercontent.com/u/2596943?v=4",
			"committer_date": "2022-10-27T11:47:41Z",
			"committer_email": "test@example.com",
			"committer_html_url": "https://github.com/yamash723",
			"committer_name": "Shuhei Yamashita",
			"committer_url": "https://api.github.com/users/yamash723",
			"committer_username": "yamash723",
			"created_at": "2022-11-01T11:08:06Z",
			"hash": "633785770b4ac4e5a0acc80395bd6a015486c6a0",
			"html_url": "https://github.com/yamash723/wakago/commit/633785770b4ac4e5a0acc80395bd6a015486c6a0",
			"human_readable_date": "Oct 27, 2022",
			"human_readable_natural_date": "5 days ago",
			"human_readable_total": "27 mins",
			"human_readable_total_with_seconds": "27 mins 37 secs",
			"id": "c8a42886-6ee5-4d5e-af8a-9acb30a94ea2",
			"is_author_found": true,
			"message": "[add] Editors",
			"ref": "refs/heads/main",
			"total_seconds": 1657,
			"truncated_hash": "6337857",
			"url": "https://api.github.com/repos/yamash723/wakago/commits/633785770b4ac4e5a0acc80395bd6a015486c6a0"
		},
		"project": {
			"badge": "badge",
			"color": null,
			"created_at": "2022-10-22T07:20:27Z",
			"has_public_url": false,
			"human_readable_last_heartbeat_at": "Nov 1, 2022, 8:05 PM JST",
			"id": "a3e655ff-cd1f-4309-819d-b9d0cfa73abc",
			"last_heartbeat_at": "2022-11-01T11:05:17Z",
			"name": "wakago",
			"repository": {
				"badge": "badge",
				"created_at": "2022-11-01T11:07:32Z",
				"default_branch": "main",
				"description": "description",
				"fork_count": 0,
				"full_name": "yamash723/wakago",
				"homepage": "homepage",
				"html_url": "https://github.com/yamash723/wakago",
				"id": "99875969-e08a-41d6-9129-1c0907b310e8",
				"image_icon_url": "https://wakatime.com/static/img/integrations/github.png",
				"is_fork": false,
				"is_private": true,
				"last_synced_at": "2022-11-01T11:08:06Z",
				"modified_at": "2022-11-01T11:08:06Z",
				"name": "wakago",
				"provider": "github",
				"star_count": 0,
				"url": "https://api.github.com/repos/yamash723/wakago",
				"urlencoded_name": "wakago",
				"wakatime_project_name": "wakago",
				"watch_count": 0
			},
			"url": "/projects/wakago",
			"urlencoded_name": "wakago"
		},
		"status": "pending_update"
	}`

	hash := "633785770b4ac4e5a0acc80395bd6a015486c6a0"
	url := "https://wakatime.com/api/v1/users/current/projects/project/commits/" + hash
	branch := "branch"

	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, dummyResponse))

	client := NewClient(nil)
	res, err := client.CommitsService.Get(context.Background(), "current", "project", hash, &branch)

	if err != nil {
		t.Fatal(err)
	}

	expected := Commit{
		Branch: "main",
		Commit: CommitDetail{
			AuthorAvatarUrl:               "https://avatars.githubusercontent.com/u/2596943?v=4",
			AuthorDate:                    time.Date(2022, 10, 27, 11, 46, 04, 0, time.UTC),
			AuthorEmail:                   "test@example.com",
			AuthorHtmlUrl:                 "https://github.com/yamash723",
			AuthorId:                      null.NewString("author_id", true),
			AuthorName:                    "Shuhei Yamashita",
			AuthorUrl:                     "https://api.github.com/users/yamash723",
			AuthorUsername:                "yamash723",
			Branch:                        "main",
			CommitterAvatarUrl:            "https://avatars.githubusercontent.com/u/2596943?v=4",
			CommitterDate:                 time.Date(2022, 10, 27, 11, 47, 41, 0, time.UTC),
			CommitterEmail:                "test@example.com",
			CommitterHtmlUrl:              "https://github.com/yamash723",
			CommitterName:                 "Shuhei Yamashita",
			CommitterUrl:                  "https://api.github.com/users/yamash723",
			CommitterUsername:             "yamash723",
			CreatedAt:                     time.Date(2022, 11, 1, 11, 8, 6, 0, time.UTC),
			Hash:                          "633785770b4ac4e5a0acc80395bd6a015486c6a0",
			HtmlUrl:                       "https://github.com/yamash723/wakago/commit/633785770b4ac4e5a0acc80395bd6a015486c6a0",
			HumanReadableDate:             "Oct 27, 2022",
			HumanReadableNaturalDate:      "5 days ago",
			HumanReadableTotal:            "27 mins",
			HumanReadableTotalWithSeconds: "27 mins 37 secs",
			Id:                            "c8a42886-6ee5-4d5e-af8a-9acb30a94ea2",
			IsAuthorFound:                 true,
			Message:                       "[add] Editors",
			Ref:                           "refs/heads/main",
			TotalSeconds:                  1657,
			TruncatedHash:                 "6337857",
			Url:                           "https://api.github.com/repos/yamash723/wakago/commits/633785770b4ac4e5a0acc80395bd6a015486c6a0",
		},
		Project: CommitProject{
			Badge:                        null.NewString("badge", true),
			CreatedAt:                    time.Date(2022, 10, 22, 7, 20, 27, 0, time.UTC),
			HasPublicUrl:                 false,
			HumanReadableLastHeartbeatAt: "Nov 1, 2022, 8:05 PM JST",
			Id:                           "a3e655ff-cd1f-4309-819d-b9d0cfa73abc",
			LastHeartbeatAt:              time.Date(2022, 11, 1, 11, 5, 17, 0, time.UTC),
			Name:                         "wakago",
			Repository: CommitRepository{
				Badge:               null.NewString("badge", true),
				CreatedAt:           time.Date(2022, 11, 1, 11, 7, 32, 0, time.UTC),
				DefaultBranch:       "main",
				Description:         null.NewString("description", true),
				ForkCount:           0,
				FullName:            "yamash723/wakago",
				Homepage:            null.NewString("homepage", true),
				HtmlUrl:             "https://github.com/yamash723/wakago",
				Id:                  "99875969-e08a-41d6-9129-1c0907b310e8",
				ImageIconUrl:        "https://wakatime.com/static/img/integrations/github.png",
				IsFork:              false,
				IsPrivate:           true,
				LastSyncedAt:        null.NewTime(time.Date(2022, 11, 1, 11, 8, 6, 0, time.UTC), true),
				ModifiedAt:          null.NewTime(time.Date(2022, 11, 1, 11, 8, 6, 0, time.UTC), true),
				Name:                "wakago",
				Provider:            "github",
				StarCount:           0,
				Url:                 "https://api.github.com/repos/yamash723/wakago",
				UrlencodedName:      "wakago",
				WakatimeProjectName: "wakago",
				WatchCount:          0,
			},
			Url:            "/projects/wakago",
			UrlencodedName: "wakago",
		},
		Status: "pending_update",
	}

	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.EqualValues(t, &expected, res)
}
