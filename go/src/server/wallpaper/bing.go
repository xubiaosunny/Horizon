package wallpaper

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type WapllpStruct struct {
	Date      string
	FileName  string
	Copyright string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Bing(c *gin.Context) {
	idx := c.DefaultQuery("idx", "0")
	n := c.DefaultQuery("n", "10")

	bingPath := "../../../dist/bing"
	data := []WapllpStruct{}

	db, err := sql.Open("sqlite3", bingPath+"/bing.db")
	checkErr(err)

	stmt, err := db.Prepare("SELECT * FROM cn_bing order by date desc limit ?, ?;")
	checkErr(err)

	rows, err := stmt.Query(idx, n)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var date string
		var copyright string
		if err := rows.Scan(&date, &copyright); err != nil {
			checkErr(err)
		}
		// fmt.Printf("date=%s copyright=%s\n", date, copyright)
	}
	if err := rows.Err(); err != nil {
		checkErr(err)
	}

	files, _ := ioutil.ReadDir(bingPath)
	fmt.Println(reverse(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fileSuffix := path.Ext(file.Name())
			if fileSuffix == ".jpg" {
				data = append(data, WapllpStruct{fileSuffix, file.Name()})
			}
		}
	}
	fmt.Println(data)
	c.JSON(http.StatusOK, data)
}
