package main

import "fmt"
import gootloader "./gootloader"

func main() {

	bl := Bootloader{
		name:     "myBootloader",
		arch:     1,
		modeBit:  16,
		loaddest: 0x7C00,
	}

	bl.getName("Bootloader name is : ")
	bl.print("Hello BIOS my name is Mohammad Hosein Balkhani create of Gootloader.")
	bl.print("Hello BIOS my name is Mohammad Hosein Balkhani create of Gootloader2.")
	cInst := bl.create()

	fmt.Println(cInst)

	ioutil.WriteFile("./"+bl.name+".asm", []byte(cInst), 0644)

}
