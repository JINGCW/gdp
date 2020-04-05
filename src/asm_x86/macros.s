.386
.model flat,stdcall

mPutchar MACRO
    ECHO "Expanding the mPutchar macro"
ENDM

.data
.code
main PROC
    mPutchar
main ENDP
