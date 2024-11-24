package polkavm

// Step represents a Programs steps metadata
type Step struct {
	Skip        uint32
	Instruction OpCode
	Data        []byte
	GasCost     uint64
}

// RA returns the register A value
func (s *Step) RA() uint8 {
	return 0 // @TODO
}

// RB returns the register B value
func (s *Step) RB() uint8 {
	return 0 // @TODO
}

// RD returns the register D value
func (s *Step) RD() uint8 {
	return 0 // @TODO
}

// @TODO pass index
func (s *Step) Immediate() uint32 {
	length := min(len(s.Data), 4)
	if length == 0 {
		return 0
	}

	value := uint32(0)
	for idx, i := range s.Data {
		value = value | uint32(i)<<(8*idx) // @TODO double check
	}

	shift := (4 - length) * 8

	return uint32(int32(value<<shift) >> shift)
}

func (s *Step) Immediate2() uint32 { return 0 }

// Program represents a "parsed" and executable pvm program
type Program struct {
	Blob, Code, BitMask []byte

	JumpTableEntrySize uint8
	JumpTable          []byte

	// parsed and cached steps
	steps []*Step
}

// StepAt returns all information for instruction at idx
func (p *Program) StepAt(idx uint32) (*Step, error) {
	if len(p.steps) <= int(idx) {
		return nil, ErrInvalidStep
	}

	return p.steps[idx], nil
}
