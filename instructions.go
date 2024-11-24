package polkavm

import (
	"errors"
	"math"
)

// InstructionFunc represents a PVM instruction handler
type InstructionFunc = func(ctx Context, s *Step) error

var instructions = map[OpCode]InstructionFunc{
	OpTrap: func(Context, *Step) error {
		return errors.New("panic")
	},
	OpFallthrough: func(ctx Context, s *Step) error {
		ctx.State.pc += s.Skip + 1
		return nil
	},
	OpStoreImmU8: func(ctx Context, s *Step) error {
		addr, value := s.Immediate2()

		err := ctx.State.Memory.Write(addr, []byte{uint8(value)}) // @TODO check
		if err != nil {
			return err
		}

		ctx.State.pc += s.Skip + 1

		return nil
	},
	OpStoreImmU16: func(ctx Context, s *Step) error {
		//addr, value := s.Immediate2()
		//
		// @Todo encoding
		//err := ctx.State.Memory.Write(addr, uint16(value)) // @TODO check
		//if err != nil {
		//	return err
		//}

		ctx.State.pc += s.Skip + 1

		return nil
	},
	OpStoreImmU32: func(ctx Context, s *Step) error {
		//addr, value := s.Immediate2()
		//
		// @Todo encoding
		//err := ctx.State.Memory.Write(addr, uint16(value)) // @TODO check
		//if err != nil {
		//	return err
		//}

		ctx.State.pc += s.Skip + 1

		return nil
	},
	OpJump: func(ctx Context, s *Step) error {
		offset := s.Immediate()

		// @TODO check offset validity
		ctx.State.pc += offset
		return nil
	},
	OpAdd: func(ctx Context, s *Step) error {
		ra := ctx.State.Registers[s.RA()]
		rb := ctx.State.Registers[s.RB()]

		ctx.State.Registers[s.RD()] = ra + rb

		ctx.State.pc += s.Skip + 1

		return nil // @TODO
	},
	OpSub: func(ctx Context, s *Step) error {
		ra := ctx.State.Registers[s.RA()]
		rb := ctx.State.Registers[s.RB()]

		ctx.State.Registers[s.RD()] = ra - rb

		ctx.State.pc += s.Skip + 1

		return nil // @TODO
	},
	OpAnd: func(ctx Context, s *Step) error {
		ra := ctx.State.Registers[s.RA()]
		rb := ctx.State.Registers[s.RB()]

		ctx.State.Registers[s.RD()] = ra & rb

		ctx.State.pc += s.Skip + 1

		return nil // @TODO
	},
	OpXor: func(ctx Context, s *Step) error {
		ra := ctx.State.Registers[s.RA()]
		rb := ctx.State.Registers[s.RB()]

		ctx.State.Registers[s.RD()] = ra ^ rb

		ctx.State.pc += s.Skip + 1

		return nil // @TODO
	},
	OpOr: func(ctx Context, s *Step) error {
		ra := ctx.State.Registers[s.RA()]
		rb := ctx.State.Registers[s.RB()]

		ctx.State.Registers[s.RD()] = ra | rb

		ctx.State.pc += s.Skip + 1

		return nil // @TODO
	},
	OpMul: func(ctx Context, s *Step) error {
		ra := ctx.State.Registers[s.RA()]
		rb := ctx.State.Registers[s.RB()]

		ctx.State.Registers[s.RD()] = ra * rb

		ctx.State.pc += s.Skip + 1

		return nil // @TODO
	},
	OpDivU: func(ctx Context, s *Step) error {
		ra := ctx.State.Registers[s.RA()]
		rb := ctx.State.Registers[s.RB()]

		if rb == 0 {
			ctx.State.Registers[s.RD()] = math.MaxUint32
		} else {
			ctx.State.Registers[s.RD()] = ra / rb
		}

		ctx.State.pc += s.Skip + 1

		return nil
	},
	OpRemU: func(ctx Context, s *Step) error {
		ra := ctx.State.Registers[s.RA()]
		rb := ctx.State.Registers[s.RB()]

		if rb == 0 {
			ctx.State.Registers[s.RD()] = ra
		} else {
			ctx.State.Registers[s.RD()] = ra % rb
		}

		ctx.State.pc += s.Skip + 1

		return nil
	},
}
