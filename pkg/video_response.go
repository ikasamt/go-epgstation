package pkg

//
type VideoResponse struct {
	Videos []Video `json:"recorded"`
	Total  int     `json:"total"`
}
