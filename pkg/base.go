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
