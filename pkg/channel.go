package pkg

import (
	"encoding/json"
	"log"
)

type Channel struct {
	ID            int64  `json:"id"`
	ServiceID     int    `json:"serviceId"`
	NetworkID     int    `json:"networkId"`
	Name          string `json:"name"`
	HasLogoData   bool   `json:"hasLogoData"`
	ChannelType   string `json:"channelType"`
	ChannelTypeID int    `json:"channelTypeId"`
	Channel       string `json:"channel"`
	Type          int    `json:"type"`
}

func APIGETChannel() (response []Channel) {
	ba := APIGet(`channels`, map[string]interface{}{})
	if err := json.Unmarshal(ba, &response); err != nil {
		log.Println(err)
	}
	return
}
