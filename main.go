package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		// On montre comment utiliser le programme
		fmt.Println("golang-vm [image-file1] ...")
		os.Exit(2)
	}

	for i := 0; i < len(args); i++ {
		if err := readImage(); err != nil {
			fmt.Printf("Impossible de charger l'image: %s\n", args[i])
			os.Exit(1)
		}
	}

	// Un condition flag doit tout le temps être présent donc on l'initialise au début du programme
	reg[R_COND] = FL_ZRO

	// Position de départ du PC (Program Count)
	// 0x3000 est la valeur par défaut
	const PC_START = 0x3000
	reg[R_PC] = PC_START

	for {
		reg[R_PC]++
		instr := memRead(reg[R_PC])
		op := instr >> 12

	inner:
		switch op {
		case OP_ADD:
			Add(instr)
			break inner
		case OP_AND:
			And(instr)
			break inner
		case OP_BR:
			Branch(instr)
			break inner
		case OP_JMP:
			Jump(instr)
			break inner
		case OP_JSR:
			JumpToSubroutine(instr)
			break inner
		case OP_LD:
			Load(instr)
			break inner
		case OP_LDI:
			LoadIndirect(instr)
			break inner
		case OP_LDR:
			LoadRegister(instr)
			break inner
		case OP_LEA:
			LoadEffectiveAddress(instr)
			break inner
		case OP_NOT:
			Not(instr)
			break inner
		case OP_ST:
			Store(instr)
			break inner
		case OP_STI:
			StoreIndirect(instr)
			break inner
		case OP_STR:
			StoreRegister(instr)
			break inner
		case OP_TRAP:
			ExecuteTrap(instr)
			break inner
		case OP_RTI:
		case OP_RES:
		default:
			fmt.Printf("Opération non prise en charge")
			os.Exit(1)
			break inner
		}
	}
}

func memRead(address uint16) uint16 {
	return 0
}

func writeMem(address uint16, value uint16) {
}

func readImage() error {
	return nil
}
