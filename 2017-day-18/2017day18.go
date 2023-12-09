package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	SND int8 = iota
	SET
	ADD
	MUL
	MOD
	RCV
	JGZ
)

type Instruction struct {
	operation int8
	arg1      string
	arg2      string
}

func (i Instruction) String() string {
	var fmtStr string
	switch i.operation {
	case SND:
		fmtStr = "SND(%s)"
	case SET:
		fmtStr = "SET(%s)"
	case ADD:
		fmtStr = "ADD(%s)"
	case MUL:
		fmtStr = "MUL(%s)"
	case MOD:
		fmtStr = "MOD(%s)"
	case RCV:
		fmtStr = "RCV(%s)"
	case JGZ:
		fmtStr = "JGZ(%s)"
	}
	var args string
	if i.arg2 == "" {
		args = i.arg1
	} else {
		args = fmt.Sprintf("%s, %s", i.arg1, i.arg2)
	}
	return fmt.Sprintf(fmtStr, args)
}

func validArg(arg string) (string, error) {
	if len(arg) == 1 && unicode.IsLetter(rune(arg[0])) {
		return arg, nil
	}
	_, err := strconv.Atoi(arg)
	if err == nil {
		return arg, nil
	}
	return "", fmt.Errorf("'%s' not a valid arg", arg)
}

func parseInstruction(line string) (Instruction, error) {
	instruction := Instruction{}
	instructionParts := strings.Fields(line)

	switch len(instructionParts) {
	case 2:
		arg, err := validArg(instructionParts[1])
		if err != nil {
			return instruction, fmt.Errorf("invalid instruction '%s': %v", line, err)
		}
		instruction.arg1 = arg
	case 3:
		arg, err := validArg(instructionParts[1])
		if err != nil {
			return instruction, fmt.Errorf("invalid instruction '%s': %v", line, err)
		}
		instruction.arg1 = arg
		arg, err = validArg(instructionParts[2])
		if err != nil {
			return instruction, fmt.Errorf("invalid instruction '%s': %v", line, err)
		}
		instruction.arg2 = arg
	default:
		return instruction, fmt.Errorf("invalid instruction '%s'", line)
	}

	switch instructionParts[0] {
	case "snd":
		instruction.operation = SND
	case "set":
		instruction.operation = SET
	case "add":
		instruction.operation = ADD
	case "mul":
		instruction.operation = MUL
	case "mod":
		instruction.operation = MOD
	case "rcv":
		instruction.operation = RCV
	case "jgz":
		instruction.operation = JGZ
	default:
		return instruction, fmt.Errorf("unknown operation '%s'", instructionParts[0])
	}
	return instruction, nil
}

func isRegisterName(arg string) bool {
	return len(arg) == 1 && unicode.IsLetter(rune(arg[0]))
}

func lookupOrParse(arg string, registers map[string]int) (int, error) {
	var value int
	var err error
	if isRegisterName(arg) {
		val, exists := registers[arg]
		if !exists {
			return value, errors.New("argument is a register but it's not set")
		}
		return val, nil
	}
	value, err = strconv.Atoi(arg)
	if err != nil {
		return value, fmt.Errorf("argument is a value but couldn't parse: %v", err)
	}
	return value, nil
}

func executeInstructionsPart1(instructions []Instruction) {
	fmt.Println("\nStart execution...")
	registers := make(map[string]int)
	nextInstructionIndex := 0

	lastPlayedFrequency := 0
	for nextInstructionIndex >= 0 && nextInstructionIndex < len(instructions) {
		nextInstruction := instructions[nextInstructionIndex]
		fmt.Printf("%d: %v\n", nextInstructionIndex, nextInstruction)
		fmt.Printf("Registers: %v\n", registers)
		switch nextInstruction.operation {
		case SND:
			freq, err := lookupOrParse(nextInstruction.arg1, registers)
			if err != nil {
				log.Fatalf("Bad instruction %s at index %d: %v\n", nextInstruction, nextInstructionIndex, err)
			}
			fmt.Printf("<))) Playind sound at frequency %d\n", freq)
			lastPlayedFrequency = freq
		case SET:
			err := updateRegister(nextInstruction.arg1, nextInstruction.arg2, registers, func(a, b int) int { return b })
			if err != nil {
				log.Fatalf("Bad instruction %s at index %d: %v\n", nextInstruction, nextInstructionIndex, err)
			}
		case ADD:
			err := updateRegister(nextInstruction.arg1, nextInstruction.arg2, registers, func(a, b int) int { return a + b })
			if err != nil {
				log.Fatalf("Bad instruction %s at index %d: %v\n", nextInstruction, nextInstructionIndex, err)
			}
		case MUL:
			err := updateRegister(nextInstruction.arg1, nextInstruction.arg2, registers, func(a, b int) int { return a * b })
			if err != nil {
				log.Fatalf("Bad instruction %s at index %d: %v\n", nextInstruction, nextInstructionIndex, err)
			}
		case MOD:
			err := updateRegister(nextInstruction.arg1, nextInstruction.arg2, registers, func(a, b int) int {
				return a % b
			})
			if err != nil {
				log.Fatalf("Bad instruction %s at index %d: %v\n", nextInstruction, nextInstructionIndex, err)
			}
		case RCV:
			val, err := lookupOrParse(nextInstruction.arg1, registers)
			if err != nil {
				log.Fatalf("Bad instruction %s at index %d: %v\n", nextInstruction, nextInstructionIndex, err)
			}
			if val > 0 {
				fmt.Printf("Recovering frequency: %d\n", lastPlayedFrequency)
				// this is the thing we are looking, so exit immediately
				os.Exit(0)
			}
		case JGZ:
			val, err := lookupOrParse(nextInstruction.arg1, registers)
			if err != nil {
				log.Fatalf("Bad instruction %s at index %d: %v\n", nextInstruction, nextInstructionIndex, err)
			}
			jumpVal, err := lookupOrParse(nextInstruction.arg2, registers)
			if err != nil {
				log.Fatalf("Bad instruction %s at index %d: %v\n", nextInstruction, nextInstructionIndex, err)
			}
			if val > 0 {
				nextInstructionIndex += jumpVal
				continue
			}
		}
		nextInstructionIndex++
	}
}

func updateRegister(registerName, arg string, registers map[string]int, f func(int, int) int) error {
	val, err := lookupOrParse(arg, registers)
	if err != nil {
		return err
	}
	if isRegisterName(registerName) {
		registers[registerName] = f(registers[registerName], val)
	} else {
		return fmt.Errorf("invalid register name '%s'", registerName)
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("bad arguments: exactly 1 arg (path) expected")
	}
	path := os.Args[1]

	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	instructions := make([]Instruction, 0, 10)
	for s.Scan() {
		line := s.Text()

		fmt.Println(line)
		instruction, err := parseInstruction(line)
		if err != nil {
			log.Printf("Couldn't parse line: %v", err)
			continue
		}
		instructions = append(instructions, instruction)
	}

	executeInstructionsPart1(instructions)
}
