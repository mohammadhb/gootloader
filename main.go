package main

import "C"

//Bootloader data structure
type Bootloader struct {
	name             string
	memstart, memend int
	arch             int // Architecture 32 or 64
	modeBit          int // Workspace mode in bits 16 or 32
	instructions     []string
}

func (r *Bootloader) getName(title string) int {
	println(title + r.name)
	return 1
}

func (r *Bootloader) addInstruction(instruction string) int {

	r.instructions = append(r.instructions, instruction)
	return 1

}

func (r *Bootloader) create() int {

	r.addInstruction("times 510 - ($ - $$) db 0")
	r.addInstruction("dw 0xAA55")

	return 1

}

func (r *Bootloader) print(message string) int {

	var loopIns = [...]string{
		".loop	lodsb",
		"or al, al",
		"jz halt",
		"int 0x10",
		"jmp .loop",
	}

	for _, element := range loopIns {

		r.addInstruction(element)

	}

	r.addInstruction("msg: db \"" + message + "\", 0")
	r.addInstruction("halt: hlt")

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
	bl.print("Hello World")
	bl.create()

	// f := C.print()
	// println(f)
	// Output: 42

}
