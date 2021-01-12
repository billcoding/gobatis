package gobatis

// A xml mapper demo file
//
//
//<?xml version="1.0" encoding="UTF-8"?>
//<batis-mapper binding="user">
//    <update id="update">
//        update 1
//    </update>
//    <select id="select">
//        select 1
//    </select>
//</batis-mapper>

import "encoding/xml"

type mapperNode struct {
	XMLName           xml.Name           `xml:"batis-mapper"`
	Binding           string             `xml:"binding,attr"`
	MapperUpdateNodes []mapperUpdateNode `xml:"update"`
	MapperSelectNodes []mapperSelectNode `xml:"select"`
}

type mapperUpdateNode struct {
	XMLName xml.Name `xml:"update"`
	Id      string   `xml:"id,attr"`
	Text    string   `xml:",cdata"`
}

type mapperSelectNode struct {
	XMLName xml.Name `xml:"select"`
	Id      string   `xml:"id,attr"`
	Text    string   `xml:",cdata"`
}
