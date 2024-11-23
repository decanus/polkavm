package polkavm

import "errors"

var (
	ErrOutOfGas    = errors.New("out of gas")
	ErrInvalidStep = errors.New("invalid step")
)

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
	steps []*Step
}

// StepAt returns all information for instruction at idx
func (p *Program) StepAt(idx uint32) (*Step, error) {
	if len(p.steps) <= int(idx) {
		return nil, ErrInvalidStep
	}

	return p.steps[idx], nil
}

type Registers [13]uint32

// State represents the VM active state
type State struct {
	ProgramCode Program
	Registers   Registers

	gas uint64
	pc  uint32
}

// Execute executes a pvm Program using the supplied Context
func Execute(program Program, ctx Context) error {

	for {
		if ctx.State.gas <= 0 {
			return ErrOutOfGas
		}

		err := step(program, ctx)
		if err != nil {
			return err
		}

		// @TODO think about successful returns
		// @TODO host calls

	}
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
