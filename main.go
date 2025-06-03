package main

import "github.com/gin-gonic/gin"

type CatForm struct {
	Key string `json:"catName" binding:"required"`
}

type DogForm struct {
	Key string `json:"dogName" binding:"required"`
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

	router.Run()
}
