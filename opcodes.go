package polkavm

type OpCode uint8

const (
	OpTrap        = OpCode(0)
	OpFallthrough = OpCode(17)
	OpEcalli      = OpCode(78)

	OpStoreImmU8  = OpCode(62)
	OpStoreImmU16 = OpCode(79)
	OpStoreImmU32 = OpCode(38)

	OpOpJumpInd = OpCode(19)

	OpLoadImm = OpCode(4)
	OpLoadU8  = OpCode(60)
	OpLoadI8  = OpCode(74)
	OpLoadU16 = OpCode(76)
	OpLoadI16 = OpCode(66)
	OpLoadU32 = OpCode(10)

	OpStoreU8  = OpCode(71)
	OpStoreU16 = OpCode(69)
	OpStoreU32 = OpCode(22)

	OpStoreImmIndU8  = OpCode(26)
	OpStoreImmIndU16 = OpCode(54)
	OpStoreImmIndU32 = OpCode(13)

	OpLoadImmJump = OpCode(6)

	OpBranchEqImm  = OpCode(6)
	OpBranchNeImm  = OpCode(15)
	OpBranchLtUImm = OpCode(44)
	OpBranchLeUImm = OpCode(59)
	OpBranchGeUImm = OpCode(52)
	OpBranchGtUImm = OpCode(50)
	OpBranchLtSImm = OpCode(32)
	OpBranchLeSImm = OpCode(46)
	OpBranchGeSImm = OpCode(45)
	OpBranchGtSImm = OpCode(53)

	OpMoveReg = OpCode(82)
	OpSbrk    = OpCode(87)

	OpStoreIndU8  = OpCode(16)
	OpStoreIndU16 = OpCode(29)
	OpStoreIndU32 = OpCode(3)
	OpLoadIndU8   = OpCode(11)
	OpLoadIndI8   = OpCode(21)
	OpLoadIndU16  = OpCode(37)
	OpLoadIndI16  = OpCode(33)
	OpLoadIndU32  = OpCode(1)

	OpAddImm = OpCode(2)
	OpAndImm = OpCode(18)
	OpXorImm = OpCode(31)
	OpOrImm  = OpCode(49)

	OpMulImm        = OpCode(35)
	OpMulUpperSSImm = OpCode(65)
	OpMulUpperUUImm = OpCode(63)

	OpSetLtUImm      = OpCode(27)
	OpSetLtSImm      = OpCode(56)
	OpShloLImm       = OpCode(9)
	OpShloRImm       = OpCode(14)
	OpSharRImm       = OpCode(25)
	OpNegAddImm      = OpCode(40)
	OpSetGtUImm      = OpCode(39)
	OpSetGtSImm      = OpCode(61)
	OpShloLImmAlt    = OpCode(75)
	OpShloRImmAlt    = OpCode(72)
	OpSharRImmAlt    = OpCode(80)
	OpCmovIzImm      = OpCode(85)
	OpCmovNzImm      = OpCode(86)
	OpBranchEq       = OpCode(24)
	OpBranchNe       = OpCode(30)
	OpBranchLtU      = OpCode(47)
	OpBranchLtS      = OpCode(48)
	OpBranchGeU      = OpCode(41)
	OpBranchGeS      = OpCode(43)
	OpLoadImmJumpInd = OpCode(42)
	OpAdd            = OpCode(8)
	OpSub            = OpCode(20)
	OpAnd            = OpCode(23)
	OpXor            = OpCode(28)
	OpOr             = OpCode(12)
	OpMul            = OpCode(84)
	OpMulUpperSS     = OpCode(67)
	OpMulUpperUU     = OpCode(57)
	OpMulUpperSU     = OpCode(81)
	OpDivU           = OpCode(68)
	OpDivS           = OpCode(64)
	OpRemU           = OpCode(73)
	OpRemS           = OpCode(70)
	OpSetLtU         = OpCode(36)
	OpSetLtS         = OpCode(58)
	OpShloL          = OpCode(55)
	OpShloR          = OpCode(51)
	OpSharR          = OpCode(77)
	OpCmovIz         = OpCode(83)
	OpOpCmovNz       = OpCode(84)
)

var instructions = map[OpCode]InstructionFunc{
	OpAdd: add,
	OpSub: sub,
	OpAnd: and,
	OpXor: xor,
	OpOr:  or,
	OpMul: mul,
}
