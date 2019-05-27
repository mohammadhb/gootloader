# Gootloader
Create bootloaders with Go.

## Getting Started

Start with creating a Bootloader with your config :
```go

bl := Bootloader{
    name:     "myBootloader",
    arch:     1,      // Bootloader Architecture
    modeBit:  16,     // Bootloader Bit Mode
    loaddest: 0x7C00, // Destination of loading of bootloader in Memory
}
  
```

Maybe Print a Message :
```go

message := "Dear BIOS, Load me into ram and give the control flow to me please"
bl.print(message)

```
Create Your Bootloader :
> This generates `bl.name`.asm file
```go

bl.create();//AKA final file for bootloader

```

## Emulate and Testings the Bootloader

After you create your bootloader using `.create()`
You can use

```bash
nasm boot.asm -f bin -o boot.bin && qemu-system-i386 -fda boot.bin
```

> Caution : You must install nasm and qemu in your device to be able to emulate your bootloader (qemu-system-i386 is x86 arch in this example)
