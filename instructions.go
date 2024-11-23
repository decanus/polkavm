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

func trap(Context, RegisterIndexes, uint32) error {
	return errors.New("panic")
}

func fallthroughExec(ctx Context, _ RegisterIndexes, skip uint32) error {
	ctx.State.pc += skip + 1
	return nil
}

func add(ctx Context, ri RegisterIndexes, skip uint32) error {
	ra := ctx.State.Registers[ri.A]
	rb := ctx.State.Registers[ri.B]

	ctx.State.Registers[ri.D] = ra + rb

	ctx.State.pc += skip + 1

	return nil // @TODO
}

func sub(ctx Context, ri RegisterIndexes, skip uint32) error {
	ra := ctx.State.Registers[ri.A]
	rb := ctx.State.Registers[ri.B]

	ctx.State.Registers[ri.D] = ra - rb

	ctx.State.pc += skip + 1

	return nil // @TODO
}

func and(ctx Context, ri RegisterIndexes, skip uint32) error {
	ra := ctx.State.Registers[ri.A]
	rb := ctx.State.Registers[ri.B]

	ctx.State.Registers[ri.D] = ra & rb

	ctx.State.pc += skip + 1

	return nil // @TODO
}

func xor(ctx Context, ri RegisterIndexes, skip uint32) error {
	ra := ctx.State.Registers[ri.A]
	rb := ctx.State.Registers[ri.B]

	ctx.State.Registers[ri.D] = ra ^ rb

	ctx.State.pc += skip + 1

	return nil // @TODO
}

func or(ctx Context, ri RegisterIndexes, skip uint32) error {
	ra := ctx.State.Registers[ri.A]
	rb := ctx.State.Registers[ri.B]

	ctx.State.Registers[ri.D] = ra | rb

	ctx.State.pc += skip + 1

	return nil // @TODO
}

func mul(ctx Context, ri RegisterIndexes, skip uint32) error {
	ra := ctx.State.Registers[ri.A]
	rb := ctx.State.Registers[ri.B]

	ctx.State.Registers[ri.D] = ra * rb

	ctx.State.pc += skip + 1

	return nil // @TODO
}

func divu(ctx Context, ri RegisterIndexes, skip uint32) error {
	ra := ctx.State.Registers[ri.A]
	rb := ctx.State.Registers[ri.B]

	if rb == 0 {
		ctx.State.Registers[ri.D] = math.MaxUint32
	} else {
		ctx.State.Registers[ri.D] = ra / rb
	}

	ctx.State.pc += skip + 1

	return nil
}

func divs(ctx Context, ri RegisterIndexes, skip uint32) error {
	// @TODO
	return nil
}

func remu(ctx Context, ri RegisterIndexes, skip uint32) error {
	ra := ctx.State.Registers[ri.A]
	rb := ctx.State.Registers[ri.B]

	if rb == 0 {
		ctx.State.Registers[ri.D] = ra
	} else {
		ctx.State.Registers[ri.D] = ra % rb
	}

	ctx.State.pc += skip + 1

	return nil
}
