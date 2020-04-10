include Macros.inc
includelib Irvine32.lib
includelib Kernel32.lib
includelib User32.lib

IF IsDefined(real_mode)
    include Irvine16.inc
ELSE
    include Irvine32.inc
ENDIF

.code
main proc
    Startup
    mWrite <"this program can be assembled to run ",0dh,0ah>
    mWrite <"in both real mode and protected mode.",0dh,0ah>
main endp
END main
