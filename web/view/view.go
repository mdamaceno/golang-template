package view

import (
    "path/filepath"

    "github.com/gin-contrib/multitemplate"
)

var viewsDir = filepath.Join("web", "view")
var defaultLayoutFile = filepath.Join(viewsDir, "layouts", "application.html")

func SetViews() multitemplate.Renderer {
    r := multitemplate.NewRenderer()

    r.AddFromFiles("home/index", defaultLayoutFile, viewsDir + "/home/index.html")
    r.AddFromFiles("users/index", defaultLayoutFile, viewsDir + "/users/index.html")

    return r
}
