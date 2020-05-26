package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
