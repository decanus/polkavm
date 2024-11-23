package polkavm

type Context struct {
	State State
	// @TODO config
}

type Program struct {
}

// InstructionAt returns the instruction at idx
func (p *Program) InstructionAt(idx uint32) (OpCode, error) {
	return 0, nil // @TODO
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

// let pc = context.state.pc
//
//	let skip = program.skip(pc)
//	let inst = program.getInstructionAt(pc: pc)
//
//	guard let inst else {
//	    return .exit(.panic(.invalidInstructionIndex))
//	}
//
//	// TODO: check after GP specified the behavior
//	if context.state.program.basicBlockIndices.contains(pc) {
//	    let blockGas = context.state.program.getBlockGasCosts(pc: pc)
//	    context.state.consumeGas(blockGas)
//	    logger.debug("consumed \(blockGas) gas for block at pc: \(pc)")
//	}
//
//	logger.debug("executing \(inst)", metadata: ["skip": "\(skip)", "pc": "\(context.state.pc)"])
//
//	return inst.execute(context: context, skip: skip)
func step(program Program, ctx Context) error {
	pc := ctx.State.pc

	i, err := program.InstructionAt(pc)
	if err != nil {
		// @TODO RETURN ERR
		return nil
	}
}
