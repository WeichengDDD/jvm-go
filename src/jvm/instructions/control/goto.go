package control

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

//goto跳转指令
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
