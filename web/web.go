// Package web provides static assets, templates, components
// and their associated data structures for rendering.
package web

import (
	"embed"
	"io/fs"
)

//go:embed static/*.css static/*.html static/*.js
var static embed.FS
var Static, _ = fs.Sub(static, "static")
