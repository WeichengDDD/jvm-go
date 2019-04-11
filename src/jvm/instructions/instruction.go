package instructions

import (
	"jvm/rtda"
	"jvmgo/instructions/base"
)

type Instruction interface {
	FetchOperands(reader *base.BytecodeReader)
	Execute(frame *rtda.Frame)
}

//无操作数指令
type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *base.BytecodeReader) {
	panic("implement me")
}

//跳转指令
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *base.BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

//读取变量表指令
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadInt16())
}
