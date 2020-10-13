package main

import (
	"fmt"
	"github.com/sqrtcat/easy_go/lib"
	"github.com/startdusk/goutil/hash"
)

func Hello() {
	lib.Hello()
}

func Md5() {
	fmt.Println(hash.StringMd5("hello go"))
}
