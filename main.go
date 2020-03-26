package main

import (
	"fmt"

	f "github.com/idawud/gobyexample/funcs"
	ms "github.com/idawud/gobyexample/morestring"
)

func main() {
	fmt.Println(ms.ReverseRunes("Pack"))
	fmt.Println(f.Vals())
	cl := f.Closures()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())
}
