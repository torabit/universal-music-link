package internal

import (
	"fmt"
	"net/url"
	"strings"
)

type MusicServices int

const (
	UnknownService MusicServices = iota
	ITunesService
	SpotifyService
)

var serviceMap = map[string]MusicServices{
	"open.spotify.com": SpotifyService,
	"music.apple.com":  ITunesService,
	"itunes.apple.com": ITunesService,
}

type URLElements struct {
	url.URL
}

func NewURLElements(urlString string) *URLElements {
	u, err := url.Parse(urlString)

	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return &URLElements{URL: *u}
	}

	return &URLElements{URL: *u}
}

func (u *URLElements) extractID() string {
	// URLの構成に厳密に依存しているため、URLの形式の変更には注意
	var id string

	// 現在ではitunesもspotifyもpathの最後の要素がIDになっている
	parts := strings.Split(u.Path, "/")
	idIndex := len(parts) - 1
	id = parts[idIndex]

	queries := u.Query()
	if len(queries) != 0 { // もしクエリパラメーターが存在するなら、そちらのIDを優先して返す
		var queryValues []string
		for _, value := range queries {
			queryValues = append(queryValues, value...)
		}
		id = queryValues[0]
	}

	return id
}

func (u *URLElements) getServiceName() MusicServices {
	host := u.Host

	serviceName, ok := serviceMap[host]
	if !ok {
		return UnknownService
	}

	return serviceName
}
