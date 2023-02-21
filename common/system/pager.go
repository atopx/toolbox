package system

type Page struct {
	Index       int `json:"index"`
	Size        int `json:"size"`
	TotalCount  int `json:"total_count"`
	FilterCount int `json:"filter_count"`
}
