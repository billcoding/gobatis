package gobatis

import (
	"fmt"
	"testing"
)

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func (u *User) Show() {
	fmt.Println(fmt.Sprintf("%d:%s", u.Id, u.Name))
}

func TestTemplateSelect(t *testing.T) {
	batis := Default().RegisterDS("_", dsn).ShowSql(true).Init()
	users := batis.Mapper("user2").Select("query").Prepare(map[string]interface{}{
		"id":   "",
		"name": "",
	}).Exec().List(new(User))

	for _, user := range users {
		user.(*User).Show()
	}
}

func TestTemplateInsert(t *testing.T) {
	batis := Default().RegisterDS("_", dsn).ShowSql(true).Init()
	users := make([]*User, 5)
	users[0] = &User{Name: "张三"}
	users[1] = &User{Name: "李四"}
	users[2] = &User{Name: "王五"}
	users[3] = &User{Name: "赵六"}
	users[4] = &User{Name: "田七"}
	batis.Mapper("user2").Update("insert").Prepare(users).Exec()
}
