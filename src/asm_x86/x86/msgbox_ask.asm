INCLUDE Irvine32.inc
includelib Irvine32.lib
INCLUDELIB Kernel32.lib
INCLUDELIB User32.lib

.data
caption BYTE "survey completed",0
question byte "thank you for completing the survey."
    byte 0dh,0ah
    byte "would you like to receive the results?",0

.code
main PROC
    mov ebx,offset caption
    mov edx,offset question
    call MsgBoxAsk
    exit
main ENDP
END main