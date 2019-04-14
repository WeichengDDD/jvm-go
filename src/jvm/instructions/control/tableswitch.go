package control

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

//tableswitch，根据索引跳转控制指令
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

//生成跳转表
func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

//查找并跳转
func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	offset := self.defaultOffset
	if index >= self.low && self.high >= index {
		offset = self.jumpOffsets[index-self.low]
	}

	base.Branch(frame, int(offset))
}
