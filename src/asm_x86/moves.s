.data
array DWORD 10000h,20000h,30000h
thesum DWORD ?

.code
main PROC
    mov esi,OFFSET array
    mov ecx,LENGTHOF array
    call array_sum
    mov thesum,eax
main ENDP

array_sum PROC
    push esi
    push ecx
    mov eax,0
L1:
    add eax,[esi]
    add esi,TYPE DWORD
    loop L1
    pop ecx
    pop esi
    ret
array_sum ENDP
END main
