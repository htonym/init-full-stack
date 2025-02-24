package api

import "net/http"

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="icon" type="image/x-icon" href="/static/img/icon.svg">
		<title>Document</title>
	</head>
	<body>
		<h1>Home Page</h1>
	</body>
	</html>
	`))
}
