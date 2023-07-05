package config

import (
    "path/filepath"

    "github.com/gin-gonic/gin"
)

func SetupAssets(r *gin.Engine) *gin.Engine {
    r.LoadHTMLGlob(filepath.Join("web", "resources", "views", "layouts", "*.html"))

    return r
}
