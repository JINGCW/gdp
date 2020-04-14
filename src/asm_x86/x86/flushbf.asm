
;.model small,stdcall
include flushbf.inc
.model small,stdcall


.code
FlushBuffer PROC
;flush the standard input buffer.
;receives: nothing. returns: nothing
.data
onebyte db ?
.code
;    pusha
    push ax
    push bx
    push cx
    push dx
    push di
l1:
    mov ah,3fh ;read file/device
    mov bx,0 ;keyboard handle
    mov cx,1 ;one byte
    mov dx,offset onebyte ;save it here
    int 21h ;call MS-DOS
    cmp onebyte,0ah ;end of lien yet?
    jne l1 ;no: read another

;    popa

    pop ax
    pop bx
    pop cx
    pop dx
    pop di

    ret
FlushBuffer ENDP
END