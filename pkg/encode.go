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
	ID         string `json:"id"`
	Name       string `json:"name"`
	Recordedid int    `json:"recordedId"`
	Program    struct {
		ID                 int    `json:"id"`
		Programid          int64  `json:"programId"`
		Channelid          int64  `json:"channelId"`
		Channeltype        string `json:"channelType"`
		Startat            int64  `json:"startAt"`
		Endat              int64  `json:"endAt"`
		Name               string `json:"name"`
		Description        string `json:"description"`
		Extended           string `json:"extended"`
		Genre1             int    `json:"genre1"`
		Genre2             int    `json:"genre2"`
		Genre3             int    `json:"genre3"`
		Genre4             int    `json:"genre4"`
		Genre5             int    `json:"genre5"`
		Genre6             int    `json:"genre6"`
		Videotype          string `json:"videoType"`
		Videoresolution    string `json:"videoResolution"`
		Videostreamcontent int    `json:"videoStreamContent"`
		Videocomponenttype int    `json:"videoComponentType"`
		Audiosamplingrate  int    `json:"audioSamplingRate"`
		Audiocomponenttype int    `json:"audioComponentType"`
		Ruleid             int    `json:"ruleId"`
		Recording          bool   `json:"recording"`
		Protection         bool   `json:"protection"`
		Filesize           int64  `json:"filesize"`
		Istmp              bool   `json:"isTmp"`
		Hasthumbnail       bool   `json:"hasThumbnail"`
		Original           bool   `json:"original"`
		Filename           string `json:"filename"`
	} `json:"program"`
	Mode int `json:"mode"`
}

// +jam ../clefs/struct_slicer.go
type EncodedQueue struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RecordedID int    `json:"recordedId"`
	Program    struct {
		ID                 int    `json:"id"`
		Programid          int64  `json:"programId"`
		Channelid          int64  `json:"channelId"`
		Channeltype        string `json:"channelType"`
		Startat            int64  `json:"startAt"`
		Endat              int64  `json:"endAt"`
		Name               string `json:"name"`
		Description        string `json:"description"`
		Extended           string `json:"extended"`
		Genre1             int    `json:"genre1"`
		Genre2             int    `json:"genre2"`
		Genre3             int    `json:"genre3"`
		Genre4             int    `json:"genre4"`
		Genre5             int    `json:"genre5"`
		Genre6             int    `json:"genre6"`
		Videotype          string `json:"videoType"`
		Videoresolution    string `json:"videoResolution"`
		Videostreamcontent int    `json:"videoStreamContent"`
		Videocomponenttype int    `json:"videoComponentType"`
		Audiosamplingrate  int    `json:"audioSamplingRate"`
		Audiocomponenttype int    `json:"audioComponentType"`
		Ruleid             int    `json:"ruleId"`
		Recording          bool   `json:"recording"`
		Protection         bool   `json:"protection"`
		Filesize           int64  `json:"filesize"`
		Istmp              bool   `json:"isTmp"`
		Hasthumbnail       bool   `json:"hasThumbnail"`
		Original           bool   `json:"original"`
		Filename           string `json:"filename"`
	} `json:"program"`
	Mode int `json:"mode"`
}

type EncodedResponse struct {
	Queue    []EncodedQueue `json:"queue"`
	Encoding Encoding       `json:"encoding"`
}

func APIGETEncoded() (response EncodedResponse) {
	ba := APIGet(`encode`, map[string]interface{}{})
	if err := json.Unmarshal(ba, &response); err != nil {
		log.Println(err)
	}
	return
}
