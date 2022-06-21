package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/upload", uploadFile)
	router.GET("/download", downloadFile)
	router.Run()
}

func uploadFile(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)

	year := time.Now().Year()
	month := time.Now().Month().String()
	day := time.Now().Day()

	path := strconv.Itoa(year) + "/" + month + "/" + strconv.Itoa(day) + "/"

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	ctx.SaveUploadedFile(file, path+file.Filename)
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", path+file.Filename))
}

func downloadFile(ctx *gin.Context) {
	path := ctx.Query("path")
	file, _ := os.Open(path)
	defer file.Close()
	
	ctx.File(path)
}
