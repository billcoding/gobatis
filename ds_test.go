package gobatis

import (
	"fmt"
	"testing"
)

func TestDs(t *testing.T) {
	b := Default()
	b.Init()
	m := b.Mapper("user")
	time := ""
	m.Select("selectSimple").Exec().Single(&time)
	fmt.Println(time)
}
