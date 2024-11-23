package polkavm

import (
	"errors"
	"math"
)

type RegisterIndexes struct {
	A, B, D uint8
}

// InstructionFunc represents a PVM instruction handler
type InstructionFunc = func(ctx Context, ri RegisterIndexes, skip uint32) error

var instructions = map[OpCode]InstructionFunc{
	OpTrap: func(Context, RegisterIndexes, uint32) error {
		return errors.New("panic")
	},
	OpFallthrough: func(ctx Context, _ RegisterIndexes, skip uint32) error {
		ctx.State.pc += skip + 1
		return nil
	},
	OpAdd: func(ctx Context, ri RegisterIndexes, skip uint32) error {
		ra := ctx.State.Registers[ri.A]
		rb := ctx.State.Registers[ri.B]

		ctx.State.Registers[ri.D] = ra + rb

		ctx.State.pc += skip + 1

		return nil // @TODO
	},
	OpSub: func(ctx Context, ri RegisterIndexes, skip uint32) error {
		ra := ctx.State.Registers[ri.A]
		rb := ctx.State.Registers[ri.B]

		ctx.State.Registers[ri.D] = ra - rb

		ctx.State.pc += skip + 1

		return nil // @TODO
	},
	OpAnd: func(ctx Context, ri RegisterIndexes, skip uint32) error {
		ra := ctx.State.Registers[ri.A]
		rb := ctx.State.Registers[ri.B]

		ctx.State.Registers[ri.D] = ra & rb

		ctx.State.pc += skip + 1

		return nil // @TODO
	},
	OpXor: func(ctx Context, ri RegisterIndexes, skip uint32) error {
		ra := ctx.State.Registers[ri.A]
		rb := ctx.State.Registers[ri.B]

		ctx.State.Registers[ri.D] = ra ^ rb

		ctx.State.pc += skip + 1

		return nil // @TODO
	},
	OpOr: func(ctx Context, ri RegisterIndexes, skip uint32) error {
		ra := ctx.State.Registers[ri.A]
		rb := ctx.State.Registers[ri.B]

		ctx.State.Registers[ri.D] = ra | rb

		ctx.State.pc += skip + 1

		return nil // @TODO
	},
	OpMul: func(ctx Context, ri RegisterIndexes, skip uint32) error {
		ra := ctx.State.Registers[ri.A]
		rb := ctx.State.Registers[ri.B]

		ctx.State.Registers[ri.D] = ra * rb

		ctx.State.pc += skip + 1

		return nil // @TODO
	},
	OpDivU: func(ctx Context, ri RegisterIndexes, skip uint32) error {
		ra := ctx.State.Registers[ri.A]
		rb := ctx.State.Registers[ri.B]

		if rb == 0 {
			ctx.State.Registers[ri.D] = math.MaxUint32
		} else {
			ctx.State.Registers[ri.D] = ra / rb
		}

		ctx.State.pc += skip + 1

		return nil
	},
	OpRemU: func(ctx Context, ri RegisterIndexes, skip uint32) error {
		ra := ctx.State.Registers[ri.A]
		rb := ctx.State.Registers[ri.B]

		if rb == 0 {
			ctx.State.Registers[ri.D] = ra
		} else {
			ctx.State.Registers[ri.D] = ra % rb
		}

		ctx.State.pc += skip + 1

		return nil
	},
}

func decodeImmediate(data []byte) uint32 {
	length := min(len(data), 4)
	if length == 0 {
		return 0
	}

	value := 0
	for idx, i := range data {
		value = value | int(i<<(8*idx)) // @TODO double check
	}

	shift := (4 - length) * 8

	return (value << shift) >> shift
}
