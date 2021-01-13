package gobatis

// Helper struct
type Helper struct {
	batis   *Batis //Bundle batis
	binding string //XML binding
	id      string //XML node id
	ds      string //DS
}

// NewHelper return new helper
func NewHelper(binding, id string) *Helper {
	return NewHelperWithDS(binding, id, "")
}

// NewHelperWithDS return new helper with DS
func NewHelperWithDS(binding, id, ds string) *Helper {
	return NewHelperWithBatisAndDS(Default(), binding, id, ds)
}

// NewHelperWithBatis return new helper with DS
func NewHelperWithBatis(batis *Batis, binding, id string) *Helper {
	return NewHelperWithBatisAndDS(batis, binding, id, "")
}

// NewHelperWithBatisAndDS return new helper with DS
func NewHelperWithBatisAndDS(batis *Batis, binding, id, ds string) *Helper {
	return &Helper{
		batis:   batis,
		binding: binding,
		id:      id,
		ds:      ds,
	}
}

// Select return query
func (h *Helper) Select() *SelectMapper {
	return h.batis.Mapper(h.binding).SelectWithDS(h.id, h.ds)
}

// Update return update
func (h *Helper) Update() *UpdateMapper {
	return h.batis.Mapper(h.binding).UpdateWithDS(h.id, h.ds)
}
