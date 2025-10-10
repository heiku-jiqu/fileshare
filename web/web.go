// Package web provides static assets, templates, components
// and their associated data structures for rendering.
package web

import (
	"embed"
	"html/template"
)

//go:embed static/*.css static/*.js
var Static embed.FS

//go:embed template
var t embed.FS

var Index = template.Must(template.ParseFS(t, "template/base.tmpl.html", "template/components/*", "template/index.html"))
var Login = template.Must(template.ParseFS(t, "template/base.tmpl.html", "template/components/*", "template/login.html"))
var Upload = template.Must(template.ParseFS(t, "template/base.tmpl.html", "template/components/*", "template/upload.html"))
