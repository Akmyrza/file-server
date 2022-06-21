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
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.Error(err)
	}

	fmt.Print("asd")
	path := makePath(file.Filename)

	er := os.MkdirAll(path, os.ModePerm)
	if er != nil {
		ctx.Error(er)
	}

	ctx.SaveUploadedFile(file, path)
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", path))
}

func makePath(filename string) string {
	year := time.Now().Year()
	month := time.Now().Month().String()
	day := time.Now().Day()
	//currentTime := time.Now()
	return strconv.Itoa(year) + "/" + month + "/" + strconv.Itoa(day) + "/" + filename
	//	return string(currentTime.Format("2000/01/31"))
}

func downloadFile(ctx *gin.Context) {
	path := ctx.Query("path")
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	ctx.File(path)
}
