package rtda

//栈帧
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
}

func newFrame(maxLocals uint, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) Thread() *Thread {
	return self.thread
}
