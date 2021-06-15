package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Routing(){
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("../public", "./public")

	r.GET("/", func(c *gin.Context) {
		files,err := ioutil.ReadDir("./public")
		if err != nil {
			print(err.Error())
		}
		c.HTML(http.StatusOK, "main.html", gin.H{
			"title":"Home Page",
			"test":files,
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			print(err.Error())
			c.String(400, fmt.Sprintf("error while reading out multipart"))
		}
		files,_ := form.File["file"]
		data,_ := ioutil.ReadDir("./public")

		for _,file := range data {
			for _, upfile := range files {
				if upfile.Filename == file.Name() {
					filename := strings.Split(upfile.Filename, ".")
					upfile.Filename = filename[0] + strconv.Itoa(len(data)) + "." + filename[1]
				}
			}
		}

		for _,file := range files {
			err := c.SaveUploadedFile(file, "./public/"+file.Filename)
			if err != nil {
				print(err.Error())
			}
		}
		c.Redirect(301, "/")
	})

	r.GET("/download/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, name)
		c.File("./public/hello.html")
		return
	})

	r.Run()
}
