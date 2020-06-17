package pkg

import (
	"fmt"
	"time"
)

// Recorded
// +jam ../clefs/struct_slicer.go
type Video struct {
	ID                 int      `json:"id"`
	ProgramID          int64    `json:"programId"`
	ChannelID          int64    `json:"channelId"`
	ChannelType        string   `json:"channelType"`
	StartAt            int64    `json:"startAt"`
	EndAt              int64    `json:"endAt"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	Extended           string   `json:"extended"`
	Genre1             int      `json:"genre1"`
	Genre2             int      `json:"genre2"`
	Genre3             int      `json:"genre3"`
	Genre4             int      `json:"genre4"`
	Genre5             int      `json:"genre5"`
	Genre6             int      `json:"genre6"`
	VideoType          string   `json:"videoType"`
	VideoResolution    string   `json:"videoResolution"`
	VideoStreamContent int      `json:"videoStreamContent"`
	VideoComponentType int      `json:"videoComponentType"`
	AudioSamplingRate  int      `json:"audioSamplingRate"`
	AudioComponentType int      `json:"audioComponentType"`
	RuleID             int      `json:"ruleId"`
	Recording          bool     `json:"recording"`
	Protection         bool     `json:"protection"`
	IsTmp              bool     `json:"isTmp"`
	HasThumbnail       bool     `json:"hasThumbnail"`
	Original           bool     `json:"original"`
	Filename           string   `json:"filename"`
	Encoded            Encodeds `json:"encoded"`
}

func (x Video) Duration() int {
	return int((x.EndAt - x.StartAt) / 1000)
}

func (x Video) Start() string {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return time.Unix(x.StartAt/1000, 0).In(jst).Format("2006-01-02 15:04")
}

func (x Video) End() string {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return time.Unix(x.EndAt/1000, 0).In(jst).Format("15:04")
}

func (x Video) VideoURL() string {
	if len(x.Encoded) > 0 {
		//return fmt.Sprintf("/video/%d", x.ID)
		return fmt.Sprintf(ServerHost+"/api/recorded/%d/file?encodedId=%d", x.ID, x.Encoded[0].EncodedID)
	}
	return ""
}

func (x Video) ChannelName() string {
	for _, c := range CachedChannel {
		if c.ID == x.ChannelID {
			return c.Name
		}
	}
	return ""
}
