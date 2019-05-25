enum Bitmode {

    Bit16=16,
    Bit32=32

};


void generate_header(enum Bitmode bitmode=Bit16,int bl_program_dest_address=0x7C00){
                      
    asm
    (
        "org %0\n"
        "bits %1\n"
        "times 510 - ($ - $$) db 0\n"
        "dw 0xAA55"
        :
        : "r" ((int)bitmode),"r" (bl_program_dest_address)
    );

}

int main(){

    enum Bitmode bitmode = Bit16;
    generate_header(bitmode);

}