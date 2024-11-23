package polkavm

type Context struct {
	State State
	// @TODO config
}

type Step struct {
	Skip            uint32
	Instruction     OpCode
	RegisterIndexes RegisterIndexes
	GasCost         uint64
}

type Program struct {
}

// StepAt returns all information for instruction at idx
func (p *Program) StepAt(idx uint32) (*Step, error) {
	return nil, nil // @TODO
}

type Registers [13]uint32

// State represents the VM active state
type State struct {
	ProgramCode Program
	Registers   Registers

	gas uint64
	pc  uint32
}

func Execute() {

}

func step(program Program, ctx Context) error {
	pc := ctx.State.pc

	s, err := program.StepAt(pc)
	if err != nil {
		// @TODO RETURN ERR
		return nil
	}

	ctx.State.gas -= s.GasCost

	fn := instructions[s.Instruction]

	return fn(ctx, s.RegisterIndexes, s.Skip)
}
