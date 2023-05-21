package internal

import (
	"log"
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

type MusicServiceURL struct {
	url.URL
	service MusicServices
}

func NewMusicServiceURL(urlString string) *MusicServiceURL {
	u, err := url.Parse(urlString)

	if err != nil {
		log.Fatal(err)
	}

	serviceName := getServiceName(u)

	if serviceName == UnknownService {
		log.Fatalf("%v is Unknown Music Service", u.Host)
	}

	return &MusicServiceURL{URL: *u, service: serviceName}
}

func (u *MusicServiceURL) extractID() string {
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

func getServiceName(u *url.URL) MusicServices {
	host := u.Host

	serviceName, ok := serviceMap[host]
	if !ok {
		return UnknownService
	}

	return serviceName
}
