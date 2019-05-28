package gootloader

import "C"
import "fmt"

// import "io/ioutil"
import "strconv"

//Bootloader data structure
type Bootloader struct {
	Name         string
	Arch         int // Architecture 32 or 64
	ModeBit      int // Workspace mode in bits 16 or 32
	Instructions []string
	Loaddest     int
}

func (r *Bootloader) getName(title string) int {
	println(title + r.Name)
	return 1
}
func (r *Bootloader) addInstruction(instruction string) int {

	r.Instructions = append(r.Instructions, instruction)
	return 1

}
func (r *Bootloader) create() string {

	fmt.Println("org " + strconv.Itoa(r.Loaddest))
	r.Instructions = append([]string{"org " + strconv.Itoa(r.Loaddest) + "\n"}, r.Instructions...)
	r.Instructions = append([]string{"bits " + strconv.Itoa(r.ModeBit) + "\n"}, r.Instructions...)

	r.addInstruction("halt: hlt\n")

	r.addInstruction("times 510 - ($ - $$) db 0\n")
	r.addInstruction("dw 0xAA55")

	cInst := ""

	for _, element := range r.Instructions {

		cInst += element

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
