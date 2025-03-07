package api

import "net/http"

type HomePageData struct {
	BaseData
	Title string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	data := HomePageData{
		BaseData: BaseData{
			User: "John Doe",
		},
		Title: "Home",
	}

	template := NewTemplate()
	template.Render(w, "home", data)
}
