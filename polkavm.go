package polkavm

type Context struct {
	State State
	// @TODO config
}

type Program struct {
}

func (p *Program) SkipAt(idx uint32) uint32 {

}

// InstructionAt returns the instruction at idx
func (p *Program) InstructionAt(idx uint32) (OpCode, error) {
	return 0, nil // @TODO
}

// RegisterIndexesAt returns the register for instruction at  indexes at idx
func (p *Program) RegisterIndexesAt(idx uint32) RegisterIndexes {

}

// GasCostAt returns the gas cost for instruction at index at idx
func (p *Program) GasCostAt(idx uint32) uint64 {

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

	i, err := program.InstructionAt(pc)
	if err != nil {
		// @TODO RETURN ERR
		return nil
	}

	ctx.State.gas -= program.GasCostAt(pc)
	ri := program.RegisterIndexesAt(pc)

	switch i {
	case OpAdd:
		return add(ctx, ri, program.SkipAt(pc))
	}

	return nil
}
