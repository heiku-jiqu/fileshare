package template

import "embed"

//go:embed *.css *.html *.js
var Website embed.FS
