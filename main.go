package main

import (
	"fmt"

	f "github.com/idawud/gobyexample/funcs"
	// ms "github.com/idawud/gobyexample/morestring"
	shapes "github.com/idawud/gobyexample/structs"
)

func main() {
	/*fmt.Println(ms.ReverseRunes("Pack"))
	fmt.Println(f.Vals())
	cl := f.Closures()
	fmt.Println(cl())
	fmt.Println(cl())
	*/
	fmt.Println(f.Plus([]int{1,2}))
	rect := shapes.Rectangle{2.7, 6.6}
	fmt.Println(rect.Area())
}
