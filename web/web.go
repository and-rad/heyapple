package web

import (
	"embed"
	"html/template"
)

var (
	//go:embed templates
	Templates embed.FS
)

var (
	App   = template.Must(template.ParseFS(Templates, "templates/app.html"))
	Home  = template.Must(template.ParseFS(Templates, "templates/home.html"))
	Login = template.Must(template.ParseFS(Templates, "templates/login.html"))
)
