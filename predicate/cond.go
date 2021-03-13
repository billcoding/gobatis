package predicate

import (
	"fmt"
	"strings"
)

type (
	Cond interface {
		SQL() (string, []interface{})
	}
	eq struct {
		column Column
		op     string
		value  interface{}
	}
	notEq struct {
		column Column
		op     string
		value  interface{}
	}
	like struct {
		column Column
		op     string
		value  interface{}
	}
	notLike struct {
		column Column
		op     string
		value  interface{}
	}
	leftLike struct {
		column Column
		op     string
		value  interface{}
	}
	notLeftLike struct {
		column Column
		op     string
		value  interface{}
	}
	rightLike struct {
		column Column
		op     string
		value  interface{}
	}
	notRightLike struct {
		column Column
		op     string
		value  interface{}
	}
	instr struct {
		column Column
		op     string
		value  interface{}
	}
	notInstr struct {
		column Column
		op     string
		value  interface{}
	}
	gt struct {
		column Column
		op     string
		value  interface{}
	}
	notGt struct {
		column Column
		op     string
		value  interface{}
	}
	gtEq struct {
		column Column
		op     string
		value  interface{}
	}
	notGtEq struct {
		column Column
		op     string
		value  interface{}
	}
	lt struct {
		column Column
		op     string
		value  interface{}
	}
	notLt struct {
		column Column
		op     string
		value  interface{}
	}
	ltEq struct {
		column Column
		op     string
		value  interface{}
	}
	notLtEq struct {
		column Column
		op     string
		value  interface{}
	}
	betweenAnd struct {
		column Column
		op     string
		left   interface{}
		right  interface{}
	}
	notBetweenAnd struct {
		column Column
		op     string
		left   interface{}
		right  interface{}
	}
	in struct {
		column Column
		op     string
		values []interface{}
	}
	notIn struct {
		column Column
		op     string
		values []interface{}
	}
)

