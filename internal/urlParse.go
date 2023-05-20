package internal

import (
	"fmt"
	"net/url"
	"strings"
)

type URLElements struct {
	schema   string
	host     string
	path     string
	rawQuery string
	fragment string
	user     *url.Userinfo
	query    url.Values
}

func newURLElements(urlString string) *URLElements {
	u, err := url.Parse(urlString)

	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return &URLElements{}
	}

	c := &URLElements{
		schema:   u.Scheme,
		host:     u.Host,
		path:     u.Path,
		rawQuery: u.RawQuery,
		fragment: u.Fragment,
		user:     u.User,
		query:    u.Query(),
	}

	return c
}

func (u *URLElements) extractID() string {
	// 現在ではitunesもspotifyもpathの最後の要素がIDになっている
	// URLの構成に厳密に依存しているため、URLの形式の変更には注意

	var id string

	parts := strings.Split(u.path, "/")
	idIndex := len(parts) - 1
	id = parts[idIndex]

	if len(u.query) != 0 { // もしクエリパラメーターが存在するなら、そちらのIDを優先して返す
		var queryValues []string
		for _, value := range u.query {
			queryValues = append(queryValues, value...)
		}
		id = queryValues[0]
	}

	return id
}
