package polkavm

type RegisterIndexes struct {
	A, B, D uint8
}

// InstructionFunc represents a PVM instruction handler
type InstructionFunc = func(ctx Context, ri RegisterIndexes, skip uint32) error

func ra(data []byte) uint8 {
	return uint8(min(uint64(data[0])&0x1111, 12)) // @TODO double check
}

func rb(data []byte) uint8 {
	return min(data[0]>>4, 12)
}

func rd(data []byte) uint8 {
	return min(data[1], 12)
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
