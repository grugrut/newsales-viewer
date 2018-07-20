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

// NisshinScrape retrieve information from nisshin
func NisshinScrape(ctx context.Context) {
	var targetURL = "https://www.nissin.com/jp/news/pressrelease/"
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

	doc.Find("#ns_main section.ns-posts--news ul.ns-posts-list li").Each(func(i int, s *goquery.Selection) {
		date := s.Find("div.ns-posts-list-article-info em").Text()
		product := strings.TrimSpace(s.Find("div.ns-posts-list-article-description a div h3").Text())
		imgURL, _ := s.Find("div.ns-posts-list-article-description a figure span img").Attr("src")
		proURL, _ := s.Find("div.ns-posts-list-article-description a").Attr("href")
		proURL = "https://www.nissin.com" + proURL
		maker := s.Find("div.ns-posts-list-article-info span:nth-child(2) a").Text()

		// 日清食品の場合は末尾が 「日発売)」であるかどうかで新商品情報かどうかを判断する
		if strings.HasSuffix(product, "日発売)") {
			prList := strings.SplitAfter(product, "」")
			if len(prList) >= 2 {
				productName := strings.Join(prList[0:len(prList)-1], "")
				tmpPDate := prList[len(prList)-1]

				df1 := "(2006年1月2日発売)"
				df2 := "(1月2日発売)"

				productDate, err := time.Parse(df1, tmpPDate)
				if err != nil {
					productDate, err = time.Parse(df2, tmpPDate)
					if err != nil {
						log.Infof(ctx, err.Error())
					}
				}

				df3 := "2006.1.2"
				pressDate, err := time.Parse(df3, date)

				if productDate.Year() == 0 {
					var y int
					if productDate.Month() >= pressDate.Month() {
						y = pressDate.Year()
					} else {
						y = pressDate.Year() + 1
					}
					productDate = time.Date(y, productDate.Month(), productDate.Day(), 0, 0, 0, 0, time.Local)
				}
				key, product := model.MakeProduct(productName, maker, productDate, proURL, imgURL, "カップラーメン")
				model.UpsertProduct(ctx, key, product)
			}
		}
	})
}
