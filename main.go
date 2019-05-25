package main

//#include "bootloader/inc.c"
import "C"

//Bootloader data structure
type Bootloader struct {
	name             string
	memstart, memend int
	arch             int
	modeBit          int // Workspace mode in bits 16 or 32

}

func (r *Bootloader) getName(title string) int {
	println(title + r.name)
	return 1
}

func (r *Bootloader) create(title string) int {
	println(title + r.name)
	return 1
}

func main() {

	bl := Bootloader{
		name:     "bl1",
		memstart: 1,
		memend:   2,
		arch:     1,
		modeBit:  16,
	}

	bl.getName("Bootloader name is : ")

	f := C.print()
	println(f)
	// Output: 42
}
