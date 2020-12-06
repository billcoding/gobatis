package gobatis

//Define Helper struct
type Helper struct {
	batis   *Batis //Bundle batis
	binding string //XML binding
	id      string //XML node id
	ds      string //DS
}

//Return new helper
func NewHelper(binding, id string) *Helper {
	return NewHelperWithDS(binding, id, "")
}

//Return new helper with DS
func NewHelperWithDS(binding, id, ds string) *Helper {
	b := Default()
	if !b.inited {
		panic("The batis not yet initialized")
	}
	return &Helper{
		batis:   b,
		binding: binding,
		id:      id,
		ds:      ds,
	}
}

//Return query
func (h *Helper) Select() *selectMapper {
	return h.batis.Mapper(h.binding).SelectWithDS(h.id, h.ds)
}

//Return update
func (h *Helper) Update() *updateMapper {
	return h.batis.Mapper(h.binding).UpdateWithDS(h.id, h.ds)
}
