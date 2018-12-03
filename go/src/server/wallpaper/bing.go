package wallpaper

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type WapllpStruct struct {
	Date      string
	FileSrc   string
	Copyright string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Bing(c *gin.Context) {
	idx := c.DefaultQuery("idx", "0")
	n := c.DefaultQuery("n", "1")

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
		ff, _ := os.Open(bingPath + "/" + date + ".jpg")
		defer ff.Close()
		sourcebuffer := make([]byte, 500000)
		n, _ := ff.Read(sourcebuffer)
		sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])

		data = append(data, WapllpStruct{date, sourcestring, copyright})
	}
	if err := rows.Err(); err != nil {
		checkErr(err)
	}

	// files, _ := ioutil.ReadDir(bingPath)
	// for _, file := range files {
	// 	if file.IsDir() {
	// 		continue
	// 	} else {
	// 		fileSuffix := path.Ext(file.Name())
	// 		if fileSuffix == ".jpg" {
	// 			data = append(data, WapllpStruct{fileSuffix, file.Name()})
	// 		}
	// 	}
	// }
	c.JSON(http.StatusOK, data)
}
