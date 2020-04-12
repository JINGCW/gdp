
.model small,stdcall
.stack 200h

.data
buffer db 80 DUP(?)
string db "This is a string$"
message db "hello, world$"

.code
start:
    mov ax,@data
    mov ds,ax
    ;mov ax,SEG buffer
    ;mov ds,ax
    mov ah,2
    mov dl,'*'
    int 21h
    call func9
    ;.exit
ret_func9:
;    mov ah,4ch
    call func_40h
ret_func_40h:
    mov ah,4ch
    mov al,0
    int 21h
;.exit
func9:
    mov ah,9
    mov dx,offset string
    int 21h
    jmp ret_func9
func_40h:
    mov ah,40h
    mov bx,1
;    mov cx,LENGTHOF message
    mov cx,12
    mov dx,offset message
    int 21h
    jmp ret_func_40h
END

;========================================
;.model small
;.stack 100h
;.386
;
;.data
;message BYTE "hello, world!",0dh,0ah
;
;.code
;main PROC
;    mov ax,@data
;    mov ds,ax
;
;    mov ah,40h
;    mov bx,1
;    mov cx,sizeof message
;    mov dx,offset message
;    int 21h
;
;    .exit
;main ENDP
;END main