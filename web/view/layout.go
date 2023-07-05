package view

import (
    "path/filepath"

    "github.com/gin-gonic/gin"
)

func SetView(r *gin.Engine, view string, layout_optional ...string) {
    layout := "application"

    if len(layout_optional) > 0 {
        layout = layout_optional[0]
    }

    lp := filepath.Join("web", "view", "layouts", layout + ".html")
    vp := filepath.Join("web", "view", view + ".html")

    r.LoadHTMLFiles(lp, vp)
}
