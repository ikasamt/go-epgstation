package pkg

import (
	"fmt"
	_ "image/jpeg"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func JPEGIndexHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "image/jpeg")
	c.Status(200)

	id_ := c.Param("ID")
	rawurl := fmt.Sprintf(ServerHost+"/api/recorded/%s/thumbnail", id_)
	resp, err := http.Get(rawurl)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = io.Copy(c.Writer, resp.Body)
	defer resp.Body.Close()
}

func VideoServeHandler(c *gin.Context) {
	key := c.Query("Key")
	fileName := filepath.Join(HDDBase, key)
	http.ServeFile(c.Writer, c.Request, fileName)
}

func APIRemoveHandler(c *gin.Context) {
	idStr := c.Param("ID")
	id_, _ := strconv.Atoi(idStr)
	err, response := APIRemoveRecorded(id_)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"err": err})
		return
	}

	c.JSON(200, response)
	return
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func APIEncodedHandler(c *gin.Context) {
	videos := CachedResponse.Videos.Where(func(x Video) bool {
		return len(x.Encoded) == 0
	})

	videosIDs := []int{}
	response := APIGETEncoded()

	for _, item := range response.RunningItems {
		videosIDs = append(videosIDs, item.Recorded.ID)
	}

	for _, item := range response.WaitItems {
		videosIDs = append(videosIDs, item.Recorded.ID)
	}

	cnt := len(response.WaitItems)

	for _, v := range videos {
		if cnt > 10 {
			c.JSON(200, cnt)
			return
		}

		if !contains(videosIDs, v.ID) {
			APIPOSTEncoded(v.ID)
			cnt++
		}
	}

	return
}
