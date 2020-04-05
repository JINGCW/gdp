
.code
Fac PROC
    move ecx,5
    move eax,0
    call Facs
Fac ENDP

Facs PROC
    push ebp
    mov ebp,esp
    mov eax,[ebp+8]
    cmp eax,0
    ja L1
    mov eax,1
    jmp L2
L1:
    dec eax
    push eax
    call Facs
ReturnFact:
    mov ebx,[ebp+8]
    mul ebx
L2:
    pop ebp
    ret 4
Facs ENDP
