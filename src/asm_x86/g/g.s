#include "textflag.h"
#include "g_tls.h"

DATA moffset+0(SB)/8,$48
GLOBL moffset(SB),RODATA,$8

DATA goffset_v1_14_1+0(SB)/8,$152
GLOBL goffset_v1_14_1(SB),RODATA,$8

TEXT ·GPTR(SB),NOSPLIT,$8-0
    //*g.
    //#define get_tls(r) MOVQ TLS,r
    //#define g(r) 0(r)(TLS*1)
    get_tls(CX)
    MOVQ g(CX), AX//move g to AX
    MOVQ AX, gID+0(FP)
    RET

TEXT ·GID(SB),NOSPLIT,$8
    CALL ·GPTR(SB)
    //get_tls(CX)
    //MOVQ g(CX), AX//move g to AX
    MOVQ 0(SP),AX //g.
    //MOVQ ·_offset_gid(SB),CX//ref _offset_gid mem var
    MOVQ goffset_v1_14_1(SB),CX
    LEAQ 0(AX)(CX*1),BX
    //MOVQ 0(AX)(CX*1),BX
    MOVQ (BX),AX
    MOVQ AX,ret+0(FP)
    RET

TEXT ·goRoutinePtr(SB),NOSPLIT,$0-8
	get_tls(CX)
	MOVQ g(CX), AX
	MOVQ AX, goid+0(FP)
	RET
