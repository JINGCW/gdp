include Irvine32.inc
includelib Irvine32.lib
includelib Kernel32.lib
includelib User32.lib

TAB=9;ascii code of tab key

.code
main PROC
    call Randomize;initialize random generator
    call Rand1
    call Rand2
    exit
main ENDP

Rand1 PROC
    mov ecx,10;loop 10 times
L1:call Random32;generate random int
    call WriteDec;output by unsigned decimal format
    mov al,TAB;水平制表符
    call WriteChar;输出制表符
    loop L1

    call Crlf
    ret
Rand1 ENDP

Rand2 PROC
    ;generate 10 pseudo-random integers from -50 to +49
    mov ecx,10;loop 10 times
L1:mov eax,100;values 0-99
    call RandomRange;generate random int
    sub eax,50;values -50 to +49
    call WriteInt ;write signed decimal
    mov al,TAB ;horizontal tab
    call WriteChar ;write the tab
    loop L1

    call Crlf
    ret
Rand2 ENDP

END main
