package gobatis

type Page struct {
	Offset    int           `json:"offset"`
	Size      int           `json:"size"`
	TotalRows int           `json:"total_rows"`
	TotalPage int           `json:"total_page"`
	List      []interface{} `json:"list"`
}

type PageMap struct {
	Offset    int                      `json:"offset"`
	Size      int                      `json:"size"`
	TotalRows int                      `json:"total_rows"`
	TotalPage int                      `json:"total_page"`
	List      []map[string]interface{} `json:"list"`
}
