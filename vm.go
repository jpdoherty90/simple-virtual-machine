package vm

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
//
func compute(memory []byte) {

	registers := [3]byte{8, 0, 0} // PC, R1 and R2
	pc := registers[0]

	// Keep looping, like a physical computer's clock
	for {

		op := memory[pc]

		// decode and execute
		switch op {
		case Load:
			registers[memory[pc+1]] = memory[memory[pc+2]]
		case Store:
			memory[memory[pc+2]] = registers[memory[pc+1]]
		case Add:
			registers[memory[pc+1]] += registers[memory[pc+2]]
		case Sub:
			registers[memory[pc+1]] -= registers[memory[pc+2]]
		case Halt:
			return
		case Addi:
			registers[memory[pc+1]] += memory[pc+2]
		case Subi:
			registers[memory[pc+1]] -= memory[pc+2]
		case Jump:
			pc = memory[pc+1]
			continue
		case Beqz:
			if registers[memory[pc+1]] == 0x00 {
				pc += memory[pc+2]
			}
		}
		pc += 3

	}

}
