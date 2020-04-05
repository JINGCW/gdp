#include "textflag.h"
#include "g_tls.h"

//func get_gid()int64
//get_tls(r) MOVQ TLS,r
//g(r) 0(r)(TLS*1)
TEXT ·Get_gid(SB),NOSPLIT,$0-8
    get_tls(CX)
    MOVQ g(CX), AX
    MOVQ ·offset(SB), BX
    LEAQ 0(AX)(BX*1), DX
    MOVQ (DX), AX
    MOVQ AX, ret+0(FP)
    RET
