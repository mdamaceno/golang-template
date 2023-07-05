package config

import (
    "github.com/gin-gonic/gin"
)

func SetupDefaultMiddleware(r *gin.Engine) *gin.Engine {
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    return r
}
