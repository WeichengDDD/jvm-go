package extended

import "jvmgo/instructions/base"

//索引4字节的goto指令
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *interface{}) {
	base.Branch(frame, self.offset)
}
