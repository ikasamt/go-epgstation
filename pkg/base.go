package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var ServerHost string
var HDDBase string

var CachedResponse VideoResponse
var CachedChannel []Channel

func APIGetRecorded(limit int, offset int) (response VideoResponse) {
	ba := APIGet(`recorded`, map[string]interface{}{"limit": limit, "offset": offset})
	if err := json.Unmarshal(ba, &response); err != nil {
		log.Println(err)
	}
	return response
}

func SetCacheVideoResponse() {
	tmp := APIGetRecorded(1, 0)
	if CachedResponse.Total != tmp.Total {
		CachedResponse = APIGetRecorded(100000, 0)
	}
}

type DeleteResponse struct {
	Code int
}

func APIRemoveRecorded(key int) (err error, response DeleteResponse) {
	// curl -X DELETE "/api/recorded/1529" -H "accept: application/json"
	rawurl := fmt.Sprintf("%s/api/recorded/%d", ServerHost, key)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", rawurl, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return err, response
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(byteArray, &response); err != nil {
		return err, response
	}
	return err, response
}
