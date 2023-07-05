package view

import (
    "path/filepath"

    "github.com/gin-gonic/gin"
)

func SetupLayout(r *gin.Engine, view string, layout string) {
    if layout == "" {
        layout = "application"
    }

    lp := filepath.Join("web", "view", "layouts", layout + ".html")
    vp := filepath.Join("web", "view", "pages", view + ".html")

    r.LoadHTMLFiles(lp, vp)
}
