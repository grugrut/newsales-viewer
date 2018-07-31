package newsales

import (
	"encoding/json"
	"github.com/grugrut/newsales-viewer/model"
	"github.com/grugrut/newsales-viewer/scraper"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/api/", api)
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

func api(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "call api: method: %v, path: %v", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/api/product":
			products, err := model.FetchAllProduct(ctx)
			if err != nil {
				log.Errorf(ctx, err.Error())
				http.Redirect(w, r, "/", http.StatusBadRequest)
			}
			result, err := json.Marshal(products)
			if err != nil {
				log.Errorf(ctx, err.Error())
				http.Redirect(w, r, "/", http.StatusBadRequest)
			}
			w.Write(result)
		}
	}
}
