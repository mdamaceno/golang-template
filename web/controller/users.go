package controller

import (
    "github.com/gin-gonic/gin"
)

func UsersController(r *gin.Engine, action string) func(c *gin.Context) {
    return map[string]func(c *gin.Context) {
        "index": usersIndex,
    }[action]
}

func usersIndex(c *gin.Context) {
    c.HTML(200, "users/index", gin.H{
        "title": "Main website | Users",
        "name": "John Doe",
    })
}
