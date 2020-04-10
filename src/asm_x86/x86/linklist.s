INCLUDE Irvine32.inc
includelib Irvine32.lib
INCLUDELIB Kernel32.lib
INCLUDELIB User32.lib
include Macros.inc


list_node struct
    val DWORD ?
    next DWORD ?
list_node ends

total_nodes=15
NULL=0
counter=0

.data
linked_list LABEL PTR list_node
REPEAT total_nodes
    counter=counter+1
    list_node<counter,($+counter*SIZEOF(list_node))>
ENDM
list_node<0,0>

.code
main PROC
    mov esi,offset linked_list
nextnode:
    mov eax,(list_node PTR [esi]).next
    cmp eax,NULL
    je quit
    mov eax,(list_node PTR [esi]).val
    call WriteDec
    mWrite <" ",0dh,0ah>
    mov esi,(list_node PTR [esi]).next
    jmp nextnode
quit:
    EXIT
main ENDP
END main