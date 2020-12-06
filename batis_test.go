package gobatis

import (
	"encoding/json"
	"fmt"
	"testing"
)

//A DDL sql for this testing

/**
CREATE DATABASE test;

USE test;

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

*/

const dsn = "root:123@tcp(127.0.0.1:3306)/test"

func TestInsert(t *testing.T) {
	var b = Default().Init().DSN(dsn)
	b.Config.PrintSql = true
	userMapper := b.Mapper("user")
	fmt.Println(userMapper.Update("insert").Exec("bdsfsdfill"))
}

func TestDelete(t *testing.T) {
	var batis = Default().Init().DSN(dsn)
	fmt.Println(batis.Mapper("user").Update("delete").Exec(3))
}

func TestUpdate(t *testing.T) {
	var batis = Default().Init().DSN(dsn)
	fmt.Println(batis.Mapper("user").Update("update").Exec("updated", 6))
}

func TestSelectSimple(t *testing.T) {
	var batis = Default().Init().DSN(dsn)
	userMapper := batis.Mapper("user")
	var time string
	userMapper.Select("selectSimple").Exec().Scan(&time)
	fmt.Println("time =", time)
}

func TestSelectStruct(t *testing.T) {
	var batis = Default().Init().DSNWithDialect(SQLite, "./user.db")
	type User struct {
		Id   int    `db:"id" json:"id"`
		Name string `db:"name" json:"name"`
	}
	list := batis.Mapper("user").Select("selectStruct").Exec().List(new(User))
	bytes, _ := json.Marshal(list)
	fmt.Println(string(bytes))
}

func TestTx(t *testing.T) {
	var mapper = Default().Init().DSN(dsn).TxMapper("user")
	mapper.Update("insert").Exec("zhangsan")
	mapper.Update("insert").Exec("lisi")
	mapper.Update("insert").Exec("wangwu")
	mapper.Update("delete").Exec(1)
	mapper.Update("delete").Exec(2)
	mapper.Update("update").Exec("updated", 3)
	mapper.Commit()
}
