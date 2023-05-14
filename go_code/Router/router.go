package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var IsLogin bool = false

func SetRouter(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "status Ok",
		})
	})
	r.POST("/login", func(ctx *gin.Context) {
		cookies := ctx.Request.Cookies()
		err := ctx.Request.ParseForm()
		if err != nil {
			log.Println("Data de las auth formas no es valido")
		}
		for key, value := range ctx.Request.PostForm {
			fmt.Println(key, value)
		}
	})
}
