; MsgBox demo                     (msgbox1.asm)
;ml /coff msgbox1.asm /link /subsystem:windows

Include Irvine32.inc
Includelib Irvine32.lib 
Includelib kernel32.lib 
Includelib user32.lib
 
 
.data
	msg BYTE "Plus_RE!",0
	msgTitle BYTE "Test MsgBox",0
	
.code 
main PROC
 
	mov edx, OFFSET msg
	mov ebx, OFFSET msgTitle
	call MsgBox
	
	exit
	
main ENDP
END main
