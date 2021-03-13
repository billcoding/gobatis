package predicate

import "fmt"

type (
	Sort interface{ SQL() string }
	asc  struct{ column Column }
	desc struct{ column Column }
)

var (
	Asc  = func(c Column) Sort { return &asc{column: c} }
	Desc = func(c Column) Sort { return &desc{column: c} }
)

func (s *asc) SQL() string {
	return fmt.Sprintf("t.%v ASC", s.column)
}

func (s *desc) SQL() string {
	return fmt.Sprintf("t.%v DESC", s.column)
}
