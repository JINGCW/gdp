include Irvine32.inc
includelib Irvine32.lib
includelib Kernel32.lib
includelib User32.lib

walk_max=50
start_x=25
start_y=25

drunkard_walk STRUCT
    path COORD walk_max DUP(<0,0>)
    pathsUsed WORD 0
drunkard_walk ENDS

display_position PROTO currx:WORD,curry:WORD

.data
awalk drunkard_walk<>

.code
main PROC
 mov esi,offset awalk
 call take_drunken_walk
 exit
main ENDP

take_drunken_walk PROC
    LOCAL currx:WORD, curry:WORD
    pushad

    mov edi,esi
    add edi,offset drunkard_walk.path
    mov ecx,walk_max
    mov currx, start_x
    mov curry,start_y
again:
    mov ax,currx
    mov (COORD PTR [edi]).X,ax
    mov ax,curry
    mov (COORD PTR [edi]).Y,ax

    INVOKE display_position,currx,curry

    mov eax,4 ;choose a position
    call RandomRange
    .IF eax==0 ;north
        dec curry
    .ELSEIF eax==1 ;south
        inc curry
    .ELSEIF eax==2 ;west
        dec currx
    .ELSE ;east
        inc currx
    .ENDIF

    add edi,TYPE COORD
    loop again
finish:
    mov (drunkard_walk PTR [esi]).pathsUsed, walk_max
    popad
    ret
take_drunken_walk ENDP

display_position PROC currx:WORD,curry:WORD
.data
comma BYTE ",",0

.code
    pushad
    movzx eax,currx
    call WriteDec
    mov edx,offset comma
    call WriteString
    movzx eax,curry
    call WriteDec
    call Crlf
    popad
    ret
display_position ENDP

END main