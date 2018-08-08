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

// SanyoScrape retrieve information from nisshin
func SanyoScrape(ctx context.Context) {
	var targetURL = "http://www.sanyofoods.co.jp/"
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

	doc.Find("#ContentProductHeadline table.TBL-ProductsHeadline tbody tr").Each(func(i int, s *goquery.Selection) {
		// TODO sanyoのページにあわせてスクレイピングする
		imgURL, _ := s.Find("td.TD-ProductsImg img").Attr("src")
		if strings.HasSuffix(imgURL, "_72.jpg") {
			imgURL = strings.TrimSuffix(imgURL, "_72.jpg") + ".jpg"
		}
		imgURL = targetURL + imgURL
		proURL, _ := s.Find("td.TD-ProductsTitle p.P-ProductsHeadlineTitle a").Attr("href")
		proURL = targetURL + proURL
		maker := "サンヨー食品"
		productName := strings.TrimSpace(s.Find("td.TD-ProductsTitle p.P-ProductsHeadlineTitle").Text())
		dataText := strings.TrimSpace(s.Find("td.TD-ProductsTitle p.P-ProductsHeadlineDate2").Text())
		if strings.HasSuffix(dataText, " 発売") {
			dataText = strings.TrimSuffix(dataText, " 発売")
			productDate, err := time.Parse("2006.1.2", dataText)
			if err != nil {
				log.Infof(ctx, err.Error())
			}
			key, product := model.MakeProduct(productName, maker, productDate, proURL, imgURL, "カップラーメン")
			model.UpsertProduct(ctx, key, product)
		}
	})
}
