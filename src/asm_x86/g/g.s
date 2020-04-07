#include "textflag.h"
#include "g_tls.h"

TEXT ·Gstruct(SB),NOSPLIT,$0
    //#define get_tls(r) MOVQ TLS,r
    //#define g(r) 0(r)(TLS*1)
    get_tls(CX)
    MOVQ g(CX), AX//move g to AX

TEXT ·goRoutinePtr(SB),NOSPLIT,$0-8
	get_tls(CX)
	MOVQ g(CX), AX
	MOVQ AX, goid+0(FP)
	RET
