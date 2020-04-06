#include "textflag.h"

//func Output(int)(int,int,int)
TEXT ·Output(SB),$8-48
    MOVQ 24(SP),DX
    //MOVQ 16(SP),DX
    MOVQ DX,ret3+24(FP)
    MOVQ arg0+16(SP),BX
    MOVQ BX,ret2+16(FP)
    MOVQ arg0+0(FP),AX
    MOVQ AX,ret1+8(FP)
    RET

// func output(a,b int) int
TEXT ·Output2(SB), NOSPLIT, $24-8
    //MOVQ a+0(FP), DX // arg a
    MOVQ a+16(SP), DX // arg a
    //MOVQ DX, 0(SP) // arg x
    MOVQ DX, argx-24(SP) // arg x
    MOVQ b+8(FP), CX // arg b
    MOVQ CX, 8(SP) // arg y
    CALL ·Add_x86(SB) // 在调用 add 之前，已经把参数都通过物理寄存器 SP 搬到了函数的栈顶
    MOVQ 16(SP), AX // add 函数会把返回值放在这个位置
    MOVQ AX, ret+16(FP) // return result
    RET
TEXT ·Add_x86(SB),NOSPLIT,$0-24
    MOVQ a+0(FP),AX
    MOVQ b+8(FP),BX
    ADDQ BX,AX
    MOVQ AX,ret+16(FP)
    RET

//func ArraySum(arr []int64)int64
TEXT ·ArraySum(SB),NOSPLIT,$0-32
    MOVQ $0,SI
    MOVQ arr0+0(FP),BX
    MOVQ len+8(FP),CX
    INCQ CX
start:
    DECQ CX
    JZ done
    ADDQ (BX),SI
    ADDQ $8,BX
    JMP start
done:
     // 返回地址是 24 是怎么得来的呢？
        // 可以通过 go tool compile -S math.go 得知
        // 在调用 sum 函数时，会传入三个值，分别为:
        // slice 的首地址、slice 的 len， slice 的 cap
        // 不过我们这里的求和只需要 len，但 cap 依然会占用参数的空间
        // 就是 16(FP)
    MOVQ SI,ret+24(FP)
    RET
