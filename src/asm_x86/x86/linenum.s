
mul32 MACRO op1,op2,product
    IFIDNI <op2>,<EAX>
        linenum TEXTEQU %(@LINE)
        ECHO ---------------------
        ECHO * Error on line linenum: EAX cannot be the second
        ECHO * argument when invoking the mul32 macro.
        ECHO ----------------------
    EXITM
    ENDIF

    PUSH EAX
    MOV EAX,op1
    MUL op2
    MOV product,EAX
    POP EAX
ENDM