var (
	Conds            = func(conds ...Cond) []Cond { return conds }
	Eq               = func(c Column, v interface{}) Cond { return &eq{value: v, op: "AND", column: c} }
	NotEq            = func(c Column, v interface{}) Cond { return &notEq{value: v, op: "AND", column: c} }
	Like             = func(c Column, v interface{}) Cond { return &like{value: v, op: "AND", column: c} }
	NotLike          = func(c Column, v interface{}) Cond { return &notLike{value: v, op: "AND", column: c} }
	LeftLike         = func(c Column, v interface{}) Cond { return &leftLike{value: v, op: "AND", column: c} }
	NotLeftLike      = func(c Column, v interface{}) Cond { return &notLeftLike{value: v, op: "AND", column: c} }
	RightLike        = func(c Column, v interface{}) Cond { return &rightLike{value: v, op: "AND", column: c} }
	NotRightLike     = func(c Column, v interface{}) Cond { return &notRightLike{value: v, op: "AND", column: c} }
	Instr            = func(c Column, v interface{}) Cond { return &instr{value: v, op: "AND", column: c} }
	NotInstr         = func(c Column, v interface{}) Cond { return &notInstr{value: v, op: "AND", column: c} }
	Gt               = func(c Column, v interface{}) Cond { return &gt{value: v, op: "AND", column: c} }
	NotGt            = func(c Column, v interface{}) Cond { return &notGt{value: v, op: "AND", column: c} }
	GtEq             = func(c Column, v interface{}) Cond { return &gtEq{value: v, op: "AND", column: c} }
	NotGtEq          = func(c Column, v interface{}) Cond { return &notGtEq{value: v, op: "AND", column: c} }
	Lt               = func(c Column, v interface{}) Cond { return &lt{value: v, op: "AND", column: c} }
	NotLt            = func(c Column, v interface{}) Cond { return &notLt{value: v, op: "AND", column: c} }
	LtEq             = func(c Column, v interface{}) Cond { return &ltEq{value: v, op: "AND", column: c} }
	NotLtEq          = func(c Column, v interface{}) Cond { return &notLtEq{value: v, op: "AND", column: c} }
	BetweenAnd       = func(c Column, l, r interface{}) Cond { return &betweenAnd{left: l, right: r, op: "AND", column: c} }
	NotBetweenAnd    = func(c Column, l, r interface{}) Cond { return &notBetweenAnd{left: l, right: r, op: "AND", column: c} }
	In               = func(c Column, vs ...interface{}) Cond { return &in{values: vs, op: "AND", column: c} }
	NotIn            = func(c Column, vs ...interface{}) Cond { return &notIn{values: vs, op: "AND", column: c} }
	AndEq            = func(c Column, v interface{}) Cond { return &eq{value: v, op: "AND", column: c} }
	AndNotEq         = func(c Column, v interface{}) Cond { return &notEq{value: v, op: "AND", column: c} }
	AndLike          = func(c Column, v interface{}) Cond { return &like{value: v, op: "AND", column: c} }
	AndNotLike       = func(c Column, v interface{}) Cond { return &notLike{value: v, op: "AND", column: c} }
	AndLeftLike      = func(c Column, v interface{}) Cond { return &leftLike{value: v, op: "AND", column: c} }
	AndNotLeftLike   = func(c Column, v interface{}) Cond { return &notLeftLike{value: v, op: "AND", column: c} }
	AndRightLike     = func(c Column, v interface{}) Cond { return &rightLike{value: v, op: "AND", column: c} }
	AndNotRightLike  = func(c Column, v interface{}) Cond { return &notRightLike{value: v, op: "AND", column: c} }
	AndInstr         = func(c Column, v interface{}) Cond { return &instr{value: v, op: "AND", column: c} }
	AndNotInstr      = func(c Column, v interface{}) Cond { return &notInstr{value: v, op: "AND", column: c} }
	AndGt            = func(c Column, v interface{}) Cond { return &gt{value: v, op: "AND", column: c} }
	AndNotGt         = func(c Column, v interface{}) Cond { return &notGt{value: v, op: "AND", column: c} }
	AndGtEq          = func(c Column, v interface{}) Cond { return &gtEq{value: v, op: "AND", column: c} }
	AndNotGtEq       = func(c Column, v interface{}) Cond { return &notGtEq{value: v, op: "AND", column: c} }
	AndLt            = func(c Column, v interface{}) Cond { return &lt{value: v, op: "AND", column: c} }
	AndNotLt         = func(c Column, v interface{}) Cond { return &notLt{value: v, op: "AND", column: c} }
	AndLtEq          = func(c Column, v interface{}) Cond { return &ltEq{value: v, op: "AND", column: c} }
	AndNotLtEq       = func(c Column, v interface{}) Cond { return &notLtEq{value: v, op: "AND", column: c} }
	AndBetweenAnd    = func(c Column, l, r interface{}) Cond { return &betweenAnd{left: l, right: r, op: "AND", column: c} }
	AndNotBetweenAnd = func(c Column, l, r interface{}) Cond { return &notBetweenAnd{left: l, right: r, op: "AND", column: c} }
	AndIn            = func(c Column, vs ...interface{}) Cond { return &in{values: vs, op: "AND", column: c} }
	AndNotIn         = func(c Column, vs ...interface{}) Cond { return &notIn{values: vs, op: "AND", column: c} }
	OrEq             = func(c Column, v interface{}) Cond { return &eq{value: v, op: "OR", column: c} }
	OrNotEq          = func(c Column, v interface{}) Cond { return &notEq{value: v, op: "OR", column: c} }
	OrLike           = func(c Column, v interface{}) Cond { return &like{value: v, op: "OR", column: c} }
	OrNotLike        = func(c Column, v interface{}) Cond { return &notLike{value: v, op: "OR", column: c} }
	OrLeftLike       = func(c Column, v interface{}) Cond { return &leftLike{value: v, op: "OR", column: c} }
	OrNotLeftLike    = func(c Column, v interface{}) Cond { return &notLeftLike{value: v, op: "OR", column: c} }
	OrRightLike      = func(c Column, v interface{}) Cond { return &rightLike{value: v, op: "OR", column: c} }
	OrNotRightLike   = func(c Column, v interface{}) Cond { return &notRightLike{value: v, op: "OR", column: c} }
	OrInstr          = func(c Column, v interface{}) Cond { return &instr{value: v, op: "OR", column: c} }
	OrNotInstr       = func(c Column, v interface{}) Cond { return &notInstr{value: v, op: "OR", column: c} }
	OrGt             = func(c Column, v interface{}) Cond { return &gt{value: v, op: "OR", column: c} }
	OrNotGt          = func(c Column, v interface{}) Cond { return &notGt{value: v, op: "OR", column: c} }
	OrGtEq           = func(c Column, v interface{}) Cond { return &gtEq{value: v, op: "OR", column: c} }
	OrNotGtEq        = func(c Column, v interface{}) Cond { return &notGtEq{value: v, op: "OR", column: c} }
	OrLt             = func(c Column, v interface{}) Cond { return &lt{value: v, op: "OR", column: c} }
	OrNotLt          = func(c Column, v interface{}) Cond { return &notLt{value: v, op: "OR", column: c} }
	OrLtEq           = func(c Column, v interface{}) Cond { return &ltEq{value: v, op: "OR", column: c} }
	OrNotLtEq        = func(c Column, v interface{}) Cond { return &notLtEq{value: v, op: "OR", column: c} }
	OrBetweenAnd     = func(c Column, l, r interface{}) Cond { return &betweenAnd{left: l, right: r, op: "OR", column: c} }
	OrNotBetweenAnd  = func(c Column, l, r interface{}) Cond { return &notBetweenAnd{left: l, right: r, op: "OR", column: c} }
	OrIn             = func(c Column, vs ...interface{}) Cond { return &in{values: vs, op: "OR", column: c} }
	OrNotIn          = func(c Column, vs ...interface{}) Cond { return &notIn{values: vs, op: "OR", column: c} }
)

