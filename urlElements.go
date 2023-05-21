package internal

import (
	"fmt"
	"net/url"
	"strings"
)

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
	// 現在ではitunesもspotifyもpathの最後の要素がIDになっている
	// URLの構成に厳密に依存しているため、URLの形式の変更には注意

	var id string

	parts := strings.Split(u.Path, "/")
	idIndex := len(parts) - 1
	id = parts[idIndex]

	if len(u.Query()) != 0 { // もしクエリパラメーターが存在するなら、そちらのIDを優先して返す
		var queryValues []string
		for _, value := range u.Query() {
			queryValues = append(queryValues, value...)
		}
		id = queryValues[0]
	}

	return id
}
