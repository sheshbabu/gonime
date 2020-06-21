package handler

import (
	"animenow/anilist"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("./templates/*/*.html"))

type ListPage struct {
	Title string
	List  anilist.AnimeList
}

func main() {
	http.HandleFunc("/", HandleListPage)
	http.ListenAndServe(":8080", nil)
}

func HandleListPage(w http.ResponseWriter, r *http.Request) {
	list := anilist.GetPopularAnime()
	tmpl.ExecuteTemplate(w, "index.html", ListPage{Title: "Hello", List: list})
}
