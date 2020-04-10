include Irvine32.inc
includelib Irvine32.lib
includelib Kernel32.lib
includelib User32.lib

.data
captionW BYTE "Warning",0
warningMsg BYTE "The current operation may take years "
            BYTE "to complete.",0

captionQ BYTE "Question",0
questionMsg BYTE "A matching user account wat not found."
            BYTE 0dh,0ah,"Do you wish to continue?",0

captionC BYTE "INformation",0
infoMsg BYTE "Select Yes to save a backup file "
        BYTE "before continuing,",0dh,0ah
        BYTE "or click Cancel to stop the operation",0

captionH BYTE "Cannot View User List",0
haltMsg BYTE "This operatioin not support.",0

.code
main PROC
   INVOKE MessageBox,NULL,ADDR warningMsg,
            ADDR captionW,MB_OK+MB_ICONEXCLAMATION
   EXIT
main ENDP
END main
