package model

import (
	"fmt"
	site "gews/model/enum"
)

type News struct {
	Link        string
	Headline    string
	SubHeadling string
	ImageLink   string
}

type NewsSite struct {
	SiteName string
	site.Site
	News    []News
	Visible bool
}

func (newsSite *NewsSite) DownloadImages() {
	for _, news := range newsSite.News {
		fmt.Printf("Baixar: %s", news.ImageLink)
	}
}
