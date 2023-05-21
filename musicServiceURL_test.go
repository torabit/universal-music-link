package internal

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNewMusicServiceURL(t *testing.T) {
	u, error := NewMusicServiceURL("https://www.youtube.com/")
	blankMusicServiceURL := &MusicServiceURL{}

	if error == nil {
		t.Fatal("expected an error")
	}

	if reflect.DeepEqual(u, *blankMusicServiceURL) {
		t.Fatalf("expected: %v; got: %v", *blankMusicServiceURL, u)
	}
}

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
			u, _ := NewMusicServiceURL(tt.input)
			actual := u.extractID()

			if actual != tt.expected {
				t.Fatalf("extractID() = %v; want %v,", actual, tt.expected)
			}
		})
	}
}

func TestGetServiceName(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected MusicServices
	}{
		"itunes":  {"https://music.apple.com/us/artist/kaho-nakamura/586860631", ITunesService},
		"spotify": {"https://open.spotify.com/artist/0illCOhPkFBykngmCWos6u", SpotifyService},
		"unknown": {"https://www.youtube.com/watch?v=Z-48u_uWMHY", UnknownService},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			u, _ := url.Parse(tt.input)
			actual := getServiceName(u)

			if actual != tt.expected {
				t.Fatalf("getServiceName() = %v; want %v", actual, tt.expected)
			}
		})
	}
}
