package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

var engine *gin.Engine

func main() {
	config()
	engine = gin.Default()
	err := engine.SetTrustedProxies(nil)
	handleError(err)
	if err != nil {
		return
	}
	log.SetFlags(0)
	log.SetPrefix("[GIN-GONIC] ")
	log.SetOutput(gin.DefaultWriter)
	router()
	err = os.Mkdir("uploads", os.FileMode(0777))
	engine.Static("/uploads", "/uploads")
	err = engine.Run(fmt.Sprintf(":%v", Configuration.Port))
	handleError(err)
}

func router() {
	engine.GET("/upload", func(context *gin.Context) {
		var fileBind FileBind

		err := context.ShouldBind(&fileBind)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		file := fileBind.File
		code := randomString()
		err = context.SaveUploadedFile(file, fmt.Sprintf("/uploads/%v", code))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code": code,
		})
	})
}

func randomString() string {
	rand.Seed(time.Now().UnixMilli())
	length := Configuration.KeyLenght
	letters := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type FileBind struct {
	Name string                `form:"name"`
	File *multipart.FileHeader `form:"file" binding:"required"`
}
