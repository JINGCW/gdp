include Irvine32.inc
includelib Irvine32.lib
includelib Kernel32.lib
includelib User32.lib

.data
HEAP_START=2000000;2MB
HEAP_MAX=400000000;400MB
BLOCK_SIZE=500000;0.5MB
hHeap HANDLE ?;heap pointer
pData DWORD ?;data block pointer

str1 byte 0dh,0ah,"memory allcation failed",0dh,0ah,0

.code
main PROC
    invoke HeapCreate,0,HEAP_START,HEAP_MAX
    .if eax==0
    call WriteWindowsMsg
    call Crlf
    jmp quit
    .else
    mov hHeap,eax
    .endif
    mov ecx,2000
l1:
    call allocate_block
    .if Carry?
    mov edx,offset str1
    call WriteString
    jmp quit
    .else
    mov al,'.'
    call WriteChar
    .endif
    ;call free_block
    loop l1
quit:
    invoke HeapDestroy,hHeap
    .if eax==NULL
    call WriteWindowsMsg
    call Crlf
    .endif
    invoke ExitProcess,0
main ENDP

allocate_block PROC uses ecx
    invoke HeapAlloc,hHeap,HEAP_ZERO_MEMORY,BLOCK_SIZE
    .if eax==NULL
        stc;CF=1
    .else
        mov pData,eax
        clc;CF=0
    .endif
    ret
allocate_block ENDP

free_block PROC uses ecx
    invoke HeapFree,hHeap,0,pData
    ret
free_block ENDP

END main
