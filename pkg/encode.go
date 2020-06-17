package pkg

// +jam ../clefs/struct_slicer.go
type Encoded struct {
	EncodedID int    `json:"encodedId"`
	Name      string `json:"name"`
	Filename  string `json:"filename"`
	Filesize  int64  `json:"filesize"`
}

// +jam ../clefs/struct_slicer.go
type EncodedProgram struct {
	ID                 int    `json:"id"`
	ProgramID          int64  `json:"programId"`
	ChannelID          int64  `json:"channelId"`
	ChannelType        string `json:"channelType"`
	StartAt            int64  `json:"startAt"`
	EndAt              int64  `json:"endAt"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Extended           string `json:"extended"`
	Genre1             int    `json:"genre1"`
	Genre2             int    `json:"genre2"`
	Genre3             int    `json:"genre3"`
	Genre4             int    `json:"genre4"`
	Genre5             int    `json:"genre5"`
	Genre6             int    `json:"genre6"`
	VideoType          string `json:"videoType"`
	VideoResolution    string `json:"videoResolution"`
	VideoStreamContent int    `json:"videoStreamContent"`
	VideoComponentType int    `json:"videoComponentType"`
	AudioSamplingRate  int    `json:"audioSamplingRate"`
	AudioComponentType int    `json:"audioComponentType"`
	RuleID             int    `json:"ruleId"`
	Recording          bool   `json:"recording"`
	Protection         bool   `json:"protection"`
	Filesize           int64  `json:"filesize"`
	IsTmp              bool   `json:"isTmp"`
	HasThumbnail       bool   `json:"hasThumbnail"`
	Original           bool   `json:"original"`
	Filename           string `json:"filename"`
}

type EncodedEncoding struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	RecordedID int             `json:"recordedId"`
	Program    EncodedPrograms `json:"program"`
	Mode       int             `json:"mode"`
}

type EncodedQueue struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	RecordedID int             `json:"recordedId"`
	Program    EncodedPrograms `json:"program"`
	Mode       int             `json:"mode"`
}
