package gobatis

//Define page interface
type page struct {
	Offset    int           //The page offset
	Size      int           //The page size
	TotalRows int           //The total record size
	TotalPage int           //The total page number
	List      []interface{} //The return data
}
