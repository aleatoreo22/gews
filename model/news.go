package model

import "fmt"

type News struct {
	Link        string
	Headline    string
	SubHeadling string
	ImageLink   string
}

type NewsSite struct {
	Site string
	News []News
	Visible bool
}

func (newsSite *NewsSite) DownloadImages() {
	for _, news := range newsSite.News {
		fmt.Printf("Baixar: %s", news.ImageLink)
	}
}
