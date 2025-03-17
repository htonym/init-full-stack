package api

import "net/http"

type LoginPageData struct {
	BaseData
	Title string
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	data := LoginPageData{
		BaseData: BaseData{
			HideNavbar: true,
		},
		Title: "Login",
	}

	NewTemplate().Render(w, "login", data)
}
