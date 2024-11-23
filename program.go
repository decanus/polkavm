package polkavm

// Step represents a Programs steps metadata
type Step struct {
	Skip            uint32
	Instruction     OpCode
	Data            []byte
	RegisterIndexes RegisterIndexes // @TODO may not need to parse indexes as they are in data?
	GasCost         uint64
}

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
