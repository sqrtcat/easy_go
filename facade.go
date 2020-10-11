package easy_go

import (
	"fmt"
	"github.com/sqrtcat/easy_go/lib"
	"github.com/startdusk/goutil/hash"
)

func Hello() {
	lib.Hello()
}

func Md5() {
	fmt.Println("md5tips: " + hash.StringMd5("hello go"))
}
