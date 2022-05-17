package pkg

//
type VideoResponse struct {
	Videos Videos `json:"records"`
	Total  int    `json:"total"`
}
