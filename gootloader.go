package gootloader

import "C"
import "fmt"
// import "io/ioutil"
import "strconv"

//Bootloader data structure
type Bootloader struct {
	name             string
	arch             int // Architecture 32 or 64
	modeBit          int // Workspace mode in bits 16 or 32
	instructions     []string
	loaddest 		 int
}

func (r *Bootloader) getName(title string) int {
	println(title + r.name)
	return 1
}
func (r *Bootloader) addInstruction(instruction string) int {

	r.instructions = append(r.instructions, instruction)
	return 1

}
func (r *Bootloader) create() string {
	
	fmt.Println("org "+strconv.Itoa(r.loaddest))
	r.instructions = append([]string{"org "+strconv.Itoa(r.loaddest)+"\n"},r.instructions...)
	r.instructions = append([]string{"bits "+strconv.Itoa(r.modeBit)+"\n"},r.instructions...)
   
	r.addInstruction("halt: hlt\n")

	r.addInstruction("times 510 - ($ - $$) db 0\n")
	r.addInstruction("dw 0xAA55")

	cInst := ""

	for _,element := range r.instructions {

		cInst+=element

	}

	return cInst

}

func (r *Bootloader) print(message string) int {

	var loopIns = [...]string{
		".loop	lodsb\n",
		"or al, al\n",
		"jz halt\n",
		"int 0x10\n",
		"jmp .loop\n",
	}

	for _, element := range loopIns {

		r.addInstruction(element)

	}

	r.addInstruction("msg: db \"" + message + "\", 0\n")
	

	return 1

}