package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tent/hawk-go"
	"github.com/tent/tent-client-go"
)

var meta *tent.MetaPost
var client *tent.Client

func discover() []*request {
	var err error
	meta, err = tent.Discover(os.Args[1])
	maybePanic(err)
	client = &tent.Client{Servers: meta.Servers}
	return getRequests()
}

func createApp() *hawk.Credentials {
	post := tent.NewAppPost(&tent.App{
		Name: "liz-tent",
		URL:  "http://127.0.0.1:8000",
		Types: tent.AppTypes{
			Write: []string{"all"},
		},
		Scopes:      []string{"permissions"},
		RedirectURI: "http://127.0.0.1:8000/oauth",
	})
	err := client.CreatePost(post)
	maybePanic(err)
	client.Credentials, _, err = post.LinkedCredentials()
	maybePanic(err)
	oauthURL := meta.Servers[0].URLs.OAuthURL(post.ID, "d173d2bb868a")
	fmt.Println(oauthURL)
	var code string
	fmt.Scanln(&code)
	creds, err := client.RequestAccessToken(code)
	maybePanic(err)
	client.Credentials = creds
	return creds
}

func statusPost(text string) *tent.Post {
	return &tent.Post{
		Type:    "https://tent.io/types/status/v0#",
		Content: []byte(fmt.Sprintf(`{"text": "%s"}`, text)),
	}
}

func newPost(text string) *request {
	err := client.CreatePost(statusPost(text))
	maybePanic(err)
	return getRequests()[0]
}

type stringReader struct{ *strings.Reader }

func (r stringReader) Len() int64 { return int64(r.Reader.Len()) }

func maybePanic(err error) {
	if err != nil {
		if resErr, ok := err.(*tent.ResponseError); ok && resErr.TentError != nil {
			fmt.Println(resErr.TentError)
		}
		panic(err)
	}
}
