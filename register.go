package main

const (
	R_R0 uint16 = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC
	R_COND
	R_COUNT

	OP_BR   uint16 = iota /* branch */
	OP_ADD                /* add  */
	OP_LD                 /* load */
	OP_ST                 /* store */
	OP_JSR                /* jump register */
	OP_AND                /* bitwise and */
	OP_LDR                /* load register */
	OP_STR                /* store register */
	OP_RTI                /* unused */
	OP_NOT                /* bitwise not */
	OP_LDI                /* load indirect */
	OP_STI                /* store indirect */
	OP_JMP                /* jump */
	OP_RES                /* reserved (unused) */
	OP_LEA                /* load effective address */
	OP_TRAP               /* execute trap */

	FL_POS = 1 << 0 // Résultat positif
	FL_ZRO = 1 << 1 // Résultat nul
	FL_NEG = 1 << 2 // Résultat négatif
)

// Tout comme la mémoire, on stocke les registres dans un tableau
var reg = make([]uint16, R_COUNT)
