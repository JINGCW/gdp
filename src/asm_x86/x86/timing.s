include Irvine32.inc
includelib Irvine32.lib
includelib Kernel32.lib
includelib User32.lib

.data
OUTER_LOOP_COUNT=3
start_time DWORD ?
msg1 BYTE "please wait...",0dh,0ah,0
msg2 BYTE "elapsed milliseconds: ",0

.code
main PROC
    mov edx,offset msg1
    call WriteString
    ;save start time
    call GetMSeconds
    mov start_time,eax
    ;start outer loop
    mov ecx,OUTER_LOOP_COUNT
L1:call inner_loop
    loop L1
    call GetMSeconds
    ;calculate executive time
    sub eax,start_time
    ;show executive time
    mov edx,offset msg2
    call WriteString
    call WriteDec
    call Crlf
    exit

inner_loop PROC
    push ecx
    mov ecx,0FFFFFFFh
L1: mul eax
    mul eax
    mul eax
    loop L1
    pop ecx
    ret
inner_loop ENDP
main ENDP

END main