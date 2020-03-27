package pointers

//PtrFunc func with pointers
func PtrFunc(valptr *int) int {
	*valptr = 22
	return *valptr
}
