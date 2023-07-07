package view

import (
    "path/filepath"
    "github.com/gin-contrib/multitemplate"
)

var templateMap = map[string]string {
    "home/index": "application",
    "users/index": "application",
}

var viewDir = filepath.Join("web", "view")

func SetViews() multitemplate.Renderer {
    r := multitemplate.NewRenderer()

    for viewName, layout := range templateMap {
        r.AddFromFiles(viewName, viewDir + "/layouts/" + layout + ".html", viewDir + "/" + viewName + ".html")
    }

    return r
}
