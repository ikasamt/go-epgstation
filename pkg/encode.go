package pkg

import (
	"encoding/json"
	"log"
)

// +jam ../clefs/struct_slicer.go
type Encoded struct {
	EncodedID int    `json:"encodedId"`
	Name      string `json:"name"`
	Filename  string `json:"filename"`
	Filesize  int64  `json:"size"`
}

// +jam ../clefs/struct_slicer.go
type Encoding struct {
	ID       int     `json:"id"`
	Mode     string  `json:"mode"`
	Percent  float64 `json:"percent"`
	Log      string  `json:"log"`
	Recorded Video   `json:"recorded"`
}

type EncodedResponse struct {
	WaitItems    []Encoding `json:"waitItems"`
	RunningItems []Encoding `json:"runningItems"`
}

func APIGETEncoded() (response EncodedResponse) {
	ba := APIGet(`encode`, map[string]interface{}{"isHalfWidth": "true"})
	if err := json.Unmarshal(ba, &response); err != nil {
		log.Println(err)
	}
	return
}
