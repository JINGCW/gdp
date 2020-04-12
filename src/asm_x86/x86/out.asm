INCLUDE flushbf.inc
INCLUDELIB flushbf.lib

.model small,stdcall
.stack 200h

count=80
KEYBOARD STRUC
    maxinput db count
    inputcount db ?
    buffer db count dup(?)
KEYBOARD ENDS

.data
char db ?
keydb KEYBOARD <>
intputbuffer db 127 dup(?)
bytesread dw ?

.code
main PROC
    ;.startup
    mov ax,@data
    mov ds,ax
    ;call function
;    call func1
;    call func0ah
    call FlushBuffer ;flush buffer character

    call func3fh
ret_func1:
ret_func0ah:
ret_func0bh:
ret_func3fh:
    ;.exit
;    mov ah,2
;    mov dl,keydb.inputcount
;    int 21h

    mov ah,4ch
    mov al,0
    int 21h

func1:
    mov ah,1
    int 21h
    mov char,al
    jmp ret_func1
func0ah:
    mov ah,0ah
    mov dx,offset keydb
    int 21h
;    mov dl,(KEYBOARD PTR [dx]).inputcount
;    mov ah,2
;    mov dl,(KEYBOARD [dx]).inputcount
;    int 21h
    jmp ret_func0ah
func0bh:
    mov ah,0bh
    int 21h
    cmp al,0
    je skip

skip:
    jmp ret_func0bh
func3fh PROC
    mov ah,3fh
    mov bx,0 ;file handle,0->keyboard
    mov cx,127 ;max bytes to read
    mov dx,offset intputbuffer ;ds:dx address of input buffer
    int 21h
    mov bytesread,ax
;    ret_func3fh
func3fh ENDP
main ENDP
END main