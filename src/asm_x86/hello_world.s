#include "textflag.h"

DATA text<>+0(SB)/8,$"hello wo"
DATA text<>+8(SB)/8,$"rld!\n"
GLOBAL text<>(SB),NOPTR,$16

TEXT ·PrintHelloWorld(SB),$56-0
    NO_LOCAL_POINTERS
    MOVQ $1,fd-56(SB)
    MOVQ $text<>+0(SB),AX
    MOVQ AX,ptr-48(SP)
    MOVQ $13,len-40(SP)
    MOVQ $16,cap-32(SP)
    CALL syscall.Write(SB)
    RET
