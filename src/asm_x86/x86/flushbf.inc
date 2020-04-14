.nolist

.model small.stdcall
.stack 200h
.386

;----------------------------------------
; Procedure Prototypes
;----------------------------------------
;extern FlushBuffer; flush input buffer
FlushBuffer PROTO; flush input buffer

includelib flushbf.lib

.list
