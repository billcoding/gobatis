package gobatis

import (
	"fmt"
	"testing"
)

func TestPage(t *testing.T) {
	type XUser struct {
		Id   int    `db:"id"`
		Name string `db:"name"`
	}

	p := Default().Init().DSN(dsn).Mapper("user").Select("queryPage").Page(new(XUser), 1, 3)
	fmt.Println(p.TotalRows)
	fmt.Println(p.TotalPage)
	for _, l := range p.List {
		fmt.Println(l.(*XUser).Name)
	}
}
