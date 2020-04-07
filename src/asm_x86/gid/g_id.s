#include "textflag.h"
#include "g_tls.h"

//func get_gid()int64
//get_tls(r) MOVQ TLS,r
//g(r) 0(r)(TLS*1)
TEXT ·Get_gid(SB),NOSPLIT,$0
    get_tls(CX)
    MOVQ g(CX), AX//move g to AX
    //#inlude "go_asm.h"
    //MOVQ g_m(AX),BX//move g.m into BX
    //MOVQ ·offset(SB), BX
    MOVQ $152, BX
    LEAQ 0(AX)(BX*1), DX

    //LEAQ 0+152(AX),BX
    //MOVQ (BX),CX
    //MOVQ CX,ret1+8(FP)
    //MOVQ $0x02,ret1+8(FP)

    MOVQ (DX), AX
    MOVQ AX, ret+0(FP)
    //MOVQ DX,ret1+8(FP)
    RET
