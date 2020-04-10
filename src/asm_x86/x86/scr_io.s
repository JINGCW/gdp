include Irvine32.inc
includelib Irvine32.lib
includelib Kernel32.lib
includelib User32.lib

.data
COUNT=4
blue_text_on_gray=blue+(lightGray*16)
default_color=lightGray+(black*16)

arrayD SDWORD 12345678h,1A4B2000H,3434H,7AB9H
prompt BYTE "enter a 32-bit signed integer: ",0

.code
main PROC
    MOV EAX,blue_text_on_gray
    CALL SetTextColor
    CALL Clrscr
    MOV ESI,OFFSET arrayD
    MOV EBX,TYPE arrayD
    MOV ECX,LENGTHOF arrayD
    CALL DumpMem
    CALL Crlf
    MOV ECX,COUNT
L1:
    MOV EDX,OFFSET prompt
    CALL WriteString
    CALL ReadInt
    CALL Crlf
    CALL WriteInt
    CALL Crlf
    CALL WriteHex
    CALL Crlf
    CALL WriteBin
    CALL Crlf
    Loop L1
    CALL WaitMsg ;"press any key..."
    MOV EAX,default_color
    CALL SetTextColor
    CALL Clrscr
    EXIT
main ENDP
END main
