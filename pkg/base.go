package pkg

import (
	"encoding/json"
	"log"
)

var ServerHost string
var HDDBase string

var CachedResponse VideoResponse
var CachedChannel []Channel

func APIGetRecorded(limit int, offset int) (response VideoResponse) {
	ba := APIGet(`recorded`, map[string]interface{}{"limit": limit, "offset": offset})
	if err := json.Unmarshal(ba, &response); err != nil {
		log.Fatal(err)
	}
	return response
}

func SetCacheVideoResponse() {
	tmp := APIGetRecorded(1, 0)
	if CachedResponse.Total != tmp.Total {
		CachedResponse = APIGetRecorded(10000, 0)
	}
}


type DeleteResponse struct {
       Code int
}

func APIRemoveRecorded(key int) (err error, response DeleteResponse) {
       // curl -X DELETE "/api/recorded/1529" -H "accept: application/json"
       rawurl := fmt.Sprintf(ServerHost+"/api/recorded/%d", key)

       client := &http.Client{}
       log.Println(rawurl)
       req, err := http.NewRequest("DELETE", rawurl, nil)
       if err != nil {
               fmt.Println(err)
               return
       }

       resp, err := client.Do(req)
       if err != nil {
               return err, response
       }
       defer resp.Body.Close()

       byteArray, _ := ioutil.ReadAll(resp.Body)
}
