bits 16
org 31744

.loop	lodsb
or al, al
jz halt
int 0x10
jmp .loop
msg: db "Hello BIOS my name is Mohammad Hosein Balkhani create of Gootloader.", 0

.loop	lodsb
or al, al
jz halt
int 0x10
jmp .loop
msg: db "Hello BIOS my name is Mohammad Hosein Balkhani create of Gootloader2.", 0

halt: hlt
times 510 - ($ - $$) db 0
dw 0xAA55