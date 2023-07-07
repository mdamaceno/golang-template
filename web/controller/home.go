package controller

import (
    "github.com/gin-gonic/gin"
)

func HomeController(r *gin.Engine, action string) func(c *gin.Context) {
    return map[string]func(c *gin.Context) {
        "index": homeIndex,
    }[action]
}

func homeIndex(c *gin.Context) {
    c.HTML(200, "home/index", gin.H{
        "title": "Main website | Home",
        "name": "John Doe",
    })
}
