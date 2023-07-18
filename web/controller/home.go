package controller

import (
    "github.com/gin-gonic/gin"
)

var HomeController = map[string]func(c *gin.Context){
    "index": func (c *gin.Context) {
        c.HTML(200, "home/index", gin.H{
            "title": "Main website | Home",
            "name": "John Doe",
        })
    },
}
