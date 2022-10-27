package wakago

import (
	"context"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestEditors_Get(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	dummyResponse := `
	{
		"data": [
			{
				"color": "#fd27bc",
				"history_url": "https://github.com/wakatime/adobe-xd-wakatime/commits/master",
				"id": "adobe-xd",
				"name": "Adobe XD",
				"released": true,
				"repository": "https://github.com/wakatime/adobe-xd-wakatime",
				"version": "1.0.0",
				"version_url": "https://raw.githubusercontent.com/wakatime/adobe-xd-wakatime/master/manifest.json",
				"website": ""
			},
			{
				"color": "#99cd00",
				"history_url": "https://github.com/wakatime/jetbrains-wakatime/blob/master/META-INF/plugin.xml#L30",
				"id": "android-studio",
				"name": "Android Studio",
				"released": true,
				"repository": "https://github.com/wakatime/jetbrains-wakatime",
				"version": "14.1.1",
				"version_url": "",
				"website": "https://developer.android.com/sdk/index.html"
			}
		]
	}`

	url := "https://wakatime.com/api/v1/editors"
	expectedQuery := "unreleased=true"
	httpmock.RegisterResponderWithQuery("GET", url, expectedQuery, httpmock.NewStringResponder(200, dummyResponse))

	client := NewClient(nil)
	opts := EditorsGetOptions{Unreleased: true}
	res, err := client.EditorsService.Get(context.Background(), &opts)

	if err != nil {
		t.Fatal(err)
	}

	expected := Editors{
		Data: []EditorsData{
			{
				ID:         "adobe-xd",
				Name:       "Adobe XD",
				Color:      "#fd27bc",
				Website:    "",
				Repository: "https://github.com/wakatime/adobe-xd-wakatime",
				Version:    "1.0.0",
				VersionURL: "https://raw.githubusercontent.com/wakatime/adobe-xd-wakatime/master/manifest.json",
				HistoryURL: "https://github.com/wakatime/adobe-xd-wakatime/commits/master",
				Released:   true,
			},
			{
				ID:         "android-studio",
				Name:       "Android Studio",
				Color:      "#99cd00",
				Website:    "https://developer.android.com/sdk/index.html",
				Repository: "https://github.com/wakatime/jetbrains-wakatime",
				Version:    "14.1.1",
				VersionURL: "",
				HistoryURL: "https://github.com/wakatime/jetbrains-wakatime/blob/master/META-INF/plugin.xml#L30",
				Released:   true,
			},
		},
	}

	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.EqualValues(t, &expected, res)
}
