package scraper

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/grugrut/newsales-viewer/model"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"strings"
	"time"
)

// AcecookScrape retrieve information from nisshin
func AcecookScrape(ctx context.Context) {
	var targetURL = "https://www.acecook.co.jp/news/index.html"
	client := urlfetch.Client(ctx)

	resp, err := client.Get(targetURL)
	if err != nil {
		log.Infof(ctx, err.Error())
		return
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Infof(ctx, err.Error())
		return
	}

	doc.Find("#Main div.news-date dl").Each(func(i int, s *goquery.Selection) {
		proURL, _ := s.Find("dd a").Attr("href")
		proURL = "http://www.acecook.co.jp" + proURL
		imgURL := ""
		maker := "エースコック"
		dataText := strings.TrimSpace(s.Find("dd a").Text())
		if strings.HasSuffix(dataText, "　新発売") {
			dataText = strings.TrimSuffix(dataText, "　新発売")
			parts := strings.SplitAfter(dataText, "　")
			productName := strings.Join(parts[0:len(parts)-1], "")
			tmpDate := parts[len(parts)-1]
			productDate, err := time.Parse("2006/1/2", tmpDate)
			if err != nil {
				log.Infof(ctx, err.Error())
			}

			key, product := model.MakeProduct(productName, maker, productDate, proURL, imgURL, "カップラーメン")
			model.UpsertProduct(ctx, key, product)
		}
	})
}
