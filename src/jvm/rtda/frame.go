package rtda

//栈帧
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func newFrame(maxLocals uint, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
