package controller

import (
    "github.com/gin-gonic/gin"
)

var UsersController = map[string]func(c *gin.Context){
    "index": func (c *gin.Context) {
        c.HTML(200, "users/index", gin.H{
            "title": "Main website | Users",
            "name": "John Doe",
        })
    },
}
