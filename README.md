# gobatis

## Introduction

A sample db tools with binding between sqls and xml nodes, similar like mybatis style.

## Features

- Pure native
- No third dependencies
- Multiple DataSource supports
- Multiple Database supports
- Dynamic SQL supports & Func (Using text/template)

## Supports database

1. MySQL(Default dialect)
2. SQLite/SQLite3
3. MSSQL/SQLServer

## Install

- GOPATH

```
mkdir -p $GOPATH/src/github.com/billcoding/gobatis

cd $GOPATH/src/github.com/billcoding

git clone https://github.com/billcoding/gobatis.git gobatis
```

- Go mod

```
require github.com/billcoding/gobatis latest
```

## Usage

- Insert

```
userMapper := Default().Init().DSN(dsn).Mapper("user")
err := userMapper.Update("insert").Exec("inserted")
```

- Delete

```
userMapper := Default().Init().DSN(dsn).Mapper("user")
err := userMapper.Update("delete").Exec(1)
```

- Update

```
userMapper := Default().Init().DSN(dsn).Mapper("user")
err := userMapper.Update("update").Exec("updated", 1)
```

- Select Simple

```
var batis = Default().Init().DSN(dsn)
userMapper := batis.Mapper("user")
userMapper.Select("selectSimple").Exec().Call(func(rows *sql.Rows) {
    if rows.Next() {
        t := ""
        rows.Scan(&t)
        fmt.Printf("time is %v\n", t)
    }
})

```

- Select Struct

```
var batis = Default().Init().DSN(dsn)
type User struct {
    Id   int    `db:"id"`
    Name string `db:"name"`
}
userList := batis.Mapper("user").Select("selectStruct").Exec().List(new(User))
})
```

## XML Definition

- `CUD` defines `Update` node

```xml
<?xml version="1.0" encoding="UTF-8"?>
<batis-mapper binding="user">

    <update id="insert">
        insert into user(`name`) values (?)
    </update>

    <update id="delete">
        delete from user where id = ?
    </update>

    <update id="update">
        update user set name = ? where id = ?
    </update>

</batis-mapper>
```

- `R` defines `Select` node

```xml
<?xml version="1.0" encoding="UTF-8"?>
<batis-mapper binding="user">

    <select id="select">
        select * from user
    </update>

</batis-mapper>
```

- Template defines

```xml
<?xml version="1.0" encoding="UTF-8"?>
<batis-mapper binding="user2">

    <select id="query">
        select * from user as u where 1 = 1
        {{ if ne .id "" }}
            and u.id = {{ .id }}
        {{ end }}
        {{ if ne .name "" }}
            and u.name = '{{ .name }}'
        {{ end }}
    </select>

    <update id="insert">
       insert into user(name) values
        {{ range $index, $element := . }}{{ if gt $index 0 }},{{ end }} ('{{$element.Name}}'){{ end }}
    </update>

</batis-mapper>
```

## Multiple DataSource Support

- Register DataSource

```
batis.RegisterDS(DSNAME, DSN)
```

- Select DataSource

```
mapper.SelectDS(DSNAME)
```

## Transaction Support

- Begin tx

```
batis.TxMapper(BINDING)
```

- Commit tx

```
txMapper.Commit()
```

- Rollback tx

```
txMapper.Rollback()
```

## Env Support

- Show SQL

```
BATIS_SHOW_SQL

e.g.

BATIS_SHOW_SQL=1|0
BATIS_SHOW_SQL=on|ON
BATIS_SHOW_SQL=true|TRUE
```

- Mapper path

```
BATIS_MAPPER_PATH

e.g.

BATIS_MAPPER_PATH=/tmp/myapp/mapper
```

- Dsn

```
BATIS_DSN

e.g.

1. BATIS_DSN=root:123@tcp(192.168.1.8:3306)/test
2. BATIS_DSN=_,root:123@tcp(192.168.1.8:3306)/test
3. BATIS_DSN=master,root:123@tcp(192.168.1.8:3306)/test|slave,root:123@tcp(192.168.1.9:3306)/test
```