func (c *eq) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v = ?", c.op, c.column), []interface{}{c.value}
}

func (c *notEq) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s NOT t.%v = ?", c.op, c.column), []interface{}{c.value}
}

func (c *like) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v %s", c.op, c.column, `LIKE CONCAT('%', ?, '%')`), []interface{}{c.value}
}

func (c *notLike) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v %s", c.op, c.column, `NOT LIKE CONCAT('%', ?, '%')`), []interface{}{c.value}
}

func (c *leftLike) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v %s", c.op, c.column, `LIKE CONCAT('%', ?)`), []interface{}{c.value}
}

func (c *notLeftLike) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v %s", c.op, c.column, `NOT LIKE CONCAT('%', ?)`), []interface{}{c.value}
}

func (c *rightLike) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v %s", c.op, c.column, `LIKE CONCAT(?, '%')`), []interface{}{c.value}
}

func (c *notRightLike) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v %s", c.op, c.column, `NOT LIKE CONCAT(?, '%')`), []interface{}{c.value}
}

func (c *instr) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s INSTR(t.%v, ?) > 0", c.op, c.column), []interface{}{c.value}
}

func (c *notInstr) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s NOT INSTR(t.%v, ?) > 0", c.op, c.column), []interface{}{c.value}
}

func (c *gt) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v > ?", c.op, c.column), []interface{}{c.value}
}

func (c *notGt) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s NOT t.%v > ?", c.op, c.column), []interface{}{c.value}
}

func (c *gtEq) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v >= ?", c.op, c.column), []interface{}{c.value}
}

func (c *notGtEq) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s NOT t.%v >= ?", c.op, c.column), []interface{}{c.value}
}

func (c *lt) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v < ?", c.op, c.column), []interface{}{c.value}
}

func (c *notLt) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s NOT t.%v < ?", c.op, c.column), []interface{}{c.value}
}

func (c *ltEq) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v <= ?", c.op, c.column), []interface{}{c.value}
}

func (c *notLtEq) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s NOT t.%v <= ?", c.op, c.column), []interface{}{c.value}
}

func (c *betweenAnd) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s (t.%v BETWEEN ? AND ?)", c.op, c.column), []interface{}{c.left, c.right}
}

func (c *notBetweenAnd) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s (NOT t.%v BETWEEN ? AND ?)", c.op, c.column), []interface{}{c.left, c.right}
}

func (c *in) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s t.%v IN (%s)", c.op, c.column, strings.TrimLeft(strings.Repeat(",?", len(c.values)), ",")), c.values
}

func (c *notIn) SQL() (string, []interface{}) {
	return fmt.Sprintf(" %s NOT t.%v IN (%s)", c.op, c.column, strings.TrimLeft(strings.Repeat(",?", len(c.values)), ",")), c.values
}
