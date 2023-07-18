package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

var UsersController = map[string]func(c *gin.Context){
    "index": func (c *gin.Context) {
        c.HTML(http.StatusOK, "users/index", gin.H{
            "title": "Index | Users",
            "name": "John Doe",
        })
    },

    "show": func (c *gin.Context) {
        c.HTML(http.StatusOK, "users/show", gin.H{
            "title": "Show | Users",
            "name": "John Doe",
        })
    },

    "new": func (c *gin.Context) {
        c.HTML(http.StatusOK, "users/new", gin.H{
            "title": "New | Users",
            "name": "John Doe",
        })
    },

    "edit": func (c *gin.Context) {
        c.HTML(http.StatusOK, "users/edit", gin.H{
            "title": "Edit | Users",
            "name": "John Doe",
        })
    },

    "create": func (c *gin.Context) {
        c.JSON(http.StatusCreated, nil)
    },

    "update": func (c *gin.Context) {
        c.JSON(http.StatusOK, nil)
    },

    "destroy": func (c *gin.Context) {
        c.JSON(http.StatusNoContent, nil)
    },
}
