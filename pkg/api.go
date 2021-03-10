package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func APIGet(path_ string, params map[string]interface{}) []byte {
	paramsList := []string{}
	for k, v := range params {
		var s string
		switch v.(type) {
		case int:
			s = fmt.Sprintf("%s=%d", k, v)
		default:
			s = fmt.Sprintf("%s=%s", k, v)
		}
		paramsList = append(paramsList, s)
	}
	queryStr := strings.Join(paramsList, "&")

	rawurl := fmt.Sprintf("%s/api/%s?%s", ServerHost, path_, queryStr)
	log.Println(rawurl)
	resp, err := http.Get(rawurl)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	return byteArray
}

func APIPOSTEncoded(ID int) int {
	rawurl := fmt.Sprintf("%s/api/recorded/%d/encode", ServerHost, ID)
	values, _ := json.Marshal(gin.H{
		"mode":                         0,
		"isOutputTheOriginalDirectory": true,
		"delTs":                        true})

	log.Println(rawurl)

	resp, err := http.Post(rawurl, "application/json", bytes.NewBuffer(values))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
