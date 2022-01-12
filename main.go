package main

import (
	"log"

	"github.com/casbin/casbin/v2"
	entadapter "github.com/casbin/ent-adapter"
)

func main() {
	a, err := entadapter.NewAdapter("postgres", "host=127.0.0.1 user=postgres password=123456 dbname=casbin port=5432 sslmode=disable TimeZone=UTC")
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer("./model.conf", a)
	if err != nil {
		panic(err)
	}
	sub := "alice" // 想要访问资源的用户。
	obj := "data2" // 将被访问的资源。
	act := "read"  // 用户对资源执行的操作。

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		panic(err)
	}

	if ok == true {
		log.Println("alice can read data2")
		// 允许alice读取data1
	} else {
		log.Println("alice can not read data2")
		// 拒绝请求，抛出异常
	}

}
