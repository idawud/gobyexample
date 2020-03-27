package main

import (
	intface "github.com/idawud/gobyexample/intface"
	// ptr "github.com/idawud/gobyexample/pointers"
)

func main() {
	/*fmt.Println(ms.ReverseRunes("Pack"))
	fmt.Println(f.Vals())
	cl := f.Closures()
	fmt.Println(cl())
	fmt.Println(cl())

	fmt.Println(f.Plus([]int{1,2}))
	rect := struct.Rectangle{2.7, 6.6}
	fmt.Println(rect.Area())

	i := 2
	fmt.Println("Before: ", i)
	fmt.Println(ptr.PtrFunc(&i))
	fmt.Println("After: ", i)

	dawud := struc.NewPerson("dawud", 24)
	fmt.Println(dawud.GetDetails())
	*/

	r := intface.NewRect(2, 3)
	c := intface.NewCircle(2.2)
	intface.Measure(r)
	intface.Measure(c)
}
