include Irvine32.inc
includelib Irvine32.lib
includelib Kernel32.lib
includelib User32.lib

.data
array_size=1000
fill_val EQU 0ffh

hHeap HANDLE ?
parray dword ?
newHeap DWORD ?
str1 BYTE "heap size is: ",0

.code
main PROC
    invoke GetProcessHeap
    .if eax==NULL
    call WriteWindowsMsg
    jmp quit
    .else
    mov hHeap,eax
    .endif

    call allocate_array
    jnc arrayOk ;faield (CF=1)?
    call WriteWindowsMsg
    call Crlf
    jmp quit
arrayOk:
    call fill_array
    call display_array
    call Crlf
    ;free the array
    invoke HeapFree,hHeap,0,parray
quit:
    INVOKE ExitProcess,0
main ENDP

allocate_array PROC USES eax
;dynamiclly allocates space for the array.
;receives:eax=handle to the program heap
;returns: CF=0 if the memory allocation succeeds
    invoke HeapAlloc,hHeap,HEAP_ZERO_MEMORY,array_size
    .if eax==NULL
        stc ;return with CF=1
    .else
        mov parray,eax
        clc ;return with CF=0
    .endif
    ret
allocate_array ENDP

fill_array PROC USES ecx edx esi
;fills all array position with a single character.
;receives:nothing
;returns:nothing
    mov ecx,array_size;loop counter
    mov esi,parray ;pointer to the array
l1: mov byte ptr [esi],fill_val
    inc esi
    loop l1
    ret
fill_array ENDP

display_array PROC USES eax ebx ecx esi
    mov ecx,array_size
    mov esi,parray
l1: mov al,[esi];get a byte
    mov ebx,type byte
    call WriteHexB;display it
    inc esi;next location
    loop l1
    ret
display_array ENDP
END main
