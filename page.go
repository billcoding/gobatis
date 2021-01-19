package gobatis

type page struct {
	Offset    int           `json:"offset"`
	Size      int           `json:"size"`
	TotalRows int           `json:"total_rows"`
	TotalPage int           `json:"total_page"`
	List      []interface{} `json:"list"`
}
