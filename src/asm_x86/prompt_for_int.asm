INCLUDE Irvine32.inc

.code
prompt_for_ints PROC
    arraysize EQU [ebp+16]
    ptrarray EQU [ebp+12]
    ptrprompt EQU [ebp+8]
    enter 0,0
    pushad
    mov ecx,arraysize
    cmp ecx,0
    jle L2
    mov edx,ptrprompt
    mov esi,ptrarray
L1:
    call WriteString
    call ReadInt
    call Crlf
    mov [esi],eax
    add esi,4
    loop L1
L2:
    popad
    leave
    ret 12
prompt_for_ints ENDP
END
