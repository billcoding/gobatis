package gobatis

import (
	"testing"
)

func TestLog(t *testing.T) {
	b := Default().Init().RegisterDS("_", dsn)
	b.Mapper("user3").Select("select").Exec()
	b.Mapper("user3").Update("update").Exec()
}
