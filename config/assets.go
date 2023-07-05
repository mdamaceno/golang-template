package config

import (
    "github.com/gin-gonic/gin"
)

func SetupLayouts(r *gin.Engine) *gin.Engine {
    r.Static("/public", "./public")

    return r
}
