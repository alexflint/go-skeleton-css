package skeletoncss

import (
	"embed"
	"net/http"
)

//go:embed normalize.css
//go:embed skeleton.css
//go:embed fontfaces.css
//go:embed Raleway-Variable.ttf
//go:embed Raleway-Italic-Variable.ttf
var Assets embed.FS

// FileServer is a http handler that serves files. You must strip everything up to
// and including the last "/" in the path, like this:
//
//	http.Handle("/css/", http.StripPrefix("/css/", skeletoncss.FileServer))
//
// Then, reference it in your HTML like this:
//
//	<link rel="stylesheet" href="css/fontfaces.css">
//	<link rel="stylesheet" href="css/normalize.css">
//	<link rel="stylesheet" href="css/skeleton.css">
var FileServer = http.FileServer(http.FS(Assets))
