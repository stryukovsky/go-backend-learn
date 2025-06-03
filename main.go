package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CatForm struct {
	Key string `json:"catName" binding:"required"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/asciiJSON", func(ctx *gin.Context) {
		result := map[string]any{
			"lang": "Россия",
		}
		ctx.AsciiJSON(200, result)
	})

	router.POST("/animal", func(ctx *gin.Context) {
		cat := CatForm{}
		if err := ctx.ShouldBind(&cat); err == nil {
			ctx.JSON(200, gin.H{
				"message": "cat over there!",
			})
		}
	})

	router.POST("/login", func(ctx *gin.Context) {
		user := Login{}
		ctx.Bind(&user)
		if user.Username == "Dima" && user.Password == "qwerty" {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "successfully authenticated",
			})
		} else {
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "Cannot enter",
			})
		}
	})

	router.Run()
}
