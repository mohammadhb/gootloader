package main

//#include "bootloader/inc.c"
import "C"

//Bootloader data structure
type Bootloader struct {
	name             string
	memstart, memend int
}

func main() {
	f := C.print()
	println(f)
	// Output: 42
}
