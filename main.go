package handler

import (
	"animenow/anilist"
	"html/template"
	"net/http"
	"strings"
)

var tmpl = template.Must(template.ParseGlob("./templates/*/*.html"))

type ListPage struct {
	List anilist.AnimeList
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/anime/") {
		id := strings.TrimPrefix(r.URL.Path, "/anime/")
		id = strings.TrimSuffix(id, "/")
		detail := anilist.GetAnimeDetail(id)
		tmpl.ExecuteTemplate(w, "detail.html", detail)
	} else {
		list := anilist.GetPopularAnime()
		tmpl.ExecuteTemplate(w, "index.html", ListPage{List: list})
	}
}
