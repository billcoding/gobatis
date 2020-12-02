package gobatis

import "fmt"

var namedParamPrefix = "@"
var namedParamSuffix = "@"

//Define namedParam struct
type NamedParam struct {
	name    string //param name
	replace string //replace name
	val     string //param val
}

//New namedParam
func NewNamedParam(name string, val interface{}) *NamedParam {
	if name == "" {
		batis.Error("namedParam 's name is empty")
		return nil
	}
	return &NamedParam{
		name:    name,
		replace: namedParamPrefix + name + namedParamSuffix,
		val:     fmt.Sprintf("%v", val),
	}
}
