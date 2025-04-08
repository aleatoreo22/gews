package tupiService

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	UrlBase = "https://www.tupi.fm/"
)

type Client struct {
	News news
}

func New() Client {
	client := Client{}
	client.News = news{Client: &client}
	return client
}

func (client *Client) getHtml(url string) (*goquery.Document, error) {
	if url == "" {
		url = UrlBase
	}
	request, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request.Request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	html, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	return html, nil
}
