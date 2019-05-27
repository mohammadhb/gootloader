# gootloader
Create bootloaders with Go.


## Usage

After you create your bootloader using `.create()`
You can use

```bash
nasm boot.asm -f bin -o boot.bin && qemu-system-i386 -fda boot.bin
```

> Caution : You must install nasm and qemu in your device to be able to emulate your bootloader (qemu-system-i386 is x86 arch in this example)
