package newsales

import (
	"github.com/grugrut/newsales-viewer/scraper"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/tasks/crawling", crawlTask)
}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/base.html", "view/main.html"))
	tmpl.Execute(w, nil)
}

func crawlTask(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "crawling start")
	scraper.NisshinScrape(ctx)
	log.Infof(ctx, "crawling end")
}
