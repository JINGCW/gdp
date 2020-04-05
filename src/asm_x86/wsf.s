INCLUDE Irvine32.inc
EXTERN WriteStackFrame@12:PROC

main PROC
    mov eax,0EAEAEAEAH
    MOV EBX,0EBEBEBEBH
    INVOKE myproc,111h,222h
    exit
main ENDP

myproc PROC USES eax,ebx,
    x:DWORD,y:DWORD
    LOCAL a:DWORD,b:DWORD

    PARAMS=2
    LOCALS=2
    SAVED_REGS=2
    mov a,0AAAAH
    mov b,0BBBBH
    INVOKE WriteStackFrame,PARAMS,LOCALS,SAVED_REGS
