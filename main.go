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
		if err := readImage; err != nil {
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
		instr := memRead(reg[R_PC] + 1)
		op := instr >> 12

	inner:
		switch op {
		}
	}
}

func memRead(instr uint16) uint16 {
	return 0
}

func readImage() error {
	return nil
}
