package internal

import (
	"testing"
)

func TestExtractID(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected string
	}{
		"itunesURL":           {"https://music.apple.com/jp/album/chandler-bing-single/1683089826", "1683089826"},
		"itunesURLWithQuery":  {"https://music.apple.com/us/album/mint-land/1611115294?i=1611115480", "1611115480"},
		"spotifyURL":          {"https://open.spotify.com/album/47fyjjJdkUJrpShiPEjwI7", "47fyjjJdkUJrpShiPEjwI7"},
		"spotifyURLWithQuery": {"https://open.spotify.com/track/3OMbfxFSrlAUBcZ3IRqeqz?si=296f572149844a31", "296f572149844a31"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			u := NewURLElements(tt.input)
			actual := u.extractID()

			if actual != tt.expected {
				t.Errorf("extractID() = %v; want %v,", actual, tt.expected)
			}
		})
	}
}
