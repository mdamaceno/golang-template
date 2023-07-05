package controller

import (
    "github.com/gin-gonic/gin"

    "app/web/view"
)

func HomeIndex(r *gin.Engine) func(*gin.Context) {
    return func(c *gin.Context) {
        view.SetupLayout(r, "home", "application")
        c.HTML(200, "home.html", gin.H{
            "title": "Main website",
            "name": "John Doe",
        })
    }
}
