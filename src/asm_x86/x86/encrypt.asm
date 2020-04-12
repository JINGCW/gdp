.model small,stdcall
.stack 200h
.386

XORVAL=239

.code
main PROC
    mov ax,@data
    mov ds,ax
l1:
    mov ah,6
    mov dl,0ffh
    int 21h
    jz l2
    xor al,XORVAL
    mov ah,6
    mov dl,al
    int 21h
;    jmp l1
l2:
    mov ah, 4cH
    int 21h
;.exit not work
;    mov ah,4ch
;    mov al,0
;    int 21h
;    mov ah,0
;    int 21h
main ENDP
END main
