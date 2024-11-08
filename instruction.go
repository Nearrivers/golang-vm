package main

// Fonction d'extension de signe. Cela me sera utile dans le cas des instructions en mode immédiat.
// En mode immédiant, la seconde opérande fait 5 bits de longs et elle sera additionnée à un nombre qui en fait 16.
// Il faut donc étendre la seconde opérande
func signExtend(x uint16, bitCount int) uint16 {
	if (x>>(bitCount-1))&1 == 1 {
		x |= (0xFFFF << bitCount)
	}
	return x
}

// Fonction qui va mettre à jour le flag de condition en fonction du signe du nombre
// dans le registre passé en paramètre
func updateFlags(register uint16) {
	if reg[register] == 0 {
		reg[R_COND] = FL_ZRO
		return
	}

	if reg[register]>>15 == 1 { // Un 1 dans le bit de poids fort indique un nombre négatif
		reg[R_COND] = FL_NEG
		return
	}

	reg[R_COND] = FL_POS
}

func Add(instr uint16) {
	// Registre de destination (DR). Contient l'adresse du registre où l'on va stocker le résultat
	// de l'opération. Selon la spec LC-3, le DR est sur 3 bits à partir de la 9e position.
	// On se positionne donc à la 9e position et on applique un masque de 0x7 soit 0111 afin de récupérer que les
	// 3 derniers bits
	destinationRegister := (instr >> 9) & 0x7
	firstOperand := (instr >> 6) & 0x7
	isItImmediateMode := (instr>>5)&0x1 == 1

	if isItImmediateMode {
		imm5 := signExtend(instr&0x1F, 5)
		reg[destinationRegister] = reg[firstOperand] + imm5
	} else {
		secondOperand := instr & 0x7
		reg[destinationRegister] = reg[firstOperand] + reg[secondOperand]
	}

	updateFlags(destinationRegister)
}

func LoadIndirect(instr uint16) {
	destinationRegister := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	reg[destinationRegister] = memRead(memRead(reg[R_PC] + pcOffset))
	updateFlags(destinationRegister)
}

func And(instr uint16) {
	destinationRegister := (instr >> 9) & 0x7
	firstOperand := (instr >> 6) & 0x7
	fifthBit := (instr >> 5) & 0x1

	if fifthBit == 0 {
		secondOperand := instr & 0x7
		reg[destinationRegister] = firstOperand & secondOperand
	} else {
		imm5 := signExtend(instr&0x1F, 5)
		reg[destinationRegister] = reg[firstOperand] & imm5
	}

	updateFlags(destinationRegister)
}

func Branch(instr uint16) {
	pcOffset := signExtend(instr&0x1FF, 9)
	condFlag := (instr >> 9) & 0x7

	if condFlag != 0 && reg[R_COND] != 0 {
		reg[R_PC] += pcOffset
	}
}

func Jump(instr uint16) {
	baseRegister := (instr >> 6) & 0x7
	reg[R_PC] = baseRegister
}

func JumpToSubroutine(instr uint16) {
	reg[R_R7] = reg[R_PC]
	eleventhBit := (instr >> 11) & 0x1

	if eleventhBit == 0 {
		baseRegister := (instr >> 6) & 0x7
		reg[R_PC] = baseRegister
	} else {
		pcOffset := signExtend(instr&0x7FF, 11)
		reg[R_PC] += pcOffset
	}
}

func Load(instr uint16) {
	destinationRegister := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	reg[destinationRegister] = memRead(reg[R_PC] + pcOffset)
	updateFlags(destinationRegister)
}

func LoadRegister(instr uint16) {
	destinationRegister := (instr >> 9) & 0x7
	baseRegister := (instr >> 6) & 0x7
	offset := signExtend(instr&0x3F, 6)
	reg[destinationRegister] = memRead(reg[baseRegister] + offset)
	updateFlags(destinationRegister)
}

func LoadEffectiveAddress(instr uint16) {
	destinationRegister := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	reg[destinationRegister] = reg[R_PC] + pcOffset
	updateFlags(destinationRegister)
}

func Not(instr uint16) {
	destinationRegister := (instr >> 9) & 0x7
	operand := (instr >> 6) & 0x7
	reg[destinationRegister] = ^operand
	updateFlags(destinationRegister)
}

func Store(instr uint16) {
	destinationRegister := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	writeMem(reg[R_PC]+pcOffset, destinationRegister)
}

func StoreIndirect(instr uint16) {
	destinationRegister := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	writeMem(memRead(reg[R_PC]+pcOffset), destinationRegister)
}

func StoreRegister(instr uint16) {
	destinationRegister := (instr >> 9) & 0x7
	baseRegister := (instr >> 6) & 0x7
	pcOffset := signExtend(instr&0x3F, 6)
	writeMem(baseRegister+pcOffset, destinationRegister)
}

func ExecuteTrap(instr uint16) {
	reg[R_R7] = reg[R_PC]
}
