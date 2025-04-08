package tupiService

import (
	"fmt"
	"gews/model"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type news struct {
	Client *Client
}

func (n *news) Get() ([]model.News, error) {
	html, err := n.Client.getHtml("")
	if err != nil {
		return nil, err
	}
	var newsArr []model.News
	html.Find("a").Each(func(i int, htmlSelection *goquery.Selection) {
		if !tryValidNews(htmlSelection) {
			return
		}
		link, _ := htmlSelection.Attr("href")
		for _, news := range newsArr {
			if link == news.Link {
				return
			}
		}
		news, err := n.Client.News.GetByUrl(link)
		if err != nil {
			_ = fmt.Errorf("%v", err.Error())
			return
		}
		if news.Headline == "" || news.ImageLink == "" || news.SubHeadling == "" {
			println(link)
		}
		newsArr = append(newsArr, news)
	})
	return newsArr, nil
}

func (n *news) GetByUrl(link string) (model.News, error) {
	html, err := n.Client.getHtml(link)
	if err != nil {
		return model.News{}, err
	}
	news := extractNews(html)
	news.Link = link
	return news, nil
}

func extractNews(html *goquery.Document) model.News {
	news := model.News{}
	news.Headline = extractHeadline(html)
	news.SubHeadling = extractSubHeadline(html)
	news.ImageLink = extractImage(html)
	return news
}

func extractImage(html *goquery.Document) string {
	selection := html.Find(`div[id="mvp-post-feat-img"]`)
	imageHtml := selection.Find("img")
	imageLink, exist := imageHtml.Attr("data-lazy-src")
	if !exist {
		imageLink, _ = imageHtml.Attr("src")
	}
	re := regexp.MustCompile(`(-\d+x\d+)(\.\w+)$`)
	imageLink = re.ReplaceAllString(imageLink, "$2")
	return imageLink
}

func extractSubHeadline(html *goquery.Document) string {
	selection := html.Find(`header[id="mvp-post-head"]`)
	subHeadline := ""
	selection.Find("span").Each(func(i int, s *goquery.Selection) {
		if i == 1 {
			subHeadline = s.Text()
		}
	})
	return prettifyText(subHeadline)
}

func extractHeadline(html *goquery.Document) string {
	selection := html.Find(`header[id="mvp-post-head"]`)
	headline := selection.Find("h1").Text()
	return prettifyText(headline)
}

func tryValidNews(html *goquery.Selection) bool {
	link, _ := html.Attr("href")
	hasValidLink := !strings.Contains(link, "webstories") || link == ""
	hasHeadline := html.Find("h2").Text() != ""
	_, hasImage := html.Find("img").Attr("src")
	return hasHeadline && hasImage && hasValidLink
}

func prettifyText(text string) string {
	text = strings.TrimSpace(strings.ReplaceAll(text, "\n", ". "))
	return removeFinalPeriod(text)
}

func removeFinalPeriod(text string) string {
	finalCaracter := text[len(text)-1:]
	if strings.Contains("!.?", finalCaracter) {
		text = text[:len(text)-1]
	}
	return text
}
