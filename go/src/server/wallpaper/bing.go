package wallpaper

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path"
	// "fmt"
)

func Bing (c *gin.Context) {
	bingPath := "../../../dist/bing"
	data := []string{}
	files, _ := ioutil.ReadDir(bingPath)
    for _, file := range files {
        if file.IsDir() {
            continue
        } else {
			fileSuffix := path.Ext(file.Name())
			if (fileSuffix == ".jpg") {
				data = append(data, file.Name())
			}
        }
	}
	c.JSON(http.StatusOK, data)
}