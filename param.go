package gobatis

import "fmt"

var paramPrefix = "@"
var paramSuffix = "@"

//Define namedParam struct
type Param struct {
	name    string //param name
	replace string //replace name
	val     string //param val
}

//New Param
func NewParam(name string, val interface{}) *Param {
	if name == "" {
		return nil
	}
	return &Param{
		name:    name,
		replace: paramPrefix + name + paramSuffix,
		val:     fmt.Sprintf("%v", val),
	}
}
