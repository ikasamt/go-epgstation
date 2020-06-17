package pkg

//
type VideoResponse struct {
	Videos Videos `json:"recorded"`
	Total  int    `json:"total"`
}
